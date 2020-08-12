package models

type Test struct {
	title         string
	actor         string
	environment   string
	id            string
	name          string
	postCondition string
	priority      string
	preconditions []string
	procedures    []string
}
