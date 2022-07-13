package repository

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (Repo *UserRepo) GetFile(ctx context.Context, url string) ([]byte, error) {
	response, err := http.Get(url)
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

	return responseData, nil
}
