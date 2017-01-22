package ovw

import (
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// MergeOverwrite overwrite map
func MergeOverwrite(to, from, dst interface{}) error {
	toMap := structs.Map(to)
	toStruct := structs.New(to)
	fromMap := structs.Map(from)
	fromStruct := structs.New(from)
	for k, v := range fromMap {
		_, ok := toMap[k]
		if !ok {
			return errors.Errorf("no key: %s", k)
		}
		toField := toStruct.Field(k)
		fromField := fromStruct.Field(k)
		if overwriteable(toField, fromField) {
			toMap[k] = v
		}
	}
	if err := mapstructure.Decode(toMap, dst); err != nil {
		return errors.Wrap(err, "faield to decode")
	}
	return nil
}

func overwriteable(to, from *structs.Field) bool {
	if to.Kind().String() == "bool" {
		return true
	}
	switch {
	case to.IsZero() && !from.IsZero():
		return true
	case !to.IsZero() && !from.IsZero():
		return true
	case !to.IsZero() && from.IsZero():
		return false
	}
	return true
}
