package services

import (
	"context"
	"fmt"

	"github.com/riskykurniawan15/xarch/domain/alquran/models"
	"github.com/riskykurniawan15/xarch/domain/alquran/repository"
	"github.com/riskykurniawan15/xarch/helpers/errors"
)

type AlquranService struct {
	AlquranRepo *repository.AlquranRepo
}

func NewAlquranService(Repo *repository.AlquranRepo) *AlquranService {
	return &AlquranService{
		Repo,
	}
}

func (svc *AlquranService) GetListChapter(ctx context.Context) (*[]models.Chapter, *errors.ErrorResponse) {
	data, err := svc.AlquranRepo.GetChapter(ctx)
	if data != nil && err == nil {
		return data, nil
	}

	data, err = svc.AlquranRepo.GetChapterAPi(ctx)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	svc.AlquranRepo.SaveChapter(ctx, data)

	return data, nil
}

func (svc *AlquranService) GetDetailChapter(ctx context.Context, ID int) (*models.Chapter, *errors.ErrorResponse) {
	if ID < 1 || ID > 114 {
		return nil, errors.BadRequest.NewError(fmt.Errorf("chapter min 1 and max 114"))
	}

	data, err := svc.AlquranRepo.GetDetailChapter(ctx, ID)
	if data != nil && err == nil {
		return data, nil
	}

	data, err = svc.AlquranRepo.GetDetailChapterAPi(ctx, ID)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	svc.AlquranRepo.SaveDetailChapter(ctx, ID, data)

	return data, nil
}

func (svc *AlquranService) GetListVerse(ctx context.Context, ID int) (*[]models.ChapterVerse, *errors.ErrorResponse) {
	if ID < 1 || ID > 114 {
		return nil, errors.BadRequest.NewError(fmt.Errorf("chapter min 1 and max 114"))
	}

	data, err := svc.AlquranRepo.GetChapterVerse(ctx, ID)
	if data != nil && err == nil {
		return data, nil
	}

	data, err = svc.AlquranRepo.GetVerseAPi(ctx, ID)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	svc.AlquranRepo.SaveChapterVerse(ctx, ID, data)

	return data, nil
}

func (svc *AlquranService) GetDetailVerse(ctx context.Context, ID, VerseNumber int) (*models.Verse, *errors.ErrorResponse) {
	if ID < 1 || ID > 114 {
		return nil, errors.BadRequest.NewError(fmt.Errorf("chapter min 1 and max 114"))
	}

	data, err := svc.AlquranRepo.GetDetailVerse(ctx, ID, VerseNumber)
	if data != nil && err == nil {
		return data, nil
	}

	data, err = svc.AlquranRepo.GetDetailVerseAPi(ctx, ID, VerseNumber)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	svc.AlquranRepo.SaveDetailVerse(ctx, ID, VerseNumber, data)

	return data, nil
}
