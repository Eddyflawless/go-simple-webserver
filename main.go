package main

import(
    "fmt" 
    "log"
    "net/http"
	"go-webserver/auth"
	"os/exec"
)

const (
	port = ":8085"
)

type callback func (w http.ResponseWriter, r *http.Request)

func notFoundHttpError( message string, w http.ResponseWriter) { 
	http.Error(w, message, http.StatusNotFound)
}

func badRequestHttpError( message string, w http.ResponseWriter) { 
	http.Error(w, message, http.StatusBadRequest)
}

func checkPath( path string, w http.ResponseWriter, r *http.Request) bool{

	if path != r.URL.Path {
		notFoundHttpError("Path not found", w )
		return false
	}
	return true
}

func formHttpHandler(w http.ResponseWriter, r *http.Request){

	if rs := checkHttpMethod("POST", w, r); !rs {
		return
	}

	file_url := r.FormValue("file_url")

	fmt.Fprintf(w, fmt.Sprintf("Downloading...%v",file_url))

}

func runCFScriptHttpHandler(w http.ResponseWriter, r * http.Request){

	out, err := exec.Command("ls", "-l").Output()

	if err != nil {
        log.Fatal(err)
    }

	fmt.Printf("output: %q\n", string(out))

}

func loginHttpHandler(w http.ResponseWriter, r *http.Request){

	if rs := checkHttpMethod("GET",w,r); rs {

		return

	}

	if rs := checkHttpMethod("POST",w,r); rs {

		return
		
	}

	http.Error(w,"Method not supported",http.StatusForbidden)

}

func checkHttpMethod( method string,  w http.ResponseWriter, r *http.Request) bool{

	if method != r.Method {
		notFoundHttpError("Path not found", w )
		return false
	}
	return true
}

func homeHttpHandler(name string) http.HandlerFunc {

	fn := func(w http.ResponseWriter, r *http.Request){

		if rs := checkHttpMethod("GET", w, r); !rs {
			return
		}
		fmt.Printf("Hello world %v \n",w)
	}

	return fn
}

func checkHeaderStatus(w http.ResponseWriter, r *http.Request){

	badRequestHttpError("Malformed request",w)
}

func guestMiddleware(next http.HandlerFunc, validateFunc callback ) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request){
		//middle ware logic goes here 
		validateFunc(w, r)
		//
		fmt.Printf("middleware logic runned....\n")
		next.ServeHTTP(w,r)
	}
}





func main(){

	auth.SomeMessage()

	mux := http.NewServeMux()

    file_server := http.FileServer(http.Dir("./static"))
	
    mux.Handle("/", file_server)
	mux.HandleFunc("/home", guestMiddleware(homeHttpHandler("hello world"), checkHeaderStatus ) )
	mux.HandleFunc("/download", formHttpHandler )

	log.Printf("Listening on %v \n",port)
    log.Fatal(http.ListenAndServe(port,mux))

}