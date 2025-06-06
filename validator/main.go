package validator

import "regexp"

func Domain(domain string) bool {
	var regex = regexp.MustCompile(`^(?i)[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?(\.[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?)+$`)
	return regex.MatchString(domain)
}
