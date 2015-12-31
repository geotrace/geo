package geo

// Polygon описывает полигон. Первый элемент массива описывает точки многоугольника, которые
// являются внешним контуром. Все остальные — описывают изъятия из данного многоугольника.
type Polygon [][]Point

// NewPolygon возвращает новое описание многоугольника, состоящего из заданных точек (без изъятий).
func NewPolygon(points ...Point) Polygon {
	p1, p2 := points[0], points[len(points)-1]
	if p1.Longitude() != p2.Longitude() || p1.Latitude() != p2.Latitude() {
		points = append(points, p1)
	}
	return Polygon{points}
}

// Exclude добавляет к полигону описание многоугольника для исключения области из основного
// полигона. Но проверки, что данные многоугольники вообще пересекаются не происходит, поэтому
// нужно быть внимательнее при использовании данной функции.
func (p *Polygon) Exclude(points ...Point) {
	if p == nil {
		return
	}
	p1, p2 := points[0], points[len(points)-1]
	if p1.Longitude() != p2.Longitude() || p1.Latitude() != p2.Latitude() {
		points = append(points, p1)
	}
	*p = append(*p, points)
}

// Geo возвращает описание полигона в формате GeoJSON
func (p Polygon) Geo() *GeoJSON {
	if len(p) == 0 {
		return nil
	}
	return &GeoJSON{
		Type:        "Polygon",
		Coordinates: p,
	}
}
