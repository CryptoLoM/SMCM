package main

import (
	"fmt"
	"time"
)

const FIELD_SIZE = 179
const ARRAY_SIZE = FIELD_SIZE/64 + 1

// POLYNOMIAL_DEGREES визначає степені полінома
var POLYNOMIAL_DEGREES = [5]int{0, 1, 2, 4, 179}

// Тип для елемента поля
type FieldElement struct {
	coefficients [ARRAY_SIZE]uint64
}

// Новий елемент поля 
func NewFieldElement() *FieldElement {
	return &FieldElement{}
}

// Конструктор з рядка
func NewFieldElementFromString(binary string) *FieldElement {
	element := NewFieldElement()
	length := len(binary)
	for i := 0; i < length; i++ {
		if binary[i] == '1' {
			element.coefficients[(length-1-i)/64] |= 1 << ((length-1-i) % 64)
		}
	}
	return element
}


func Zero() *FieldElement {
	return NewFieldElement()
}

func One() *FieldElement {
	element := NewFieldElement()
	element.coefficients[0] = 1
	return element
}

// Перевірка на нульовий елемент
func (a *FieldElement) IsZero() bool {
    for i := 0; i < ARRAY_SIZE; i++ {
        if a.coefficients[i] != 0 {
            return false
	
        }
		
    }
    return true
}

// Операція додавання (XOR)
func (a *FieldElement) Add(b *FieldElement) *FieldElement {
	result := NewFieldElement()
	for i := 0; i <= FIELD_SIZE / 64; i++ {
		result.coefficients[i] = a.coefficients[i] ^ b.coefficients[i]
	}
	return result
}

// Операція піднесення до квадрату (циклічний зсув вправо)
func (a *FieldElement) Square() *FieldElement {
	var tempResult [ARRAY_SIZE * 2]uint64

    for i := 0; i < FIELD_SIZE; i++ {
        bit := (a.coefficients[i/64] >> uint(i%64)) & 1
        tempResult[(2*i)/64] |= bit << uint((2*i)%64)
    }

    // Зменшуємо поліном і повертаємо результат
    return a.ReducePolynomial(tempResult[:]) // Передаємо як зріз
}


// Операція множення через матрицю Λ
func (a *FieldElement) Mul(b *FieldElement) *FieldElement {
	var result [ARRAY_SIZE * 2]uint64
	for i := 0; i < FIELD_SIZE; i++ {
		for j := 0; j < FIELD_SIZE; j++ {
			if (a.coefficients[i/64]>>uint(i%64))&1 == 1 && (b.coefficients[j/64]>>uint(j%64))&1 == 1 {
				result[(i+j)/64] ^= 1 << uint((i+j)%64)
			}
		}
	}
	return a.ReducePolynomial(result[:])
}

// Знаходження ступеня полінома
func PolynomialDegree(poly []uint64) int {
	for i := len(poly)*64 - 1; i >= 0; i-- {
		if (poly[i/64]>>uint(i%64))&1 == 1 {
			return i
		}
	}
	return 0
}

// Операція зменшення полінома
func (a *FieldElement) ReducePolynomial(poly []uint64) *FieldElement {
	// Знайдемо старший коефіцієнт
	higher_deg := PolynomialDegree(poly)
	for higher_deg >= FIELD_SIZE {
		for _, deg := range POLYNOMIAL_DEGREES {
			index := higher_deg + deg - FIELD_SIZE
			poly[index/64] ^= 1 << uint(index%64)
		}
		higher_deg = PolynomialDegree(poly)
	}

	result := NewFieldElement()
	for i := 0; i < ARRAY_SIZE; i++ {
		result.coefficients[i] = poly[i]
	}
	return result
}


// Піднесення до степеня
func (a *FieldElement) Power(bit *FieldElement) *FieldElement {
	
	result := One()
	result.coefficients[0] = 1
	for i := FIELD_SIZE - 1; i >= 0; i-- {
		if (bit.coefficients[i/64]>>uint(i%64))&1 == 1 {
			result = result.Mul(a)
		}
		if i > 0{
			result = result.Square()
		} 
		
	}
	return result
}

// Зворотній елемент через схему Горнера
func (a *FieldElement) Inv() *FieldElement {
	result := NewFieldElement()
	result.coefficients[0] = 1
	for i := 1; i < FIELD_SIZE; i++ {
		a = a.Square()
		result = result.Mul(a)
	}
	return result
}

// Обчислення сліду
func (*FieldElement) Trace() *FieldElement {

	temp := NewFieldElement()
	result := NewFieldElement()
	result.coefficients[0] = temp.coefficients[0]
	for i := 1; i < FIELD_SIZE; i++ {
		temp = temp.Square()
		result = result.Add(temp)
	}
	return result
}

// Перетворення в рядок
func (a *FieldElement) ToBinaryString() string {
	binary := ""
	for i := FIELD_SIZE - 1; i >= 0; i-- {
		if (a.coefficients[i/64]>>uint(i%64))&1 == 1 {
			binary += "1"
		} else {
			binary += "0"
		}
	}
	return binary
}

// Друк елемента поля
func (a *FieldElement) String() string {
	return a.ToBinaryString()
}


func measureTime(operation func() *FieldElement, iterations int) int64 {
	start := time.Now() // Start time
	for i := 0; i < iterations; i++ {
		operation() // Execute the operation multiple times
	}
	duration := time.Since(start) // Total time for all iterations
	return duration.Nanoseconds() / int64(iterations) // Average time per operation in nanoseconds
}


func CheckAdditiveIdentity(a *FieldElement) bool {
	zero := Zero()
	return !a.Add(zero).IsZero() && !a.IsZero()
}

func CheckMultiplicativeIdentity(a *FieldElement) bool {
	one := One()
	return a.Mul(one).String() == a.String()
}

func CheckDistributivity(a, b, c *FieldElement) bool {
	left := a.Add(b).Mul(c)
	right := a.Mul(c).Add(b.Mul(c))
	return left.String() == right.String()
}

func CheckAssociativityAddition(a, b, c *FieldElement) bool {
	left := a.Add(b).Add(c)
	right := a.Add(b.Add(c))
	return left.String() == right.String()
}

func CheckAssociativityMultiplication(a, b, c *FieldElement) bool {
	left := (a.Mul(b)).Mul(c)
	right := a.Mul(b.Mul(c))
	return left.String() == right.String()
}

func CheckCommutativityAddition(a, b *FieldElement) bool {
	return a.Add(b).String() == b.Add(a).String()
}

func CheckCommutativityMultiplication(a, b *FieldElement) bool {
	return a.Mul(b).String() == b.Mul(a).String()
}



func TestFieldProperties() {
	// Create some field elements for testing
	a := NewFieldElementFromString("01000110101101011100100100011111110001000111010111101110011101110111001011011101001010101100000111011001111101100000001001100001110001100110000111010010110001101100100001011001100110110010000")
	b := NewFieldElementFromString("01000001011001110010111010111100110000100111010010000011110001001011100011011110001101101000010011001001101100000110000100100010001011110010101001010110001101110001010000010110100001111111000")
	c := NewFieldElementFromString("11011111111010000111101011101010101001101001000101000111111010111011000001110101100111100011110111111101010000001011110100101000001100001111011110001000001100100010111110000101010010000010110")


	// Check Additive Identity
	fmt.Println("Check Additive Identity:", CheckAdditiveIdentity(a))

	// Check Multiplicative Identity
	fmt.Println("Check Multiplicative Identity:", CheckMultiplicativeIdentity(a))

	// Check Distributivity
	fmt.Println("Check Distributivity:", CheckDistributivity(a, b, c))

	// Check Associativity of Addition
	fmt.Println("Check Associativity of Addition:", CheckAssociativityAddition(a, b, c))

	// Check Associativity of Multiplication
	fmt.Println("Check Associativity of Multiplication:", CheckAssociativityMultiplication(a, b, c))

	// Check Commutativity of Addition
	fmt.Println("Check Commutativity of Addition:", CheckCommutativityAddition(a, b))

	// Check Commutativity of Multiplication
	fmt.Println("Check Commutativity of Multiplication:", CheckCommutativityMultiplication(a, b))

}


func main() {
	// Створення елементів поля з рядків
	f := NewFieldElementFromString("01010010010001011011001101101010110110101100010001100101011010010111010110000111101000011111010011000110000111001000001010101001000110101100010100000010010001110011100111010100011")
	g := NewFieldElementFromString("10000110111010100001100011001011110110101101100101111101000110000100110001111100100100111001110100000100111110000100101110101111010100101100000001000101101111000100101011011101111")
	h := NewFieldElementFromString("01010100001110100010110010010111100110010101010110101001001001100111111111100011001110101010000001110011010000000001111011011100100111001101111010011010100010000000100011010001100")
	
	fmt.Println("A = ", f)
	fmt.Println("B = ", g)
	fmt.Println("N = ", h)

	add := f.Add(g)
	fmt.Println("A+B = ", add)
	
	mul := f.Mul(g)
	fmt.Println("A*B = ", mul)

	square := f.Square()
	fmt.Println("A^2 = ", square)

	inv := f.Inv()
	fmt.Println("A^-1 = ", inv)

	power := f.Power(h)
	fmt.Println("A^N = ", power)

	trace := f.Trace()
	fmt.Println("Tr(A) = ", trace)


	addTime := measureTime(func() *FieldElement {
		return f.Add(g) // Додавання
	},1000000)
	mulTime := measureTime(func() *FieldElement {
		return f.Mul(g) // Множення
	},10)
	squareTime := measureTime(func() *FieldElement {
		return f.Square() // Піднесення до квадрату
	},100)
	powerTime := measureTime(func() *FieldElement {
		return f.Power(h) // Піднесення до степеня
	},1)
	invTime := measureTime(func() *FieldElement {
		return f.Inv() // обернений елемент
	},1)
	traceTime := measureTime(func() *FieldElement {
		return f.Trace() // Додавання
	},1)

	// Виведення результатів вимірювання часу
	fmt.Printf("Час виконання додавання: %v\n", addTime)
	fmt.Printf("Час виконання множення: %v\n", mulTime)
	fmt.Printf("Час виконання піднесення до квадрату: %v\n", squareTime)
	fmt.Printf("Час виконання піднесення до степеня: %v\n", powerTime)
	fmt.Printf("Час виконання знаходження оберненого: %v\n", invTime)
	fmt.Printf("Час виконання знаходження сліду: %v\n", traceTime)

	TestFieldProperties()
}
