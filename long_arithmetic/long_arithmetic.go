package main

import (
	"fmt"
	"strconv"
	"time"
)

// Максимальна кількість 32-бітних слів для числа
var maxWords = 4

// Додавання двох беззнакових великих чисел
func addBigNumbers(a, b []uint32) []uint32 {
    result := make([]uint32, maxWords)
    carry := uint32(0)
    for i := 0; i < maxWords; i++ {
        sum := a[i] + b[i] + carry
        result[i] = sum
        if sum < a[i] || (carry == 1 && sum == a[i]) {
            carry = 1
        } else {
            carry = 0
        }
    }
    return result
}

// Віднімання двох беззнакових великих чисел
func subtractBigNumbers(a, b []uint32) []uint32 {
	result := make([]uint32, maxWords)
	borrow := uint32(0)
	for i := 0; i < maxWords; i++ {
		if a[i] < b[i]+borrow {
			result[i] = a[i] + 0xFFFFFFFF - b[i] - borrow 
			borrow = 1
		} else {
			result[i] = a[i] - b[i] - borrow
			borrow = 0
		}
	}
	return result
}



func multiplyBigNumbers(a, b []uint32) []uint32 {
    result := make([]uint32, maxWords*2)
    for i := 0; i < maxWords; i++ {
        carry := uint64(0)
        for j := 0; j < maxWords; j++ {
            if i+j < len(result) {
                product := uint64(a[i])*uint64(b[j]) + uint64(result[i+j]) + carry
                result[i+j] = uint32(product & 0xFFFFFFFF)
                carry = product >> 32
            }
        }
        if i+maxWords < len(result) {
            result[i+maxWords] += uint32(carry)
        }
    }
    return result[:maxWords]
}






// Піднесення до квадрату
func squareBigNumber(a []uint32) []uint32 {
	return multiplyBigNumbers(a, a)
}

func powerBigNumber(base []uint32, exponent uint32) []uint32 {
    result := make([]uint32, maxWords)
    result[0] = 1
    tempBase := make([]uint32, maxWords)
    copy(tempBase, base)

    for exponent > 0 {
        if exponent&1 == 1 {
            result = multiplyBigNumbers(result, tempBase)[:maxWords]
        }
        tempBase = multiplyBigNumbers(tempBase, tempBase)[:maxWords]
        exponent >>= 1
    }
    return result
}


// Функція ділення з залишком
func divideBigNumbers(a, b []uint32) ([]uint32, []uint32) {
	quotient := make([]uint32, maxWords)
	remainder := make([]uint32, maxWords)
	copy(remainder, a)

	for i := maxWords*32 - 1; i >= 0; i-- {
		remainder = shiftLeft(remainder, 1)
		if compareBigNumbers(remainder, b) >= 0 {
			remainder = subtractBigNumbers(remainder, b)
			setBit(quotient, i)
		}
	}
	return quotient, remainder
}

// Функція для встановлення біта на позиції `bitPosition` у числі `number`
func setBit(number []uint32, bitPosition int) {
	wordIndex := bitPosition / 32
	bitIndex := bitPosition % 32
	number[wordIndex] |= (1 << bitIndex)
}

func shiftLeft(a []uint32, shift uint) []uint32 {
	result := make([]uint32, maxWords)
	for i := maxWords - 1; i >= 0; i-- {
		if i-int(shift)/32 >= 0 {
			result[i] = a[i-int(shift)/32] << (shift % 32)
		}
		if i-1-int(shift)/32 >= 0 && shift%32 > 0 {
			result[i] |= a[i-1-int(shift)/32] >> (32 - shift%32)
		}
	}
	return result
}

func shiftRight(a []uint32, shift uint) []uint32 {
	result := make([]uint32, maxWords)
	for i := 0; i < maxWords; i++ {
		if i+int(shift)/32 < maxWords {
			result[i] = a[i+int(shift)/32] >> (shift % 32)
		}
		if i+1+int(shift)/32 < maxWords && shift%32 > 0 {
			result[i] |= a[i+1+int(shift)/32] << (32 - shift%32)
		}
	}
	return result
}

// Знаходження номера старшого ненульового біта
func highestSetBit(a []uint32) int {
	for i := maxWords - 1; i >= 0; i-- {
		for j := 31; j >= 0; j-- {
			if a[i]&(1<<j) != 0 {
				return i*32 + j
			}
		}
	}
	return -1 // Повертає -1, якщо число дорівнює нулю
}

// Перетворення числа у десяткову систему
func toDecimal(number []uint32) string {
	decimal := uint64(0)
	for i := len(number) - 1; i >= 0; i-- {
		decimal = (decimal << 32) + uint64(number[i])
	}
	return strconv.FormatUint(decimal, 10)
}

// Перетворення числа у двійкову систему
func toBinary(number []uint32) string {
	binary := ""
	for i := len(number) - 1; i >= 0; i-- {
		binary += fmt.Sprintf("%032b", number[i])
	}
	return binary
}

// Перетворення числа у шістнадцяткову систему
func toHex(number []uint32) string {
	hex := ""
	for i := len(number) - 1; i >= 0; i-- {
		hex += fmt.Sprintf("%08X", number[i])
	}
	return hex
}

// Перевірка тотожності (a + b) * c = c * (a + b) = a * c + b * c
func testIdentity1(a, b, c []uint32) bool {
	sum := addBigNumbers(a, b)
	left := multiplyBigNumbers(sum, c)

	ac := multiplyBigNumbers(a, c)
	bc := multiplyBigNumbers(b, c)
	right := addBigNumbers(ac, bc)

	return compareBigNumbers(left, right) == 0
}



// Функція для порівняння двох великих чисел
func compareBigNumbers(a, b []uint32) int {
	for i := maxWords - 1; i >= 0; i-- {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	return 0
}

// Функція для вимірювання часу
func measureRepeatedTime(repeats int, f func()) time.Duration {
	start := time.Now()
	for i := 0; i < repeats; i++ {
		f()
	}
	return time.Since(start)
}


// Основна функція
func main() {
	// Вхідні числа
	a := []uint32{0x12345678, 0x23456789, 0x34567890, 0x45678901}
	b := []uint32{0x98765432, 0x87654321, 0x76543210, 0x6543210F}
	c := []uint32{0x11111111, 0x22222222, 0x33333333, 0x44444444}


	// Тести на тотожність
	fmt.Println("Перевірка тотожностей:")
	fmt.Printf("(a + b) * c = a * c + b * c : %v\n", testIdentity1(a, b, c))
	

	// Замір часу для кожної операції
	repeats := 100000

	fmt.Println("\nЧас виконання операцій:")
	addTime := measureRepeatedTime(repeats, func() { addBigNumbers(a, b) })
	fmt.Printf("Додавання: %v\n", addTime)

	subtractTime := measureRepeatedTime(repeats, func() { subtractBigNumbers(a, b) })
	fmt.Printf("Віднімання: %v\n", subtractTime)

	multiplyTime := measureRepeatedTime(repeats, func() { multiplyBigNumbers(a, b) })
	fmt.Printf("Множення: %v\n", multiplyTime)

	// Ділення
	divideTime := measureRepeatedTime(repeats, func() {
		quotient, remainder := divideBigNumbers(a, b)
		_ = quotient
		_ = remainder
	})
	fmt.Printf("Ділення: %v\n", divideTime)

	// Піднесення до квадрату
	squareTime := measureRepeatedTime(repeats, func() { squareBigNumber(a) })
	fmt.Printf("Піднесення до квадрату: %v\n", squareTime)

	// Піднесення до степеня
	powerTime := measureRepeatedTime(repeats, func() { powerBigNumber(a, 10) })
	fmt.Printf("Піднесення до степеня: %v\n", powerTime)

	// Зсув вліво та вправо
	shiftLeftTime := measureRepeatedTime(repeats, func() { shiftLeft(a, 1) })
	fmt.Printf("Зсув вліво: %v\n", shiftLeftTime)

	shiftRightTime := measureRepeatedTime(repeats, func() { shiftRight(a, 1) })
	fmt.Printf("Зсув вправо: %v\n", shiftRightTime)

	// Вивід номера старшого ненульового біта
	fmt.Printf("Старший ненульовий біт в a: %d\n", highestSetBit(a))
	fmt.Printf("Старший ненульовий біт в b: %d\n", highestSetBit(b))
	fmt.Printf("Старший ненульовий біт в c: %d\n", highestSetBit(c))

	// Вивід конвертацій
	fmt.Println("\nКонвертації чисел:")
	fmt.Printf("Число a у двійковій системі: %s\n", toBinary(a))
	fmt.Printf("Число a у десятковій системі: %s\n", toDecimal(a))
	fmt.Printf("Число a у шістнадцятковій системі: %s\n", toHex(a))

	fmt.Printf("Число b у двійковій системі: %s\n", toBinary(b))
	fmt.Printf("Число b у десятковій системі: %s\n", toDecimal(b))
	fmt.Printf("Число b у шістнадцятковій системі: %s\n", toHex(b))
}
