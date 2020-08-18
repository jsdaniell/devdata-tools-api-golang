package rules

import (
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/models"
)

type Values struct {
	ChildrenName string
	StructType interface{}
}

func SetValues(childrenName string, structType interface{}) Values {
	val := Values{}

	val.ChildrenName = childrenName
	val.StructType = structType

	return val
}

var s = map[string]Values{
	"testsGroups":    SetValues("tests", models.Test{}),
	"useCasesGroups":  SetValues("useCases", models.UseCase{}),
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
		return s[suiteType].ChildrenName, nil
	} else {
		return "" ,fmt.Errorf("invalid %q suite type required", suiteType)
	}
}

func GetInterfaceOfSuite(suiteType string) (interface{}, error) {
	if _, ok := s[suiteType]; ok {
		return s[suiteType].StructType, nil
	} else {
		return "" ,fmt.Errorf("invalid %q suite type required", suiteType)
	}
}
