package middlewares

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/v420v/qrmarkapi/common"
)

var (
	logNo int
	mu    sync.Mutex
)

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1

	mu.Unlock()

	return no
}

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func (m *QrmarkAPIMiddleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		origin := req.Header.Get("Origin")
		if origin == "http://ibukiqrmark.com" || origin == "https://ibukiqrmark.com" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else if origin == "http://127.0.0.1" {
			w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "https://ibukiqrmark.com")
		}
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Origin, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		traceID := newTraceID()

		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			log.Printf("[%d]response to Preflight Request", traceID)
			return
		}

		log.Printf("[%d]%s %s", traceID, req.RequestURI, req.Method)

		ctx := req.Context()
		ctx = context.WithValue(ctx, common.TraceIDKey{}, traceID)
		req = req.WithContext(ctx)

		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}
