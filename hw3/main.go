package main

import (
	"fmt"
)

type hero struct {
	name     string
	haveCode bool
}

var Hero hero

func printScene(sceneDesc string) {
	fmt.Print("\n> ", sceneDesc, "\n\n")
}

func printWays(chooses []nextScene) {
	var str string
	for i, choose := range chooses {
		str += fmt.Sprint(i+1) + ") " + string(choose.action) + "\n"
	}
	str += " \n0) " + string(Exit) + "\n\n> "

	fmt.Print(str)
}

func getSceneId(scanIndex int, scenes []nextScene) sceneId {
	index := scanIndex - 1
	scenesLen := len(scenes)

	if index > scenesLen-1 {
		index = scenesLen - 1
	}

	if index < 0 {
		return Exit
	}

	scene := scenes[index].nextSceneId

	return scene
}

func itemsModify(scene sceneId) {
	if scene == Flashlight {
		Hero.haveCode = true
	}

	if scene == Start {
		Hero.haveCode = false
	}
}

func updateScenes(scene sceneId, scenes scenesIn) scenesIn {
	if Hero.haveCode && scene == Tent {
		*scenes[Tent].nextScenes = append(*scenes[Tent].nextScenes, nextScene{StrongBoxOpen, "Ввести код зі стіни печери"})
		Hero.haveCode = false
	}

	return scenes
}

func main() {
	Hero.name = "Стівен"
	Hero.haveCode = false
	scenes := InitScenes()
	currSceneId := Start
	var scanIndex int

	for {
		printScene(scenes[currSceneId].desc)
		printWays(*scenes[currSceneId].nextScenes)
		fmt.Scanf("%d", &scanIndex)
		currSceneId = getSceneId(scanIndex, *scenes[currSceneId].nextScenes)
		if currSceneId == Exit {
			fmt.Printf("Гру завершено, бувай, %s.\n", Hero.name)
			break
		}
		itemsModify(currSceneId)
		scenes = updateScenes(currSceneId, scenes)
	}
}
