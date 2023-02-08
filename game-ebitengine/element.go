package main

type Element interface {
	Size() (width, height int)
	Coordinate() (x, y float64)
}

func checkElementCollision(elementA, elementB Element) bool {
	width, height := elementA.Size()
	left, top := elementA.Coordinate()
	right, bottom := left+float64(width), top+float64(height)

	// 左上角
	w, h := elementB.Size()
	x, y := elementB.Coordinate()
	if left < x && x < right && top < y && y < bottom {
		return true
	}

	// 右上角
	x += float64(w)
	if left < x && x < right && top < y && y < bottom {
		return true
	}

	// 右下角
	y += float64(h)
	if left < x && x < right && top < y && y < bottom {
		return true
	}

	// 左下角
	x -= float64(w)
	if left < x && x < right && top < y && y < bottom {
		return true
	}

	return false
}
