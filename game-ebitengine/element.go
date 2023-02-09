package main

type Element interface {
	Size() (width, height int)
	Coordinate() (x, y float64)
}

func checkElementCollision(elementA, elementB Element) bool {
	widthA, heightA := elementA.Size()
	leftA, topA := elementA.Coordinate()

	widthB, heightB := elementB.Size()
	leftB, topB := elementB.Coordinate()

	top, bottom := max(topA, topB), min(topA+float64(heightA), topB+float64(heightB))
	left, right := max(leftA, leftB), min(leftA+float64(widthA), leftB+float64(widthB))

	return top <= bottom && left <= right
}

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
