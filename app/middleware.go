package app

import (
	"github.com/ilie20088/go-web-app-boilerplate/utils"
	"net/http"
	"go.uber.org/zap"
)

func LoggingMiddleWare(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//r.Context().
		//utils.ZapLogger.Core().

		//utils.Logger.Debug("A request is done", r.URL.String())
		h.ServeHTTP(w, r)
		utils.ZapLogger.Info("Request intro", zap.String("URL", r.URL.String()))
		utils.ZapLogger.Info("Request out")
	})
}


func init() {
	utils.Logger.Info("Middleware")
}