package user

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(
	w http.ResponseWriter, 
	r *http.Request,
) {
    if r.Method != http.MethodPost {
        http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
        return
    }

    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(
			w,
			"Invalid JSON: " + err.Error(),
			http.StatusBadRequest,
		)
        return
    }

    userResponse, err := h.service.Register(r.Context(), &req)
    if err != nil {
        handleError(w, err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(userResponse); 
		err != nil {
        http.Error(
			w,
			"Failed to encode response",
		    http.StatusInternalServerError,
		)
    }

}

func (h *Handler) GetUser(
	w http.ResponseWriter,
	r *http.Request,
) {
    if r.Method != http.MethodGet {
        http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
        return
    }

    path := strings.Trim(r.URL.Path, "/")
    segments := strings.Split(path, "/")

    if len(segments) < 4 {
        http.Error(
			w,
			"Invalid URL path",
			http.StatusBadRequest,
		)
        return
    }

    userID := segments[3]

    userResponse, err := h.service.GetUser(
						 	r.Context(),
						 	userID,
							)
    if err != nil {
        handleError(w, err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(userResponse); 
		err != nil {
        http.Error(
			w,
			"Failed to encode response",
			 http.StatusInternalServerError,
		)
    }
}

func (h *Handler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
    path := strings.Trim(r.URL.Path, "/")
    segments := strings.Split(path, "/")

    if len(segments) < 3 ||
		segments[0] != "api" ||
		segments[1] != "v1"||
		segments[2] != "users" {

		http.Error(w, "Not found", http.StatusNotFound)
        return
    }
 
    if len(segments) == 4 &&
		segments[3] == "register" &&
		r.Method == http.MethodPost {
        h.Register(w, r)
        return
    }

    if len(segments) == 4 && r.Method == http.MethodGet {
        h.GetUser(w, r)
        return
    }

    http.Error(w, "Not found", http.StatusNotFound)
}

func handleError(w http.ResponseWriter, err error) {
    http.Error(
		w,
		err.Error(),
		http.StatusInternalServerError,
	)
}














































