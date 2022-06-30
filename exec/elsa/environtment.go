package elsa

import (
	"io/ioutil"
	"log"
)

const EnvTemplate string = `#HTTP SERVER
USING_SECURE=false #true or false
SERVER=0.0.0.0
PORT=1999

#DATABASE
DB_DRIVER= #MYSQL
DB_USER=root
DB_PASS=
DB_SERVER=localhost
DB_PORT=3306
DB_NAME=xarch
DB_MAX_IDLE_CON=10
DB_MAX_OPEN_CON=100
DB_MAX_LIFE_TIME=5 #minute

#REDIS
RDB_ADDRESS=127.0.0.1
RDB_PORT=6379
RDB_USER=
RDB_PASS=
RDB_DB_DEFAULT=0

#EMAIL
EMAIL_HOST=""
EMAIL_PORT=
EMAIL_NAME=""
EMAIL_EMAIL=""
EMAIL_PASSWORD=""

#KAFKA
KAFKA_SERVER=localhost
KAFKA_PORT=9092
KAFKA_CONSUMER_GROUP=

#TOPIC_KAFKA
TOPIC_EMAIL_VERIFIED=""

#OTHER
ALQURAN_API="https://quranlci.com/api/"

#JWT
JWT_SECRET_KEY="SecretKey"
JWT_EXPIRED=24  #IN HOURS`

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
