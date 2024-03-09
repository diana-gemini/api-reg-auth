package handler

import (
	"forum/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// func (h *Handler) getUserFromContext(r *http.Request) *types.User {
// 	user, ok := r.Context().Value(ctxKey).(*types.User)

// 	if !ok {
// 		fmt.Printf("Error Context")
// 		return nil
// 	}
// 	return user
// }
