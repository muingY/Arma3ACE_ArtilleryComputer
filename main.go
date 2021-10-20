package main

import (
	InputEvent "ArtilleryComputer/inputEvent"
	"ArtilleryComputer/interactCore"
	"os"
	"strconv"
)

const (
	firstMenuScene = iota
	infoScene
	fcuM109A6Scene
	fcuM109A6ResultScene
)
func initializeScene(sceneList *[]interactCore.Scene, len int) {
	var newScene interactCore.Scene

	// firstMenuScene
	newScene.DisplayBuffer[0] = "    [ARTILLERY COMPUTER]    "
	newScene.DisplayBuffer[1] = "----------------------------"
	newScene.DisplayBuffer[2] = "                            "
	newScene.DisplayBuffer[3] = ">M109A6 FCU                 "
	newScene.DisplayBuffer[4] = "                            "
	newScene.DisplayBuffer[5] = ">INFOMATION                 "
	newScene.DisplayBuffer[6] = "                            "
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

	// fcuM109A6Scene
	newScene.DisplayBuffer[0] = "  M109A6 FIRE CONTROL UNIT  "
	newScene.DisplayBuffer[1] = "----------------------------"
	newScene.DisplayBuffer[2] = "DISTANCE   :                "
	newScene.DisplayBuffer[3] = "MY ALT     :                "
	newScene.DisplayBuffer[4] = "TARGET ALT :                "
	newScene.DisplayBuffer[5] = "                            "
	newScene.DisplayBuffer[6] = "<BACK                  NEXT>"
	newScene.SceneControlFunc = fcuM109A6SceneManage
	*sceneList = append(*sceneList, newScene)

	// fceResultScene
	newScene.DisplayBuffer[0] = "           RESULT           "
	newScene.DisplayBuffer[1] = "----------------------------"
	newScene.DisplayBuffer[2] = "MILL LOW   :                "
	newScene.DisplayBuffer[3] = "MILL HIGH  :                "
	newScene.DisplayBuffer[4] = "                            "
	newScene.DisplayBuffer[5] = "                            "
	newScene.DisplayBuffer[6] = "<BACK                       "
	newScene.SceneControlFunc = fcuM109A6ResultSceneManage
	*sceneList = append(*sceneList, newScene)
}

func firstMenuSceneManage(sceneData *interactCore.Scene) int {
	input := InputEvent.GetKeyEvent()

	switch input {
	case InputEvent.KEY_Button_l2:
		return fcuM109A6Scene
	case InputEvent.KEY_Button_l4:
		return infoScene
	case InputEvent.KEY_Button_r5:
		os.Exit(0)
	default:
		return firstMenuScene
	}
	return -1
}
func infoSceneManage(sceneData *interactCore.Scene) int {
	input := InputEvent.GetKeyEvent()

	switch input {
	case InputEvent.KEY_Button_l5:
		return firstMenuScene
	default:
		return infoScene
	}
	return -1
}
var fcuDistance, fcuMyAlt, fcuTargetAlt float64
func fcuM109A6SceneManageDisplay(sceneData *interactCore.Scene, mode int) {
	interactCore.SetEmptyStrForScene(&sceneData.DisplayBuffer[2])
	interactCore.SetEmptyStrForScene(&sceneData.DisplayBuffer[3])
	interactCore.SetEmptyStrForScene(&sceneData.DisplayBuffer[4])
	interactCore.InsertStringForScene(&sceneData.DisplayBuffer[2], "DISTANCE   :", 0, 0)
	interactCore.InsertStringForScene(&sceneData.DisplayBuffer[2], strconv.FormatFloat(fcuDistance, 'f', -1, 64)+"m", 1, 0)
	interactCore.InsertStringForScene(&sceneData.DisplayBuffer[3], "MY ALT     :", 0, 0)
	interactCore.InsertStringForScene(&sceneData.DisplayBuffer[3], strconv.FormatFloat(fcuMyAlt, 'f', -1, 64) + "m", 1, 0)
	interactCore.InsertStringForScene(&sceneData.DisplayBuffer[4], "TARGET ALT :", 0, 0)
	interactCore.InsertStringForScene(&sceneData.DisplayBuffer[4], strconv.FormatFloat(fcuTargetAlt, 'f', -1, 64) + "m", 1, 0)
	if mode == 0 {
		interactCore.SetEmptyStrForScene(&sceneData.DisplayBuffer[2])
		interactCore.InsertStringForScene(&sceneData.DisplayBuffer[2], "DISTANCE   :", 0, 0)
		interactCore.InsertStringForScene(&sceneData.DisplayBuffer[2], strconv.FormatFloat(fcuDistance, 'f', -1, 64)+"m<", 1, 0)
	} else if mode == 1 {
		interactCore.SetEmptyStrForScene(&sceneData.DisplayBuffer[3])
		interactCore.InsertStringForScene(&sceneData.DisplayBuffer[3], "MY ALT     :", 0, 0)
		interactCore.InsertStringForScene(&sceneData.DisplayBuffer[3], strconv.FormatFloat(fcuMyAlt, 'f', -1, 64)+"m<", 1, 0)
	} else if mode == 2 {
		interactCore.SetEmptyStrForScene(&sceneData.DisplayBuffer[4])
		interactCore.InsertStringForScene(&sceneData.DisplayBuffer[4], "TARGET ALT :", 0, 0)
		interactCore.InsertStringForScene(&sceneData.DisplayBuffer[4], strconv.FormatFloat(fcuTargetAlt, 'f', -1, 64)+"m<", 1, 0)
	}
	interactCore.DisplayScene(sceneData)
}
func fcuM109A6SceneManage(sceneData *interactCore.Scene) int {
	inputMode := -1
	fcuM109A6SceneManageDisplay(sceneData, inputMode)

	input := InputEvent.GetKeyEvent()
	for {
		switch input {
		case InputEvent.KEY_Button_r1: // input distance
			inputMode = 0
			fcuM109A6SceneManageDisplay(sceneData, inputMode)
		case InputEvent.KEY_Button_r2: // input my alt
			inputMode = 1
			fcuM109A6SceneManageDisplay(sceneData, inputMode)
		case InputEvent.KEY_Button_r3: // input target alt
			inputMode = 2
			fcuM109A6SceneManageDisplay(sceneData, inputMode)
		case InputEvent.KEY_Button_l5:
			return firstMenuScene
		case InputEvent.KEY_Button_r5:
			return fcuM109A6ResultScene
		}
		input = InputEvent.GetKeyEvent()
		if (input >= InputEvent.KEY_Num_0 && input <= InputEvent.KEY_Num_9) || input == InputEvent.KEY_Backspace {
			switch inputMode {
			case 0:
				if input == InputEvent.KEY_Backspace {
					fcuDistance = float64(int(fcuDistance / 10))
				} else {
					fcuDistance *= 10
					fcuDistance += float64(input - InputEvent.KEY_Num_0)
				}
			case 1:
				if input == InputEvent.KEY_Backspace {
					fcuMyAlt = float64(int(fcuMyAlt / 10))
				} else {
					fcuMyAlt *= 10
					fcuMyAlt += float64(input - InputEvent.KEY_Num_0)
				}
			case 2:
				if input == InputEvent.KEY_Backspace {
					fcuTargetAlt = float64(int(fcuTargetAlt / 10))
				} else {
					fcuTargetAlt *= 10
					fcuTargetAlt += float64(input - InputEvent.KEY_Num_0)
				}
			}
			fcuM109A6SceneManageDisplay(sceneData, inputMode)
		}
	}
	return -1
}
func fcuM109A6ResultSceneManage(sceneData *interactCore.Scene) int {
	//...

	input := InputEvent.GetKeyEvent()

	switch input {
	case InputEvent.KEY_Button_l5:
		return fcuM109A6Scene
	default:
		return fcuM109A6ResultScene
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
		currentScene = sceneList[currentScene].SceneControlFunc(&sceneList[currentScene])
	}
}
