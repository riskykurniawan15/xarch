package repository

import (
	"context"
	"mime/multipart"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/riskykurniawan15/xarch/helpers/md5"
)

var (
	folder string = "xarch/user_profile"
)

func (repo UserRepo) StoreFile(ctx context.Context, file multipart.File) (string, error) {
	filename := md5.Encrypt(time.Now().String())
	cld, _ := cloudinary.NewFromParams(repo.cfg.Cloudinary.CLOUD_NAME, repo.cfg.Cloudinary.API_KEY, repo.cfg.Cloudinary.API_SECRET)
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: filename,
		Folder:   folder,
	})
	return filename + "." + resp.Format, err
}

func (repo UserRepo) DestroyFile(ctx context.Context, filename string) error {
	file := strings.Split(filename, ".")
	cld, _ := cloudinary.NewFromParams(repo.cfg.Cloudinary.CLOUD_NAME, repo.cfg.Cloudinary.API_KEY, repo.cfg.Cloudinary.API_SECRET)
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: folder + "/" + file[0],
	})
	return err
}

func (repo UserRepo) RetrieveFile(ctx context.Context, filename string) (*admin.AssetResult, error) {
	file := strings.Split(filename, ".")
	cld, _ := cloudinary.NewFromParams(repo.cfg.Cloudinary.CLOUD_NAME, repo.cfg.Cloudinary.API_KEY, repo.cfg.Cloudinary.API_SECRET)
	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{
		PublicID: folder + "/" + file[0],
	})

	return resp, err
}
