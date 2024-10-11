package long_arithmetic

import (
	"math/big"
	"testing"
)

func TestAddBigInts(t *testing.T) {
	a := big.NewInt(123456789)
	b := big.NewInt(987654321)
	result, _ := AddBigInts(a, b)
	expected := big.NewInt(1111111110)
	if result.Cmp(expected) != 0 {
		t.Errorf("Додавання не пройшло. Очікувано %v, отримано %v", expected, result)
	}
}

func TestSubtractBigInts(t *testing.T) {
	a := big.NewInt(987654321)
	b := big.NewInt(123456789)
	result, _ := SubtractBigInts(a, b)
	expected := big.NewInt(864197532)
	if result.Cmp(expected) != 0 {
		t.Errorf("Віднімання не пройшло. Очікувано %v, отримано %v", expected, result)
	}
}

func TestMultiplyBigInts(t *testing.T) {
	a := big.NewInt(12345)
	b := big.NewInt(6789)
	result, _ := MultiplyBigInts(a, b)
	expected := big.NewInt(83810205)
	if result.Cmp(expected) != 0 {
		t.Errorf("Множення не пройшло. Очікувано %v, отримано %v", expected, result)
	}
}

func TestDivideBigInts(t *testing.T) {
	a := big.NewInt(987654321)
	b := big.NewInt(123456789)
	result, _ := DivideBigInts(a, b)
	expected := big.NewInt(8)
	if result.Cmp(expected) != 0 {
		t.Errorf("Ділення не пройшло. Очікувано %v, отримано %v", expected, result)
	}
}

func TestShiftLeft(t *testing.T) {
	a := big.NewInt(1024)
	result, _ := ShiftLeft(a, 2) // зсув на 2 біти (множення на 4)
	expected := big.NewInt(4096)
	if result.Cmp(expected) != 0 {
		t.Errorf("Зсув вліво не пройшов. Очікувано %v, отримано %v", expected, result)
	}
}

func TestShiftRight(t *testing.T) {
	a := big.NewInt(1024)
	result, _ := ShiftRight(a, 2) // зсув на 2 біти (ділення на 4)
	expected := big.NewInt(256)
	if result.Cmp(expected) != 0 {
		t.Errorf("Зсув вправо не пройшов. Очікувано %v, отримано %v", expected, result)
	}
}

func TestHighestNonZeroBit(t *testing.T) {
	a := big.NewInt(1024)
	result, _ := HighestNonZeroBit(a)
	expected := 10 // оскільки 1024 = 2^10
	if result != expected {
		t.Errorf("Номер старшого ненульового біта неправильний. Очікувано %v, отримано %v", expected, result)
	}
}

func TestHexConversion(t *testing.T) {
	a := big.NewInt(255)
	hexStr := ToHexString(a)
	expectedHex := "ff"
	if hexStr != expectedHex {
		t.Errorf("Конвертація в шістнадцяткову систему неправильна. Очікувано %v, отримано %v", expectedHex, hexStr)
	}

	parsedNum, _ := FromHexString("ff")
	if parsedNum.Cmp(a) != 0 {
		t.Errorf("Конвертація з шістнадцяткового рядка неправильна. Очікувано %v, отримано %v", a, parsedNum)
	}
}
