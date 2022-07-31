package did

import (
	"bufio"
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

func AuthLogin() (ed25519.PrivateKey, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---Login---")
	fmt.Print("Enter DID: ")
	did, _ := reader.ReadString('\n')
	did = strings.Replace(did, "\n", "", -1)

	var priv ed25519.PrivateKey
	if _, err := os.Stat(DidFilename(did)); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No existing keystore file for this DID. Creating one now.")
		fmt.Print("Private Key: ")
		privStr, _ := reader.ReadString('\n')
		privStr = strings.Replace(privStr, "\n", "", -1)
		priv = []byte(privStr)
		fmt.Print("Create Password: ")
		pw, _ := reader.ReadString('\n')
		pw = strings.Replace(pw, "\n", "", -1)
		_ = WriteKeystore(did, []byte(priv), pw)
		fmt.Printf("Keystore has been created. %s is now the node's default signer.\n", did)
	} else {
		fmt.Print("Enter Password: ")
		pw, _ := reader.ReadString('\n')
		pw = strings.Replace(pw, "\n", "", -1)
		priv, err = ReadKeystore(did, pw)
		if err != nil {
			return nil, err
		}
	}

	return priv, nil
}

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

	block, err := x509.EncryptPEMBlock(rand.Reader, "PRIVATE KEY", bytes, []byte(pw), x509.PEMCipher3DES)
	if err != nil {
		return err
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

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New(".pem file contained no blocks")
	}

	privRaw, err := x509.DecryptPEMBlock(block, []byte(pw))
	priv, err := x509.ParsePKCS8PrivateKey(privRaw)
	if err != nil {
		return nil, err
	}

	edPriv, ok := priv.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("Type of key was not ed25519")
	}

	return edPriv, nil
}

func DeleteKeystore(did string) error {
	return os.Remove(DidFilename(did))
}
