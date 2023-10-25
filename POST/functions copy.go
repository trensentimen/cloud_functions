package trensentimen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	be_trensentimen "github.com/trensentimen/be_trensentimen"
)

func init() {
	functions.HTTP("Trensentimen2", TrensentimenPost2)
}

func TrensentimenPost2(w http.ResponseWriter, r *http.Request) {
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

	// response := be_trensentimen.GCFPostHandler("PASETOPRIVATEKEY", "MONGOSTRING", "trensentimen", "user", r)
	// fmt.Println(w, response)

	response := be_trensentimen.GCFPostHandler("PASETOPRIVATEKEY", "MONGOSTRING", "trensentimen", "user", r)

	// Mengirimkan respons JSON kepada klien
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

	// a := `&{0xc000368bd0 0xc0000e2500 0xc0000850c0 0x4ca660 false false false false {{} 0} {0 0} 0xc000085100 {0xc00036e1c0 map[] false false} map[Access-Control-Allow-Origin:[https://trensentimen.github.io]] true 0 -1 0 false false false [] {{} 0} [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc0003f80e0 {{} 0}} {"status":true,"token":"v4.public.eyJleHAiOiIyMDIzLTEwLTE4VDEyOjI4OjU0WiIsImlhdCI6IjIwMjMtMTAtMThUMTA6Mjg6NTRaIiwiaWQiOiJkYW5pIiwibmJmIjoiMjAyMy0xMC0xOFQxMDoyODo1NFoifeXUSmW7z6AuE0cQ46ygS4Lf-66BEjJ-4QhkLhae0iin-DB_K8mRnaONo1c7EwCxh4snrJluUuNT21ryXV2Cpw4","message":"Selamat Datang"}"timestamp: "2023-10-18T10:28:54.361108Z`

	type ResponseJson struct {
		Status  bool   `json:"status"`
		Token   string `json:"token"`
		Message string `json:"message"`
	}

	var responseJson ResponseJson

	// Mengurai JSON ke dalam struktur data Response
	if err := json.Unmarshal([]byte(response), &response); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Mengakses token dari struktur Response
	token := responseJson.Token

	// Set Cookie
	expiration := time.Now().Add(24 * time.Hour) // Ganti sesuai kebutuhan
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiration,
		Secure:   true, // Hanya melalui HTTPS
		HttpOnly: true, // Tidak dapat diakses melalui JavaScript
		SameSite: http.SameSiteStrictMode,
		Path:     "/", // Sesuaikan dengan path yang sesuai
	}
	http.SetCookie(w, &cookie) // Mengirim cookie ke peramban klien

}
