package main

import (
	"testing"
)


// Тест віднімання
func TestSubtractBigNumbers(t *testing.T) {
	a := []uint32{1, 0, 0, 0}
	b := []uint32{1, 0, 0, 0}
	expected := []uint32{0, 0, 0, 0}
	result := subtractBigNumbers(a, b)
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Очікувано %08X, отримано %08X", expected[i], result[i])
		}
	}
}

// Тест множення
func TestMultiplyBigNumbers(t *testing.T) {
	a := []uint32{0xFFFFFFFF, 0, 0, 0}
	b := []uint32{2, 0, 0, 0}
	expected := []uint32{0xFFFFFFFE, 1, 0, 0}
	result := multiplyBigNumbers(a, b)
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Очікувано %08X, отримано %08X", expected[i], result[i])
		}
	}
}

// Тест функції ділення великих чисел
func TestDivideBigNumbers(t *testing.T) {
	a := []uint32{0x2468ACF0, 0x468ACF12, 0x68ACF120, 0x8ACF1202}
	b := []uint32{0x12345678}
	expectedQuotient := []uint32{0x2}
	expectedRemainder := []uint32{0}
	quotient, remainder := divideBigNumbers(a, b)

	if !reflect.DeepEqual(quotient[:maxWords], expectedQuotient) {
		t.Errorf("Expected quotient %v, but got %v", expectedQuotient, quotient)
	}
	if !reflect.DeepEqual(remainder[:maxWords], expectedRemainder) {
		t.Errorf("Expected remainder %v, but got %v", expectedRemainder, remainder)
	}
}

// Тест піднесення до квадрату
func TestSquareBigNumber(t *testing.T) {
	a := []uint32{0x2}
	expected := []uint32{0x4}
	result := squareBigNumber(a)[:maxWords]

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Тест піднесення до степеня
func TestPowerBigNumber(t *testing.T) {
	base := []uint32{0x2}
	exponent := uint32(3)
	expected := []uint32{0x8}
	result := powerBigNumber(base, exponent)[:maxWords]

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Тест функції зсуву вліво
func TestShiftLeft(t *testing.T) {
	a := []uint32{0x1}
	shift := uint(1)
	expected := []uint32{0x2}
	result := shiftLeft(a, shift)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Тест функції зсуву вправо
func TestShiftRight(t *testing.T) {
	a := []uint32{0x2}
	shift := uint(1)
	expected := []uint32{0x1}
	result := shiftRight(a, shift)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Тест знаходження номера старшого ненульового біта
func TestHighestSetBit(t *testing.T) {
	a := []uint32{0x0, 0x0, 0x0, 0x80000000}
	expected := 127
	result := highestSetBit(a)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Тест конвертації в десяткову систему
func TestToDecimal(t *testing.T) {
	a := []uint32{0x12345678, 0x23456789}
	expected := "97732674312987896"
	result := toDecimal(a)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Тест конвертації в двійкову систему
func TestToBinary(t *testing.T) {
	a := []uint32{0x12345678}
	expected := "00010010001101000101011001111000"
	result := toBinary(a)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Тест конвертації в шістнадцяткову систему
func TestToHex(t *testing.T) {
	a := []uint32{0x12345678}
	expected := "12345678"
	result := toHex(a)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
