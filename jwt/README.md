# jwt

```go
import "github.com/euforic/pkg-go/jwt"
```

## Index

- [type JWT](<#type-jwt>)
  - [func New(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT](<#func-new>)
  - [func (j JWT) CreatAndSign(ttl time.Duration, claims jwt.MapClaims) (string, error)](<#func-jwt-creatandsign>)
  - [func (j JWT) Create(ttl time.Duration, claims jwt.MapClaims) (*Token, error)](<#func-jwt-create>)
  - [func (j JWT) Parse(token string, validate bool) (*Token, error)](<#func-jwt-parse>)
  - [func (j JWT) Sign(token *Token) (string, error)](<#func-jwt-sign>)
- [type Token](<#type-token>)
  - [func NewToken(t *jwt.Token) *Token](<#func-newtoken>)
  - [func (t Token) Claims() (value jwt.MapClaims, exists bool)](<#func-token-claims>)
  - [func (t Token) Get(key string) (value interface{}, exists bool)](<#func-token-get>)
  - [func (t Token) GetBool(key string) (b bool)](<#func-token-getbool>)
  - [func (t Token) GetDuration(key string) (d time.Duration)](<#func-token-getduration>)
  - [func (t Token) GetFloat64(key string) (f64 float64)](<#func-token-getfloat64>)
  - [func (t Token) GetInt(key string) (i int)](<#func-token-getint>)
  - [func (t Token) GetInt64(key string) (i64 int64)](<#func-token-getint64>)
  - [func (t Token) GetSlice(key string) []interface{}](<#func-token-getslice>)
  - [func (t Token) GetString(key string) (s string)](<#func-token-getstring>)
  - [func (t Token) GetStringSlice(key string) (ss []string)](<#func-token-getstringslice>)
  - [func (t Token) GetTime(key string) (tm time.Time)](<#func-token-gettime>)
  - [func (t Token) GetUint(key string) (ui uint)](<#func-token-getuint>)
  - [func (t Token) GetUint64(key string) (ui64 uint64)](<#func-token-getuint64>)
  - [func (t Token) ParseError() error](<#func-token-parseerror>)


## type JWT

JWT

```go
type JWT struct {
    // contains filtered or unexported fields
}
```

### func New

```go
func New(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT
```

New creates a new instance of JWT util

### func \(JWT\) CreatAndSign

```go
func (j JWT) CreatAndSign(ttl time.Duration, claims jwt.MapClaims) (string, error)
```

CreateAndSign a new jwt token

### func \(JWT\) Create

```go
func (j JWT) Create(ttl time.Duration, claims jwt.MapClaims) (*Token, error)
```

Create generates a new jwt token string

### func \(JWT\) Parse

```go
func (j JWT) Parse(token string, validate bool) (*Token, error)
```

Parse takes in a jwt token string parses it\, validates it and return a \*Token

### func \(JWT\) Sign

```go
func (j JWT) Sign(token *Token) (string, error)
```

Create generates a new jwt token string

## type Token

```go
type Token struct {
    *jwt.Token
    // contains filtered or unexported fields
}
```

### func NewToken

```go
func NewToken(t *jwt.Token) *Token
```

NewToken creates a new instance of Token from a \*jwt\.Token

### func \(Token\) Claims

```go
func (t Token) Claims() (value jwt.MapClaims, exists bool)
```

### func \(Token\) Get

```go
func (t Token) Get(key string) (value interface{}, exists bool)
```

### func \(Token\) GetBool

```go
func (t Token) GetBool(key string) (b bool)
```

GetBool returns the value associated with the key as a boolean\.

### func \(Token\) GetDuration

```go
func (t Token) GetDuration(key string) (d time.Duration)
```

GetDuration returns the value associated with the key as a duration\.

### func \(Token\) GetFloat64

```go
func (t Token) GetFloat64(key string) (f64 float64)
```

GetFloat64 returns the value associated with the key as a float64\.

### func \(Token\) GetInt

```go
func (t Token) GetInt(key string) (i int)
```

GetInt returns the value associated with the key as an integer\.

### func \(Token\) GetInt64

```go
func (t Token) GetInt64(key string) (i64 int64)
```

GetInt64 returns the value associated with the key as an integer\.

### func \(Token\) GetSlice

```go
func (t Token) GetSlice(key string) []interface{}
```

GetSlice returns the value associated with the key as a slice of interface\{\}\.

### func \(Token\) GetString

```go
func (t Token) GetString(key string) (s string)
```

GetString returns the value associated with the key as a string\.

### func \(Token\) GetStringSlice

```go
func (t Token) GetStringSlice(key string) (ss []string)
```

GetStringSlice returns the value associated with the key as a slice of strings\.

### func \(Token\) GetTime

```go
func (t Token) GetTime(key string) (tm time.Time)
```

GetTime returns the value associated with the key as time\.

### func \(Token\) GetUint

```go
func (t Token) GetUint(key string) (ui uint)
```

GetUint returns the value associated with the key as an unsigned integer\.

### func \(Token\) GetUint64

```go
func (t Token) GetUint64(key string) (ui64 uint64)
```

GetUint64 returns the value associated with the key as an unsigned integer\.

### func \(Token\) ParseError

```go
func (t Token) ParseError() error
```
