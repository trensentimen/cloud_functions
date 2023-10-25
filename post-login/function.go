package trensen_login

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	be_trensentimen "github.com/trensentimen/be_trensen"
)

func init() {
	functions.HTTP("post_login", TrensentimenLogin)
}

func TrensentimenLogin(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://jscroot.github.io")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://jscroot.github.io")
	fmt.Fprintf(w, be_trensentimen.GCFPostHandler("PASETOPRIVATEKEY", "MONGOSTRING", "trensentimen", "user", r))

}
