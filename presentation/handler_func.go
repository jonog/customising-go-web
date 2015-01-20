

import "net/http"

// START OMIT
func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// END OMIT