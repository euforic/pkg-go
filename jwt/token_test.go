package jwt

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestNewToken(t *testing.T) {
	type args struct {
		t *jwt.Token
	}
	tests := []struct {
		name string
		args args
		want *Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewToken(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewToken(%v) = %v, want %v", tt.args.t, got, tt.want)
			}
		})
	}
}

func TestToken_ParseError(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if err := tr.ParseError(); (err != nil) != tt.wantErr {
				t.Errorf("Token.ParseError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToken_Claims(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	tests := []struct {
		name       string
		fields     fields
		wantValue  jwt.MapClaims
		wantExists bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			gotValue, gotExists := tr.Claims()
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Token.Claims() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotExists != tt.wantExists {
				t.Errorf("Token.Claims() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestToken_Get(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantValue  interface{}
		wantExists bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			gotValue, gotExists := tr.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Token.Get(%v) gotValue = %v, want %v", tt.args.key, gotValue, tt.wantValue)
			}
			if gotExists != tt.wantExists {
				t.Errorf("Token.Get(%v) gotExists = %v, want %v", tt.args.key, gotExists, tt.wantExists)
			}
		})
	}
}

func TestToken_GetString(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantS  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotS := tr.GetString(tt.args.key); gotS != tt.wantS {
				t.Errorf("Token.GetString(%v) = %v, want %v", tt.args.key, gotS, tt.wantS)
			}
		})
	}
}

func TestToken_GetBool(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantB  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotB := tr.GetBool(tt.args.key); gotB != tt.wantB {
				t.Errorf("Token.GetBool(%v) = %v, want %v", tt.args.key, gotB, tt.wantB)
			}
		})
	}
}

func TestToken_GetInt(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantI  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotI := tr.GetInt(tt.args.key); gotI != tt.wantI {
				t.Errorf("Token.GetInt(%v) = %v, want %v", tt.args.key, gotI, tt.wantI)
			}
		})
	}
}

func TestToken_GetInt64(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantI64 int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotI64 := tr.GetInt64(tt.args.key); gotI64 != tt.wantI64 {
				t.Errorf("Token.GetInt64(%v) = %v, want %v", tt.args.key, gotI64, tt.wantI64)
			}
		})
	}
}

func TestToken_GetUint(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantUi uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotUi := tr.GetUint(tt.args.key); gotUi != tt.wantUi {
				t.Errorf("Token.GetUint(%v) = %v, want %v", tt.args.key, gotUi, tt.wantUi)
			}
		})
	}
}

func TestToken_GetUint64(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUi64 uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotUi64 := tr.GetUint64(tt.args.key); gotUi64 != tt.wantUi64 {
				t.Errorf("Token.GetUint64(%v) = %v, want %v", tt.args.key, gotUi64, tt.wantUi64)
			}
		})
	}
}

func TestToken_GetFloat64(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantF64 float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotF64 := tr.GetFloat64(tt.args.key); gotF64 != tt.wantF64 {
				t.Errorf("Token.GetFloat64(%v) = %v, want %v", tt.args.key, gotF64, tt.wantF64)
			}
		})
	}
}

func TestToken_GetTime(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantTm time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotTm := tr.GetTime(tt.args.key); !reflect.DeepEqual(gotTm, tt.wantTm) {
				t.Errorf("Token.GetTime(%v) = %v, want %v", tt.args.key, gotTm, tt.wantTm)
			}
		})
	}
}

func TestToken_GetDuration(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantD  time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotD := tr.GetDuration(tt.args.key); gotD != tt.wantD {
				t.Errorf("Token.GetDuration(%v) = %v, want %v", tt.args.key, gotD, tt.wantD)
			}
		})
	}
}

func TestToken_GetSlice(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if got := tr.GetSlice(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Token.GetSlice(%v) = %v, want %v", tt.args.key, got, tt.want)
			}
		})
	}
}

func TestToken_GetStringSlice(t *testing.T) {
	type fields struct {
		Token      *jwt.Token
		parseError error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantSs []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Token{
				Token:      tt.fields.Token,
				parseError: tt.fields.parseError,
			}
			if gotSs := tr.GetStringSlice(tt.args.key); !reflect.DeepEqual(gotSs, tt.wantSs) {
				t.Errorf("Token.GetStringSlice(%v) = %v, want %v", tt.args.key, gotSs, tt.wantSs)
			}
		})
	}
}
