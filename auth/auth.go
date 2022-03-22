package auth

import(
	"github.com/gorilla/sessions"
	"fmt"
	"net/http"
)

var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func GetCookie(name string) string{

	return ""
}

func IsUserAuthenticated(w http.ResponseWriter, r *http.Request){

	session, _ := store.Get(r, "cookie-name")

	// check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	// Print secret messages
	fmt.Fprintf(w, "The cake is a lie")

}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r,w)

}

func SomeMessage(){
	fmt.Printf("some message displayed \n")
}