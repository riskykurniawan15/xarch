package engine

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/interfaces/http/routers"
	"github.com/riskykurniawan15/xarch/logger"
)

func Start(wg *sync.WaitGroup, cfg config.Config, log logger.Logger, svc *domain.Service) {
	e := routers.Routers(svc, cfg, log)

	go func() {
		e.HideBanner = true
		if err := e.Start(fmt.Sprintf("%s:%d", cfg.Http.Server, cfg.Http.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
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
	} else {
		log.Info("shutting down the server")
	}

	wg.Done()
}
