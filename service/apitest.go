package service

import (
    "net/http"
	"time"
    "github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        formatter.JSON(w, http.StatusOK, struct {
            Name      	string `json:"name"`
			Date		string	`json:"date"`
        }{Name: "T-Machine", Date: time.Now().Format("2006-01-02 15:04")})
    }
}
