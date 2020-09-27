package key_gen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func GeneratePair(c *Config) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, c.bits)
	if err != nil {
		return nil, nil, fmt.Errorf("generating key caused error: %w", err)
	}

	return key, &key.PublicKey, nil
}

func CreatePrivatePEM(c *Config, key *rsa.PrivateKey) ([]byte, error) {
	var err error
	block := &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Bytes:   x509.MarshalPKCS1PrivateKey(key),
	}

	if c.passphrase != "" {
		block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(c.passphrase), x509.PEMCipherAES256)
		if err != nil {
			return nil, err
		}
	}

	return pem.EncodeToMemory(block), nil
}

func CreatePublicPEM(key *rsa.PublicKey) ([]byte, error) {
	b, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return nil, fmt.Errorf("marshaling public key returned error: %w", err)
	}

	return pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: b,
	}), nil
}


func PrintKeys(c *Config, private *rsa.PrivateKey, public *rsa.PublicKey) error {
	var err error
	if private == nil {
		private, public, err = GeneratePair(c)
		if err != nil {
			return fmt.Errorf("error occured when generating keys: %w", err)
		}
	}

	privB, err := CreatePrivatePEM(c, private)
	if err != nil {
		return fmt.Errorf("error occured when creating PEM out of private key: %w", err)
	}

	pubB, err := CreatePublicPEM(public)
	if err != nil {
		return fmt.Errorf("error occured when creating PEM out of public key: %w", err)
	}

	fmt.Println("Private key:")
	fmt.Println(string(privB))

	fmt.Println("\nPublic key:")
	fmt.Println(string(pubB))

	return nil
}