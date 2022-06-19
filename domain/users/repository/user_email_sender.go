package repository

import (
	"context"
	"fmt"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/helpers/mail"
)

func (repo UserRepo) EmailVerfiedSender(ctx context.Context, user *models.User, token *models.UserToken) error {
	link := fmt.Sprintf("%s:%d/user/verif/%s",
		repo.cfg.Http.Server,
		repo.cfg.Http.Port,
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
