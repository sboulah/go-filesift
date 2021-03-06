package utils

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/webp"
)

func init() {
	fmt.Println("Loaded decode.go")
}

func DecodeWEBP(fileName string) {

	// Open File
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening WEBP file: %s\n", err)
		os.Exit(1)
	}

	// Close File
	defer file.Close()

	// Decode File
	data, err := webp.Decode(file)
	if err != nil {
		log.Fatalf("Error decoding WEBP file: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(data.ColorModel())

	// Get Hash

}

func DecodePNG(fileName string) {

	// Open File
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening PNG file: %s\n", err)
		os.Exit(1)
	}

	// Close File
	defer file.Close()

	// Decode File
	data, err := png.Decode(file)
	if err != nil {
		log.Fatalf("Error decoding PNG file: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(data.ColorModel())
}

func DecodeJPEG(fileName string) {

	// Open File
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening JPEG file: %s\n", err)
		os.Exit(1)
	}

	// Close File
	defer file.Close()

	// Decode File
	data, err := jpeg.Decode(file)
	if err != nil {
		log.Fatalf("Error decoding JPEG file: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(data.ColorModel())
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

func DecodeWEBM(fileName string) {

	// Open File
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening WEBM file: %s\n", err)
		os.Exit(1)
	}

	// Close File
	defer file.Close()
}

func DecodeMOV(fileName string) {

	// Open File
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening MOV file: %s\n", err)
		os.Exit(1)
	}

	// Close File
	defer file.Close()
}

func DecodeMP4(fileName string) {

	// Open File
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening MP4 file: %s\n", err)
		os.Exit(1)
	}

	// Close File
	defer file.Close()
}
