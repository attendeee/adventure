package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/attendeee/adventure/story"
)

func printArc(a story.Arc) {
	fmt.Printf("Title: %s\n", a.Title)

	fmt.Printf("Story: \n")
	for _, s := range a.Story {
		fmt.Printf("\t%s\n\n", s)
	}

	fmt.Printf("Options: \n")
	for i, o := range a.Options {
		fmt.Printf("%d %s\n", i, o.Text)
	}

}

func clearTerminal() {
	var cmd *exec.Cmd

	cmd = exec.Command("clear")

	cmd.Stdout = os.Stdout

	cmd.Run()
}

func StoryLoop(story story.Story) {

	a := story["intro"]

	for {
		clearTerminal()

		printArc(a)
		if len(a.Options) == 0 {
			os.Exit(0)
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
