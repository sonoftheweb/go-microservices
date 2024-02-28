package handler

import (
	"context"
	"encoding/json"
	"gateway/pb"
	"gateway/utils"
	"net/http"
	"time"

	"google.golang.org/grpc/status"
)

func (h *AuthServiceHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	grpcResp, err := h.AuthServiceClient.Register(ctx, &pb.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	})
	if err != nil {
		errorStatus, _ := status.FromError(err)
		http.Error(w, errorStatus.Message(), utils.ConvertToHTTPStatusCode(errorStatus.Code()))
		return
	}

	resp := RegisterResponse{
		Success: grpcResp.Success,
		Error:   grpcResp.Error,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
