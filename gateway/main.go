package main

import (
	"context"
	"encoding/json"
	pb "gateway/pb"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
}

var authServiceClient pb.AuthServiceClient

func main() {
	r := chi.NewRouter()

	// Set up a connection to the gRPC server when the application starts.
	conn, err := grpc.Dial(
		"auth:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth service: %v", err)
	}
	defer conn.Close()

	authServiceClient = pb.NewAuthServiceClient(conn)

	// Define your route and handler
	r.Post("/api/auth/login", loginHandler)

	// Start the HTTP server
	log.Println("api gateway running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	// Create a new context, with a timeout based on the HTTP request context
	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	grpcResp, err := authServiceClient.Login(ctx, &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		errorStatus, _ := status.FromError(err)
		http.Error(w, errorStatus.Message(), convertToHTTPStatusCode(errorStatus.Code()))
		return
	}

	// Map the gRPC response to the REST response
	resp := LoginResponse{
		Success: grpcResp.Success,
		Token:   grpcResp.Token,
		Error:   grpcResp.Error,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// This function converts a gRPC status code to the corresponding HTTP status code.
func convertToHTTPStatusCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
