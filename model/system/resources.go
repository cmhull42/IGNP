package system

// Resource - Resources table model
type Resource struct {
	ID           uint64  `json:"resource_id"`
	ResourceType uint    `json:"resource_type_id"`
	Name         string  `json:"name"`
	Rarity       float32 `json:"rarity"`
}
