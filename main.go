package main

import (
	"flag"
	"fmt"

	"github.com/attendeee/adventure/utils"
)

func main() {
	path := flag.String("path", "", "Path to story")
	flag.Parse()

	story := utils.MustGetStoryFromFile(*path)

	// fmt.Println("Contents of file", string(file))

	defer fmt.Println("adventure")

	utils.StoryLoop(story)
}
