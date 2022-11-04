package sliceutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCompare(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestCompareTrue1", args: args{[]string{"one", "sa12d"}, []string{"sa12d", "one"}}, want: true},
		{name: "TestCompareTrue2", args: args{[]string{"one", "sa12d"}, []string{"one", "sa12d"}}, want: true},
		{name: "TestCompareFalse", args: args{[]string{"one", "sa12d"}, []string{"1", "213"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedCompare(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestOrderedCompareFalse1", args: args{[]string{"one", "sa12d"}, []string{"sa12d", "one"}}, want: false},
		{name: "TestOrderedCompareTrue", args: args{[]string{"one", "sa12d"}, []string{"one", "sa12d"}}, want: true},
		{name: "TestOrderedCompareFalse2", args: args{[]string{"one", "sa12d"}, []string{"1", "213"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderedCompare(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("OrderedCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestContainsTrue1", args: args{[]string{"one", "sa12d"}, "one"}, want: true},
		{name: "TestContainsFalse", args: args{[]string{"one", "sa12d"}, "1"}, want: false},
		{name: "TestContainsTrue2", args: args{[]string{"one", "sa12d"}, "sa12d"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	in1 := make([]string, 0)
	in1 = append(in1, "2")
	in1 = append(in1, "4.2")
	in1 = append(in1, "1")
	out1 := make([]string, 0)
	out1 = append(out1, "1")
	out1 = append(out1, "4.2")
	out1 = append(out1, "2")

	in2 := make([]string, 0)
	in2 = append(in2, "1")
	in2 = append(in2, "2")
	in2 = append(in2, "3")
	out2 := make([]string, 0)
	out2 = append(out2, "3")
	out2 = append(out2, "2")
	out2 = append(out2, "1")

	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Reverse1", args: args{s: in1}, want: out1},
		{name: "Reverse2", args: args{s: in2}, want: out2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); !cmp.Equal(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFastContains(t *testing.T) {
	type args struct {
		s          []string
		searchterm string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FastContains(tt.args.s, tt.args.searchterm); got != tt.want {
				t.Errorf("FastContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkContains10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Contains([]string{"one", "ten"}, "ten")
	}
}

func BenchmarkFastContains10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FastContains([]string{"one", "ten"}, "ten")
	}
}

func GroupContains() bool {
	tt := map[string]int{}
	for _, v := range []string{"one", "10", "3"} {
		tt[v] = 0
	}

	for _, v := range []string{"10", "3"} {
		if _, ok := tt[v]; ok != true {
			return true
		}
	}
	return false
}

func BenchmarkGroupContains10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GroupContains()
	}
}
func GroupFastContains() bool {
	search := []string{"1", "2", "3", "4"}
	terms := []string{"1", "2", "3"}
	i := len(terms) - 1
	for i >= 0 {
		r := FastContains(search, terms[i])
		if r == false {
			return false
		}
		i--
	}
	return true
}

func BenchmarkGroupFastContains10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GroupFastContains()
	}
}
