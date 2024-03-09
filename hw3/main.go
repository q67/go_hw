package main

import "fmt"

func main() {
	for {
		var name string
		var age int
		fmt.Println("Введите имя и возраст:")
		fmt.Scanf("%s %d", &name, &age)
		fmt.Println(name, age)
	}
}
