package elsa

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CreateDomain(domain_name string) error {
	if domain_name == "" {
		return fmt.Errorf("failed to create domain because domain name is null")
	}

	if _, err := os.Stat("domain/" + domain_name); os.IsNotExist(err) {
		var content_files map[string]string = map[string]string{
			"domain/" + domain_name + "/models/":     "package models",
			"domain/" + domain_name + "/repository/": TemplateRepo(domain_name),
			"domain/" + domain_name + "/services/":   TemplateService(domain_name),
		}

		for path, content := range content_files {
			if err := BuildContent(path, domain_name, content); err != nil {
				return err
			}
		}

		log.Println("Success to build domain " + domain_name)

		return nil
	}

	return fmt.Errorf("failed to create domain because domain is exists")
}

func BuildContent(path, domain_name, content string) error {
	os.MkdirAll(path, 0700)
	if err := ioutil.WriteFile(path+domain_name+".go", []byte(content), 0644); err != nil {
		return err
	}
	return nil
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
