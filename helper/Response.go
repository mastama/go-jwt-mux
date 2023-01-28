package helper

import (
	"encoding/json"
	"net/http"
)

// type Response struct {
// 	Meta Meta        `json:"meta"`
// 	Data interface{} `json:data`
// }

// type Meta struct {
// 	Message string `json:"message"`
// 	Code    int    `json:"code"`
// 	Status  string `json:"status"`
// }

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
