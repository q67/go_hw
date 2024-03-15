package main

import (
	"fmt"
)

var haveCode bool = false

func getKey(scenes map[sceneId]string) sceneId {
	keys := make([]sceneId, len(scenes))
	i := 0
	for k := range scenes {
		keys[i] = k
		i++
	}

	return keys[0]
}

func printScene(sceneDesc string) {
	fmt.Print("\n> ", sceneDesc, "\n\n")
}

func getWays(chooses scenesList) {
	var str string
	for i, choose := range chooses {
		key := getKey(choose)
		str += fmt.Sprint(i+1) + ") " + string(choose[key]) + "\n"
	}
	str += " \n0) " + string(Exit) + "\n\n> "

	fmt.Print(str)
}

func getSceneId(scanIndex int, scenes scenesList) sceneId {
	index := scanIndex - 1
	scenesLen := len(scenes)

	if index > scenesLen-1 {
		index = scenesLen - 1
	}

	if index < 0 {
		return Exit
	}

	scene := getKey(scenes[index])

	if scene == Flashlight {
		haveCode = true
	}

	if scene == Start {
		haveCode = false
	}

	return scene
}

func updateScenes(scene sceneId, scenes scenesIn) scenesIn {
	if haveCode && scene == Tent {
		*scenes[Tent].nextScenes = append(*scenes[Tent].nextScenes, map[sceneId]string{StrongBoxOpen: "Ввести код зі стіни печери"})
		haveCode = false
	}

	return scenes
}

func main() {
	var scenes map[sceneId]scene = InitScenes()
	var currSceneId sceneId = Start
	var scanIndex int

	for {
		printScene(scenes[currSceneId].desc)
		getWays(*scenes[currSceneId].nextScenes)
		fmt.Scanf("%d", &scanIndex)
		currSceneId = getSceneId(scanIndex, *scenes[currSceneId].nextScenes)
		if currSceneId == Exit {
			fmt.Println("Гру завершено.")
			break
		}
		scenes = updateScenes(currSceneId, scenes)
	}
}
