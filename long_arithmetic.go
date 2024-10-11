package long_arithmetic

import (
	"fmt"
	"math/big"
)

// Створення числа з фіксованим розміром (1024 біти)
const NumBits = 1024
const Base = 1 << 32 // Базова система B=2^32

// Додавання великих чисел
func AddBigInts(a, b *big.Int) (*big.Int, error) {
	result := new(big.Int).Add(a, b)
	return result, nil
}

// Віднімання великих чисел
func SubtractBigInts(a, b *big.Int) (*big.Int, error) {
	result := new(big.Int).Sub(a, b)
	return result, nil
}

// Множення великих чисел
func MultiplyBigInts(a, b *big.Int) (*big.Int, error) {
	result := new(big.Int).Mul(a, b)
	return result, nil
}

// Ділення великих чисел
func DivideBigInts(a, b *big.Int) (*big.Int, error) {
	if b.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("ділення на нуль")
	}
	result := new(big.Int).Div(a, b)
	return result, nil
}

// Зсув вліво (множення на 2^shift)
func ShiftLeft(a *big.Int, shift uint) (*big.Int, error) {
	result := new(big.Int).Lsh(a, shift)
	return result, nil
}

// Зсув вправо (ділення на 2^shift)
func ShiftRight(a *big.Int, shift uint) (*big.Int, error) {
	result := new(big.Int).Rsh(a, shift)
	return result, nil
}

// Знаходження старшого ненульового біта
func HighestNonZeroBit(a *big.Int) (int, error) {
	if a.Cmp(big.NewInt(0)) == 0 {
		return 0, fmt.Errorf("число дорівнює нулю")
	}
	return a.BitLen() - 1, nil
}

// Конвертація числа у шістнадцятковий рядок
func ToHexString(a *big.Int) string {
	return fmt.Sprintf("%x", a)
}

// Конвертація шістнадцяткового рядка у велике число
func FromHexString(hexStr string) (*big.Int, error) {
	result := new(big.Int)
	_, ok := result.SetString(hexStr, 16)
	if !ok {
		return nil, fmt.Errorf("неправильний формат шістнадцяткового числа")
	}
	return result, nil
}
