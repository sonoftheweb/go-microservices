package handler

import (
	"context"
	"encoding/json"
	pb "gateway/pb"
	"gateway/utils"
	"net/http"
	"time"

	"google.golang.org/grpc/status"
)

func (h *AuthServiceHandler) HandleAuthentication(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	grpcResp, err := h.AuthServiceClient.Login(ctx, &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		errorStatus, _ := status.FromError(err)
		http.Error(w, errorStatus.Message(), utils.ConvertToHTTPStatusCode(errorStatus.Code()))
		return
	}

	resp := LoginResponse{
		Success: grpcResp.Success,
		Token:   grpcResp.Token,
		Error:   grpcResp.Error,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
