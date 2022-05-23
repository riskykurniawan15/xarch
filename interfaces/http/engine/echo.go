package engine

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/interfaces/http/routers"
)

func Start(cfg config.Config, svc *domain.Service) {
	e := routers.Routers(svc)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", cfg.Http.Server, cfg.Http.Port)))

	go func() {
		if err := e.Start(fmt.Sprintf("%s:%d", cfg.Http.Server, cfg.Http.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
