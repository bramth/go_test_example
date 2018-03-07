package main

import (
	"testing"
)

//BIKIN FUNCTION TEST
func TestSum(t *testing.T) {

	//testtable
	tt := []struct {
		Name   string
		Int1   int
		Int2   int
		Result int
	}{
		{
			Name:   "Test 2+3",
			Int1:   2,
			Int2:   3,
			Result: 5,
		}, {
			Name:   "Test -200 + 100",
			Int1:   -200,
			Int2:   100,
			Result: -100,
		},
	}

	for _, tc := range tt {
		result := Sum(tc.Int1, tc.Int2)
		if result != tc.Result {
			t.Fatalf("Error di tc '%s' mintanya %d tapi dapetnya %d",
				tc.Name,
				tc.Result,
				result)

		}
	}
	t.Log("Sudah selesai jalanin TestSum")
}

func TestSubstract(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 5-3",
			args: args{
				a: 5,
				b: 3,
			},
			want: 2,
		},
		{
			name: "test 10-5",
			args: args{
				a: 10,
				b: 5,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Substract() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Log("Sudah selesai jalanin TestSubstract")
}
