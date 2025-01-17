package main

import (
	"fmt"
	"strings"
	"time"
	
	
)

const FIELD_SIZE = 179
const ARRAY_SIZE = (FIELD_SIZE / 64) + 1

type FieldElement struct {
	coefficients [ARRAY_SIZE]uint64
}

// Створення нового FieldElement 
func NewFieldElement() FieldElement {
	fe := FieldElement{}
	return fe
}

// Копіювання FieldElement
func NewFieldElementFromString(other FieldElement) FieldElement {
	fe := FieldElement{}
	for i := 0; i < ARRAY_SIZE; i++ {
		fe.coefficients[i] = other.coefficients[i]
	}
	return fe
}

// Створення FieldElement з бінарного рядка
func NewFieldElementFromBinary(binary string) FieldElement {
	other := ConvertToNumber(binary)
	return NewFieldElementFromString(other)
}

func (f FieldElement) Add(other FieldElement) FieldElement {
	result := FieldElement{}
	for i := 0; i < ARRAY_SIZE; i++ {
		result.coefficients[i] = f.coefficients[i] ^ other.coefficients[i]
	}
	return result
}

func (f FieldElement) Mul(other FieldElement) FieldElement {
    result := NewFieldElement()
    left := NewFieldElementFromString(f)
    right := NewFieldElementFromString(other)


  
    multMatrix := calculateMultiplicativeMatrix()
   
    for i := 0; i < FIELD_SIZE; i++ {
        // Формуємо temp = left * M, де M — наша матриця
        temp := NewFieldElement()
        for j := 0; j < FIELD_SIZE; j++ {
            // (left ⋅ matrix[j]) mod 2 – “скалярний добуток” із j-тою колонкою 
            bit := (left.And(multMatrix[j])).Weight() % 2
            if bit == 1 {
                temp.SetBit(j)
            }
        }

        // Тепер перевіряємо, чи temp має непарну кількість спільних одиничних бітів із right
        // Якщо так — тоді i-ий біт result є 1
		temp = temp.And(right)
        if temp.Weight()& 1 == 1 {
            result.SetBit(i)
        }

        
        left = left.LeftCycleShift()
        right = right.LeftCycleShift()
    }
    return result
}

func calculateMultiplicativeMatrix() [FIELD_SIZE]FieldElement {
	var multmatrix [FIELD_SIZE]FieldElement
	p := FIELD_SIZE*2 + 1
	
	powers := make([]int, FIELD_SIZE)
	temp := make([][]bool, FIELD_SIZE)
	for i := range temp {
		temp[i] = make([]bool, FIELD_SIZE)
	}

	powers[0] = 1
	for i := 1; i < FIELD_SIZE; i++ {
		powers[i] = (2 * powers[i-1]) % p
	}
	



    for i := 0; i < FIELD_SIZE; i++ {
		
        for j := i; j < FIELD_SIZE; j++ {
            cond1 := (powers[i] + powers[j]) % p == 1
            cond2 := (powers[i] + powers[j]) % p == p-1
            cond3 := (powers[i] - powers[j]) % p == 1
            cond4 := (powers[i] - powers[j]) % p == -1

     

            if cond1 || cond2 || cond3 || cond4 {
                temp[i][j] = true
                temp[j][i] = true
      
        }
    }
}

	for i := 0; i < FIELD_SIZE; i++ {
		var column FieldElement
		for j := 0; j < FIELD_SIZE; j++ {
			if temp[i][j] {
				column.SetBit(j)
			}
		}
		multmatrix[i] = column
	}

	return multmatrix
}


func (f FieldElement) LeftCycleShift() FieldElement {
	result := FieldElement{}
	right := FIELD_SIZE % 64

	for i := 0; i < ARRAY_SIZE - 1; i++ {
		result.coefficients[i] = (f.coefficients[i] >> 1) + (f.coefficients[i+1] << 63)
	}
	temp := f.coefficients[0] & 1
	result.coefficients[ARRAY_SIZE-1] = (f.coefficients[ARRAY_SIZE-1] >> 1) + (uint64(temp) << (right - 1))
	return result
}

func (f FieldElement) RightCycleShift() FieldElement {
	result := FieldElement{}
	right := FIELD_SIZE % 64
	temp := (f.coefficients[ARRAY_SIZE-1] >> (right - 1)) & 1
	result.coefficients[0] = (f.coefficients[0] << 1) | temp

	for i := 1; i < ARRAY_SIZE; i++ {
		result.coefficients[i] = (f.coefficients[i] << 1) | (f.coefficients[i-1] >> 63)
	}

	result.coefficients[ARRAY_SIZE-1] &= ^(uint64(1) << right)
	return result
}


func (f FieldElement) Square() FieldElement {
	return f.RightCycleShift()
}

func (f FieldElement) Trace() FieldElement {
	if (f.Weight() & 1) == 1 {
		return Zero()
	}
	return One()
}

func (f FieldElement) Equal(other FieldElement) bool {
	for i := 0; i < ARRAY_SIZE; i++ {
		if f.coefficients[i] != other.coefficients[i] {
			return false
		}
	}
	return true
}


func (f FieldElement) And(other FieldElement) FieldElement {
	result := NewFieldElement()
	for i := 0; i < ARRAY_SIZE; i++ {
		result.coefficients[i] = f.coefficients[i] & other.coefficients[i]
	}
	return result
}


func (f *FieldElement) SetBit(index int) {
	f.coefficients[index/64] ^= uint64(1) << (index % 64)
}



func (f FieldElement) GetBit(index int) int {
    return int((f.coefficients[index/64] >> (index % 64)) & 1)
}


func (f FieldElement) Weight() int {
	weight := 0
	for i := 0; i <= FIELD_SIZE/64; i++ {
		weight += popCount(f.coefficients[i])
	
	}
	return weight
}


func (f FieldElement) Power(el FieldElement) FieldElement {
	result := One()
	for i := 0; i < FIELD_SIZE; i++ {
		if el.GetBit(i) == 1 {
			result = f.Mul(result)
		}
		if i != FIELD_SIZE-1 {
			result = result.Square()
		}
	}
	return result
}


func (a FieldElement) Inv() FieldElement {
	result := One()
	for i := 0; i < FIELD_SIZE - 1; i++ {
		a= a.Square()
		result = result.Mul(a)
	}
	return result
}

func One() FieldElement {
	result := NewFieldElement()
	for i := 0; i < ARRAY_SIZE; i++ {
		result.coefficients[i] = ^uint64(0)
	}
	result.coefficients[ARRAY_SIZE-1] >>= (64 - FIELD_SIZE%64)
	return result
}

func Zero() FieldElement {
	return NewFieldElement()
}


func ConvertToNumber(binary string) FieldElement {
	result := NewFieldElement()
	for i := 0; i < len(binary); i++ {
		if binary[i] == '1' {
			result.coefficients[i/64] |= uint64(1) << (i % 64)
		}
	}
	return result
}


func ConvertToBinary(element FieldElement) string {
	var builder strings.Builder
	for i := 0; i < FIELD_SIZE; i++ {
		bit := (element.coefficients[i/64] >> (i % 64)) & 1
		if bit == 1 {
			builder.WriteByte('1')
		} else {
			builder.WriteByte('0')
		}
	}
	return builder.String()
}

func popCount(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

func measureTime(operation func() FieldElement, iterations int) int64 {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		operation()
	}
	duration := time.Since(start)
	return duration.Nanoseconds() / int64(iterations)
}


func (fe FieldElement) IsZero() bool {
	for _, coeff := range fe.coefficients {
		if coeff != 0 {
			return false
		}
	}
	return true
}

func TestAssociativity() {
    a := NewFieldElementFromBinary("1011")
    b := NewFieldElementFromBinary("1101")
    c := NewFieldElementFromBinary("0110")

 
    add1 := a.Add(b).Add(c) // (a + b) + c
    add2 := a.Add(b.Add(c)) // a + (b + c)
    if !add1.Equal(add2) {
        fmt.Println("Ассоціативність додавання НЕ виконується!")
    } else {
        fmt.Println("Ассоціативність додавання виконується.")
    }

 
    mul1 := (a.Mul(b)).Mul(c) // (a * b) * c
    mul2 := a.Mul(b.Mul(c)) // a * (b * c)
    if !mul1.Equal(mul2) {
        fmt.Println("Ассоціативність множення НЕ виконується!")
    } else {
        fmt.Println("Ассоціативність множення виконується.")
    }
}

func TestCommutativity() {
    a := NewFieldElementFromBinary("1011")
    b := NewFieldElementFromBinary("1101")


    add1 := a.Add(b)
    add2 := b.Add(a)
    if !add1.Equal(add2) {
        fmt.Println("Комутативність додавання НЕ виконується!")
    } else {
        fmt.Println("Комутативність додавання виконується.")
    }

    mul1 := a.Mul(b)
    mul2 := b.Mul(a)
    if !mul1.Equal(mul2) {
        fmt.Println("Комутативність множення НЕ виконується!")
    } else {
        fmt.Println("Комутативність множення виконується.")
    }
}

func TestDistributivity() {
    a := NewFieldElementFromBinary("1011")
    b := NewFieldElementFromBinary("1101")
    c := NewFieldElementFromBinary("0110")


    left := a.Mul(b.Add(c))    // a * (b + c)
    right := a.Mul(b).Add(a.Mul(c)) // (a * b) + (a * c)
    if !left.Equal(right) {
        fmt.Println("Дистрибутивність НЕ виконується!")
    } else {
        fmt.Println("Дистрибутивність виконується.")
    }
}

func TestSelfMultiplication() {
    a := NewFieldElementFromBinary("11001010001111001001111001110010011100110111110101000010101001000110010000100100100010111101100100100011000000100110100000011011111001111010000000101000001111101101000111001110100") 
    mulResult := a.Mul(a)
    squareResult := a.Square()

    fmt.Println("Результат множення вектора на себе:", ConvertToBinary(mulResult))
    fmt.Println("Результат піднесення до квадрату:", ConvertToBinary(squareResult))

 
    if mulResult.Equal(squareResult) {
        fmt.Println("Результати збігаються: множення вектора на себе дорівнює піднесенню до квадрату.")
    } else {
        fmt.Println("Результати НЕ збігаються: множення вектора на себе не дорівнює піднесенню до квадрату.")
    }
}

func TestWeightAndPopCount() {
 
    testVectors := []struct {
        binary    string
        expected  int
    }{
        {"11001010001111001001111001110010011100110111110101000010101001000110010000100100100010111101100100100011000000100110100000011011111001111010000000101000001111101101000111001110100", 83},
        {"11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111", 179},
        {"10101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101", 90},
        {"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0},
        {"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001", 1},
        {"10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 1},
    }

    for _, vector := range testVectors {
        element := NewFieldElementFromBinary(vector.binary)

        weight := element.Weight()

        fmt.Printf("Вектор: %s, Очікувана вага: %d, Обчислена вага: %d\n",
            vector.binary, vector.expected, weight)

        // Перевірка
        if weight != vector.expected {
            fmt.Println("Помилка: вага не відповідає очікуваному значенню!")
        } else {
            fmt.Println("Результат правильний.")
        }
    }
	fmt.Println("Попка каунта:")
    testValues := []struct {
        binary    string
        expected  int
    }{
        {
            binary:  "11001010001111001001111001110010011100110111110101000010101001000110010000100100100010111101100100100011000000100110100000011011111001111010000000101000001111101101000111001110100",
            expected: 83,
        },
        {
            binary:  "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
            expected: 179,
        },
        {
            binary:  "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
            expected: 0,
        },
    }

    for _, test := range testValues {
        count := 0
        for i := 0; i < len(test.binary); i++ {
            if test.binary[i] == '1' {
                count++
            }
        }

        fmt.Printf("Значення: %s\nОчікувано: %d, Отримано: %d\n",
            test.binary, test.expected, count)

        if count != test.expected {
            fmt.Println("Помилка: результат не відповідає очікуваному значенню!")
        } else {
            fmt.Println("Результат правильний.")
        }
    }

}

func main() {
	a := NewFieldElementFromBinary("11001010001111001001111001110010011100110111110101000010101001000110010000100100100010111101100100100011000000100110100000011011111001111010000000101000001111101101000111001110100")
	b := NewFieldElementFromBinary("11000010110101010011100111111011011101000101001101010110001100001110111100000010111110101001001101100001010101001110001111011010000010110000000110000101100111100011100001101011101")
	n := NewFieldElementFromBinary("11000010110101010011100111111011011101000101001101010110001100001110111100000010111110101001001101100001010101001110001111011010000010110000000110000101100111100011100001101011101")

	fmt.Println("a:", ConvertToBinary(a))
	fmt.Println("b:", ConvertToBinary(b))
	fmt.Println("n:", ConvertToBinary(n))

	fmt.Println("Тестування роботи ваги:")
	TestWeightAndPopCount()

	fmt.Println("")
	TestSelfMultiplication()

	fmt.Println("Тестування асоціативності:")
    TestAssociativity()

    fmt.Println("\nТестування комутативності:")
    TestCommutativity()

    fmt.Println("\nТестування дистрибутивності:")
    TestDistributivity()

	fmt.Println("a + b:", ConvertToBinary(a.Add(b)))
	fmt.Println("a * b:", ConvertToBinary(a.Mul(b)))
	fmt.Println("a ^ 2:", ConvertToBinary(a.Square()))
	fmt.Println("a ^ -1:", ConvertToBinary(a.Inv()))
	fmt.Println("a ^ n:", ConvertToBinary(a.Power(n)))
	fmt.Println("Tr(a):", ConvertToBinary(a.Trace()))
	

	addTime := measureTime(func() FieldElement {
		return a.Add(b) // Додавання
	},1)
	mulTime := measureTime(func() FieldElement {
		return a.Mul(b) // Множення
	},1)
	squareTime := measureTime(func() FieldElement {
		return a.Square() // Піднесення до квадрату
	},1)
	powerTime := measureTime(func() FieldElement {
		return a.Power(b) // Піднесення до степеня
	},1)
	invTime := measureTime(func() FieldElement {
		return a.Inv() // обернений елемент
	},1)
	traceTime := measureTime(func() FieldElement {
		return a.Trace() // Додавання
	},1)

	fmt.Printf("Час виконання додавання: %v\n", addTime)
	fmt.Printf("Час виконання множення: %v\n", mulTime)
	fmt.Printf("Час виконання піднесення до квадрату: %v\n", squareTime)
	fmt.Printf("Час виконання піднесення до степеня: %v\n", powerTime)
	fmt.Printf("Час виконання знаходження оберненого: %v\n", invTime)
	fmt.Printf("Час виконання знаходження сліду: %v\n", traceTime)
}