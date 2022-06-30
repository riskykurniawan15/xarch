package repository

import (
	"context"
	"fmt"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/helpers/mail"
)

func (repo UserRepo) EmailVerfiedSender(ctx context.Context, user *models.User, token *models.UserToken) error {
	link := fmt.Sprintf("%suser/verif/%s",
		repo.cfg.Http.URL,
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
