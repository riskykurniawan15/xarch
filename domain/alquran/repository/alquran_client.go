package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/riskykurniawan15/xarch/domain/alquran/models"
)

func (Repo *AlquranRepo) GetChapterAPi() (*[]models.Chapter, error) {
	response, err := http.Get(Repo.cfg.OTHER.AlQuranAPI + "chapter")
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed retrieve data, status code not 200")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result []models.Chapter
	err = json.Unmarshal([]byte(string(responseData)), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (Repo *AlquranRepo) GetDetailChapterAPi(ID int) (*models.Chapter, error) {
	response, err := http.Get(Repo.cfg.OTHER.AlQuranAPI + "chapter/" + fmt.Sprint(ID))
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed retrieve data, status code not 200")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result []models.Chapter
	err = json.Unmarshal([]byte(string(responseData)), &result)
	if err != nil {
		return nil, err
	}

	return &result[0], nil
}