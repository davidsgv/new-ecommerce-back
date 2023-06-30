package main

import (
	"log"
	"os"
	"seguridad/core/service"
	"seguridad/driven/mysql"
	"seguridad/driver/rest"
	"strconv"
	"time"
	"validacion"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type EnvVars struct {
	//app
	host            string
	port            string
	keyServers      string
	keyTokens       string
	TokenExpireTime time.Duration
	//debug
	ginMode string
	//DB
	connMaxLifetime time.Duration
	connMaxIdleTime time.Duration
	maxIdleConns    int
	maxOpenConns    int
	dbUrl           string
	//Logger
	logFile string
}

func main() {
	//cargar ambiente
	ambiente := cargarEnvVars()

	//emprezar a escribir los logs en un archivo si no es debug
	startLogger(ambiente.logFile, ambiente.ginMode)

	server := gin.New()
	group := server.Group("/api/")

	repo, err := mysql.NewMysqlRepository(mysql.MysqlConfig{
		ConnMaxLifetime: ambiente.connMaxLifetime,
		ConnMaxIdleTime: ambiente.connMaxIdleTime,
		MaxIdleConns:    ambiente.maxIdleConns,
		MaxOpenConns:    ambiente.maxOpenConns,
		URL:             ambiente.dbUrl,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	validator := validacion.NewValidador()
	//servicios
	//autenticacionServicio := service.NewAutenticacionServicio(repo, ambiente.TokenExpireTime, ambiente.keyTokens, validador)
	rolService := service.NewRolService(repo, validator)
	//empresaServicio := service.NewEmpresaServicio(repo, ambiente.keyServers, validador)
	//usuarioServicio := service.NewUsuarioServicio(repo, validador)

	//midleware
	//group.Use(rest.AuthorizeJWT(autenticacionServicio))

	//handlers
	//rest.NewAutenticacionHandler(server, group, autenticacionServicio)
	rest.NewRolHandler(group, rolService)
	//rest.NewEmpresasHandler(group, empresaServicio)
	//rest.NewUsuarioHandler(group, usuarioServicio)

	gin.SetMode(ambiente.ginMode)
	server.Run(ambiente.host + ":" + ambiente.port)
}

func cargarEnvVars() EnvVars {
	ambiente := EnvVars{}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("No se pudieron cargar la variables de entorno")
	}

	//cargar env vars
	ambiente.logFile = os.Getenv("LOG_FILE")
	ambiente.ginMode = os.Getenv("GIN_MODE")
	ambiente.host = os.Getenv("HOST")
	ambiente.port = os.Getenv("PORT")
	ambiente.keyServers = os.Getenv("SECURITY_KEY_SERVER")
	ambiente.keyTokens = os.Getenv("SECURITY_KEY_TOKENS")

	expireTime, err := strconv.ParseInt(os.Getenv("TOKEN_EXPIRE_TIME"), 10, 64)
	if err != nil {
		expireTime = 24
		log.Println("No se pudo establecer el tiempo de duración del token")
	}
	ambiente.TokenExpireTime = time.Duration(expireTime)

	aux := os.Getenv("ConnMaxLifetime")
	aux2, err := strconv.ParseInt(aux, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ambiente.connMaxIdleTime = time.Duration(aux2)

	aux = os.Getenv("ConnMaxIdleTime")
	aux2, err = strconv.ParseInt(aux, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ambiente.connMaxIdleTime = time.Duration(aux2)

	aux = os.Getenv("MaxIdleConns")
	aux2, err = strconv.ParseInt(aux, 10, 32)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ambiente.maxIdleConns = int(aux2)

	aux = os.Getenv("MaxOpenConns")
	aux2, err = strconv.ParseInt(aux, 10, 32)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ambiente.maxOpenConns = int(aux2)

	ambiente.dbUrl = os.Getenv("DB_URL")

	return ambiente
}

func startLogger(filePath string, debugMode string) {
	if debugMode == "debug" || debugMode == "test" {
		return
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}
