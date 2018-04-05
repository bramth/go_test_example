package main

import "testing"

func TestSum(t *testing.T) {
	tc := []struct {
		name   string
		input1 int
		input2 int
		expRes int
	}{
		{
			name:   "tc1",
			input1: 5,
			input2: 3,
			expRes: 8,
		},
		{
			name:   "tc2",
			input1: 15,
			input2: 3,
			expRes: 100,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			if Sum(tt.input1, tt.input2) != tt.expRes {
				t.Errorf("hasil penjumlahan tidak sesuai")
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				a: 5,
				b: 3,
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				a: 7,
				b: 4,
			},
			want: 3,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
