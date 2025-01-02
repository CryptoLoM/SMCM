package long_arithmetic

import (
	"fmt"
	"strings"
)

const numLen = 128
type BigInt struct {
	numberArray [numLen]uint32
	errorFlag   bool // Позначка, якщо сталося помилка
}

// zero повертає BigInt, що представляє число 0.
func zero() *BigInt {
    return &BigInt{} // Повертаємо новий BigInt, всі елементи numberArray будуть 0 за замовчуванням
}
// Конструктори
func NewBigInt() *BigInt {
	return &BigInt{}
}


func NewBigIntFromUint32(number uint32) *BigInt {
	ln := NewBigInt()
	ln.numberArray[0] = number
	return ln
}


func NewBigIntFromHex(hex string) *BigInt {
	bigInt := ParseHex(hex)
	if bigInt.errorFlag {
		return &BigInt{errorFlag: true} // Помилка позначається через errorFlag
	}
	return bigInt
}

// Перетворення з шістнадцяткового рядка
func ParseHex(hex string) *BigInt {
	result := NewBigInt()
	length := len(hex)

	for i := 0; i < length; i++ {
		digit := hexToDigit(hex[length-i-1])
		if digit < 0 {
			return &BigInt{errorFlag: true} // Неправильний символ
		}

		shift := (i % 8) * 4
		result.numberArray[i/8] |= uint32(digit) << shift
	}

	return result
}

func hexToDigit(hex byte) int {
	switch {
	case '0' <= hex && hex <= '9':
		return int(hex - '0')
	case 'a' <= hex && hex <= 'f':
		return int(hex-'a') + 10
	case 'A' <= hex && hex <= 'F':
		return int(hex-'A') + 10
	default:
		return -1
	}
}

// Перетворення в шістнадцятковий рядок
func (ln *BigInt) ToHex() string {
    if ln.errorFlag {
        return "error"
    }
    if ln.IsZero() {
        return "0"
    }
    var builder strings.Builder
    for i := numLen - 1; i >= 0; i-- {
        fmt.Fprintf(&builder, "%08x", ln.numberArray[i])
    }
    hex := builder.String()
    return strings.TrimLeft(hex, "0")
}


// Додавання
func (ln *BigInt) Add(b *BigInt) *BigInt {
    if ln.errorFlag || b.errorFlag {
        return &BigInt{errorFlag: true}
    }

    result := NewBigInt()
    carry := uint32(0)

    for i := 0; i < numLen; i++ {
        sum := uint64(ln.numberArray[i]) + uint64(b.numberArray[i]) + uint64(carry)
        result.numberArray[i] = uint32(sum)
        carry = uint32(sum >> 32) // Перенос
    }

    // Перевірити, чи залишився переніс після останнього числа
    if carry != 0 {
        return &BigInt{errorFlag: true} // Це може означати переповнення
    }

    return result
}


// Віднімання
func (ln *BigInt) Subtract(b *BigInt) *BigInt {
	if ln.errorFlag || b.errorFlag {
		return &BigInt{errorFlag: true}
	}

	result := NewBigInt()
	borrow := uint32(0)

	// Визначити, яке число більше
	if ln.Compare(b) < 0 {
		ln, b = b, ln // Поміняти місцями, якщо b > ln
	}

	for i := 0; i < numLen; i++ {
		diff := int64(ln.numberArray[i]) - int64(b.numberArray[i]) - int64(borrow)
		if diff >= 0 {
			result.numberArray[i] = uint32(diff)
			borrow = 0
		} else {
			result.numberArray[i] = uint32(diff + (1 << 32))
			borrow = 1
		}
	}

	return result
}


func (ln *BigInt) LongMulOneDigit(b uint32) *BigInt {
    result := NewBigInt()
    carry := uint32(0)

    for i := 0; i < numLen; i++ {
        temp := uint64(ln.numberArray[i])*uint64(b) + uint64(carry)
        result.numberArray[i] = uint32(temp & 0xFFFFFFFF) // Зберігаємо молодші 32 біти
        carry = uint32(temp >> 32)                               // Зберігаємо перенос
    }

    // Якщо залишився перенос, додаємо його
   
    result.numberArray[numLen-1] = uint32(carry)
    

    return result
}


func (ln *BigInt) Multiply(b *BigInt) *BigInt {
	result := NewBigInt()

    for i := 0; i < numLen; i++ {

        temp := ln.LongMulOneDigit(b.numberArray[i])
        temp = temp.LongShiftDigitsToHigh(i)
        result = result.Add(temp)
    }

    return result
}

func (ln *BigInt) LongShiftDigitsToHigh(shift int) *BigInt {
    result := NewBigInt()

    if shift >= numLen {
        return result // Якщо зсув більше довжини числа, повертаємо 0
    }

    for i := 0; i < numLen - shift; i++ {
        result.numberArray[i + shift] = ln.numberArray[i]
    }

    return result
}

func (ln *BigInt) LongShiftDigitsToLow(shift int) *BigInt {
    result := NewBigInt()

    // Якщо зсув більше або рівний довжині числа, повертаємо 0
    if shift >= len(ln.numberArray) {
        return result
    }

    // Переміщаємо елементи вправо, заповнюючи нулями зліва
    for i := len(ln.numberArray) - 1; i >= shift; i-- {
        result.numberArray[i-shift] = ln.numberArray[i]
    }

    return result
}


// Ділення з цілою частиною
func (ln *BigInt) Divide(b *BigInt) (*BigInt) {
	if ln.errorFlag || b.errorFlag || b.IsZero() {
		return &BigInt{errorFlag: true}
	}
	if b.BitLength() == 1 && b.numberArray[0] == 1{
		return ln.Copy()
	}

	quotient := NewBigInt()
	remainder := ln.Copy()

	for remainder.Compare(b) >= 0 {
		shift := remainder.BitLength() - b.BitLength()
		shifted := b.ShiftLeft(shift)

		if remainder.Compare(shifted) < 0 {
			shift--
			shifted = b.ShiftLeft(shift)
		}

		remainder = remainder.Subtract(shifted)
		quotient = quotient.Add(NewBigIntFromUint32(1).ShiftLeft(shift))
	}

	return quotient
}

// Бітова довжина числа
func (ln *BigInt) BitLength() int {
	for i := numLen - 1; i >= 0; i-- {
		if ln.numberArray[i] != 0 {
			for bit := 31; bit >= 0; bit-- {
				if ln.numberArray[i]&(1<<bit) != 0 {
					return i*32 + bit + 1
				}
			}
		}
	}
	return 0
}


func (ln *BigInt) digitLength() int {
	for  i := numLen - 1; i >= 0; i-- {
		if (ln.numberArray[i] != 0){
			return i + 1
		}
	}
	return 0
}

func (ln *BigInt) Copy() *BigInt {
	copy := NewBigInt()
	for i := 0; i < numLen; i++ {
		copy.numberArray[i] = ln.numberArray[i]
	}
	return copy
} 

// Перевірка на нуль
func (ln *BigInt) IsZero() bool {
	for _, v := range ln.numberArray {
		if v != 0 {
			return false
		}
	}
	return true
}



// Зсув вліво
func (ln *BigInt) ShiftLeft(bits int) *BigInt {
	result := NewBigInt()
	digitShift := bits / 32
	bitShift := bits % 32
	carry := uint32(0)

	for i := 0; i < numLen-digitShift; i++ {
		nextCarry := ln.numberArray[i] >> (32 - bitShift)
		result.numberArray[i+digitShift] = (ln.numberArray[i] << bitShift) | carry
		carry = nextCarry
	}

	return result
}

// RightShift виконує побітовий зсув числа вправо
func (ln *BigInt) RightShift(bits int) *BigInt {
    result := NewBigInt()
    digitShift := bits / 32
    bitShift := bits % 32

    for i := digitShift; i < numLen; i++ {
        result.numberArray[i-digitShift] = ln.numberArray[i] >> bitShift
        if i+1 < numLen {
            result.numberArray[i-digitShift] |= ln.numberArray[i+1] << (32 - bitShift)
        }
    }

    return result
}


// Порівняння
func (ln *BigInt) Compare(b *BigInt) int {
	for i := numLen - 1; i >= 0; i-- {
		if ln.numberArray[i] > b.numberArray[i] {
			return 1
		} else if ln.numberArray[i] < b.numberArray[i] {
			return -1
		}
	}
	return 0
}


func (ln *BigInt) Square() *BigInt {
    return ln.Multiply(ln)
}


func (ln *BigInt) Power(exponent *BigInt) *BigInt {
    result := NewBigIntFromUint32(1) // Початковий результат
    base := ln.Copy()               // Базове число

    for i := 0; i < exponent.BitLength(); i++ {
        if exponent.BitAt(i) == 1 {
            result = result.Multiply(base)
        }
        base = base.Multiply(base) // Квадрат бази
    }

    return result
}

func (ln *BigInt) BitAt(i int) uint32 { 
	if i < 0 || i >= ln.BitLength() { 
	 return 0 // Якщо біт за межами числа, повертаємо 0 
	} 
	block := i / 32                // Номер 32-бітного блоку 
	bit := i % 32               // Зсув у блоці 
	return (ln.numberArray[block] >> bit) & 1 // Зсув і перевірка 


	
   }
