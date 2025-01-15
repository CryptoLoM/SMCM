package long_arithmetic


// Mod виконує операцію залишку
func (ln *BigInt) Mod(n *BigInt) *BigInt {
    // Перший етап: отримуємо частку від ділення
    quotient := ln.Divide(n)

    // Другий етап: множимо частку на модуль
    product := quotient.Multiply(n)

    // Третій етап: віднімаємо результат множення від оригінального числа
    remainder := ln.Subtract(product)

    // Якщо залишок від'ємний, додаємо модуль, щоб отримати додатний залишок
    if remainder.Compare(zero()) < 0 {
        remainder = remainder.Add(n)
    }

    return remainder
}

func (ln *BigInt) ModAdd(other, n *BigInt) *BigInt {
    sum := ln.Add(other)   // Додаємо два числа
    return sum.Mod(n)      // Знаходимо залишок за модулем n
}

func (ln *BigInt) ModSubtract(other, n *BigInt) *BigInt {
    diff := ln.Subtract(other) // Віднімаємо два числа
    if diff.Compare(NewBigIntFromUint32(0)) < 0 {
        diff = diff.Add(n)    // Якщо результат від'ємний, додаємо модуль
    }
    return diff.Mod(n)       
}


func (ln *BigInt) BarrettReduction(mod *BigInt, mu *BigInt) *BigInt {
    digitLength := mod.digitLength()

    if ln.Compare(mod) < 0 {
        return ln // Якщо ln < mod, залишаємо без змін
    }

    // Використовуємо Q як тимчасову змінну
    Q := ln.LongShiftDigitsToLow(digitLength - 1)
    Q = Q.Multiply(mu)
    Q = Q.LongShiftDigitsToLow(digitLength + 1)
    Q = Q.Multiply(mod)

    // R обчислюється на основі Q
    R := ln.Subtract(Q)

    // Зменшуємо R, якщо він >= мод
    if R.Compare(mod) >= 0 {
        R = R.Mod(mod)
    }

    return R
}


func (*BigInt) Mu(mod *BigInt) *BigInt {
    if mod == nil || len(mod.numberArray) == 0 {
        panic("Modulo is not initialized correctly")
    }

    digitLength := mod.digitLength()
    

    beta := NewBigIntFromUint32(1)
    beta = beta.LongShiftDigitsToHigh(digitLength * 2)


    mu := beta.Divide(mod)
   

    if mu == nil || len(mu.numberArray) == 0 || mu.IsZero() {
        panic("Calculated mu is invalid")
    }

    return mu
}

func (ln *BigInt) ModMultiply(other, mod *BigInt) *BigInt {
    mu := ln.Mu(mod)
    product := ln.Multiply(other)
    result := product.BarrettReduction(mod,mu)

    return result
}


// ModSquare виконує піднесення числа до квадрату за модулем
func (bn *BigInt) ModSquare(mod *BigInt) *BigInt {
    square := bn.Multiply(bn)
    
    // Обчислення залишку за модулем
    return square.Mod(mod)
}


func (ln *BigInt) LongModPowerBarrett(exponent *BigInt, mod *BigInt) *BigInt {
    if mod == nil || mod.IsZero() {
        panic("Modulo cannot be nil or zero")
    }
    if ln.errorFlag || exponent.errorFlag || mod.errorFlag {
        panic("Error in input BigInt")
    }


    mu := ln.Mu(mod)

    // Зменшуємо базу перед початком
    base := ln.BarrettReduction(mod,mu)
    result := NewBigIntFromUint32(1) // Ініціалізуємо результат

    // Експоненціація через Барретта
    for i := exponent.BitLength() - 1; i >= 0; i-- {
        result = result.Multiply(result) // Піднесення до квадрата
        result = result.BarrettReduction(mod, mu) // Використовуємо mu для швидшого зменшення

        if exponent.getBit(i) == 1 { // Якщо біт експоненти дорівнює 1
            result = result.Multiply(base)
            result = result.BarrettReduction(mod, mu) // Використовуємо попередньо обчислене mu
        }
    }

    return result
}


func (ln *BigInt) getBit(pos int) int {
    if pos < 0 || pos >= numLen*32 {
        return 0
    }
    return int((ln.numberArray[pos/32] >> (pos % 32)) & 1)
}



// Карацуба для множення великих чисел
func (ln *BigInt) KaratsubaMultiply(a, b *BigInt) *BigInt {
    // Базовий випадок: якщо довжина числа = 1, множимо прямо
    if a.digitLength() == 1 || b.digitLength() == 1 {
        return a.Multiply(b)
    }

    // Знаходимо розмір половини числа
    aLen := a.digitLength()
    bLen := b.digitLength()
    maxLen := aLen
    if bLen > aLen {
        maxLen = bLen
    }
    m := maxLen / 2

    // Розбиваємо числа на дві частини
    a1 := a.LongShiftDigitsToLow(m)       // старша частина числа a
    a0 := a.Subtract(a1.LongShiftDigitsToHigh(m)) // молодша частина числа a

    b1 := b.LongShiftDigitsToLow(m)       // старша частина числа b
    b0 := b.Subtract(b1.LongShiftDigitsToHigh(m)) // молодша частина числа b

    // Рекурсивно обчислюємо три множення
    z0 := ln.KaratsubaMultiply(a0, b0) // b0 * a0
    z2 := ln.KaratsubaMultiply(a1, b1) // b1 * a1
    z1 := ln.KaratsubaMultiply(a0.Add(a1), b0.Add(b1)) // (a0 + a1) * (b0 + b1)

    // Обчислюємо середнє множення: z1 - z0 - z2
    z1 = z1.Subtract(z0).Subtract(z2)

    // Формуємо результат
    result := z2.LongShiftDigitsToHigh(2 * m)
    z1 = z1.LongShiftDigitsToHigh(m)
    result = result.Add(z1).Add(z0)

    return result
}



