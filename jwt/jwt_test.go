package jwt

import (
	"crypto/rsa"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	type args struct {
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
	}
	tests := []struct {
		name string
		args args
		want *JWT
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.privateKey, tt.args.publicKey); !cmp.Equal(got, tt.want) {
				t.Errorf("New(%v, %v) = %v, want %v", tt.args.privateKey, tt.args.publicKey, got, tt.want)
			}
		})
	}
}

func TestJWT_CreatAndSign(t *testing.T) {
	type fields struct {
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
	}
	type args struct {
		ttl    time.Duration
		claims jwt.MapClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JWT{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			got, err := j.CreatAndSign(tt.args.ttl, tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWT.CreatAndSign(%v, %v) error = %v, wantErr %v", tt.args.ttl, tt.args.claims, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JWT.CreatAndSign(%v, %v) = %v, want %v", tt.args.ttl, tt.args.claims, got, tt.want)
			}
		})
	}
}

func TestJWT_Create(t *testing.T) {
	type fields struct {
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
	}
	type args struct {
		ttl    time.Duration
		claims jwt.MapClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JWT{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			got, err := j.Create(tt.args.ttl, tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWT.Create(%v, %v) error = %v, wantErr %v", tt.args.ttl, tt.args.claims, err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("JWT.Create(%v, %v) = %v, want %v", tt.args.ttl, tt.args.claims, got, tt.want)
			}
		})
	}
}

func TestJWT_Sign(t *testing.T) {
	type fields struct {
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
	}
	type args struct {
		token *Token
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JWT{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			got, err := j.Sign(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWT.Sign(%v) error = %v, wantErr %v", tt.args.token, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JWT.Sign(%v) = %v, want %v", tt.args.token, got, tt.want)
			}
		})
	}
}

func TestJWT_Parse(t *testing.T) {
	type fields struct {
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
	}
	type args struct {
		token    string
		validate bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JWT{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			got, err := j.Parse(tt.args.token, tt.args.validate)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWT.Parse(%v, %v) error = %v, wantErr %v", tt.args.token, tt.args.validate, err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("JWT.Parse(%v, %v) = %v, want %v", tt.args.token, tt.args.validate, got, tt.want)
			}
		})
	}
}
