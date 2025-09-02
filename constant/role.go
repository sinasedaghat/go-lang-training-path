package constant

import "strings"

type Role struct {
	Value int
	Label string
}

var Roles = []Role{
	{Value: 1, Label: "SimpleUser"},
	{Value: 2, Label: "Admin"},
}

func DefaultRole() int {
	for _, role := range Roles {
		if strings.ToLower(role.Label) == "simpleuser" {
			return role.Value
		}
	}
	return 0
}
