package router

import (
	"main/controllers"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	route := strings.Trim(r.URL.Path, "/")

	switch route {
	case "":
		w.Write([]byte("Welcome to the homepage!"))
	case "setcache":
		controllers.SetCache(w, r)		
	case "getcache":
		controllers.GetCache(w, r)		
	case "deletecache":
		controllers.DeleteCache(w, r)
	case "setcapacity":
		controllers.SetCapacity(w, r)
	}
}
