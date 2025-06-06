package tStruct

import (
	"errors"
	"practices/validator"
	"strings"
)

type Website struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rate        int    `json:"rate"`
}

type websiteOption func(*Website)

func WithName(name string) websiteOption {
	return func(w *Website) {
		w.Name = name
	}
}

func WithTitle(title string) websiteOption {
	return func(w *Website) {
		w.Title = title
	}
}

func WithDescription(desc string) websiteOption {
	return func(w *Website) {
		w.Description = desc
	}
}

func WithRate(rate int) websiteOption {
	return func(w *Website) {
		w.Rate = rate
	}
}

func NewWebsite(address string, options ...websiteOption) (*Website, error) {
	if address == "" || !validator.Domain(address) {
		return nil, errors.New("you must provide valid the address")
	}

	website := Website{
		Address: address,
	}

	for _, option := range options {
		option(&website)
	}

	if website.Name == "" {
		lastIndex := strings.LastIndex(address, ".")
		if lastIndex == -1 {
			website.Name = address
		} else {
			website.Name = address[:lastIndex]
		}
	}

	return &website, nil
}

// usage
// website, err := tStruct.NewWebsite("newApp.booking.ir", tStruct.WithName("booking"), tStruct.WithTitle("new app of booking.ir"), tStruct.WithDescription(""))
// if err == nil {
// 	fmt.Println(website)
// } else {
// 	fmt.Println(err)
// }
