package ptype

import (
	"testing"
)

var (
	simpleFloat  = 1.1
	simpleInt    = 1
	simpleBool   = true
	simpleString = "simple"
)

func TestPFloat64(t *testing.T) {
	type args struct {
		num *float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "simple",
			args: args{
				num: &simpleFloat,
			},
			want: simpleFloat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PFloat64(tt.args.num); got != tt.want {
				t.Errorf("PFloat64() = %v, want %v", got, tt.want)
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
			if got := IntP(tt.args.i); *got != *tt.want {
				t.Errorf("IntP() = %v, want %v", got, tt.want)
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
			if got := Float64P(tt.args.f); *got != *tt.want {
				t.Errorf("Float64P() = %v, want %v", got, tt.want)
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
			if got := StringP(tt.args.str); *got != *tt.want {
				t.Errorf("StringP() = %v, want %v", got, tt.want)
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
			if got := PString(tt.args.str); got != tt.want {
				t.Errorf("PString() = %v, want %v", got, tt.want)
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
			if got := PInt(tt.args.i); got != tt.want {
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
			if got := BoolP(tt.args.b); *got != *tt.want {
				t.Errorf("BoolP() = %v, want %v", got, tt.want)
			}
		})
	}
}
