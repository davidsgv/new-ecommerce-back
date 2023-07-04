package validator

import (
	"errors"
	"fmt"
)

func validarLongitud(dato string) error {
	err := errors.New(fmt.Sprintf("%s no contiene texto", dato))
	if len(dato) <= 0 {
		return err
	}
	return nil
}

func esEmail(dato interface{}, tagValue string) error {
	correo := dato.(string)
	err := errors.New(fmt.Sprintf("%s no es un correo valido", correo))
	return validarRegEx(correoRegEx, correo, err)
}

func esDominio(dato interface{}, _ string) error {
	dominio := dato.(string)
	err := validarLongitud(dominio)
	if err != nil {
		return err
	}

	err = errors.New(fmt.Sprintf("%s no es un dominio valido", dominio))
	err = validarRegEx(dominioRegEx, dominio, err)
	if err == nil {
		return err
	}

	if esIp(dominio, "") != nil {
		return err
	}
	return nil
}

func esIp(dato interface{}, _ string) error {
	ip := dato.(string)
	err := validarLongitud(ip)
	if err != nil {
		return err
	}

	err = errors.New(fmt.Sprintf("%s no es una ip valida", ip))
	return validarRegEx(localIpRegEx, ip, err)
}

// func validarToken(dato interface{}, _ string) error {
// 	token := dato.(string)
// 	err := validarLongitud(token)
// 	if err != nil {
// 		return err
// 	}

// 	err = errors.New(fmt.Sprintf("%s no es un token valido", token))
// 	return validarRegEx(jwtRegEx, token, err)
// }

func esNumero(dato interface{}, _ string) error {
	str := dato.(string)
	err := errors.New(fmt.Sprintf("%s solo puede contener numeros", str))
	return validarRegEx(numeroRegEx, str, err)
}
