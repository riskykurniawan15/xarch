package elsa

import (
	"fmt"
	"io/ioutil"
)

const EnvTemplate string = `#SERVER
SERVER=0.0.0.0
PORT=1999

#DATABASE
DB_USER=root
DB_PASS=
DB_SERVER=localhost
DB_PORT=3306
DB_NAME=xarch`

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

	fmt.Println("Success to build environtment")

	return nil
}
