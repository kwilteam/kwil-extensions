package extension

import (
	"fmt"
	"math/big"
	"testing"
)

func Test_round(t *testing.T) {
	num1 := big.NewFloat(1.3)
	down := roundDown(num1)

	downNum := down.Int64()
	fmt.Println(downNum)

	num := big.NewFloat(1.7)
	up := roundUp(num)

	if down.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("roundDown(%v) = %v, want %v", num1, down, 1)
	}

	if up.Cmp(big.NewInt(2)) != 0 {
		t.Errorf("roundUp(%v) = %v, want %v", num, up, 2)
	}
}

func Test_Down(t *testing.T) {
	num := newBigFloat(1.2)
	down := roundDown(num)

	if down.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("roundDown(%v) = %v, want %v", num, down, 1)
	}
}

func Test_Up(t *testing.T) {
	num := big.NewFloat(1.9)
	up := roundUp(num)

	if up.Cmp(big.NewInt(2)) != 0 {
		t.Errorf("roundUp(%v) = %v, want %v", num, up, 2)
	}
}
