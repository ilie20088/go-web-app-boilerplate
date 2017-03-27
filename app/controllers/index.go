package controllers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("adsasdasdasd"))
}
