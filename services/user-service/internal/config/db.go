package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	host           string
	port           string
	user           string
	password       string
	dBName         string
	DataSourceName string
}

func NewDB() (*DBConfig, error) {
	c := DBConfig{}

	c.host = os.Getenv("POSTGRES_HOST")
	c.port = os.Getenv("POSTGRES_PORT")
	c.user = os.Getenv("POSTGRES_USER")
	c.password = os.Getenv("POSTGRES_PASSWORD")
	c.dBName = os.Getenv("POSTGRES_DBNAME")

	// todo: сделай обработку ошибок
	// reflect.MapOf()
	// for key, value := range c {
	// }

	// if len(host) == 0 {
	// 	return nil, errors.New("grpc host not found")
	// }

	c.DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.host, c.port, c.user, c.password, c.dBName)

	return &c, nil
}
