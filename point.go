package geo

import (
	"fmt"
	"math"
)

// Point описывает координаты географической точки: долгота, широта. Именно такая последовательность
// выбрана только потому, что это наиболее близко представлению координат на плоскости: x и y, а не
// наоборот. Вторая причина — именно такой формат точек поддерживается в MongoDB, а данная
// библиотека сделана в первую очередь именно для работы с этой базой данных.
type Point [2]float64

// NewPoint возвращает новое описание точки с указанными координатами. Если координаты выходят
// за допустимый диапазон данных для географических координат, то возникает panic.
func NewPoint(lon, lat float64) Point {
	if lon < -180 || lon > 180 {
		panic("bad longitude")
	}
	if lat < -90 || lat > 90 {
		panic("bad latitude")
	}
	return Point{lon, lat}
}

// Longitude возвращает долготу точки.
func (p Point) Longitude() float64 {
	return p[0]
}

// Latitude возвращает широту точки.
func (p Point) Latitude() float64 {
	return p[1]
}

// IsZero возвращает true, если обе координаты точки равны нулю.
func (p Point) IsZero() bool {
	return p[0] == 0 && p[1] == 0
}

// Geo возвращает представление точки в формате GeoJSON.
func (p Point) Geo() *GeoJSON {
	if p.IsZero() {
		return nil
	}
	return &GeoJSON{
		Type:        "Point",
		Coordinates: p[:],
	}
}

// MarshalJSON отдает форматированное представление JSON.
func (p Point) MarshalJSON() ([]byte, error) {
	if p.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("[%f,%f]", p.Longitude(), p.Latitude())), nil
}

const (
	EarthRadius float64 = 6378137.0 // радиус Земли в метрах
)

// Move возвращает новую точку, перемещенную от изначально на dist метров в направлении bearing
// в градусах.
func (p Point) Move(dist float64, bearing float64) Point {
	dr := dist / EarthRadius
	bearing = bearing * math.Pi / 180.0
	lon1 := p.Longitude() * math.Pi / 180.0
	lat1 := p.Latitude() * math.Pi / 180.0
	lat2_part1 := math.Sin(lat1) * math.Cos(dr)
	lat2_part2 := math.Cos(lat1) * math.Sin(dr) * math.Cos(bearing)
	lat2 := math.Asin(lat2_part1 + lat2_part2)
	lon2_part1 := math.Sin(bearing) * math.Sin(dr) * math.Cos(lat1)
	lon2_part2 := math.Cos(dr) - (math.Sin(lat1) * math.Sin(lat2))
	lon2 := lon1 + math.Atan2(lon2_part1, lon2_part2)
	lon2 = math.Mod((lon2+3*math.Pi), (2*math.Pi)) - math.Pi
	lon2 = lon2 * 180.0 / math.Pi
	lat2 = lat2 * 180.0 / math.Pi
	return NewPoint(lon2, lat2)
}

// BearingTo возвращает направление в градусах на указанную точку.
func (p Point) BearingTo(p2 Point) float64 {
	dLon := (p2.Longitude() - p.Longitude()) * math.Pi / 180.0
	lat1 := p.Latitude() * math.Pi / 180.0
	lat2 := p2.Latitude() * math.Pi / 180.0
	y := math.Sin(dLon) * math.Cos(lat2)
	x := math.Cos(lat1)*math.Sin(lat2) -
		math.Sin(lat1)*math.Cos(lat2)*math.Cos(dLon)
	return math.Atan2(y, x) * 180.0 / math.Pi
}

// Distance возвращает дистанцию между двумя точками в метрах.
func (p Point) Distance(p2 Point) float64 {
	dLon := (p2.Longitude() - p.Longitude()) * math.Pi / 180.0
	dLat := (p2.Latitude() - p.Latitude()) * math.Pi / 180.0
	lat1 := p.Latitude() * math.Pi / 180.0
	lat2 := p2.Latitude() * math.Pi / 180.0
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	return EarthRadius * 2 * math.Asin(math.Sqrt(a))
}

// Centroid возвращает цент окружности, содержащей все указанные точки.
func Centroid(points ...Point) Point {
	switch l := float64(len(points)); l {
	case 0:
		return Point{}
	case 1:
		return points[0]
	default:
		var lon, lat float64
		for _, point := range points {
			lon += point.Longitude()
			lat += point.Latitude()
		}
		return Point{lon / l, lat / l}
	}
}
