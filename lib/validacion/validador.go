package validacion

import (
	"errors"
	"fmt"
	"strings"
)

func NewValidador() *Validador {
	Validador := Validador{tags: make(map[string]funcionValidacion)}
	Validador.registrarValidacion("email", esEmail)
	Validador.registrarValidacion("correo", esEmail)
	Validador.registrarValidacion("dominio", esDominio)
	Validador.registrarValidacion("jwt", validarToken)
	Validador.registrarValidacion("id", validarId)
	Validador.registrarValidacion("min", min)
	Validador.registrarValidacion("max", max)
	return &Validador
}

// funcion de validacion,
type funcionValidacion func(dato interface{}, tagValue string) error

type Validador struct {
	tags map[string]funcionValidacion
}

// crear funcion que valide si contiene la etiqueta opcional, si viene vacio el dato se omiten validaciones
func (val *Validador) Verificar(tag string, dato interface{}) (err error) {
	rules := strings.Split(tag, ",")
	for _, rule := range rules {
		parts := strings.Split(rule, "=")
		key := parts[0]

		var value string
		if len(parts) > 1 {
			value = parts[1]
		}

		funcion := val.tags[key]
		if funcion == nil {
			mensaje := fmt.Sprintf("tag no registrada: %s", rule)
			return errors.New(mensaje)
		}
		err := funcion(dato, value)
		if err != nil {
			return err
		}
	}
	return err
}

func (val *Validador) registrarValidacion(tag string, funcion funcionValidacion) {
	val.tags[tag] = funcion
}
