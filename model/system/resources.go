package system

// Resource - Resources table model
type Resource struct {
	LocationID   uint64
	ResourceType uint
	Name         string
	Capacity     uint
}
