package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// CustomResponseWriter - обертка над http.ResponseWriter.
type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

// WriteHeader сохраняет статус ответа.
func (crw *CustomResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

// Write сохраняет размер ответа.
func (crw *CustomResponseWriter) Write(data []byte) (int, error) {
	size, err := crw.ResponseWriter.Write(data)
	crw.size += size

	return size, err
}

// NewResponseWriter - функция для создания нового CustomResponseWriter.
func NewResponseWriter(w http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{w, http.StatusOK, 0} // по умолчанию статус 200 OK
}

// GetLoggerMiddleware - Получение обертки для логирования http запросов.
func GetLoggerMiddleware(externalLogger *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(writer http.ResponseWriter, req *http.Request) {
			start := time.Now()
			wrappedWriter := NewResponseWriter(writer)

			// Логируем по окончанию выполнения ServeHTTP
			defer func() {
				// Извлекаем информацию для логирования
				duration := time.Since(start)
				statusCode := wrappedWriter.statusCode
				size := wrappedWriter.size
				clientIP := req.RemoteAddr
				method := req.Method
				url := req.URL.String()

				externalLogger.Infof(
					"\"%s %s\" from %s - %d %dB in %s",
					method, url, clientIP, statusCode, size, duration,
				)
			}()

			next.ServeHTTP(wrappedWriter, req)
		}

		return http.HandlerFunc(fn)
	}
}
