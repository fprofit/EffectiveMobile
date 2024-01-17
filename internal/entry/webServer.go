package entry

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fprofit/EffectiveMobile/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func runServer(handlers *handler.Handler, config EnvConfig, log *logrus.Logger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	if config.LOGLevel == "debug" {
		gin.SetMode(gin.DebugMode)
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.AppPort),
		Handler:        handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info("Server closed")
			} else {
				log.Error("Error occurred while running http server:", err)
			}
		}
	}()

	log.Info("Server started on port: ", s.Addr)
	return s
}

func shutdownServer(s *http.Server, log *logrus.Logger) {
	if s == nil {
		log.Error("Server is nil")
		return
	}
	log.Infoln("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Infoln("Server exiting")
}
