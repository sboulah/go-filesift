package main

import (
	"fmt"
	"image"
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

		// Open file
		file, err := os.Open("./presort/" + i.Name())
		if err != nil {
			log.Fatalf("Error opening file: %s\n", err)
			os.Exit(1)
		}

		defer file.Close()

		// Decode File & Get Type
		_, fileType, err := image.Decode(file)
		if err != nil {
			log.Fatalf("Error decoding file: %s\n", err)
			os.Exit(1)
		}

		switch fileType {
		case "png":
			fmt.Println("PNG")
			utils.DecodePNG(file.Name())
		case "jpg":
			fmt.Println("JPG")
			utils.DecodeJPG(file.Name())
		case "jpeg":
			fmt.Println("JPEG")
			utils.DecodeJPEG(file.Name())
		case "webp":
			fmt.Println("WEBP")
			utils.DecodeWEBP(file.Name())
		}

	}
}
