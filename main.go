package main

import (
	InputEvent "ArtilleryComputer/inputEvent"
	"ArtilleryComputer/interactCore"
	"os"
)

const (
	firstMenuScene = iota
	infoScene
	fcuM109A6Scene
	fireRecordScene
)
func initializeScene(sceneList *[]interactCore.Scene, len int) {
	var newScene interactCore.Scene

	// firstMenuScene
	newScene.DisplayBuffer[0] = "    [ARTILLERY COMPUTER]    "
	newScene.DisplayBuffer[1] = "----------------------------"
	newScene.DisplayBuffer[2] = "                            "
	newScene.DisplayBuffer[3] = ">M109A6 FCU                 "
	newScene.DisplayBuffer[4] = "                            "
	newScene.DisplayBuffer[5] = ">FIRE RECORD                "
	newScene.DisplayBuffer[6] = ">INFOMATION                 "
	interactCore.InsertStringForScene(&newScene.DisplayBuffer[6], "EXIT<", 1, 0)
	newScene.SceneControlFunc = firstMenuSceneManage
	*sceneList = append(*sceneList, newScene)

	// infoScene
	newScene.DisplayBuffer[0] = "         INFOMATION         "
	newScene.DisplayBuffer[1] = "----------------------------"
	newScene.DisplayBuffer[2] = "VERS : 1.0                  "
	newScene.DisplayBuffer[3] = "LANG : Go 1.17.2            "
	newScene.DisplayBuffer[4] = "DEV  : MUING                "
	newScene.DisplayBuffer[5] = "       (ungetqqq@gmail.com) "
	newScene.DisplayBuffer[6] = "<BACK                       "
	newScene.SceneControlFunc = infoSceneManage
	*sceneList = append(*sceneList, newScene)
}

func firstMenuSceneManage() int {
	input := InputEvent.GetKeyEvent()

	switch input {
	case InputEvent.KEY_Button_l2:
		return fcuM109A6Scene
	case InputEvent.KEY_Button_l4:
		return fireRecordScene
	case InputEvent.KEY_Button_l5:
		return infoScene
	case InputEvent.KEY_Button_r5:
		os.Exit(0)
	default:
		return firstMenuScene
	}
	return -1
}
func infoSceneManage() int {
	input := InputEvent.GetKeyEvent()

	switch input {
	case InputEvent.KEY_Button_l5:
		return firstMenuScene
	default:
		return infoScene
	}
	return -1
}

func main() {
	var sceneList []interactCore.Scene
	var currentScene int

	initializeScene(&sceneList, 5)
	currentScene = firstMenuScene
	interactCore.DisplayScene(&sceneList[currentScene])

	for {
		interactCore.DisplayScene(&sceneList[currentScene])
		currentScene = sceneList[currentScene].SceneControlFunc()
	}
}
