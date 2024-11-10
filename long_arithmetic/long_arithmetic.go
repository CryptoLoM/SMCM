package main
import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"errors"
)

func zero() []uint32 {
	return []uint32{0}
}

func one() []uint32 {
	return []uint32{1}
}

// Додавання двох беззнакових великих чисел
func addBigNumbers(num1, num2 []uint32) []uint32 {
	maxWords := max(len(num1), len(num2))
	result := make([]uint32, maxWords)
	carry := uint32(0)
	for i := 0; i < maxWords; i++ {
		aVal := getWord(num1, i)
		bVal := getWord(num2, i)
		sum := aVal + bVal + carry
		result[i] = sum
		if sum < aVal || (carry == 1 && sum == aVal) {
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry > 0 {
		result = append(result, carry)
	}
	return result
}


func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}


func subtractBigNumbers(num1, num2 []uint32) []uint32 {
	maxWords := max(len(num1), len(num2))
	result := make([]uint32, maxWords)
	borrow := uint32(0)
	
	// Якщо перше число менше другого (від'ємник більше зменшуваного)
	if compareBigNumbers(num1, num2) < 0 {
		// Виконуємо віднімання в зворотньому порядку і додаємо 0xFFFFFFFF для від'ємності
		result = subtractBigNumbers(num2, num1)
		return append([]uint32{}, result...)
	}

	// Основний цикл віднімання
	for i := 0; i < maxWords; i++ {
		aVal := getWord(num1, i)
		bVal := getWord(num2, i)
		if aVal < bVal+borrow {
			result[i] = aVal  - bVal - borrow
			borrow = 1
		} else {
			result[i] = aVal - bVal - borrow
			borrow = 0
		}
	}
	return trimLeadingZeros(result)
}



func compareBigNumbers(number1, number2 []uint32) int {
	// Визначаємо мінімальну довжину масивів
	n := min(len(number1), len(number2))
	
	// Починаємо з кінця масивів
	i := n - 1
	for i >= 0 && number1[i] == number2[i] {
		i--
	}

	// Якщо дійшли до -1, то числа однакові
	if i == -1 {
		return 0
	} else {
		// Порівнюємо невідповідні цифри
		if number1[i] > number2[i] {
			return 1
		} else {
			return -1
		}
	}
}


func getWord(number []uint32, index int) uint32 {
	if index < len(number) {
		return number[index]
	}
	return 0
}

// Множення великих чисел довільної довжини
func multiplyBigNumbers(num1, num2 []uint32) []uint32 {
	result := make([]uint32, len(num1)+len(num2))
	for i := 0; i < len(num1); i++ {
		carry := uint64(0)
		for j := 0; j < len(num2); j++ {
			if i+j < len(result) {
				product := uint64(num1[i])*uint64(num2[j]) + uint64(result[i+j]) + carry
				result[i+j] = uint32(product & 0xFFFFFFFF)
				carry = product >> 32
			}
		}
		if i+len(num2) < len(result) {
			result[i+len(num2)] += uint32(carry)
		}
	}
	return trimLeadingZeros(result)
}


// Ділення з використанням блочного порозрядного підходу
func divideBigNumbers(dividend, divisor []uint32) []uint32 {
	if compareBigNumbers(dividend, divisor) < 0 {
		// Якщо ділене менше за дільник, частка дорівнює 0
		return []uint32{0}
	}
	if compareBigNumbers(divisor, []uint32{0}) == 0 {
		// Обробка випадку ділення на нуль
		panic("Division by zero")
	}

	// Ініціалізація частки та залишку
	quotient := make([]uint32, len(dividend))
	remainder := make([]uint32, len(dividend))
	copy(remainder, dividend)

	// Отримуємо кількість бітів у дільнику
	k := highestSetBit(divisor)

	// Цикл для ділення
	for compareBigNumbers(remainder, divisor) >= 0 {
		// Отримуємо кількість бітів у залишку
		t := highestSetBit(remainder)

		// Оптимізоване обчислення зсуву
		shiftAmount := t - k
		result := shiftLeft(divisor, uint(shiftAmount))

		//бінарний пошук, щоб знайти максимальний зсув, який дозволяє віднімання
		for compareBigNumbers(remainder, result) < 0 && shiftAmount > 0 {
			shiftAmount--
			result = shiftLeft(divisor, uint(shiftAmount))
		}

		// Віднімемо з залишку
		remainder = subtractBigNumbers(remainder, result)

		// Встановимо біт у частці
		setBit(quotient, int(shiftAmount))
	}

	return trimLeadingZeros(quotient) 
}

func highestSetBit(h1 []uint32) int {
	for i := len(h1) - 1; i >= 0; i-- {
		for j := 31; j >= 0; j-- {
			if h1[i]&(1<<j) != 0 {
				return i*32 + j
			}
		}
	}
	return -1
}


func shiftLeft(num []uint32, shift uint) []uint32 {
	// Якщо масив порожній, повертаємо порожній масив
	if len(num) == 0 {
		return num
	}
	
	// Розраховуємо зсув на рівні слів і бітів
	wordShift := shift / 32
	bitShift := shift % 32
	
	// Визначаємо кількість елементів для результату
	// Це потрібно для того, щоб не створювати зайвих елементів
	resultSize := len(num) + int(shift/32)
	if bitShift > 0 {
		resultSize++
	}
	result := make([]uint32, resultSize)
	
	// Виконуємо зсув
	for i := 0; i < len(num); i++ {
		if bitShift == 0 {
			result[i+int(wordShift)] = num[i]
		} else {
			result[i+int(wordShift)] |= num[i] << bitShift
			if i+int(wordShift)+1 < len(result) {
				result[i+int(wordShift)+1] = num[i] >> (32 - bitShift)
			}
		}
	}
	
	return result
}

func shiftRight(num []uint32, shift uint) []uint32 {
	// Якщо масив порожній, повертаємо порожній масив
	if len(num) == 0 {
		return num
	}
	
	// Розраховуємо зсув на рівні слів і бітів
	wordShift := shift / 32
	bitShift := shift % 32
	
	// Визначаємо кількість елементів для результату
	// Для вправо ми просто робимо віднімання на кількість слів, тому розмір результату коригується
	resultSize := len(num) - int(shift/32)
	if bitShift > 0 {
		resultSize++
	}
	if resultSize < 0 {
		resultSize = 0
	}
	result := make([]uint32, resultSize)
	
	// Виконуємо зсув
	for i := len(num) - 1; i >= int(wordShift); i-- {
		if bitShift == 0 {
			result[i-int(wordShift)] = num[i]
		} else {
			result[i-int(wordShift)] |= num[i] >> bitShift
			if i-int(wordShift)-1 >= 0 {
				result[i-int(wordShift)-1] = num[i] << (32 - bitShift)
			}
		}
	}
	
	return result
}



func setBit(number []uint32, bitPosition int) {
	wordIndex := bitPosition / 32
	bitIndex := bitPosition % 32
	number[wordIndex] |= (1 << bitIndex)
}


func trimLeadingZeros(number []uint32) []uint32 {
    for len(number) > 1 && number[len(number)-1] == 0 {
        number = number[:len(number)-1]
    }
    if len(number) == 0 {
        return []uint32{0} 
    }
    return number
}


func squareBigNumber(num []uint32) []uint32 {
	return multiplyBigNumbers(num, num)
}



func powerBigNumber(base, exponent []uint32) []uint32 {
	if compareBigNumbers(exponent, []uint32{0}) == 0 {
		return []uint32{1} // Base case: anything to the power of 0 is 1
	}

	result := []uint32{1}
	tempBase := base

	// Використання двійкового подання для прискорення
	for compareBigNumbers(exponent, []uint32{0}) > 0 {
		if exponent[0]&1 == 1 { // Перевірка на непарність
			result = multiplyBigNumbers(result, tempBase)
		}
		tempBase = multiplyBigNumbers(tempBase, tempBase)
		exponent = shiftRight(exponent, 1) // Зсув на 1 біт вправо
	}

	return result
}


func toHex(number []uint32) string {
	if len(number) == 1 && number[0] == 0 {
        return "00000000"
    }

    // Якщо число є 1, повертаємо "1"
    if len(number) == 1 && number[0] == 1 {
        return "00000001"
    }
    hex := ""
    for i := len(number) - 1; i >= 0; i-- {
        hex += fmt.Sprintf("%08X", number[i])
    }
    hex = strings.TrimLeft(hex, "0")
    if hex == "" {
        return "0"
    }
    return hex
}


func parseHex(hexStr string) []uint32 {
	hexStr = strings.TrimLeft(hexStr, "0x")
	hexStr = strings.TrimLeft(hexStr, "0")
	
	if hexStr == "" {
		return []uint32{0} 
	}
	padding := 8 - len(hexStr)%8
	if padding < 8 {
		hexStr = strings.Repeat("0", padding) + hexStr
	}
	var result []uint32
	for i := 0; i < len(hexStr); i += 8 {
		word, _ := strconv.ParseUint(hexStr[i:i+8], 16, 32)
		result = append([]uint32{uint32(word)}, result...)
	}
	return result
}


func measureRepeatedTime(repeats int, f func()) time.Duration {
	start := time.Now()
	for i := 0; i < repeats; i++ {
		f()
	}
	return time.Since(start)
}


func multiplyByAddition(n int, a int) (int, error) {
	// Перевіряємо обмеження на n
	if n < 100 {
		return 0, errors.New("помилка: n повинно бути не меншим за 100")
	}

	result := 0
	for i := 0; i < n; i++ {
		result += a
	}
	return result, nil
}

// Перевірка тотожності (a + b) ⋅ c = c ⋅ (a + b) = a ⋅ c + b ⋅ c 
func testIdentity1(a, b, c []uint32) bool {
	sum := addBigNumbers(a, b)
	left := multiplyBigNumbers(sum, c)

	ac := multiplyBigNumbers(a, c)
	bc := multiplyBigNumbers(b, c)
	right := addBigNumbers(ac, bc)

	return compareBigNumbers(left, right) == 0
}

func bigNumbersEqual(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {

	zeroNumber := zero()
	oneNumber := one()

	fmt.Printf("конвертування 0: %s\n", toHex(zeroNumber))
	fmt.Printf("конвертування 1: %s\n", toHex(oneNumber))

    hexInput1 := "2"	
    hexInput2 := "e6f2c4a8d3b7c9e1f5b4a2d6e9c8f3a7b5d1c8e2f4a9d6b7c3e1f5b8a6d4c9e7b3f2a5c1d9b6f4e8a7c5b2d1f3e9b4a6c7d2f1b5e3c8a9d7f2b4e6c1a3d8b9f5c7e2a4b3f1d6c8e5b2a7f9b3c6e1d4f8a5c9e2b7d1f3a649e8fbc709682fd27b5374521000a9f7a84c1e31156eaf661db2cef3e738e9a05ed540487a805dd5098d19b5dd1eed610cff655279e2be39fb520c7713eb41258886210005a46e6de9311231b85da6d4f32c028847aa64bc04458861be442512db2056bae4a1d44d10d7013ddb5f8dcab1cc17f535d080974a219d4b0177fbf9"
	hexInput3 := "6"
	hexInput4 := "6f2c4a8d3b7c9e1f5b4a2d6e9c8f3a7b5d1c8e2f4a9d6b7c3e1f5b8a6d4c9e7b3f2a5c1d9b6f4e8a7c5b2d1f3e9b4a6c7d2f1b5e3c8a9d7f2b4e6c1a3d8b9f5c7e2a4b3f1d6c8e5b2a7f9b3c6e1d4f8a5c9e2b7d1f3a6d3a518e47135c29e70e5295c2e14e4dd0e1e9329c6788b6422b3e9b2e933e8f480fa9ebac93f72eea89b659ce5e1dee1d1626ef0d17449edf87f2d647720fd47cb811c1087ab4d4374d6e7883a46b4e1c49aa05b44863d43c66d63e2d9160ed3dbec0e1e27d8d9e3b69a65d39eeb0675e45e3e67a45d503eab92bcbc7f64f67"

	h1 := parseHex(hexInput1)
    h2 := parseHex(hexInput2)
	h3 := parseHex(hexInput3)
	h4 := parseHex(hexInput4)

	fmt.Println("Числа в нормальному виді:")


	str := ""
    for _, val := range h1 {
        str += fmt.Sprintf("%d", val)
    }
    fmt.Println("h2 в нормальному виді:", str)

	str = ""
    for _, val := range h2 {
        str += fmt.Sprintf("%d", val)
    }
    fmt.Println("h2 в нормальному виді:", str)

	str = ""
    for _, val := range h3 {
        str += fmt.Sprintf("%d", val)
    }
    fmt.Println("h2 в нормальному виді:", str)

	str = ""
    for _, val := range h4 {
        str += fmt.Sprintf("%d", val)
    }
    fmt.Println("h2 в нормальному виді:", str)
	

    fmt.Println("Перевірка тотожностей:")

	n := 100
	a := 5
	result, err := multiplyByAddition(n, a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Результат %d * %d = %d\n", n, a, result)
	}

    fmt.Printf("(a + b) ⋅ c = c ⋅ (a + b) = a ⋅ c + b ⋅ c : %v\n", testIdentity1(h1, h2, h3))

    repeats := 10000

    addTime := measureRepeatedTime(repeats, func() { addBigNumbers(h1, h2) })
    addResult := addBigNumbers(h1, h2)
    fmt.Printf("Додавання результат: %s\n", toHex(addResult))
    fmt.Printf("Час виконання (додавання): %v\n", addTime)

    subtractTime := measureRepeatedTime(repeats, func() { subtractBigNumbers(h1, h3) })
    subtractResult := subtractBigNumbers(h1, h3)
    fmt.Printf("Віднімання результат: %s\n", toHex(subtractResult))
    fmt.Printf("Час виконання (віднімання): %v\n", subtractTime)

    multiplyTime := measureRepeatedTime(repeats, func() { multiplyBigNumbers(h1, h2) })
    multiplyResult := multiplyBigNumbers(h1, h2)
    fmt.Printf("Множення результат: %s\n", toHex(multiplyResult))
    fmt.Printf("Час виконання (множення): %v\n", multiplyTime)

    squareTime := measureRepeatedTime(repeats, func() { squareBigNumber(h1) })
    squareResult := squareBigNumber(h1)
    fmt.Printf("Піднесення до квадрату результат: %s\n", toHex(squareResult))
    fmt.Printf("Час виконання (піднесення до квадрату): %v\n", squareTime)

    shiftLeftTime := measureRepeatedTime(repeats, func() { shiftLeft(h2, 1) })
    shiftLeftResult := shiftLeft(h2, 1)
    fmt.Printf("Зсув вліво результат: %s\n", toHex(shiftLeftResult))
    fmt.Printf("Час виконання (зсув вліво): %v\n", shiftLeftTime)

    // Зсув вправо
    shiftRightTime := measureRepeatedTime(repeats, func() { shiftRight(h2, 1) })
    shiftRightResult := shiftRight(h2, 1)
    fmt.Printf("Зсув вправо результат: %s\n", toHex(shiftRightResult))
    fmt.Printf("Час виконання (зсув вправо): %v\n", shiftRightTime)

    // Вивід номера старшого ненульового біта
    fmt.Printf("Старший ненульовий біт в h1: %d\n", highestSetBit(h1))
	fmt.Printf("Старший ненульовий біт в h2: %d\n", highestSetBit(h2))
 
    // Вивід конвертацій
    fmt.Println("\nКонвертації чисел:")

    fmt.Printf("Число h1 у шістнадцятковій системі: %s\n", toHex(h1))

    fmt.Printf("Число h2 у шістнадцятковій системі: %s\n", toHex(h2))

	powerHexTime := measureRepeatedTime(repeats, func() { powerBigNumber(h1, h3) })
	powerHexResult := powerBigNumber(h1, h3)
	fmt.Printf("Піднесення hex до hex результат: %s\n", toHex(powerHexResult))
	fmt.Printf("Час виконання (піднесення hex до hex): %v\n", powerHexTime)

	divideTime := measureRepeatedTime(repeats, func() {
     divideBigNumbers(h2, h4)})
    quotient := divideBigNumbers(h2, h4)
    fmt.Printf("Ділення результат : %s\n", toHex(quotient))
    fmt.Printf("Час виконання (ділення): %v\n", divideTime)
	
}