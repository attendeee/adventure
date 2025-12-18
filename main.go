package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type Story map[string]Arc

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func printArc(a Arc) {
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

func main() {
	path := flag.String("path", "", "Path to story")
	flag.Parse()

	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	var story Story

	json.Unmarshal(file, &story)

	// fmt.Println("Contents of file", string(file))

	defer fmt.Println("adventure")

	a := story["intro"]
	for {
		clearTerminal()

		printArc(a)
		if len(a.Options) == 0 {
			os.Exit(0)
		}

		fmt.Printf("Choose arc: ")

		var c int

		for {

			fmt.Scanf("%d", &c)

			if c < 0 || c >= len(a.Options) {
				continue
			} else {
				break
			}
		}

		a = story[a.Options[c].Arc]

	}

}
