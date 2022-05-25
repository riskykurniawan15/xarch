package elsa

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
		err = ioutil.WriteFile(path+domain_name+".go", []byte(TemplateRepo(domain_name)), 0644)
		if err != nil {
			os.MkdirAll(path, 0700)
		}

		path = "domain/" + domain_name + "/services/"
		os.MkdirAll(path, 0700)
		err = ioutil.WriteFile(path+domain_name+".go", []byte(TemplateService(domain_name)), 0644)
		if err != nil {
			return err
		}

		log.Println("Success to build domain " + domain_name)

		return nil
	}

	return fmt.Errorf("Failed to create domain because domain is exists")
}

func TemplateRepo(domain_name string) string {
	domain_name = strings.Title(domain_name)
	return fmt.Sprintf(`package repository

type %sRepo struct {
}

func New%sRepo() *%sRepo {
	return &%sRepo{}
}
`, domain_name, domain_name, domain_name, domain_name)
}

func TemplateService(domain_name string) string {
	domain_name = strings.Title(domain_name)
	return fmt.Sprintf(`package services

import (
	"github.com/riskykurniawan15/xarch/domain/%s/repository"
)

type %sService struct {
	%sRepo *repository.%sRepo
}

func New%sService(Repo *repository.%sRepo) *%sService {
	return &%sService{
		Repo,
	}
}
`, strings.ToLower(domain_name), domain_name, domain_name, domain_name, domain_name, domain_name, domain_name, domain_name)
}
