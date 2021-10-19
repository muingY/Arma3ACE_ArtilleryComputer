package interactCore

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

/*
* Variables
*/
var clear map[string]func()
const ScreenWidth int = 28

/*
* Constructor
*/
func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

/*
* Functions
*/
func ClearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
		value()  //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func DisplayScene(displayScene *Scene) {
	ClearScreen()
	fmt.Println("________________________________________")
	fmt.Println("    | " + displayScene.DisplayBuffer[0][:ScreenWidth] + " |    ")
	fmt.Println("[Q] | " + displayScene.DisplayBuffer[1][:ScreenWidth] + " | [I]")
	fmt.Println("[A] | " + displayScene.DisplayBuffer[2][:ScreenWidth] + " | [J]")
	fmt.Println("[Z] | " + displayScene.DisplayBuffer[3][:ScreenWidth] + " | [N]")
	fmt.Println("[W] | " + displayScene.DisplayBuffer[4][:ScreenWidth] + " | [O]")
	fmt.Println("[S] | " + displayScene.DisplayBuffer[5][:ScreenWidth] + " | [K]")
	fmt.Println("[X] | " + displayScene.DisplayBuffer[6][:ScreenWidth] + " | [M]")
	fmt.Println("________________________________________")
}

// InsertStringForScene mode = 0 is left, = 1 is right
func InsertStringForScene(str *string, insertStr string, mode int, pos int) {
	var temp string
	temp = *str
	if len(*str) != 28 {
		if len(*str) > 28 {
			temp = temp[:28]
		} else {
			for i := len(*str); i < 28; i++ {
				temp += " "
			}
		}
	}
	if mode == 0 {
		temp = temp[:pos] + insertStr + temp[pos+1:]
	} else if mode == 1 {
		pos = ScreenWidth - (pos + len(insertStr))
		temp = temp[:pos] + insertStr + temp[pos+1:]
	}
	*str = temp
}
func SetEmptyStrForScene(str *string) {
	*str = "                            "
}
