package utils

import (
	"fmt"
	"image"
	"io/fs"
	"log"
	"os"
)

func init() {
	fmt.Println("Loaded detect.go")
}

// Open file and print type to console
func Detect(fileName fs.FileInfo) {

	// Open file
	file, err := os.Open("./presort/" + fileName.Name())
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

	fmt.Println(fileType)
}
