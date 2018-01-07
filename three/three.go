package three

type spiralPoint struct {
	x int
	y int
}

func abs(value int) int {
	if value < 0 {
		return value * -1
	} else {
		return value
	}
}

func getSpiralPoint(loc int) (spiralPoint, error) {
	spiral := map[spiralPoint]int{}
	currentPoint := spiralPoint{x: 0, y: 0}
	newPoint := spiralPoint{x: 0, y: 0}
	spiral[currentPoint] = 1

	for i := 1; i < loc; i++ {
		switch nextMove(spiral, currentPoint) {
			case "down":
				newPoint = spiralPoint{x: currentPoint.x, y: currentPoint.y-1}
			case "up":
				newPoint = spiralPoint{x: currentPoint.x, y: currentPoint.y+1}
			case "right":
				newPoint = spiralPoint{x: currentPoint.x+1, y: currentPoint.y}
			case "left":
				newPoint = spiralPoint{x: currentPoint.x-1, y: currentPoint.y}
		}
		currentPoint = newPoint
		spiral[currentPoint] = i
	}
	return currentPoint, nil
}

func nextMove(spiral map[spiralPoint]int, point spiralPoint) string {
	if point.x == 0 && point.y == 0 {
		return "right"
	}

	_, hasLeft := spiral[spiralPoint{x: point.x-1, y: point.y}]
	_, hasRight := spiral[spiralPoint{x: point.x+1, y: point.y}]
	_, hasAbove := spiral[spiralPoint{x: point.x, y: point.y+1}]
	_, hasBelow := spiral[spiralPoint{x: point.x, y: point.y-1}]

	if !hasAbove && hasLeft {
		return "up"
	} else if !hasLeft && hasBelow {
		return "left"
	} else if !hasBelow && hasRight {
		return "down"
	} else if !hasRight && hasAbove {
		return "right"
	} else {
		return "err"
	}
}

func Distance(loc int) int {
	point, err := getSpiralPoint(loc)

	if err != nil {
		return 0
	}
	return abs(point.x) + abs(point.y)
}
