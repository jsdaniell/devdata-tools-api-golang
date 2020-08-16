package rules

import "fmt"

var s = map[string]string{
	"testsGroups":    "tests",
	"useCasesGroups": "useCases",
}

func ValidateExistentSuites(suiteType string) error {
	if _, ok := s[suiteType]; ok {
		return nil
	} else {
		return fmt.Errorf("invalid %q suite type required", suiteType)
	}
}

func GetChildrenNameOfSuite(suiteType string) (string, error) {
	if _, ok := s[suiteType]; ok {
		return s[suiteType], nil
	} else {
		return "" ,fmt.Errorf("invalid %q suite type required", suiteType)
	}
}
