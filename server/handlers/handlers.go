package handlers

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	resp := []byte("ok")

	w.Write(resp)
}

func HealthAgni(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	resp := []byte("agni")

	w.Write(resp)
}
