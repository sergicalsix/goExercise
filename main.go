package main

import (
	"errors"
	"fmt"
	"myproject/mylib"
	"strconv"
)

func printPart(d int) {
	fmt.Println("###")
	fmt.Println("Part", d)
	fmt.Println("###")
}

// 2-4
func checkMapKey(m map[string]int, key string) {
	if val, ok := m[key]; ok {
		fmt.Printf("%sは%d円です\n", key, val)
	} else {
		fmt.Printf("%sはありません\n", key)
	}
}

// 3-1 FizzBuzz
func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// 3-2
func fizzBuzzSwitch() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

}

// 3-3
func judgeEvenOdd(nums []int) {
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%d: even\n", num)
		} else {
			fmt.Printf("%d: odd\n", num)
		}
	}

}

// 4-1
type Person struct {
	Name string
	Sex  string
}

func (p *Person) greet() {
	if p.Sex == "man" {
		fmt.Printf("Hi,Mr.%s\n", p.Name)
	} else {
		fmt.Println("男性ではありませんn")
	}
}

// 4-2
type Greeter interface {
	greet()
}

// 5-1
func divideNumPanic(a, b int) int {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}

	}()

	if b == 0 {
		panic("division by zero")
	} else {
		return (a / b)
	}
}

// 5-2
func divideNum(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("error: division by zero")
	} else {
		return a / b, nil
	}
}

func main() {

	printPart(1)
	// answer
	fmt.Println("Hello World")
	mylib.GoodEvening()

	printPart(2)

	var (
		intValue    int     = 1
		stringValue string  = "hoge"
		floatValue  float64 = 1.0
		boolValue   bool    = true
		piFloat     float64 = 3.14
		piString    string  = "3.14"
	)

	// 2-2
	convertString := strconv.FormatFloat(piFloat, 'f', -1, 64)
	convertFloat, err := strconv.ParseFloat(piString, 64)
	if err != nil {
		return
	}

	// 2-3&3-3
	var slice []int = []int{1, 3, 66, 12}

	// 2-4
	var Shop = map[string]int{"apple": 500, "banana": 800, "orange": 700}

	// answer
	fmt.Println(intValue, stringValue, fmt.Sprintf("%f", floatValue), boolValue)
	fmt.Println("flaot -> str", convertString, "str -> flaot", convertFloat)
	fmt.Println(slice)
	checkMapKey(Shop, "apple")
	checkMapKey(Shop, "lemmon")

	printPart(3)

	// answer
	fizzBuzz()
	fizzBuzzSwitch()
	judgeEvenOdd(slice)

	printPart(4)

	p := new(Person)
	p.Name = "shibuya"
	p.Sex = "man"
	var g Greeter = &Person{Name: "shobuya", Sex: "man"}

	// answer
	p.greet()
	g.greet()

	printPart(5)

	fmt.Println("5/2")
	fmt.Println(divideNumPanic(5, 2))

	fmt.Println("5/0")
	fmt.Println(divideNumPanic(5, 0))

	fmt.Println("5/2")
	result3, err3 := divideNum(5, 2)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(result3)
	}

	fmt.Println("5/0")
	result4, err4 := divideNum(5, 0)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(result4)
	}

}
