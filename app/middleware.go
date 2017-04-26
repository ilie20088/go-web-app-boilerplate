package app

import (
	"io/ioutil"
	"net/http"

	"github.com/ilie20088/go-web-app-boilerplate/utils"
	"go.uber.org/zap"
)

// LoggingMiddleware logs the request after calling h.ServeHTTP
func LoggingMiddleware(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		if utils.LogLevel == zap.DebugLevel {
			reqBody, _ := ioutil.ReadAll(r.Body)
			utils.Logger.Debug(
				"HTTP Request",
				zap.String("method", r.Method),
				zap.String("user_agent", r.UserAgent()),
				zap.String("content_type", r.Header.Get("Content-Type")),
				zap.String("url_path", r.URL.Path),
				zap.String("url_params", r.URL.RawQuery),
				zap.String("request_body", string(reqBody)),
			)
			defer r.Body.Close()
		}
	})
}

func init() {
}
