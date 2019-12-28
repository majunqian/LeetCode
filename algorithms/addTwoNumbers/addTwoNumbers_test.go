package addTwoNumbers

import (
	"math/big"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	test := []struct {
		a, b, c string
	}{
		// normal case
		{"111", "999", "1110"},
		// big int
		{"10000000000000000000000000000001", "111", "10000000000000000000000000000112"},
	}
	for _, tt := range test {
		var (
			bigA big.Int
			bigB big.Int
		)
		bigA.SetString(tt.a, 10)
		bigB.SetString(tt.b, 10)
		lists := AddTwoNumbers(numToList(&bigA), numToList(&bigB))
		result := listToNum(lists).String()
		if tt.c != listToNum(lists).String() {
			t.Errorf("%s+%s\n expect:[%s] get[%s]\n", tt.a, tt.b, tt.c, result)
		}
	}
}

func BenchmarkAddTwoNumbers(ben *testing.B) {
	a := "10000000000000000000000000000001"
	b := "111"
	c := "10000000000000000000000000000112"

	var (
		bigA big.Int
		bigB big.Int
	)
	bigA.SetString(a, 10)
	bigB.SetString(b, 10)
	lista := numToList(&bigA)
	listb := numToList(&bigB)
	ben.ResetTimer()
	for i := 0; i < ben.N; i++ {
		lists := AddTwoNumbers(lista, listb)
		result := listToNum(lists).String()
		if c != listToNum(lists).String() {
			ben.Errorf("%s+%s\n expect:[%s] get[%s]\n", a, b, c, result)
		}
	}
}
