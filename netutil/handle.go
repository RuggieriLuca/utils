package netutil

import "net/http"

type HandleWithError func(http.ResponseWriter, *http.Request) ResponseLayout
func (hwe HandleWithError) ServeHTTP(w http.ResponseWriter, r *http.Request){
	response := hwe(w, r)
	Respond(w,response)
}