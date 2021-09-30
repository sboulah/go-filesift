package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

// Main
func main() {

	// Create a slice to store the file list
	var fileList []string

	// Read the directory and pass in a pointer to the slice
	readCurrentDir(&fileList)

	// for _, i := range fileList {
	// 	fmt.Println(i)
	// }

	// Loop through files
	for _, i := range fileList {

		// Open file
		f, err := os.Open(fmt.Sprintf("./presort/%s", i))
		fmt.Println(i)

		// Error Checking
		if err != nil {
			log.Fatalf("Error opening file: %s", err)
		}

		// Decode file
		x, err := exif.Decode(f)

		// Error Checking
		if err != nil {
			log.Fatalf("Error decoding file: %s", err)
		}

		fileMake, _ := x.Get("Make")

		fmt.Println(fileMake)
	}

}

// Read Directory
func readCurrentDir(fileList *[]string) {

	// Open Directory
	file, err := os.Open("./presort")

	// Error Checking
	if err != nil {
		log.Fatalf("Failed to open directory: %s\n", err)
	}

	// Defer the closing of the directory
	defer file.Close()

	// Read all contents in the directory
	*fileList, err = file.Readdirnames(0)

	// Error Checking
	if err != nil {
		log.Fatalf("Error reading contents: %s", err)
	}
}
