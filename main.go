package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

var clear map[string]func() //create a map for storing clear funcs

func main() {
	box := string(219)
	type (
		placeholder [5]string
	)
	zero := placeholder{
		box + box + box,
		box + " " + box,
		box + " " + box,
		box + " " + box,
		box + box + box,
	}

	one := placeholder{
		" " + " " + box,
		" " + " " + box,
		" " + " " + box,
		" " + " " + box,
		" " + " " + box,
	}

	two := placeholder{
		box + box + box,
		" " + " " + box,
		" " + box + " ",
		box + " " + " ",
		box + box + box,
	}

	three := placeholder{
		box + box + box,
		" " + " " + box,
		box + box + box,
		" " + " " + box,
		box + box + box,
	}

	four := placeholder{
		box + " " + box,
		box + " " + box,
		box + box + box,
		" " + " " + box,
		" " + " " + box,
	}

	five := placeholder{
		box + box + box,
		box + " " + " ",
		" " + box + " ",
		" " + " " + box,
		box + box + box,
	}

	six := placeholder{
		box + box + box,
		box + " " + " ",
		box + box + box,
		box + " " + box,
		box + box + box,
	}

	seven := placeholder{
		box + box + box,
		box + " " + box,
		" " + " " + box,
		" " + " " + box,
		" " + " " + box,
	}

	eight := placeholder{
		box + box + box,
		box + " " + box,
		box + box + box,
		box + " " + box,
		box + box + box,
	}

	nine := placeholder{
		box + box + box,
		box + " " + box,
		box + box + box,
		" " + " " + box,
		" " + " " + box,
	}

	dot := placeholder{
		" " + " " + " ",
		" " + box + " ",
		" " + " " + " ",
		" " + box + " ",
		" " + " " + " ",
	}

	blankDot := placeholder{
		" " + " " + " ",
		" " + " " + " ",
		" " + " " + " ",
		" " + " " + " ",
		" " + " " + " ",
	}

	digits := [...][5]string{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
		dot,
		blankDot,
	}
	for {
		screen.MoveTopLeft()
		time.Sleep(1 * time.Second)
		now := time.Now()
		hour, minute, second := now.Hour(), now.Minute(), now.Second()
		clock := [10]placeholder{
			digits[hour/10], digits[hour%10],
			dot,
			digits[minute/10], digits[minute%10],
			dot,
			digits[second/10], digits[second%10],
		}

		for line := range zero {
			for idx, digit := range clock {
				next := clock[idx][line]
				if digit == dot && second%2 == 1 {
					next = "   "
				}
				fmt.Print(next, " ")
			}
			fmt.Println()
		}
		clock[2] = blankDot
		clock[5] = blankDot
	}
}
