package did

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/multiformats/go-multibase"
	. "github.com/ockam-network/did"
)

type KeyResolver struct{}

type invalidDidError struct {
	did string
}

func (i invalidDidError) Error() string {
	return fmt.Sprintf("DID %s is invalid.", i.did)
}

type decodedPublicKey struct {
	MulticodecValue   multibase.Encoding
	RawPublicKeyBytes []byte
}

/* Returns decoded public key given base58 BTC multibase encoded key */
func decodePublicKey(multibaseVal string, options string) (decodedPublicKey, error) {
	pub := decodedPublicKey{}
	enc, data, err := multibase.Decode(multibaseVal)
	if err != nil {
		return pub, err
	}
	if enc != 'z' {
		return pub, errors.New("Multibase prefix is invalid: should be z (base58 BTC)")
	}
	pub.MulticodecValue = enc // TODO: WHAT IS THE HEADER?!?!?!?!? https://w3c-ccg.github.io/did-method-key/#decode-public-key-algorithm:~:text=decodedPublicKey%20to%20an%20empty%20object.-,Decode%20multibaseValue%20using%20the%20base58%2Dbtc%20multibase%20alphabet%20and%20set%20multicodecValue%20to%20the%20multicodec%20header%20for%20the%20decoded%20value.%20Implementers%20are%20cautioned%20to%20ensure%20that%20the%20multicodecValue%20is%20set%20to%20the%20result%20after%20performing%20varint%20decoding.,-Set%20the%20rawPublicKeyBytes%20to%20the
	pub.RawPublicKeyBytes = data
	return pub, nil
}

// Based on https://w3c-ccg.github.io/did-method-key/#signature-method-creation-algorithm
var KEY_LENGTH_TABLE = map[int]int{
	0xe7:   33,
	0xec:   32,
	0xed:   32,
	0x1200: 33,
	0x1201: 49,
	0x1202: 0, // ??
	0x1205: 0, // ??
}

func createSignature(did string, multibaseVal string, options string) (VerificationMethod, error) {
	vm := VerificationMethod{}
	pub, err := decodePublicKey(multibaseVal, options)
	if err != nil {
		return vm, err
	}
	// Check length of the decoded key
	if len(pub.RawPublicKeyBytes) != KEY_LENGTH_TABLE[int(pub.MulticodecValue)] {
		return vm, errors.New("Invalid public key length.")
	}

	vm.Id = fmt.Sprintf("%s#%s", did, pub.MulticodecValue)

	return vm, nil
}

// https://w3c-ccg.github.io/did-method-key/#document-creation-algorithm
func (k KeyResolver) Resolve(did string) ([]byte, error) {
	parsed, err := Parse(did)
	if err != nil {
		return nil, err
	}
	if parsed.Method != "key" {
		return nil, invalidDidError{did}
	}
	// version := 1

	key := parsed.ID
	if key[0] != 'z' {
		return nil, invalidDidError{did}
	}

	withFrag := fmt.Sprintf("%s#%s", did, key)

	// TODO: This is a cheap way of doing this...finish the above functions
	rawJson := fmt.Sprintf(`{
		"@context": [
		  "https://www.w3.org/ns/did/v1",
		  "https://w3id.org/security/suites/ed25519-2020/v1",
		  "https://w3id.org/security/suites/x25519-2020/v1"
		],
		"id": "%s",
		"verificationMethod": [{
		  "id": "%s",
		  "type": "Ed25519VerificationKey2020",
		  "controller": "%s",
		  "publicKeyMultibase": "%s"
		}],
		"authentication": [
		  "%s"
		],
		"assertionMethod": [
		  "%s"
		],
		"capabilityDelegation": [
		  "%s"
		],
		"capabilityInvocation": [
		  "%s"
		],
		"keyAgreement": [{
		  "id": "%s",
		  "type": "X25519KeyAgreementKey2020",
		  "controller": "%s",
		  "publicKeyMultibase": "%s"
		}]
	  }`, did, withFrag, did, key, withFrag, withFrag, withFrag, withFrag, withFrag, did, key)

	doc, err := json.Marshal(rawJson)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
