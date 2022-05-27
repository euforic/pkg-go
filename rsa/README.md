# rsa

```go
import "github.com/euforic/pkg-go/rsa"
```

## Index

- [func PublicKeyFromBytes(b []byte) (*rsa.PublicKey, error)](<#func-publickeyfrombytes>)
- [type Rsa](<#type-rsa>)
  - [func New() *Rsa](<#func-new>)
  - [func (r *Rsa) Generate() error](<#func-rsa-generate>)
  - [func (r *Rsa) ReadPrivateKey(reader io.Reader) error](<#func-rsa-readprivatekey>)
  - [func (r *Rsa) ReadPublicKey(reader io.Reader) error](<#func-rsa-readpublickey>)
  - [func (r Rsa) WritePrivateKey(w io.Writer) error](<#func-rsa-writeprivatekey>)
  - [func (r Rsa) WritePublicKey(w io.Writer) error](<#func-rsa-writepublickey>)


## func PublicKeyFromBytes

```go
func PublicKeyFromBytes(b []byte) (*rsa.PublicKey, error)
```

PublicKeyFromBytes takes in the public key bytes and decodes the public key only

## type Rsa

Rsa\.\.\.

```go
type Rsa struct {
    Private *rsa.PrivateKey
    Public  *rsa.PublicKey
}
```

### func New

```go
func New() *Rsa
```

New creates a new instance of \*Rsa

### func \(\*Rsa\) Generate

```go
func (r *Rsa) Generate() error
```

Generate gerates the public and private keys and sets them

### func \(\*Rsa\) ReadPrivateKey

```go
func (r *Rsa) ReadPrivateKey(reader io.Reader) error
```

ReadPrivateKey takes in the private key bytes and decodes and sets the private and public key

### func \(\*Rsa\) ReadPublicKey

```go
func (r *Rsa) ReadPublicKey(reader io.Reader) error
```

ReadPrivateKey takes in the public key bytes and decodes and sets the public key only

### func \(Rsa\) WritePrivateKey

```go
func (r Rsa) WritePrivateKey(w io.Writer) error
```

WritePrivateKey encodes the rsa\.Private key to the io\.Writer

### func \(Rsa\) WritePublicKey

```go
func (r Rsa) WritePublicKey(w io.Writer) error
```

WritePublicKey encodes the rsa\.Private key to the io\.Writer
