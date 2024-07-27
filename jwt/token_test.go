package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/go-cmp/cmp"
)

func TestTokenClaims(t *testing.T) {
	tests := []struct {
		name      string
		token     *jwt.Token
		want      jwt.MapClaims
		wantExist bool
	}{
		{
			name: "Valid claims",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub":   "1234567890",
				"name":  "John Doe",
				"admin": true,
			}),
			want: jwt.MapClaims{
				"sub":   "1234567890",
				"name":  "John Doe",
				"admin": true,
			},
			wantExist: true,
		},
		{
			name:      "No claims",
			token:     jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{}),
			want:      jwt.MapClaims{},
			wantExist: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got, gotExist := token.Claims()
			if !cmp.Equal(got, tt.want) || gotExist != tt.wantExist {
				t.Errorf("Claims() got = %v, want = %v, gotExist = %v, wantExist = %v, diff: %v", got, tt.want, gotExist, tt.wantExist, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGet(t *testing.T) {
	tests := []struct {
		name      string
		token     *jwt.Token
		key       string
		want      interface{}
		wantExist bool
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": "1234567890",
			}),
			key:       "sub",
			want:      "1234567890",
			wantExist: true,
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": "1234567890",
			}),
			key:       "name",
			want:      nil,
			wantExist: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got, gotExist := token.Get(tt.key)
			if !cmp.Equal(got, tt.want) || gotExist != tt.wantExist {
				t.Errorf("Get() got = %v, want = %v, gotExist = %v, wantExist = %v, diff: %v", got, tt.want, gotExist, tt.wantExist, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetString(t *testing.T) {
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  string
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": "1234567890",
			}),
			key:  "sub",
			want: "1234567890",
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": "1234567890",
			}),
			key:  "name",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetString(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetString() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetBool(t *testing.T) {
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  bool
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"admin": true,
			}),
			key:  "admin",
			want: true,
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"admin": true,
			}),
			key:  "user",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetBool(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetBool() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetInt(t *testing.T) {
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  int
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"age": 30,
			}),
			key:  "age",
			want: 30,
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"age": 30,
			}),
			key:  "height",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetInt(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetInt() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetUint(t *testing.T) {
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  uint
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"points": uint(100),
			}),
			key:  "points",
			want: uint(100),
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"points": uint(100),
			}),
			key:  "level",
			want: uint(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetUint(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetUint() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetFloat32(t *testing.T) {
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  float32
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"rating": float32(4.5),
			}),
			key:  "rating",
			want: float32(4.5),
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"rating": float32(4.5),
			}),
			key:  "score",
			want: float32(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetFloat32(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetFloat32() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetFloat64(t *testing.T) {
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  float64
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"rating": float64(4.5),
			}),
			key:  "rating",
			want: float64(4.5),
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"rating": float64(4.5),
			}),
			key:  "score",
			want: float64(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetFloat64(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetFloat64() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetTime(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  time.Time
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": now,
			}),
			key:  "exp",
			want: now,
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": now,
			}),
			key:  "iat",
			want: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetTime(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetTime() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetDuration(t *testing.T) {
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  time.Duration
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"timeout": time.Minute,
			}),
			key:  "timeout",
			want: time.Minute,
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"timeout": time.Minute,
			}),
			key:  "delay",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetDuration(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetDuration() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetSlice(t *testing.T) {
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  []interface{}
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"roles": []interface{}{"admin", "user"},
			}),
			key:  "roles",
			want: []interface{}{"admin", "user"},
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"roles": []interface{}{"admin", "user"},
			}),
			key:  "permissions",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetSlice(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetSlice() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestTokenGetStringSlice(t *testing.T) {
	tests := []struct {
		name  string
		token *jwt.Token
		key   string
		want  []string
	}{
		{
			name: "Existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"roles": []string{"admin", "user"},
			}),
			key:  "roles",
			want: []string{"admin", "user"},
		},
		{
			name: "Non-existing key",
			token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"roles": []interface{}{"admin", "user"},
			}),
			key:  "permissions",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewToken(tt.token)
			got := token.GetStringSlice(tt.key)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GetStringSlice() got = %v, want = %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
