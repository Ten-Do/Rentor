package middleware

import (
	"net/http"
	"rentor/internal/logger"
	"time"

	"github.com/go-chi/chi/middleware"
)

// LoggingMiddleware logging all HTTP requests and responses
func LoggingMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := logger.With(
			logger.Field("component", "middleware/logger"),
		)

		log.Info("logger middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				logger.Field("method", r.Method),
				logger.Field("path", r.URL.Path),
				logger.Field("remote_addr", r.RemoteAddr),
				logger.Field("user_agent", r.UserAgent()),
				logger.Field("request_id", middleware.GetReqID(r.Context())),
			)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				entry.Info("request completed",
					logger.Field("status", ww.Status()),
					logger.Field("bytes", ww.BytesWritten()),
					logger.Field("duration", time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
