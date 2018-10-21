package system

// ResourceLocation - SystemResourceLocations table
type ResourceLocation struct {
	ResourceType uint   `json:"resource_type"`
	LocationID   uint64 `json:"location_id"`
	Capacity     uint   `json:"capacity"`
}
