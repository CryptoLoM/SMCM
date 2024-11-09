package main

import (
	"testing"
)


func TestAddBigNumbers(t *testing.T) {
    a := parseHex("9")
    b := parseHex("1")
    expected := parseHex("A")
    result := addBigNumbers(a, b)

    if !bigNumbersEqual(result, expected) {
        t.Errorf("Expected %s, got %s", toHex(expected), toHex(result))
    }

	a = parseHex("7")
    b = parseHex("10")
    expected = parseHex("17")
    result = addBigNumbers(a, b)

    if !bigNumbersEqual(result, expected) {
        t.Errorf("Expected %s, got %s", toHex(expected), toHex(result))
    }

	a = parseHex("10")
    b = parseHex("7")
    expected = parseHex("17")
    result = addBigNumbers(a, b)

    if !bigNumbersEqual(result, expected) {
        t.Errorf("Expected %s, got %s", toHex(expected), toHex(result))
    }

	a = parseHex("ffff")
    b = parseHex("1")
    expected = parseHex("10000")
    result = addBigNumbers(a, b)

    if !bigNumbersEqual(result, expected) {
        t.Errorf("Expected %s, got %s", toHex(expected), toHex(result))
    }
}


func TestSubtractBigNumbers(t *testing.T) {
	
	a := parseHex("10") 
	b := parseHex("5")  
	expected := parseHex("b") 
	result := subtractBigNumbers(a, b)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("1")
	b = parseHex("2")
	expected = parseHex("1")
	result = subtractBigNumbers(a, b)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("1000")
	b = parseHex("1000")
	expected = parseHex("0")
	result = subtractBigNumbers(a, b)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("2")
	b = parseHex("10") 
	expected = parseHex("e") 
	result = subtractBigNumbers(a, b)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("12345678")
	b = parseHex("0")
	expected = parseHex("12345678")
	result = subtractBigNumbers(a, b)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("0")
	b = parseHex("12345678")
	expected = parseHex("12345678")
	result = subtractBigNumbers(a, b)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("ffffffffffffffff")
	b = parseHex("1")
	expected = parseHex("fffffffffffffffe")
	result = subtractBigNumbers(a, b)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("1000")
	b = parseHex("1")
	expected = parseHex("fff")
	result = subtractBigNumbers(a, b)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}
}


func TestMultiplyBigNumbers(t *testing.T) {
	a := parseHex("2")
	b := parseHex("3")
	expected := parseHex("6")
	result := multiplyBigNumbers(a, b)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("0")
	b = parseHex("3fadca1f3fc31412424237653938458736453207")
	expected = parseHex("0")
	result = multiplyBigNumbers(a, b)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}
}

func TestSquareBigNumber(t *testing.T) {
	a := parseHex("3")
	expected := parseHex("9")
	result := squareBigNumber(a)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("0")
	expected = parseHex("0")
	result = squareBigNumber(a)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("1")
	expected = parseHex("1")
	result = squareBigNumber(a)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("a")
	expected = parseHex("64")
	result = squareBigNumber(a)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("f")
	expected = parseHex("")
	result = squareBigNumber(a)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}
}

func TestPowerBigNumber(t *testing.T) {
	a := parseHex("7")
	exponent := parseHex("1")
	expected := parseHex("7")
	result := powerBigNumber(a, exponent)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("a")
	exponent = parseHex("3")
	expected = parseHex("3e8")
	result = powerBigNumber(a, exponent)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("1")
	exponent = parseHex("3abfc114h1ab4cd8102412fad")
	expected = parseHex("1")
	result = powerBigNumber(a, exponent)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	a = parseHex("af12124188d7baf712c48")
	exponent = parseHex("0")
	expected = parseHex("1")
	result = powerBigNumber(a, exponent)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}

	
}

func TestDivideBigNumbers(t *testing.T) {
	a := parseHex("a7f3b")
	b := parseHex("e69")
	expectedQuotient := parseHex("ba")
	quotient := divideBigNumbers(a, b)

	if !bigNumbersEqual(quotient, expectedQuotient) {
		t.Errorf("Expected quotient %v, got %v", toHex(expectedQuotient), toHex(quotient))
	}

	a = parseHex("7")
	b = parseHex("7")
	expectedQuotient = parseHex("1")
	quotient = divideBigNumbers(a, b)

	if !bigNumbersEqual(quotient, expectedQuotient) {
		t.Errorf("Expected quotient %v, got %v", toHex(expectedQuotient), toHex(quotient))
	}

	a = parseHex("a")
	b = parseHex("a")
	expectedQuotient = parseHex("1")
	quotient = divideBigNumbers(a, b)

	if !bigNumbersEqual(quotient, expectedQuotient) {
		t.Errorf("Expected quotient %v, got %v", toHex(expectedQuotient), toHex(quotient))
	}

	a = parseHex("7")
	b = parseHex("a")
	expectedQuotient = parseHex("0")
	quotient = divideBigNumbers(a, b)

	if !bigNumbersEqual(quotient, expectedQuotient) {
		t.Errorf("Expected quotient %v, got %v", toHex(expectedQuotient), toHex(quotient))
	}

}

func TestCompareBigNumbers(t *testing.T) {
	a := parseHex("a7f3b8c4d9e213f7b6a8c9d4e7f9b5c2a9d6e8b4f7a9d4c8e2b5f7c8a6d9b3e4f1c7a9e3d8b6c7f2a8b4d6c9e7f3b1d8a6c5e2f7b9d4e8c6a2f3b7c1d8e5f6a9c7b4d1f2c8e3b5a7d9b6c1f4e3d7a9c8b2d6f5b7a4c9e2f1d8c6b5a7e4f3d2c1a6b9e8f5d7c4e2b3f1a9d7b6c8e5f3a4c2b8d1")
	b := parseHex("e6f2c4a8d3b7c9e1f5b4a2d6e9c8f3a7b5d1c8e2f4a9d6b7c3e1f5b8a6d4c9e7b3f2a5c1d9b6f4e8a7c5b2d1f3e9b4a6c7d2f1b5e3c8a9d7f2b4e6c1a3d8b9f5c7e2a4b3f1d6c8e5b2a7f9b3c6e1d4f8a5c9e2b7d1f3a6")

	if compareBigNumbers(b, a) != -1 {
		t.Errorf("Expected e6f2c4a8d3b7c9e1f5b4a2d6e9c8f3a7b5d1c8e2f4a9d6b7c3e1f5b8a6d4c9e7b3f2a5c1d9b6f4e8a7c5b2d1f3e9b4a6c7d2f1b5e3c8a9d7f2b4e6c1a3d8b9f5c7e2a4b3f1d6c8e5b2a7f9b3c6e1d4f8a5c9e2b7d1f3a6 < a7f3b8c4d9e213f7b6a8c9d4e7f9b5c2a9d6e8b4f7a9d4c8e2b5f7c8a6d9b3e4f1c7a9e3d8b6c7f2a8b4d6c9e7f3b1d8a6c5e2f7b9d4e8c6a2f3b7c1d8e5f6a9c7b4d1f2c8e3b5a7d9b6c1f4e3d7a9c8b2d6f5b7a4c9e2f1d8c6b5a7e4f3d2c1a6b9e8f5d7c4e2b3f1a9d7b6c8e5f3a4c2b8d1")
	}
	if compareBigNumbers(a, b) != 1 {
		t.Errorf("Expected a7f3b8c4d9e213f7b6a8c9d4e7f9b5c2a9d6e8b4f7a9d4c8e2b5f7c8a6d9b3e4f1c7a9e3d8b6c7f2a8b4d6c9e7f3b1d8a6c5e2f7b9d4e8c6a2f3b7c1d8e5f6a9c7b4d1f2c8e3b5a7d9b6c1f4e3d7a9c8b2d6f5b7a4c9e2f1d8c6b5a7e4f3d2c1a6b9e8f5d7c4e2b3f1a9d7b6c8e5f3a4c2b8d1 > e6f2c4a8d3b7c9e1f5b4a2d6e9c8f3a7b5d1c8e2f4a9d6b7c3e1f5b8a6d4c9e7b3f2a5c1d9b6f4e8a7c5b2d1f3e9b4a6c7d2f1b5e3c8a9d7f2b4e6c1a3d8b9f5c7e2a4b3f1d6c8e5b2a7f9b3c6e1d4f8a5c9e2b7d1f3a6")
	}
	if compareBigNumbers(b, b) != 0 {
		t.Errorf("e6f2c4a8d3b7c9e1f5b4a2d6e9c8f3a7b5d1c8e2f4a9d6b7c3e1f5b8a6d4c9e7b3f2a5c1d9b6f4e8a7c5b2d1f3e9b4a6c7d2f1b5e3c8a9d7f2b4e6c1a3d8b9f5c7e2a4b3f1d6c8e5b2a7f9b3c6e1d4f8a5c9e2b7d1f3a6 = e6f2c4a8d3b7c9e1f5b4a2d6e9c8f3a7b5d1c8e2f4a9d6b7c3e1f5b8a6d4c9e7b3f2a5c1d9b6f4e8a7c5b2d1f3e9b4a6c7d2f1b5e3c8a9d7f2b4e6c1a3d8b9f5c7e2a4b3f1d6c8e5b2a7f9b3c6e1d4f8a5c9e2b7d1f3a6")
	}
}


func TestTrimLeadingZeros(t *testing.T) {
	input := parseHex("00000000FFFFFFFF")
	expected := parseHex("FFFFFFFF")
	result := trimLeadingZeros(input)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", toHex(expected), toHex(result))
	}
}



func TestBigNumbersEqual(t *testing.T) {
	tests := []struct {
		a, b     []uint32
		expected bool
	}{
		{[]uint32{1, 2, 3}, []uint32{1, 2, 3}, true},
		{[]uint32{1, 2, 3}, []uint32{1, 2, 4}, false},
		{[]uint32{1, 2, 3}, []uint32{1, 2}, false},
		{[]uint32{}, []uint32{}, true},
		{[]uint32{0}, []uint32{0}, true},
	}

	for _, tt := range tests {
		result := bigNumbersEqual(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("bigNumbersEqual(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
		}else{
			t.Logf("PASS: Input: %v, Result: %v", tt.a, result)
		}
	}
}





func TestToHex(t *testing.T) {
	// Define test cases with parsed numbers (as slices of uint32) and their expected hexadecimal string results.
	tests := []struct {
		input    []uint32 // Parsed input as a slice of uint32 numbers
		expected string   // Expected result as a hexadecimal string
	}{
		{parseHex("9ABCDEF012345678"), "9ABCDEF012345678"},
		{parseHex("12345678"), "12345678"},
		{parseHex("FFFFFFFF"), "FFFFFFFF"},
		{parseHex("0"), "0"},                    // Test with zero
		{parseHex("0000000000000000"), "0"},     // Test with leading zeros
		{parseHex("100000000"), "100000000"},    // Test with a large power of two
	}

	for _, test := range tests {
		result := toHex(test.input)
		if result != test.expected {
			t.Errorf("toHex(%v) = %s; want %s", test.input, result, test.expected)
		}else{
			t.Logf("PASS: Input: %v, Result: %v", test.input, result)
		}
	}
}



func TestParseHex(t *testing.T) {
	tests := []struct {
		hexInput string   // Input hexadecimal string
		expected []uint32 // Expected result as a slice of uint32 values
	}{
		{"9ABCDEF012345678", []uint32{0x12345678, 0x9ABCDEF0}},    // Standard case with two uint32s
		{"12345678", []uint32{0x12345678}},                        // Single uint32 value
		{"FFFFFFFF", []uint32{0xFFFFFFFF}},                        // Maximum 32-bit unsigned integer
		{"0", []uint32{0}},                                        // Zero case
		{"0000000000000000", []uint32{0}},                         // Leading zeros
		{"100000000", []uint32{0x0, 0x1}},                         // Large power of two spanning two uint32s
	}

	for _, test := range tests {
		result := parseHex(test.hexInput)
		if !bigNumbersEqual(result, test.expected) {
			t.Errorf("parseHex(%s) = %v; want %v", test.hexInput, result, test.expected)
		}else{
			t.Logf("PASS: hexInput: %v, Result: %v", test.hexInput, result)
	}
}
}


func TestShiftLeft(t *testing.T) {
	a := []uint32{1}
	expected := []uint32{0, 1}
	result := shiftLeft(a, 32)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	a = []uint32{1, 0}
	expected = []uint32{0, 1, 0} 
	result = shiftLeft(a, 32)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	a = []uint32{1, 0, 0}
	expected = []uint32{0, 1, 0, 0}
	result = shiftLeft(a, 32)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	a = []uint32{} 
	expected = []uint32{}
	result = shiftLeft(a, 1)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	a = []uint32{0} 
	expected = []uint32{0,0}
	result = shiftLeft(a, 5)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	a = []uint32{uint32('A')} 
	expected = []uint32{130,0}   
	result = shiftLeft(a, 1)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}


func TestShiftRight(t *testing.T) {
    a := []uint32{1}
	expected := []uint32{0, 1}// Перше число стане 0, друге - 1 після зсуву
	result := shiftLeft(a, 32)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}


	a = []uint32{0, 0, 1}
	expected = []uint32{0, 0, 0, 1}
	result = shiftLeft(a, 32)

	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	a = []uint32{} 
	expected = []uint32{}
	result = shiftRight(a, 1)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	a = []uint32{0} 
	expected = []uint32{0, 0}
	result = shiftRight(a, 5)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	a = []uint32{uint32('A')} 
	expected = []uint32{32, 0}   
	result = shiftRight(a, 1)
	if !bigNumbersEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}


func TestSetBit(t *testing.T) {
	tests := []struct {
		number    []uint32
		bitPos    int
		expected  []uint32
	}{
		{[]uint32{0x0}, 0, []uint32{0x1}},
		{[]uint32{0x0}, 31, []uint32{0x80000000}},
		{[]uint32{0x0, 0x0}, 32, []uint32{0x0, 0x1}},
		{[]uint32{0x1, 0x0}, 32, []uint32{0x1, 0x1}},
	}

	for _, tt := range tests {
		number := make([]uint32, len(tt.number))
		copy(number, tt.number)
		setBit(number, tt.bitPos)
		if !bigNumbersEqual(number, tt.expected) {
			t.Errorf("setBit(%v, %d) = %v; want %v", tt.number, tt.bitPos, number, tt.expected)
		}else{
			t.Logf("PASS: Input: %v, Result: %v", tt.number, number)
	}
}
}


func TestHighestSetBit(t *testing.T) {
	tests := []struct {
		input    []uint32
		expected int
	}{
		{[]uint32{}, -1},                         // Порожній вхід
		{[]uint32{0, 0, 0}, -1},                  // Усі нулі
		{[]uint32{1}, 0},                         // Найменший ненульовий біт
		{[]uint32{1 << 31}, 31},                  // Найстарший біт в одному слові
		{[]uint32{0, 0, 1 << 15}, 79},            // Ненульовий біт у третьому слові
		{[]uint32{1 << 7, 0, 0}, 7},              // Ненульовий біт у першому слові
		{[]uint32{0, 0, 1 << 31}, 95},            // Найстарший біт у третьому слові
		{[]uint32{0xFFFFFFFF, 0xFFFFFFFF}, 63},   // Усі біти встановлені у двох словах
		{[]uint32{1, 0, 1 << 31}, 95},            // Змішаний випадок
		{[]uint32{0, 1}, 32},                     // Ненульовий біт на початку другого слова
		{[]uint32{0, 0, 0, 1 << 3}, 99},          // Ненульовий біт у четвертому слові
	}

	for _, test := range tests {
		result := highestSetBit(test.input)
		if result != test.expected {
			t.Errorf("FAIL: Input: %v, Expected: %d, Got: %d", test.input, test.expected, result)
		} else {
			t.Logf("PASS: Input: %v, Result: %d", test.input, result)
		}
	}
}