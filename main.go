package main

import(
    "fmt" 
    "log"
    "net/http"
)

func notFoundHttpError( message string, w http.ResponseWriter) { 
	http.Error(w, message, http.StatusNotFound)
}

func checkPath( path string, w http.ResponseWriter, r *http.Request) bool{

	if path != r.URL.Path {
		notFoundHttpError("Path not found", w )
		return false
	}
	return true
}

func formHttpHandler(w http.ResponseWriter, r *http.Request){

	if rs := checkPath("download", w, r); !rs {
		return
	} 

	if rs := checkHttpMethod("POST", w, r); !rs {
		return
	}

	file_url := r.FormValue("file_url")

	fmt.Fprintf(w, fmt.Sprintf("Downloading...%v",file_url))

}

func checkHttpMethod( method string,  w http.ResponseWriter, r *http.Request) bool{

	// if rs := checkPath("", w, r); !rs {
	// 	return
	// } 

	if method != r.Method {
		notFoundHttpError("Path not found", w )
		return false
	}
	return true
}


func homeHttpHandler(w http.ResponseWriter, r *http.Request){

	if rs := checkHttpMethod("GET", w, r); !rs {
		return
	}
    fmt.Printf("Hello world %v \n",w)

}


func main(){

    file_server := http.FileServer(http.Dir("./static"))
    http.Handle("/", file_server)
	http.HandleFunc("/home", homeHttpHandler)
	http.HandleFunc("/download", formHttpHandler)

    log.Fatal(http.ListenAndServe(":8085",nil))

}