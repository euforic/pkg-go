package eth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

var (
	// ErrInvalidMnemonic is returned when the mnemonic is invalid.
	ErrInvalidMnemonic = errors.New("invalid mnemonic")
	// ErrInvalidPrivateKey is returned when the private key is invalid.
	ErrInvalidPrivateKey = errors.New("invalid private key")
	// ErrInvalidSignature is returned when the signature is invalid.
	ErrInvalidSignature = errors.New("invalid signature")
	// ErrInvalidAddress is returned when the address is invalid.
	ErrInvalidAddress = errors.New("invalid address")
)

// Wallet represents a cryptocurrency wallet
type Wallet struct {
	privateKey *ecdsa.PrivateKey
}

// WalletOpt is a functional option for creating a wallet.
type WalletOpt func(*Wallet) error

// WithPrivateKey creates a wallet from a private key.
func WithPrivateKey(privateKey string) WalletOpt {
	return func(w *Wallet) error {
		pk, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return fmt.Errorf("failed to convert to ECDSA %w: %w", err, ErrInvalidPrivateKey)
		}

		w.privateKey = pk

		return nil
	}
}

// WithMnemonic creates a wallet from a mnemonic.
func WithMnemonic(mnemonic string) WalletOpt {
	return func(w *Wallet) error {
		if mnemonic == "" {
			return fmt.Errorf("mnemonic is required: %w", ErrInvalidMnemonic)
		}

		// Use the provided mnemonic
		if !bip39.IsMnemonicValid(mnemonic) {
			return fmt.Errorf("invalid mnemonic: %w", ErrInvalidMnemonic)
		}

		seed := bip39.NewSeed(mnemonic, "")

		// Generate a master key from the seed
		masterKey, err := bip32.NewMasterKey(seed)
		if err != nil {
			return fmt.Errorf("failed to generate master key %w: %w", err, ErrInvalidMnemonic)
		}

		// Derive the Ethereum key (BIP-44: m/44'/60'/0'/0/0)
		purpose, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44) //nolint:mnd
		if err != nil {
			return fmt.Errorf("failed to derive purpose key %w: %w", err, ErrInvalidMnemonic)
		}

		// Ethereum uses coin type 60
		coinType, err := purpose.NewChildKey(bip32.FirstHardenedChild + 60) //nolint:mnd
		if err != nil {
			return fmt.Errorf("failed to derive coin type key %w: %w", err, ErrInvalidMnemonic)
		}

		account, err := coinType.NewChildKey(bip32.FirstHardenedChild)
		if err != nil {
			return fmt.Errorf("failed to derive account key %w: %w", err, ErrInvalidMnemonic)
		}

		change, err := account.NewChildKey(0)
		if err != nil {
			return fmt.Errorf("failed to derive change key %w: %w", err, ErrInvalidMnemonic)
		}

		addressIndex, err := change.NewChildKey(0)
		if err != nil {
			return fmt.Errorf("failed to derive address index key %w: %w", err, ErrInvalidMnemonic)
		}

		privateKey, err := crypto.ToECDSA(addressIndex.Key)
		if err != nil {
			return fmt.Errorf("failed to convert to ECDSA %w: %w", err, ErrInvalidMnemonic)
		}

		w.privateKey = privateKey

		return nil
	}
}

// New creates a new wallet from a private key.
func New(opt WalletOpt) (*Wallet, error) {
	var w Wallet

	if err := opt(&w); err != nil {
		return nil, err
	}

	return &w, nil
}

// PublicKey returns the public key of the wallet.
func (w Wallet) PublicKeyStr() string {
	return hexutil.Encode(crypto.FromECDSAPub(&w.privateKey.PublicKey))
}

// PrivateKey returns the private key of the wallet.
func (w Wallet) PrivateKeyStr() string {
	return hexutil.Encode(crypto.FromECDSA(w.privateKey))
}

// Address returns the address of the wallet.
func (w Wallet) Address() string {
	return crypto.PubkeyToAddress(w.privateKey.PublicKey).Hex()
}

// SignMessage signs a message using the provided private key.
func (w Wallet) SignMessage(message string) (string, error) {
	messageHash := accounts.TextHash([]byte(message))

	signature, err := crypto.Sign(messageHash, w.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %w", err)
	}

	signature[crypto.RecoveryIDOffset] += 27

	return hexutil.Encode(signature), nil
}

// SignTransaction signs a transaction with the provided private key and chain ID.
func (w Wallet) SignTransaction(chainID *big.Int, tx *types.Transaction) (*types.Transaction, error) {
	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, w.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	return signedTx, nil
}

// VerifySignature verifies the signature of a message.
func (w Wallet) VerifySignature(signatureHex, message string) error {
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return fmt.Errorf("failed to decode signature: %w", err)
	}

	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	messageHash := accounts.TextHash([]byte(message))

	pubKey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		return fmt.Errorf("failed to recover public key: %w", err)
	}

	fromAddress := crypto.PubkeyToAddress(*pubKey).Hex()

	if common.HexToAddress(fromAddress) != crypto.PubkeyToAddress(*pubKey) {
		return ErrInvalidSignature
	}

	return nil
}

// NewMnemonic creates a new mnemonic.
func NewMnemonic() (string, error) {
	// 256 bits of entropy is recommended for a 24-word mnemonic
	bitsOfEntropy := 256
	// Generate a new mnemonic
	entropy, err := bip39.NewEntropy(bitsOfEntropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate entropy: %w", err)
	}

	newMnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate mnemonic: %w", err)
	}

	return newMnemonic, nil
}

// SignatureAddress returns the address of the signer of a message.
func SignatureAddress(signatureHex, message string) (string, error) {
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode signature: %w", err)
	}

	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	messageHash := accounts.TextHash([]byte(message))

	pubKey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		return "", fmt.Errorf("failed to recover public key: %w", err)
	}

	return crypto.PubkeyToAddress(*pubKey).Hex(), nil
}
