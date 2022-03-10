package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/api/v1/health", health)
	http.HandleFunc("/api/v1/info", info)
	log.Println("http://127.0.0.1:8080/api/v1/health")
	log.Println("http://127.0.0.1:8080/api/v1/info")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func health(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("ok"))
}

func info(w http.ResponseWriter, r *http.Request)  {
	res := map[string]interface{}{
		"code": 0,
		"msg": "success",
		"data": "this is data field",
	}
	bytes, err := json.Marshal(res)
	if err !=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}