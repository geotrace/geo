package geo

import (
	"fmt"
	"testing"
)

func TestPolygon(t *testing.T) {
	p := NewPolygon(Point{0, 0}, Point{1, 1}, Point{1, 0})
	fmt.Println(p.Geo())
	p.Exclude(Point{0, 0}, Point{.5, .5}, Point{.5, 0})
	fmt.Println(p.Geo())
}
