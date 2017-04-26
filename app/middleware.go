package app

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/ilie20088/go-web-app-boilerplate/utils"
	"github.com/newrelic/go-agent"
	"go.uber.org/zap"
)

var authURL = "http://www.google.com"

// AuthMiddleware makes REST call to external
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := authenticate(w); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func authenticate(w http.ResponseWriter) error {
	if txn, ok := w.(newrelic.Transaction); ok {
		utils.Logger.Info("Starting external segment...")
		defer authSegment(txn).End()
	}

	response, err := http.Get(authURL)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return errors.New("Not authorized to perform this operation")
	}

	return nil
}

func authSegment(txn newrelic.Transaction) newrelic.ExternalSegment {
	return newrelic.ExternalSegment{
		StartTime: txn.StartSegmentNow(),
		URL:       authURL,
	}
}

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
