package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain/alquran/models"
)

type AlquranRepo struct {
	cfg config.Config
	RDB *redis.Client
}

func NewAlquranRepo(cfg config.Config, RDB *redis.Client) *AlquranRepo {
	return &AlquranRepo{
		cfg,
		RDB,
	}
}

func (Repo *AlquranRepo) GetChapter(ctx context.Context) (*[]models.Chapter, error) {
	res, err := Repo.RDB.Get(ctx, "chapter").Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		panic(err)
	}

	var result []models.Chapter
	err = json.Unmarshal([]byte(string(res)), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (Repo *AlquranRepo) SaveChapter(ctx context.Context, Chapter *[]models.Chapter) error {
	value, err := json.Marshal(Chapter)
	if err != nil {
		return err
	}

	err = Repo.RDB.Set(ctx, "chapter", string(value), 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func (Repo *AlquranRepo) GetDetailChapter(ctx context.Context, ID int) (*models.Chapter, error) {
	res, err := Repo.RDB.Get(ctx, "chapter_"+fmt.Sprint(ID)).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		panic(err)
	}

	var result models.Chapter
	err = json.Unmarshal([]byte(string(res)), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (Repo *AlquranRepo) SaveDetailChapter(ctx context.Context, ID int, Chapter *models.Chapter) error {
	value, err := json.Marshal(Chapter)
	if err != nil {
		return err
	}

	err = Repo.RDB.Set(ctx, "chapter_"+fmt.Sprint(ID), string(value), 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}
