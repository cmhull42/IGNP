package system

// Resource - Resources table model
type Resource struct {
	ResourceID     uint64  `json:"resource_id"`
	ResourceTypeID uint    `json:"resource_type_id"`
	Name           string  `json:"name"`
	Rarity         float32 `json:"rarity"`
}
