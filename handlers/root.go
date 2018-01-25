package handlers

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ydl/static/client/dist/", http.StatusMovedPermanently)
}
