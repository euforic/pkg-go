package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

// Rsa...
type Rsa struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
}

// New creates a new instance of *Rsa
func New() *Rsa {
	return &Rsa{}
}

// Generate gerates the public and private keys and sets them
func (r *Rsa) Generate() error {
	var err error
	// generate key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return errors.New("cannot generate RSA key")
	}
	r.Private = privateKey
	r.Public = &privateKey.PublicKey

	return nil
}

// WritePrivateKey encodes the rsa.Private key to the io.Writer
func (r Rsa) WritePrivateKey(w io.Writer) error {
	if r.Private == nil {
		return errors.New("no private key set")
	}
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(r.Private),
	}

	if err := pem.Encode(w, privateKeyBlock); err != nil {
		return fmt.Errorf("error when encode private pem: %s", err)
	}
	return nil
}

// ReadPrivateKey takes in the private key bytes and decodes and sets the private and public key
func (r *Rsa) ReadPrivateKey(reader io.Reader) error {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		return err
	}
	block, _ := pem.Decode(buf.Bytes())
	if block == nil {
		return errors.New("failed to parse PEM block containing the public key")
	}

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return errors.New("failed to parse encoded private key: " + err.Error())
	}
	r.Private = private
	r.Public = &r.Private.PublicKey

	return nil
}

// WritePublicKey encodes the rsa.Private key to the io.Writer
func (r Rsa) WritePublicKey(w io.Writer) error {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(r.Public)
	if err != nil {
		return fmt.Errorf("error when dumping publickey: %s", err)
	}

	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	err = pem.Encode(w, &publicKeyBlock)
	if err != nil {
		return fmt.Errorf("error when encode public pem: %s", err)
	}

	return nil
}

// ReadPrivateKey takes in the public key bytes and decodes and sets the public key only
func (r *Rsa) ReadPublicKey(reader io.Reader) error {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		return err
	}

	block, _ := pem.Decode(buf.Bytes())
	if block == nil {
		return errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return errors.New("failed to parse encoded public key: " + err.Error())
	}

	r.Public = pub.(*rsa.PublicKey)

	return nil
}

// PublicKeyFromBytes takes in the public key bytes and decodes the public key only
func PublicKeyFromBytes(b []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse encoded public key: " + err.Error())
	}

	return pub.(*rsa.PublicKey), nil
}
