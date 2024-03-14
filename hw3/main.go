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
