package service

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"encoding/base64"
// 	"errors"
// 	"io"
// 	"seguridad/core/domain"
// 	"seguridad/core/repository"
// 	"validacion"

// 	"gopkg.in/validator.v2"
// )

// type EmpresaServicio struct {
// 	repo      repository.IRepositoryEmpresa
// 	key       string //clave cifrar datos servidor
// 	validador *validacion.Validador
// }

// func NewEmpresaServicio(repo repository.IRepositoryEmpresa, key string, validador *validacion.Validador) *EmpresaServicio {
// 	return &EmpresaServicio{
// 		repo:      repo,
// 		key:       key,
// 		validador: validador,
// 	}
// }

// // region Servidor
// func (servicio *EmpresaServicio) CreateServidor(servidor *domain.Servidor) (*domain.Servidor, error) {
// 	//validar datos
// 	err := validator.Validate(servidor)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//cifrar datos servidor
// 	key := []byte(servicio.key)
// 	servidor.NombreBD, err = cifrar(key, servidor.NombreBD)
// 	if err != nil {
// 		return nil, err
// 	}
// 	servidor.Password, err = cifrar(key, servidor.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	servidor.Usuario, err = cifrar(key, servidor.Usuario)
// 	if err != nil {
// 		return nil, err
// 	}
// 	servidor.DireccionIp, err = cifrar(key, servidor.DireccionIp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// var nombreBD, password, usuario string
// 	// nombreBD, password, usuario, err = servicio.cifrarDatos(servidor.NombreBD, servidor.Password, servidor.Usuario)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// servidor.NombreBD = nombreBD
// 	// servidor.Password = password
// 	// servidor.Usuario = usuario

// 	//guardar servidor
// 	return servicio.repo.CreateServidor(servidor)
// }

// func (servicio *EmpresaServicio) UpdateServidor(servidor *domain.Servidor) (*domain.Servidor, error) {
// 	//validar datos
// 	// err := servicio.validator.Validate(servidor)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	//cifrar datos servidor
// 	//cifrar datos servidor
// 	var err error
// 	key := []byte(servicio.key)
// 	servidor.NombreBD, err = cifrar(key, servidor.NombreBD)
// 	if err != nil {
// 		return nil, err
// 	}
// 	servidor.Password, err = cifrar(key, servidor.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	servidor.Usuario, err = cifrar(key, servidor.Usuario)
// 	if err != nil {
// 		return nil, err
// 	}
// 	servidor.DireccionIp, err = cifrar(key, servidor.DireccionIp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return servicio.repo.UpdateServidor(servidor)
// }

// func (servicio *EmpresaServicio) GetServidores() ([]domain.Servidor, error) {
// 	return servicio.repo.GetServidores()
// }

// func (servicio EmpresaServicio) DeleteServidor(id uint) error {
// 	return servicio.repo.DeleteServidor(id)
// }

// //endregion

// // region Empresa
// func (servicio *EmpresaServicio) GetEmpresas() ([]domain.Empresa, error) {
// 	return servicio.repo.GetEmpresas()
// }

// func (servicio *EmpresaServicio) GetEmpresaById(id uint) (*domain.Empresa, error) {
// 	empresa, err := servicio.repo.GetEmpresaById(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if empresa == nil {
// 		return nil, errors.New("no existe la empresa")
// 	}

// 	// for i := 0; i < len(empresa.Servidores); i++ {
// 	// 	servidor := empresa.Servidores[i]
// 	// 	nombreBD, password, usuario, err := servicio.decifrarDatos(servidor.NombreBD, servidor.Password, servidor.Usuario)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	servidor.NombreBD = nombreBD
// 	// 	servidor.Password = password
// 	// 	servidor.Usuario = usuario
// 	// }

// 	return empresa, nil
// }

// func (servicio *EmpresaServicio) CreateEmpresa(empresa domain.Empresa) error {
// 	//validar los datos
// 	err := validator.Validate(empresa)
// 	if err != nil {
// 		return err
// 	}

// 	//insertar en bd
// 	return servicio.repo.CreateEmpresa(empresa)
// }

// func (servicio *EmpresaServicio) UpdateEmpresa(empresa *domain.Empresa) error {
// 	//actualizar los datos de la empresa
// 	// err := servicio.validator.Validate(empresa)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	return servicio.repo.UpdateEmpresa(*empresa)
// }

// func (servicio EmpresaServicio) DeleteEmpresa(id uint) error {
// 	return servicio.repo.DeleteEmpresa(id)
// }

// //endregion

// // region Funciones
// func cifrar(key []byte, message string) (encoded string, err error) {
// 	//Create byte array from the input string
// 	plainText := []byte(message)

// 	//Create a new AES cipher using the key
// 	block, err := aes.NewCipher(key)

// 	//IF NewCipher failed, exit:
// 	if err != nil {
// 		return
// 	}

// 	//Make the cipher text a byte array of size BlockSize + the length of the message
// 	cipherText := make([]byte, aes.BlockSize+len(plainText))

// 	//iv is the ciphertext up to the blocksize (16)
// 	iv := cipherText[:aes.BlockSize]
// 	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
// 		return
// 	}

// 	//Encrypt the data:
// 	stream := cipher.NewCFBEncrypter(block, iv)
// 	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

// 	//Return string encoded in base64
// 	return base64.RawStdEncoding.EncodeToString(cipherText), err
// }

// func descifrar(key []byte, secure string) (decoded string, err error) {
// 	//Remove base64 encoding:
// 	cipherText, err := base64.RawStdEncoding.DecodeString(secure)

// 	//IF DecodeString failed, exit:
// 	if err != nil {
// 		return
// 	}

// 	//Create a new AES cipher with the key and encrypted message
// 	block, err := aes.NewCipher(key)

// 	//IF NewCipher failed, exit:
// 	if err != nil {
// 		return
// 	}

// 	//IF the length of the cipherText is less than 16 Bytes:
// 	if len(cipherText) < aes.BlockSize {
// 		err = errors.New("ciphertext block size is too short")
// 		return
// 	}

// 	iv := cipherText[:aes.BlockSize]
// 	cipherText = cipherText[aes.BlockSize:]

// 	//Decrypt the message
// 	stream := cipher.NewCFBDecrypter(block, iv)
// 	stream.XORKeyStream(cipherText, cipherText)

// 	return string(cipherText), err
// }

// //endregion
