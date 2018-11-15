package service

import (
    "net/http"

    "github.com/unrolled/render"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        formatter.HTML(w, http.StatusOK, "index", struct {
            Name		string `json:"name"`
            Day		 	string `json:"day"`
        }{Name: "T-Machine", Day: "Friday"})
    }
}
