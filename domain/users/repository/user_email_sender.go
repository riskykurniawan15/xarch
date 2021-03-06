package repository

import (
	"context"
	"fmt"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/helpers/mail"
)

func (repo UserRepo) EmailForgotSender(ctx context.Context, user *models.User, token *models.UserToken) error {
	body := fmt.Sprintf(`
		Hai, %s <br><br>
		Berikut adalah token <b>%s</b> yang dapat digunakan untuk melakukan pemulihan akun
	`, user.Name, token.Token)
	err := mail.Send(repo.cfg, user.Email, "Lupa Password", body)
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepo) EmailVerfiedSender(ctx context.Context, user *models.User, token *models.UserToken) error {
	link := fmt.Sprintf("%suser/verif/%d/%s",
		repo.cfg.Http.URL,
		user.ID,
		token.Token,
	)

	body := fmt.Sprintf(`
		Hai, %s <br><br>
		Berikut adalah link verifikasi email kamu <a href="%s">%s</a>
	`, user.Name, link, link)
	err := mail.Send(repo.cfg, user.Email, "Verifikasi Email", body)
	if err != nil {
		return err
	}
	return nil
}
