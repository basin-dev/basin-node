package did

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
)

const KEYSTORE_PATH = "~/.basin/keystores/"

/* Generate and return a new DID, storing its private key in a keyfile, encoded with the given password. */
func NewPrivateKey(pw string) (string, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", err
	}

	did := fmt.Sprintf("did:key:%s", string(pub))

	err = WriteKeystore(did, priv, pw)
	if err != nil {
		return did, err
	}

	return did, nil
}

func DidFilename(did string) string {
	return KEYSTORE_PATH + strings.ReplaceAll(did, ":", "_") + ".pem"
}

func WriteKeystore(did string, priv ed25519.PrivateKey, pw string) error {
	file, err := os.Create(DidFilename(did))
	defer file.Close()
	if err != nil {
		return err
	}

	bytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return err
	}

	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: bytes,
	}

	_, err = file.Write(pem.EncodeToMemory(block))
	if err != nil {
		return err
	}

	return nil
}

func ReadKeystore(did string, pw string) (ed25519.PrivateKey, error) {
	data, err := os.ReadFile(DidFilename(did))
	if err != nil {
		return nil, err
	}

	p, _ := pem.Decode(data)
	if p == nil {
		return nil, errors.New(".pem file contained no blocks")
	}

	priv, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	if err != nil {
		return nil, err
	}

	edPriv, ok := priv.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("Type of key was not ed25519")
	}

	return edPriv, nil
}
