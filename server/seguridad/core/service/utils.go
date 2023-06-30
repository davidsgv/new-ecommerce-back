package service

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func cifrarContraseña(password string) (string, error) {
	sha := sha3.New384()
	_, err := sha.Write([]byte(password))
	if err != nil {
		return "", err
	}
	hash := fmt.Sprintf("%x", sha.Sum(nil))
	return hash, nil
}
