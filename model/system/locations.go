package system

// Location - SystemLocations table
type Location struct {
	LocationID uint64 `json:"location_id"`
	CoordX     int64  `json:"coordx"`
	CoordY     int64  `json:"coordy"`
}
