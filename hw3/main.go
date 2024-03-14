package main

import "fmt"

func printScene(sc string) {
	fmt.Print("\n> ", sc, "\n\n")
}

func getWays(chooses []sceneId) {
	var str string
	for i, line := range chooses {
		str += fmt.Sprint(i+1) + ") " + string(line) + "\n"
	}
	str += " \n0) " + string(Exit) + "\n\n> "

	fmt.Print(str)
}

func getSceneId(scanIndex int, scenes []sceneId) sceneId {
	index := scanIndex - 1
	scenesLen := len(scenes)

	if index > scenesLen {
		index = scenesLen
	}

	if index < 0 {
		return Exit
	}

	return scenes[index]
}

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
	var scenes map[sceneId]scene = InitScenes()
	var currSceneId sceneId = Start
	var scanIndex int

	for {
		printScene(scenes[currSceneId].desc)
		getWays(scenes[currSceneId].nextScenes)
		fmt.Scanf("%d", &scanIndex)
		currSceneId = getSceneId(scanIndex, scenes[currSceneId].nextScenes)
		if currSceneId == Exit {
			fmt.Println("Гру завершено.")
			break
		}

	}
}
