package router

import (
	"main/controllers"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	route := strings.Trim(r.URL.Path, "/")
        msg:="Please provide a valid method."
	
	switch route {
	case "":
		w.Write([]byte("Welcome to the homepage!"))
	case "setcache":
		if r.Method =="POST"{
		   controllers.SetCache(w, r)
		}else{
                 controllers.ResponseGenerateJson(w,r,"Failure",msg,"")
		}
				
	case "getcache":
		if r.Method =="GET"{
		   controllers.GetCache(w, r)
		}else{
                 controllers.ResponseGenerateJson(w,r,"Failure",msg,"")
		}
				
	case "deletecache":
		if r.Method =="DELETE"{
		   controllers.DeleteCache(w, r)
		}else{
                controllers.ResponseGenerateJson(w,r,"Failure",msg,"")
		}
		
	case "setcapacity":
		if r.Method =="POST"{
		   controllers.SetCapacity(w, r)
		}else{
                controllers.ResponseGenerateJson(w,r,"Failure",msg,"")
		}
		}
}
