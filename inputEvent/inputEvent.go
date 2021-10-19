package InputEvent

import (
	"github.com/mattn/go-tty"
	"log"
)

func GetKeyEvent() rune {
	var ret rune

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if r != 0 {
			ret = r
			break
		}
	}

	return ret
}
