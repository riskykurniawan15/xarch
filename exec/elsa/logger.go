package elsa

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/riskykurniawan15/xarch/logger"
)

var log_dir string = "./logger/"

func FlushLog() error {
	files, err := ioutil.ReadDir(log_dir)
	if err != nil {
		return err
	}

	counter := 0
	for _, f := range files {
		if f.Name() != "logger.go" && f.Name() != logger.Prefix()+".log" {
			e := os.Remove(log_dir + f.Name())
			if e == nil {
				counter++
				log.Println("Success to Remove log " + f.Name())
			} else {
				log.Println("Failed to Remove log " + f.Name())
			}
		}
	}

	if counter == 0 {
		log.Println("No Logs File Cleared")
	}

	return nil
}
