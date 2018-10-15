package system

// Resource - Resources table model
type Resource struct {
	// because i cant be bothered to make my csv code work with
	// anything other than strings right now
	LocationID   string //uint64
	ResourceType string //uint
	Name         string
	Capacity     string //uint
}
