package validator

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

// para implementar un singleton
var lock = &sync.Mutex{}
var instance *Validator

// funcion de validacion,
type funcionValidacion func(dato interface{}, tagValue string) error

// validador
type Validator struct {
	tags map[string]funcionValidacion
}

func Validate(tag string, value interface{}) error {
	instance := GetValidator()
	return instance.Validate(tag, value)
}

// crear funcion que valide si contiene la etiqueta opcional, si viene vacio el dato se omiten validaciones
func (val *Validator) Validate(tag string, value interface{}) (err error) {
	rules := strings.Split(tag, ",")
	for _, rule := range rules {
		parts := strings.Split(rule, "=")
		key := parts[0]

		var tagvalue string
		if len(parts) > 1 {
			tagvalue = parts[1]
		}

		funcion := val.tags[key]
		if funcion == nil {
			mensaje := fmt.Sprintf("tag no registrada: %s", rule)
			return errors.New(mensaje)
		}
		err := funcion(value, tagvalue)
		if err != nil {
			return err
		}
	}
	return err
}

// obtiene la instancia singleton
func GetValidator() *Validator {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &Validator{}
			instance.registerfuncs()
		}
	}
	return instance
}

// registra las funciones que va a tener la instacia del validador
func (instance *Validator) registerfuncs() {
	instance.tags = make(map[string]funcionValidacion)

	instance.registrarValidacion("email", esEmail)
	instance.registrarValidacion("correo", esEmail)
	instance.registrarValidacion("dominio", esDominio)
	// instance.registrarValidacion("jwt", validarToken)
	instance.registrarValidacion("id", validarId)
	instance.registrarValidacion("min", min)
	instance.registrarValidacion("max", max)
}

// registra una funcion en el validador
func (val *Validator) registrarValidacion(tag string, funcion funcionValidacion) {
	val.tags[tag] = funcion
}
