package main

import (
	"fmt"
	"smcm/long_arithmetic"
	"time"
)


func measureRepeatedTime(repeats int, f func()) time.Duration {
	start := time.Now()
	for i := 0; i < repeats; i++ {
		f()
	}
	return time.Since(start)
}


func main() {
	fmt.Printf("\n")
	repeats := 1000


	a := long_arithmetic.ParseHex("cef48dbb53b8bea5f89905703f7f9441be4abae58aab3f9eb4e8f559bdeadcd365d888e9f7c90f9f4be023818bcb71fddd800b26db081b0524c2aa09676bce6e5cb1753b6b14f3afb66063af006f6ee9f300a789be4382e216606407c14644c526e1169e3b497fae02364374538f7d94408bbc6ee98c830f10cc97af3dbd4f1ff5df67a0aead5ac2b76eb8d56b6256e8f36d5f0e0e06f6ac6c7930074f725ad0b0ec7a9c533ee87a74d10836600ce8c8ea290a400778b5de7e854f3935b3494db98b617251db5cb698e95dcfeb5c93d6454101fe440c4d5902bd88833a65827958e3c411766137685b0b8b98453fa63fb48d49ecbb6f0d9cafa999f844341d6b")
	b := long_arithmetic.ParseHex("7762eca0d63bf84694ff81b0ef356bd89760377107b18ec7d116dba9ec8771fe363dcc04dfca3a735a90ecd641c883fa7b72e1aaeb5aab15a35d5b5b4e5ef24050dc43a8c1a8133d70c45ce9ca40c03b71fd0cd90faa508968e05cccd0cc7ad1fea6b891cb65954851f6ad61e001623ab53336e082c7670b056a65017caff175f44a7658f34d6daf76dc26095a857f0d694b3e945c001894ebb2f78f9021d24fda50b10e2ca006bfc5e48d8893f363766f8af250f85cc6f676dbedd04e1c186668e15bc9bcf77da752c0f795d2afa3c140dd0ee996255ecc5f3443945b10fa519d77860f995f3dac92787ee243aa9f93b7c35eabca713481275d851353d99f88")
	c := long_arithmetic.ParseHex("c6dfa643d0298ec57474768f363e65fa456d331cae5dc9cf9f2ee1f49aa9b479786312ad63bd7a377deac26acd94686d2e50629e5936a4b122826779ab4ada6f9b8da264eb9173e6744bf3ef387a496b70811873f7a1c735a28d10d89be7917672b77476dc1f9b9981a9022bc5a511f80941c7a72e78211e2b2fb5da07d7a557c6dfa643d0298ec57474768f363e65fa456d331cae5dc9cf9f2ee1f49aa9b479786312ad63bd7a377deac26acd94686d2e50629e5936a4b122826779ab4ada6f9b8da264eb9173e6744bf3ef387a496b70811873f7a1c735a28d10d89be7917672b77476dc1f9b9981a9022bc5a511f80941c7a72e78211e2b2fb5da07d7a557")
    g := long_arithmetic.ParseHex("1000")
	h := long_arithmetic.ParseHex("3") 

	
	fmt.Printf("a в hex: %s\n", a.ToHex())
	fmt.Printf("\n")
	fmt.Printf("b в hex: %s\n", b.ToHex())
	fmt.Printf("\n")

	sum := a.Add(b)
	AddTime := measureRepeatedTime(repeats, func() { a.Add(b) })
	fmt.Printf("Додавання: a + b = %s\n", sum.ToHex())
	fmt.Printf("Час виконання : %v\n", AddTime)
	fmt.Printf("\n")

	diff := a.Subtract(b)
	SubTime := measureRepeatedTime(repeats, func() { a.Subtract(b) })
	fmt.Printf("Віднімання: a - b = %s\n", diff.ToHex())
	fmt.Printf("Час виконання : %v\n", SubTime) 
	fmt.Printf("\n")

	product := a.Multiply(b)
	MultiplyTime := measureRepeatedTime(repeats, func() { a.Multiply(b) })
	fmt.Printf("Множення: a * b = %s\n", product.ToHex())
	fmt.Printf("Час виконання : %v\n", MultiplyTime)
	fmt.Printf("\n")

	quotient := a.Divide(b)
	DivideTime := measureRepeatedTime(repeats, func() { a.Divide(b) })
	fmt.Printf("Ділення: a / b = %s,", quotient.ToHex())
	fmt.Printf("Час виконання : %v\n", DivideTime)
    fmt.Printf("\n")

	square := a.Square()
	SquareTime := measureRepeatedTime(repeats, func() { a.Square() })
	fmt.Printf("Піднесення до квадрату: a ^ 2 = %s\n", square.ToHex())
	fmt.Printf("Час виконання : %v\n", SquareTime)
    fmt.Printf("\n")

	exp := g.Power(h)
	ExpTime := measureRepeatedTime(repeats, func() { g.Power(h) })
	fmt.Printf("Піднесення багаторозрядного числа до степеня: e ^ d = %s\n", exp.ToHex())
	fmt.Printf("Час виконання : %v\n", ExpTime)
    fmt.Printf("\n")

	left := a.ShiftLeft(1)
	LeftTime := measureRepeatedTime(repeats, func() { a.ShiftLeft(1) })
	fmt.Printf("Зсув вліво = %s\n", left.ToHex())
	fmt.Printf("Час виконання : %v\n", LeftTime)
	fmt.Printf("\n")

	right := a.RightShift(1)
	RightTime := measureRepeatedTime(repeats, func() { a.RightShift(1) })
	fmt.Printf("Зсув вправо = %s\n", right.ToHex())
	fmt.Printf("Час виконання : %v\n", RightTime)
	fmt.Printf("\n")

    gcd := a.GCD(b)
    GCDTime := measureRepeatedTime(1, func() { a.GCD(b) })
    fmt.Printf("GCD(a, b) = %s\n", gcd.ToHex())
    fmt.Printf("Час виконання : %v\n", GCDTime)
    fmt.Printf("\n")

    lcm := a.LCM(b)
    LCMTime := measureRepeatedTime(1, func() { a.LCM(b) })
    fmt.Printf("НСК: LCM(a, b) = %s\n", lcm.ToHex())
    fmt.Printf("Час виконання : %v\n", LCMTime)
    fmt.Printf("\n")

	modSum := a.ModAdd(b, c) 
    ModAddTime := measureRepeatedTime(1, func() { a.ModAdd(b, c) })
    fmt.Printf("Додавання за модулем: (a + b) mod n = %s\n", modSum.ToHex())
    fmt.Printf("Час виконання : %v\n", ModAddTime)
    fmt.Printf("\n")

    modDiff := a.ModSubtract(b, c) 
    ModSubtractTime := measureRepeatedTime(1, func() { a.ModSubtract(b, c) })
    fmt.Printf("Віднімання за модулем: (a - b) mod n = %s\n", modDiff.ToHex())
    fmt.Printf("Час виконання : %v\n", ModSubtractTime)
    fmt.Printf("\n")

    modProduct := a.ModMultiply(b, c) 
    ModMultiplyTime := measureRepeatedTime(1, func() { a.ModMultiply(b, c) })
    fmt.Printf("Множення за модулем: (a * b) mod n = %s\n", modProduct.ToHex())
    fmt.Printf("Час виконання : %v\n", ModMultiplyTime)
    fmt.Printf("\n")

	ModSquareTime := measureRepeatedTime(1, func() { a.ModSquare(c) })
    modsquare := a.ModSquare(c)
    fmt.Printf("Піднесення до квадрату: (a^2) mod n = %s\n", modsquare.ToHex())
    fmt.Printf("Час виконання : %v\n", ModSquareTime)
    fmt.Printf("\n")

	ln := new(long_arithmetic.BigInt)
    KaratsubaTime := measureRepeatedTime(1, func() { ln.KaratsubaMultiply(a,b) })
	mult := ln.KaratsubaMultiply(a,b)
	fmt.Printf("Множення за Карацубою: (a * b) = %s\n", mult.ToHex())
	fmt.Printf("Час виконання: %v\n", KaratsubaTime)
	fmt.Printf("\n")

    BarrettExpTime := measureRepeatedTime(1, func() { a.LongModPowerBarrett(b, c) })
    power := a.LongModPowerBarrett(b, c)
    fmt.Printf("Піднесення до степеня за модулем (редукція Баррета): (a^exp) mod n = %s\n", power.ToHex())
    fmt.Printf("Час виконання: %v\n", BarrettExpTime)
    fmt.Printf("\n")
}