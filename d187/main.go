package main

import "log"

// Coordinate contains the X and Y coordinate of a rectangle point
type Coordinate struct {
	XCoordinate int
	YCoordinate int
}

// Dimension has the length and height of a Rectangle
type Dimension struct {
	Length int
	Height int
}

// Rectangle is the rectangle struct
type Rectangle struct {
	TopLeft    Coordinate
	Dimensions Dimension
}

// GetBottomRightCoordinates returns the bottom right coordinates given
// The TopLeft and Dimension of a rectangle
func (r *Rectangle) GetBottomRightCoordinates() Coordinate {
	bottomRight := Coordinate{
		XCoordinate: r.TopLeft.XCoordinate + r.Dimensions.Length,
		YCoordinate: r.TopLeft.YCoordinate + r.Dimensions.Height,
	}
	return bottomRight
}

// Overlaps checks to see if rectangle r overlaps rectangle r1
func (r *Rectangle) Overlaps(r1 Rectangle) bool {
	if r.TopLeft.XCoordinate <= r1.GetBottomRightCoordinates().YCoordinate {
		if r.TopLeft.YCoordinate <= r1.GetBottomRightCoordinates().YCoordinate {
			return true
		}
		return false
	}
	return false
}

func main() {
	rectangles := []Rectangle{
		NewRectangle(1, 4, 3, 3),
		NewRectangle(-1, 3, 2, 1),
		NewRectangle(0, 5, 4, 3),
	}
	log.Printf("Overlapping: %v\n", RectanglesOverlap(rectangles))
}

// RectanglesOverlap checks to see if any rectangles overlap
func RectanglesOverlap(rectangles []Rectangle) bool {
	for i := 0; i < len(rectangles); i++ {
		for j := 1; j < len(rectangles)-1; j++ {
			if rectangles[i].Overlaps(rectangles[j]) {
				return true
			}
		}
	}
	return false
}

// NewRectangle creates a new rectangle
func NewRectangle(topLeftX, topLeftY, dimensionLength, dimensionHeight int) Rectangle {
	return Rectangle{
		TopLeft: Coordinate{
			XCoordinate: topLeftX,
			YCoordinate: topLeftY,
		},
		Dimensions: Dimension{
			Length: dimensionLength,
			Height: dimensionHeight,
		},
	}
}
