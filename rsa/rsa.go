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

var (
	// ErrNoPrivateKey is returned when no private key is set
	ErrNoPrivateKey = errors.New("no private key set")
	// ErrNoPublicKey is returned when no public key is set
	ErrNoPublicKey = errors.New("no public key set")
	// ErrFailedToParse is returned when the key cannot be parsed
	ErrFailedToParse = errors.New("failed to parse key")
)

// Rsa struct to hold the public and private keys
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
	recommendedKeySize := 2048
	privateKey, err := rsa.GenerateKey(rand.Reader, recommendedKeySize)
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}
	r.Private = privateKey
	r.Public = &privateKey.PublicKey

	return nil
}

// WritePrivateKey encodes the rsa.Private key to the io.Writer
func (r Rsa) WritePrivateKey(w io.Writer) error {
	if r.Private == nil {
		return ErrNoPrivateKey
	}
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(r.Private),
	}

	if err := pem.Encode(w, privateKeyBlock); err != nil {
		return fmt.Errorf("error when encode private pem: %w", err)
	}

	return nil
}

// ReadPrivateKey takes in the private key bytes and decodes and sets the private and public key
func (r *Rsa) ReadPrivateKey(reader io.Reader) error {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		return fmt.Errorf("error when reading private key: %w", err)
	}
	block, _ := pem.Decode(buf.Bytes())
	if block == nil {
		return fmt.Errorf("failed to parse PEM block containing the private key: %w", ErrFailedToParse)
	}

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse encoded private key: %w", err)
	}
	r.Private = private
	r.Public = &r.Private.PublicKey

	return nil
}

// WritePublicKey encodes the rsa.Private key to the io.Writer
func (r Rsa) WritePublicKey(w io.Writer) error {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(r.Public)
	if err != nil {
		return fmt.Errorf("error when marshal public key: %w", err)
	}

	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	err = pem.Encode(w, &publicKeyBlock)
	if err != nil {
		return fmt.Errorf("error when encode public pem: %w", err)
	}

	return nil
}

// ReadPrivateKey takes in the public key bytes and decodes and sets the public key only
func (r *Rsa) ReadPublicKey(reader io.Reader) error {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		return fmt.Errorf("error when reading public key: %w", err)
	}

	block, _ := pem.Decode(buf.Bytes())
	if block == nil {
		return fmt.Errorf("failed to parse PEM block containing the public key: %w", ErrFailedToParse)
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse encoded public key: %w", err)
	}

	var ok bool
	r.Public, ok = pub.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("failed to parse encoded public key: %w", ErrFailedToParse)
	}

	return nil
}

// PublicKeyFromBytes takes in the public key bytes and decodes the public key only
func PublicKeyFromBytes(b []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key: %w", ErrFailedToParse)
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse encoded public key: %w", err)
	}

	var ok bool
	key, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to parse encoded public key: %w", ErrFailedToParse)
	}

	return key, nil
}
