package tStruct

import (
	"errors"
	"regexp"
	"strings"
)

type Website struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rate        int    `json:"rate"`
}

func NewWebsite(address, name string) (*Website, error) {
	if address == "" || !validDomain(address) {
		return nil, errors.New("you must provide valid the address")
	}

	if name == "" {
		lastIndex := strings.LastIndex(address, ".")
		if lastIndex == -1 {
			name = address
		} else {
			name = address[:lastIndex]
		}
	}
	return &Website{
		Address: address,
		Name:    name,
	}, nil
}

func validDomain(domain string) bool {
	var regex = regexp.MustCompile(`^(?i)[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?(\.[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?)+$`)
	return regex.MatchString(domain)
}
