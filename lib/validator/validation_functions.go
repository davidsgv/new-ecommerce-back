package validator

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

// valida si un id es valido para la base de datos
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
		return fmt.Errorf(InvalidId, dato)
	}
	return nil
}

// valida si un string tiene unos caracteres minimos
func min(dato interface{}, tagValue string) error {
	str := dato.(string)
	chars, err := strconv.ParseInt(tagValue, 10, 64)
	if err != nil {
		log.Fatal("Debe proveer un string para validar la etiqueta min")
	}

	if len(str) < int(chars) {
		return fmt.Errorf(MinString, dato, tagValue)
	}
	return nil
}

// valida si un string tiene un maximo de caracteres
func max(dato interface{}, tagValue string) error {
	str := dato.(string)
	chars, err := strconv.ParseInt(tagValue, 10, 64)
	if err != nil {
		log.Fatal("Debe proveer un string para validar la etiqueta min")
	}

	if len(str) > int(chars) {
		return fmt.Errorf(MinString, dato, tagValue)
	}
	return nil
}
