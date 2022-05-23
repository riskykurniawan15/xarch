package elsa

import (
	"fmt"
	"io/ioutil"
	"os"
)

func BuildDomain(domain_name string) error {
	if domain_name == "" {
		return fmt.Errorf("Failed to create domain because domain name is null")
	}

	if _, err := os.Stat("domain/" + domain_name); os.IsNotExist(err) {
		path := "domain/" + domain_name + "/models/"
		os.MkdirAll(path, 0700)
		err := ioutil.WriteFile(path+domain_name+".go", []byte("package models"), 0644)
		if err != nil {
			os.MkdirAll(path, 0700)
		}

		path = "domain/" + domain_name + "/repository/"
		os.MkdirAll(path, 0700)
		err = ioutil.WriteFile(path+domain_name+".go", []byte("package repository"), 0644)
		if err != nil {
			os.MkdirAll(path, 0700)
		}

		path = "domain/" + domain_name + "/services/"
		os.MkdirAll(path, 0700)
		err = ioutil.WriteFile(path+domain_name+".go", []byte("package services"), 0644)
		if err != nil {
			return err
		}

		fmt.Println("Success to build domain " + domain_name)

		return nil
	}

	return fmt.Errorf("Failed to create domain because domain is exists")
}
