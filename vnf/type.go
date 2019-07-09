package vnf

// Type is a Virtual Network Function
type Type struct {
	Name     string
	Capacity int
}

var types = []Type{
	{"vIDS", 250},
	{"vNAT", 500},
	{"vFW", 500},
	{"vDPI", 500},
}

// Len return total number of VNFs
func Len() int {
	return len(types)
}

// Get return VNF type that is associated with given index
func Get(i int) Type {
	return types[i]
}
