package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/riskykurniawan15/xarch/domain/alquran/models"
)

func (Repo *AlquranRepo) GetChapterAPi(ctx context.Context) (*[]models.Chapter, error) {
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

func (Repo *AlquranRepo) GetDetailChapterAPi(ctx context.Context, ID int) (*models.Chapter, error) {
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

func (Repo *AlquranRepo) GetVerseAPi(ctx context.Context, ID int) (*[]models.ChapterVerse, error) {
	response, err := http.Get(Repo.cfg.OTHER.AlQuranAPI + "chapter/" + fmt.Sprint(ID) + "/verse")
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

	var result []models.ChapterVerse
	err = json.Unmarshal([]byte(string(responseData)), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (Repo *AlquranRepo) GetDetailVerseAPi(ctx context.Context, ID, VerseNumber int) (*models.Verse, error) {
	response, err := http.Get(Repo.cfg.OTHER.AlQuranAPI + "chapter/" + fmt.Sprint(ID) + "/verse/" + fmt.Sprint(VerseNumber))
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

	var result []models.Verse
	err = json.Unmarshal([]byte(string(responseData)), &result)
	if err != nil {
		return nil, err
	}

	return &result[0], nil
}
