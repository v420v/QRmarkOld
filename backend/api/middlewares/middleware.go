package middlewares

import "github.com/v420v/qrmarkapi/services"

type QrmarkAPIMiddleware struct {
	service *services.QrmarkAPIService
}

func NewMiddleware(service *services.QrmarkAPIService) *QrmarkAPIMiddleware {
	return &QrmarkAPIMiddleware{
		service: service,
	}
}
