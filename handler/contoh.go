package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func WriteJSONResponse(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonData, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func ContohHandler(w http.ResponseWriter, r *http.Request) {
	zParam := r.PathValue("z")
	var x int = 2
	var y int = 2
	var sum int = x + y
	z, err := strconv.Atoi(zParam)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "z harus diisi dengan angka",
		})
		return
	}
	if z != sum {
		WriteJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Jawaban salah",
		})
		return
	}
	WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Jawaban benar",
	})
}
