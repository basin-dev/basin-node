package did

import (
	"context"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"

	"github.com/ucan-wg/go-ucan"
)

func tokenFilename(url string) string {
	home, _ := os.UserHomeDir()
	return home + "/.basin/ucan/root_" + strings.ReplaceAll(url, "basin://", "") + ".pem"
}

type PersistantTokenStore struct {
	pw string
}

func (s PersistantTokenStore) PutToken(ctx context.Context, key string, rawToken string) error {
	home, _ := os.UserHomeDir()
	os.Chdir(home)
	os.Mkdir(".basin", os.ModePerm)
	os.Chdir(".basin")
	os.Mkdir("ucan", os.ModePerm)
	file, err := os.Create(tokenFilename(key))
	defer file.Close()
	if err != nil {
		return err
	}

	// Can a .pem file encode arbitrary strings? Is this acceptable practice?
	block, err := x509.EncryptPEMBlock(rand.Reader, "PRIVATE KEY", []byte(rawToken), []byte(s.pw), x509.PEMCipher3DES)
	if err != nil {
		return err
	}

	_, err = file.Write(pem.EncodeToMemory(block))
	if err != nil {
		return err
	}

	return nil
}
func (s PersistantTokenStore) RawToken(ctx context.Context, key string) (rawToken string, err error) {
	data, err := os.ReadFile(tokenFilename(key))
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return "", fmt.Errorf(".pem file for resource %s contained no blocks", key)
	}

	tokenBytes, err := x509.DecryptPEMBlock(block, []byte(s.pw))
	if err != nil {
		return "", err
	}

	return string(tokenBytes), nil
}
func (s PersistantTokenStore) DeleteToken(ctx context.Context, key string) (err error) {
	return os.Remove(tokenFilename(key))
}
func (s PersistantTokenStore) ListTokens(ctx context.Context, offset, limit int) (results []ucan.RawToken, err error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (s PersistantTokenStore) GetParsed(ctx context.Context, url string) (*ucan.Token, error) {
	raw, err := s.RawToken(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("Error getting raw token from store: %w\n", err)
	}

	return ParseToken([]byte(raw))
}

var (
	memStore        ucan.TokenStore
	PersistantStore PersistantTokenStore
	caps            ucan.NestedCapabilities
)

func init() {
	memStore = ucan.NewMemTokenStore()
	caps = ucan.NewNestedCapabilities("SUPER", "WRITE", "READ")
}

func CreateRootUcanToken(resourceUrl string, privKey crypto.PrivKey, did string) (*ucan.Token, error) {
	source, err := ucan.NewPrivKeySource(privKey)
	if err != nil {
		return nil, fmt.Errorf("Error creating new private key source for UCAN: %w\n", err)
	}

	att := ucan.Attenuations{
		{Cap: caps.Cap("SUPER"), Rsc: ucan.NewStringLengthResource("basin", resourceUrl)},
		{Cap: caps.Cap("WRITE"), Rsc: ucan.NewStringLengthResource("basin", resourceUrl)},
		{Cap: caps.Cap("READ"), Rsc: ucan.NewStringLengthResource("basin", resourceUrl)},
	}
	return source.NewOriginToken(did, att, nil, time.Now(), time.Date(9999999999999, 1, 1, 1, 1, 1, 1, time.UTC))
}

func CreateAttenuatedToken(ctx context.Context, resourceUrl string, privKey crypto.PrivKey, audienceDid string, expiration time.Time, actions []string) (*ucan.Token, error) {
	source, err := ucan.NewPrivKeySource(privKey)
	if err != nil {
		return nil, fmt.Errorf("Error creating new private key source for UCAN: %w\n", err)
	}

	att := ucan.Attenuations{}
	// Only READ and WRITE for now
	for _, action := range actions {
		if action == "READ" || action == "WRITE" {
			att = append(att, ucan.Attenuation{Cap: caps.Cap(action), Rsc: ucan.NewStringLengthResource("basin", resourceUrl)})
		}
	}

	parentToken, err := PersistantStore.GetParsed(ctx, resourceUrl)
	if err != nil {
		return nil, fmt.Errorf("Error reading parent token from UCAN store: %w\n", err)
	}

	return source.NewAttenuatedToken(parentToken, audienceDid, att, nil, time.Now(), expiration)
}

func StartPersistantTokenStore(pw string) *PersistantTokenStore {
	PersistantStore = PersistantTokenStore{pw}
	return &PersistantStore
}

func ParseToken(raw []byte) (*ucan.Token, error) {
	store := ucan.NewMemTokenStore()

	// ngl, no idea what this is...
	ac := func(m map[string]interface{}) (ucan.Attenuation, error) {
		var (
			cap string
			rsc ucan.Resource
		)
		for key, vali := range m {
			val, ok := vali.(string)
			if !ok {
				return ucan.Attenuation{}, fmt.Errorf(`expected attenuation value to be a string`)
			}

			if key == ucan.CapKey {
				cap = val
			} else {
				rsc = ucan.NewStringLengthResource(key, val)
			}
		}

		return ucan.Attenuation{
			Rsc: rsc,
			Cap: caps.Cap(cap),
		}, nil
	}

	// Instantiate new token parser. What is the attenuation constructor (ac)?
	p := ucan.NewTokenParser(ac, ucan.StringDIDPubKeyResolver{}, store.(ucan.CIDBytesResolver))

	// Ingest raw token and verify
	res, err := p.ParseAndVerify(context.Background(), string(raw))
	if err != nil {
		log.Fatal(err)
	}

	return res, nil
}
