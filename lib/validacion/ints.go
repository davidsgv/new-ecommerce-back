package validacion

import (
	"fmt"
	"reflect"
)

func validarId(dato interface{}, _ string) error {
	value := reflect.ValueOf(dato)
	kind := value.Kind().String()
	var id int64

	switch kind {
	case "uint", "uint8", "uint16", "uint32", "uint64":
		id = int64(value.Uint())
	case "int", "int8", "int16", "int32", "int64":
		id = value.Int()
	}

	//var id int64 //= reflect.ValueOf(dato).Interface().(int64)
	if id <= 0 {
		return fmt.Errorf("%d no es un id valido", id)
	}
	return nil
}
