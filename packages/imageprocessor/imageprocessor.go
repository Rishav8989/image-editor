// imageprocessor/imageprocessor.go
package imageprocessor

import (
	"fmt"

	"github.com/disintegration/imaging"
)

// ApplyBlur applies blur effect to the image with the specified blur level
func ApplyBlur(srcImagePath, outputImagePath string, blurLevel float64) error {
	srcImage, err := imaging.Open(srcImagePath)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}

	// Apply the blur effect with the specified blur level
	dstImage := imaging.Blur(srcImage, blurLevel)

	// Save the resulting image to a new file
	if err := imaging.Save(dstImage, outputImagePath); err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	return nil
}

// ResizeTo resizes the image to the specified width and height using the Lanczos filter.
func ResizeTo(srcImagePath, outputImagePath string, width, height int) error {
	srcImage, err := imaging.Open(srcImagePath)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}

	// Resize the image to the specified width and height using Lanczos filter
	dstImage := imaging.Resize(srcImage, width, height, imaging.Lanczos)

	// Save the resulting image to a new file
	if err := imaging.Save(dstImage, outputImagePath); err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	return nil
}

// ScaleToFit scales down the image to fit the specified width and height bounding box.
func ScaleToFit(srcImagePath, outputImagePath string, width, height int) error {
	srcImage, err := imaging.Open(srcImagePath)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}

	// Scale down the image to fit the specified width and height bounding box
	dstImage := imaging.Fit(srcImage, width, height, imaging.Lanczos)

	// Save the resulting image to a new file
	if err := imaging.Save(dstImage, outputImagePath); err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	return nil
}

// ResizeAndFill resizes and crops the image to fill the specified width and height area.
func ResizeAndFill(srcImagePath, outputImagePath string, width, height int) error {
	srcImage, err := imaging.Open(srcImagePath)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}

	// Resize and crop the image to fill the specified width and height area
	dstImage := imaging.Fill(srcImage, width, height, imaging.Center, imaging.Lanczos)

	// Save the resulting image to a new file
	if err := imaging.Save(dstImage, outputImagePath); err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	return nil
}
