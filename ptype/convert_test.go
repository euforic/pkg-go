package ptype

import (
	"testing"
	"time"
)

var (
	simpleFloat  = 1.1
	simpleInt    = 1
	simpleBool   = true
	simpleString = "simple"
	simpleTime   = time.Time{}
)

func TestPFloat64(t *testing.T) {
	type args[T any] struct {
		name string
		v    *T
		want T
	}

	tests := []args[float64]{
		{
			name: "simple",
			v:    &simpleFloat,
			want: simpleFloat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Deref(tt.v); got != tt.want {
				t.Errorf("Deref(float64) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntP(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "basic",
			args: args{
				i: simpleInt,
			},
			want: &simpleInt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ptr(tt.args.i); *got != *tt.want {
				t.Errorf("Ptr(*int) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64P(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want *float64
	}{
		{
			name: "simple",
			args: args{
				f: simpleFloat,
			},
			want: &simpleFloat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ptr(tt.args.f); *got != *tt.want {
				t.Errorf("Ptr(*float64) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringP(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "simple",
			args: args{
				str: simpleString,
			},
			want: &simpleString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ptr(tt.args.str); *got != *tt.want {
				t.Errorf("Ptr(*string) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPString(t *testing.T) {
	type args struct {
		str *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				str: &simpleString,
			},
			want: simpleString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Deref(tt.args.str); got != tt.want {
				t.Errorf("Deref(string) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPInt(t *testing.T) {
	type args struct {
		i *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple",
			args: args{
				i: &simpleInt,
			},
			want: simpleInt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Deref(tt.args.i); got != tt.want {
				t.Errorf("PInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolP(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		{
			name: "simple",
			args: args{
				b: simpleBool,
			},
			want: &simpleBool,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ptr(tt.args.b); *got != *tt.want {
				t.Errorf("Ptr(*boolean) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeP(t *testing.T) {
	type args struct {
		v time.Time
	}
	tests := []struct {
		name string
		args args
		want *time.Time
	}{
		{
			name: "simple",
			args: args{
				v: simpleTime,
			},
			want: &simpleTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ptr(tt.args.v); *got != *tt.want {
				t.Errorf("Ptr(*time.Time) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPTime(t *testing.T) {
	type args struct {
		v *time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "simple",
			args: args{
				v: &simpleTime,
			},
			want: simpleTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Deref(tt.args.v); got != tt.want {
				t.Errorf("Ptr(time.Time) = %v, want %v", got, tt.want)
			}
		})
	}
}
