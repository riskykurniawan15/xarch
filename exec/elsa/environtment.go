package elsa

import (
	"io/ioutil"
	"log"
)

const EnvTemplate string = `#SERVER
SERVER=0.0.0.0
PORT=1999

#DATABASE
DB_USER=root
DB_PASS=
DB_SERVER=localhost
DB_PORT=3306
DB_NAME=xarch

#REDIS
RDB_ADDRESS=127.0.0.1
RDB_PORT=6379
RDB_PASS=
RDB_DB_DEFAULT=0

#JWT
JWT_SECRET_KEY="SecretKey"
JWT_EXPIRED=24  #IN HOURS

#OTHER
ALQURAN_API="https://quranlci.com/api/"`

func BuildEnvirontment() error {
	data, err := ioutil.ReadFile(".env.example")
	if err != nil {
		data = []byte(EnvTemplate)
	}
	// Write data to dst
	err = ioutil.WriteFile(".env", data, 0644)
	if err != nil {
		return err
	}

	log.Println("Success to build environtment")

	return nil
}
