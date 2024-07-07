package apierrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/v420v/qrmarkapi/common"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *APIError
	if !errors.As(err, &appErr) {
		appErr = &APIError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := common.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int
	switch appErr.ErrCode {
	case ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	case Unauthorizated:
		statusCode = http.StatusUnauthorized
	case NotMatchUser:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
