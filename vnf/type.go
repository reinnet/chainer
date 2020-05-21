package vnf

// Type is a Virtual Network Function
type Type struct {
	Name string
}

// Types returns all available VNF types.
func Types() []Type {
	return []Type{
		{"vIDS"},
		{"vNAT"},
		{"vFW"},
		{"vDPI"},
	}
}
