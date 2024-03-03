package main

import "fmt"

const (
	date      int    = 1995
	txt1      string = "txttxttxt"
	loopLimit int    = 5
)

type profession string

func (p profession) String() string {
	return fmt.Sprintf("message on String method, %s", string(p))
}

func probes() {

	var (
		digit1    int
		digit2    int
		somebool1 bool
	)

	fmt.Printf("Digit values: %d %d, bool value %t\n", digit1, digit2, somebool1)
	fmt.Println("hello who are you?")

	for i := 0; i < loopLimit; i++ {
		fmt.Printf("odd loop iteration %d, some calc %d\n", i, 1+i*i)
	}

	num1 := float32(-11)
	num2 := float32(2.3)
	var num3 uint32 = 3
	var num4 byte = 3
	var num5 uint8 = 44
	num6 := num4 == num5 // diff types no problem
	num7 := num1 / num2  // must be one type

	prof1 := profession("driver")
	fmt.Printf("\"%%s\" expect string but: %s\n", prof1)
	fmt.Printf("\"%%T\" is type of value: %T <-- and --> %T\n", prof1, txt1)
	fmt.Printf("expect bool: %t\n", num6)
	fmt.Printf("result must be uint type: %d\n", num3-4)
	fmt.Printf("division result %.3f\n", num7)
	fmt.Printf("default format 'v': %v %v\n", num7, txt1)

}
