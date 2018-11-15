package service

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
//codegangsta/negroni
func NewServer() *negroni.Negroni {
	//unrolled/render
	formatter := render.New(render.Options {
		Directory:  "templates",
        Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	//gorilla/mux
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}


	
	//js
	mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	//template
	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	//form
	mx.HandleFunc("/", Submit).Methods("POST")
	//unknown 501
	mx.HandleFunc("/unknown", NotImplemented)
	//static file
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
	
}