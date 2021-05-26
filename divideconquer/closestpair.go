package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	X, Y float64
}

func dist(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.X - p2.X, 2) + math.Pow(p1.Y - p2.Y, 2))
}

func bruteForce(P []Point, n int) float64 {
	min := math.Inf(1)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := dist(P[i], P[j])
			if d < min {
				min = d
			}
		}
	}
	return min
}

func stripClosest(strip []Point, size int, d float64) float64 {
	min := d
	// sort by Y
	sort.Slice(strip, func(i, j int) bool {
		return strip[i].Y < strip[j].Y
	})

	for i := 0; i < size; i++ {
		for j := i + 1; j < size && (strip[j].Y - strip[i].Y) < min; j++ {
			nowD := dist(strip[i], strip[j])
			if nowD < min {
				min = nowD
			}
		}
	}
	return min
}

func closestUtil(P []Point, n int) float64 {
	if n <= 3 {
		return bruteForce(P, n)
	}

	mid := n / 2
	midPoint := P[mid]

	Pl := P[:mid]
	Pr := P[mid:]

	dl := closestUtil(Pl, mid)
	dr := closestUtil(Pr, n - mid)

	d := math.Min(dl, dr)

	strip := make([]Point, 0)
	for i := 0; i < n; i++ {
		if math.Abs(P[i].X - midPoint.X) < d {
			strip = append(strip, P[i])
		}
	}

	return math.Min(d, stripClosest(strip, len(strip), d))
}

func closestPair(P []Point, n int) float64 {
	// sort by X
	sort.Slice(P, func(i, j int) bool {
		return P[i].X < P[j].X
	})

	return closestUtil(P, n)
}

func main() {
	P := []Point{
		{29, 2}, {0, 19}, {4, 94}, {51, 45}, {80, 47},
		{44, 34}, {28, 80}, {14, 54}, {43, 18}, {71, 44},
		{86, 6}, {33, 97}, {49, 94}, {23, 37}, {55, 90},
		{19, 96}, {97, 57}, {93, 43}, {84, 79}, {23, 81},
	}

	minVal := closestPair(P, len(P))
	fmt.Println(minVal)
	// result: 5.09902
}
