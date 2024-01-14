package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

const (
	gridWidth   = 3
	gridHeight  = 9
	imageWidth  = 400
	imageHeight = 225
)

func main() {
	// Create a new blank image for the grid
	grid := image.NewRGBA(image.Rect(0, 0, gridWidth*imageWidth, gridHeight*imageHeight))

	// Iterate over the images in the "frames" folder
	for i := 1; i <= gridWidth*gridHeight; i++ {
		fileName := fmt.Sprintf("frames/%04d.jpeg", i)
		img, err := loadImage(fileName)
		if err != nil {
			fmt.Printf("Error loading image %s: %v\n", fileName, err)
			continue
		}

		// Calculate the position for the current image
		col := (i - 1) % gridWidth
		row := (i - 1) / gridWidth

		// Draw the current image onto the grid
		draw.Draw(grid, image.Rect(col*imageWidth, row*imageHeight, (col+1)*imageWidth, (row+1)*imageHeight), img, image.Point{0, 0}, draw.Over)
	}

	// Save the grid image to a file
	saveFileName := "grid.jpeg"
	saveGridImage(grid, saveFileName)
	fmt.Printf("Grid image saved to %s\n", saveFileName)
}

func loadImage(fileName string) (image.Image, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func saveGridImage(img image.Image, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		fmt.Printf("Error encoding image to %s: %v\n", fileName, err)
	}
}
