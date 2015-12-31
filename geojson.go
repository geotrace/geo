package geo

// GeoJSON описывает представление географических примитивов в виде GeoJSON.
type GeoJSON struct {
	Type        string
	Coordinates interface{}
}
