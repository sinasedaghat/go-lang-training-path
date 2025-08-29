package constant

type Role struct {
	Value int
	Label string
}

var Roles = []Role{
	{Value: 1, Label: "SimpleUser"},
	{Value: 2, Label: "Admin"},
}
