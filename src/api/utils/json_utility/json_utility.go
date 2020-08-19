package json_utility

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
)

func StructToLowerCaseJson(data interface{}) (interface{}, error){

	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var r interface{}

	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	spew.Dump(r)

	return r, nil
}
