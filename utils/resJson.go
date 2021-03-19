package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// ubahKeByte, err := json.Marshal(data)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }

	// w.Write([]byte(ubahKeByte))

	json.NewEncoder(w).Encode(data)
}

func IsError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return true
	}
	return false
}
