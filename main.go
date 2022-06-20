package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sboulah/go-filesift/utils"

	_ "golang.org/x/image/webp"
)

// Main
func main() {

	// Presort Directory
	err := os.Mkdir("./presort", 0755)
	if err != nil {
		fmt.Println("Directory already exists!")
	}

	// Get Files
	fmt.Print("Place all files into the \"presort\" direcotry! Then press enter to continue...")
	fmt.Scanln()

	// Check Directory
	preDir, err := ioutil.ReadDir("./presort")
	if err != nil {
		log.Fatalf("Error opening directory: %s", err)
	}

	// Sorted Directory
	err = os.Mkdir("./sorted", 0755)
	if err != nil {
		fmt.Println("Directory already exists!")
	}

	// Loop through files
	for _, i := range preDir {
		utils.Detect(i)
	}
}
