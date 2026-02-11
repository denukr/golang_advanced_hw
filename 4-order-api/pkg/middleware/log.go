package middleware

import (
	"golang-adv/4-order-api/pkg/logs"
	"net/http"

	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	Logger *logs.Logger
}

func NewLogMiddleware() *LogMiddleware {
	l := logs.NewLogger()
	return &LogMiddleware{
		Logger: l,
	}
}

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrapperWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (lm *LogMiddleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wr := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		lm.Logger.WithFields(logrus.Fields{
			"Method":     r.Method,
			"Host":       r.Host,
			"Path":       r.URL.Path,
			"Stats code": wr.StatusCode,
		}).Info("Request successful")
		next.ServeHTTP(wr, r)
	})
}
