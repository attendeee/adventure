package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/attendeee/adventure/story"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Cyan = "\033[36m"

func printArc(a story.Arc) {
	fmt.Printf("%sTitle:%s %s\n", Red, Reset, a.Title)

	fmt.Printf("%sStory:%s \n", Red, Reset)
	for _, s := range a.Story {
		fmt.Printf("\t%s\n\n", s)
	}

	fmt.Printf("%sOptions:%s \n", Cyan, Reset)
	for i, o := range a.Options {
		fmt.Printf("%s%d%s %s\n", Cyan, i, Reset, o.Text)
	}

}

func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	cmd.Run()
}

func StoryLoop(story story.Story) {

	a := story["introduction"]

	for {
		clearTerminal()

		printArc(a)

		if len(a.Options) == 0 {
			return
		}

		var c int

		for {

			fmt.Printf("Choose arc: ")

			fmt.Scanf("%d", &c)

			if c < 0 || c >= len(a.Options) {
				fmt.Println("Invalid value is provided")
				continue
			} else {
				break
			}
		}

		a = story[a.Options[c].Arc]

	}
}

func MustGetStoryFromFile(path string) story.Story {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var story story.Story

	json.Unmarshal(file, &story)

	return story

}
