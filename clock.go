package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/JakeMakesStuff/color"
)

// multidimensional arrays consisting of 10 "number arrays" for each number from 0-9
// each number array has 7 strings, 1 for each line
var fonts = map[string][10][7]string{
	"block": {
		{"█████", "█   █", "█   █", "█   █", "█   █", "█   █", "█████"},
		{"███  ", "  █  ", "  █  ", "  █  ", "  █  ", "  █  ", "█████"},
		{"█████", "    █", "    █", "█████", "█    ", "█    ", "█████"},
		{"█████", "    █", "    █", "█████", "    █", "    █", "█████"},
		{"█   █", "█   █", "█   █", "█████", "    █", "    █", "    █"},
		{"█████", "█    ", "█    ", "█████", "    █", "    █", "█████"},
		{"█████", "█    ", "█    ", "█████", "█   █", "█   █", "█████"},
		{"█████", "    █", "    █", "    █", "    █", "    █", "    █"},
		{"█████", "█   █", "█   █", "█████", "█   █", "█   █", "█████"},
		{"█████", "█   █", "█   █", "█████", "    █", "    █", "█████"},
	},
	"pipe": {
		{"╔═══╗", "║   ║", "║   ║", "║   ║", "║   ║", "║   ║", "╚═══╝"},
		{"══╗  ", "  ║  ", "  ║  ", "  ║  ", "  ║  ", "  ║  ", "══╩══"},
		{"╔═══╗", "    ║", "    ║", "╔═══╝", "║    ", "║    ", "╚════"},
		{"════╗", "    ║", "    ║", "════╣", "    ║", "    ║", "════╝"},
		{"║   ║", "║   ║", "║   ║", "╚═══╣", "    ║", "    ║", "    ║"},
		{"╔════", "║    ", "║    ", "╚═══╗", "    ║", "    ║", "════╝"},
		{"╔═══╗", "║    ", "║    ", "╠═══╗", "║   ║", "║   ║", "╚═══╝"},
		{"════╗", "    ║", "    ║", "    ║", "    ║", "    ║", "    ║"},
		{"╔═══╗", "║   ║", "║   ║", "╠═══╣", "║   ║", "║   ║", "╚═══╝"},
		{"╔═══╗", "║   ║", "║   ║", "╚═══╣", "    ║", "    ║", "╚═══╝"},
	},
	"line": {
		{"┏━━━┓", "┃   ┃", "┃   ┃", "┃   ┃", "┃   ┃", "┃   ┃", "┗━━━┛"},
		{"━━┓  ", "  ┃  ", "  ┃  ", "  ┃  ", "  ┃  ", "  ┃  ", "━━┻━━"},
		{"┏━━━┓", "    ┃", "    ┃", "┏━━━┛", "┃    ", "┃    ", "┗━━━━"},
		{"━━━━┓", "    ┃", "    ┃", "━━━━┫", "    ┃", "    ┃", "━━━━┛"},
		{"╻   ╻", "┃   ┃", "┃   ┃", "┗━━━┫", "    ┃", "    ┃", "    ╹"},
		{"┏━━━━", "┃    ", "┃    ", "┗━━━┓", "    ┃", "    ┃", "━━━━┛"},
		{"┏━━━┓", "┃    ", "┃    ", "┣━━━┓", "┃   ┃", "┃   ┃", "┗━━━┛"},
		{"━━━━┓", "    ┃", "    ┃", "    ┃", "    ┃", "    ┃", "    ╹"},
		{"┏━━━┓", "┃   ┃", "┃   ┃", "┣━━━┫", "┃   ┃", "┃   ┃", "┗━━━┛"},
		{"┏━━━┓", "┃   ┃", "┃   ┃", "┗━━━┫", "    ┃", "    ┃", "┗━━━┛"},
	},
}

func clear() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear") // Linux and Mac
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	// for command line flags
	// flag name, default value, description
	foreground := flag.String("color", "white", "Foreground color of the time output (white, red, yellow, green, mint, cyan, teal, blue, purple, magenta, violet, pink, black, grey, gray)")

	// font
	font := flag.String("font", "block", "█block, ║pipe, ┃line")

	// if this flag is included (--seconds) then seconds are enabled
	seconds := flag.Bool("seconds", false, "Enable Seconds display")

	// if this flag is included (--static) then the clock does not update
	static := flag.Bool("static", false, "Makes the clock print once without updating")

	flag.Parse() // parse flags

	if _, ok := fonts[*font]; !ok { // if the font exists
		fmt.Println("Invalid font. Using default")
		*font = "block"
	}

	switch *foreground {
	case "red":
		color.Set(color.FgHiRed)
	case "yellow":
		color.Set(color.FgHiYellow)
	case "green", "mint", "manjaro":
		color.Set(color.FgHiGreen)
	case "cyan", "teal":
		color.Set(color.FgHiCyan)
	case "blue":
		color.Set(color.FgHiBlue)
	case "purple", "magenta", "violet", "pink":
		color.Set(color.FgHiMagenta)
	case "black":
		color.Set(color.FgBlack)
	case "grey", "gray":
		color.Set(color.FgHiBlack)
	default:
		color.Set(color.FgWhite)
	}

	format := "15:04" // 24h
	if *seconds {
		// if seconds flag true add seconds
		format += ":05"
	}

	printTime(format, *font) // print time instantly on app start

	if !*static {
		ticker := time.NewTicker(1 * time.Second) // run ticker once per second
		for range ticker.C {                      // every time it triggers
			printTime(format, *font) // print time
		}
	}
}

func printTime(format string, font string) {
	clear() // clear screen
	// get current time, formatted as 24h, then remove colons and
	// separate it into a slice of letters
	timeArr := strings.Split(strings.ReplaceAll(time.Now().Format(format), ":", ""), "")

	for i := 0; i < 7; i++ { // for each line
		var line string
		for j := 0; j < len(timeArr); j++ { // for each digit
			// get index j of split time and conv to string
			digit, err := strconv.Atoi(timeArr[j])
			if err != nil { // if error
				panic(err)
			}

			// add a dot to lines 3 and 5 only if it is after the second digit
			// or seconds is enabled and it is after the fourth digit
			dot := " "
			if j == 1 || (j == 3 && len(timeArr) == 6) {
				if i == 2 || i == 4 {
					dot = " █ "
				} else {
					dot = "   "
				}
			}

			// add i line for the digit plus dots if applicable
			// to the variable for the whole line
			line += fonts[font][digit][i] + dot
		}
		fmt.Println(line) // print the whole line
	}
}
