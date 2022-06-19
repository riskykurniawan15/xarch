package services

import (
	"context"

	"github.com/riskykurniawan15/xarch/domain/alquran/models"
	"github.com/riskykurniawan15/xarch/domain/alquran/repository"
)

type AlquranService struct {
	AlquranRepo *repository.AlquranRepo
}

func NewAlquranService(Repo *repository.AlquranRepo) *AlquranService {
	return &AlquranService{
		Repo,
	}
}

func (svc *AlquranService) GetListChapter(ctx context.Context) (*[]models.Chapter, error) {
	data, err := svc.AlquranRepo.GetChapter(ctx)
	if err != nil {
		return nil, err
	}

	if data != nil {
		return data, nil
	}

	data, err = svc.AlquranRepo.GetChapterAPi(ctx)
	if err != nil {
		return nil, err
	}

	err = svc.AlquranRepo.SaveChapter(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *AlquranService) GetDetailChapter(ctx context.Context, ID int) (*models.Chapter, error) {
	data, err := svc.AlquranRepo.GetDetailChapter(ctx, ID)
	if err != nil {
		return nil, err
	}

	if data != nil {
		return data, nil
	}

	data, err = svc.AlquranRepo.GetDetailChapterAPi(ctx, ID)
	if err != nil {
		return nil, err
	}

	err = svc.AlquranRepo.SaveDetailChapter(ctx, ID, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *AlquranService) GetListVerse(ctx context.Context, ID int) (*[]models.ChapterVerse, error) {
	data, err := svc.AlquranRepo.GetChapterVerse(ctx, ID)
	if err != nil {
		return nil, err
	}

	if data != nil {
		return data, nil
	}

	data, err = svc.AlquranRepo.GetVerseAPi(ctx, ID)
	if err != nil {
		return nil, err
	}

	err = svc.AlquranRepo.SaveChapterVerse(ctx, ID, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
