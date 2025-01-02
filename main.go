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

	a := long_arithmetic.ParseHex("fed11a84df48288672aed240ccecd3eb9b0748135c8047a1900b66552efd4a958e1061959b3a5e7f487957513e4d067eb38c15db62ade93d3c273ef4bff5d2c9f77e477848ade9bf4ce0237ccc134bc8b924bd0345e8ebee82a9626a2b7d8bb875aee46b30f28361704becc061ad818229ef2411010fa12ab69db488846159c3072bcac4509a29b52e6e9627b93ea3f3ac8c1102dd0c0dda21d12c99dc094d0a0549edaae03b0e5eeef4a0d24d7e457c336ec01ec00362cf4d99e57e261cf1978f13052f4d2d67a827a728652c6892b2295c3752cf70fd75171000ffef654a6e507434baa613e0089237a7a87d60f70f7f93adb4ad3b3d565971a3c4025a7cd0")
	b := long_arithmetic.ParseHex("7762eca0d63bf84694ff81b0ef356bd89760377107b18ec7d116dba9ec8771fe363dcc04dfca3a735a90ecd641c883fa7b72e1aaeb5aab15a35d5b5b4e5ef24050dc43a8c1a8133d70c45ce9ca40c03b71fd0cd90faa508968e05cccd0cc7ad1fea6b891cb65954851f6ad61e001623ab53336e082c7670b056a65017caff175f44a7658f34d6daf76dc26095a857f0d694b3e945c001894ebb2f78f9021d24fda50b10e2ca006bfc5e48d8893f363766f8af250f85cc6f676dbedd04e1c186668e15bc9bcf77da752c0f795d2afa3c140dd0ee996255ecc5f3443945b10fa519d77860f995f3dac92787ee243aa9f93b7c35eabca713481275d851353d99f88")
	c := long_arithmetic.ParseHex("d3a518e47135c29e70e5295c2e14e4dd0e1e9329c6788b6422b3e9b2e933e8f480fa9ebac93f72eea89b659ce5e1dee1d1626ef0d17449edf87f2d647720fd47cb811c1087ab4d4374d6e7883a46b4e1c49aa05b44863d43c66d63e2d9160ed3dbec0e1e27d8d9e3b69a65d39eeb0675e45e3e67a45d503eab92bcbc7f64f671")
	d := long_arithmetic.ParseHex("100")
	e := long_arithmetic.ParseHex("d1f")

	
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


	exp := e.Power(d)
	ExpTime := measureRepeatedTime(repeats, func() { e.Power(d) })
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

	modSum := a.ModAdd(b, c) 
    ModAddTime := measureRepeatedTime(1, func() { a.ModAdd(b, c) })
    fmt.Printf("Додавання за модулем: (a + b) mod n = %s\n", modSum.ToHex())
    fmt.Printf("Час виконання : %v\n", ModAddTime)
    fmt.Printf("\n")

    // Віднімання за модулем
    modDiff := a.ModSubtract(b, c) 
    ModSubtractTime := measureRepeatedTime(1, func() { a.ModSubtract(b, c) })
    fmt.Printf("Віднімання за модулем: (a - b) mod n = %s\n", modDiff.ToHex())
    fmt.Printf("Час виконання : %v\n", ModSubtractTime)
    fmt.Printf("\n")

    // Множення за модулем
    modProduct := a.ModMultiply(b, c) 
    ModMultiplyTime := measureRepeatedTime(1, func() { a.ModMultiply(b, c) })
    fmt.Printf("Множення за модулем: (a * b) mod n = %s\n", modProduct.ToHex())
    fmt.Printf("Час виконання : %v\n", ModMultiplyTime)
    fmt.Printf("\n")

    // Найбільший спільний дільник (НСД)
    gcd := a.GCD(b)
    GCDTime := measureRepeatedTime(1, func() { a.GCD(b) })
    fmt.Printf("GCD(a, b) = %s\n", gcd.ToHex())
    fmt.Printf("Час виконання : %v\n", GCDTime)
    fmt.Printf("\n")

    // Найменше спільне кратне (НСК)
    lcm := a.LCM(b)
    LCMTime := measureRepeatedTime(1, func() { a.LCM(b) })
    fmt.Printf("НСК: LCM(a, b) = %s\n", lcm.ToHex())
    fmt.Printf("Час виконання : %v\n", LCMTime)
    fmt.Printf("\n")

	ModSquareTime := measureRepeatedTime(1, func() { a.ModSquare(c) })
    modsquare := a.ModSquare(c)
    fmt.Printf("Піднесення до квадрату: (a^2) mod n = %s\n", modsquare.ToHex())
    fmt.Printf("Час виконання : %v\n", ModSquareTime)
    fmt.Printf("\n")


    BarrettExpTime := measureRepeatedTime(1, func() { e.LongModPowerBarrett(d, c) })
    power := e.LongModPowerBarrett(d, c)
    fmt.Printf("Піднесення до степеня за модулем (редукція Баррета): (a^exp) mod n = %s\n", power.ToHex())
    fmt.Printf("Час виконання : %v\n", BarrettExpTime)
    fmt.Printf("\n")
}

