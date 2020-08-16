package rules

import "strings"

func DocNameByTitle(title string) string {
	return strings.ToLower(strings.Replace(title ," ", "-", -1))
}
