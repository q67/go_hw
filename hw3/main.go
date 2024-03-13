package main

import "fmt"

func main() {
	// game start
	// promt name of person
	// start first scene

	// infinity loop move from scene to scene

	// every scene: descripe place, list of chooses
	// (every choose contain id of next scene)

	// if finish then only one choose is start of game

	// for {
	// 	var name string
	// 	var age int
	// 	fmt.Println("Введите имя и возраст:")
	// 	fmt.Scanf("%s %d", &name, &age)
	// 	fmt.Println(name, age)
	// }
	InitScenes()

	var scene1 = *Scenes[scene1wake]

	fmt.Println(scene1)
}
