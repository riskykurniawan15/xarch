package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/riskykurniawan15/xarch/domain/alquran/models"
)

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

func (Repo *AlquranRepo) GetChapterVerse(ctx context.Context, ID int) (*[]models.ChapterVerse, error) {
	res, err := Repo.RDB.Get(ctx, "chapter_"+fmt.Sprint(ID)+"_verse").Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		panic(err)
	}

	var result []models.ChapterVerse
	err = json.Unmarshal([]byte(string(res)), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (Repo *AlquranRepo) SaveChapterVerse(ctx context.Context, ID int, ChapterVerse *[]models.ChapterVerse) error {
	value, err := json.Marshal(ChapterVerse)
	if err != nil {
		return err
	}

	err = Repo.RDB.Set(ctx, "chapter_"+fmt.Sprint(ID)+"_verse", string(value), 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}
