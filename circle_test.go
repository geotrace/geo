package geo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCircleGeo(t *testing.T) {
	circle := Circle{
		Center: NewPoint(37.57351, 55.715084),
		Radius: 500,
	}
	fmt.Println(circle.Geo())
}

func TestCircleJSON(t *testing.T) {
	circle := Circle{
		Center: NewPoint(37.57351, 55.715084),
		Radius: 500.639843,
	}
	data, err := json.Marshal(circle)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
	data, err = json.MarshalIndent(circle.Geo(), "", "\t")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
}
