#Fortis IT - Kurs Devops  https://www.youtube.com/@fortisit5443
#Server przykładowy w GOLANG
package main

import (
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, req *http.Request) {
	result := "Fortis IT - Szkolenia! - " + os.Getenv("VAR") +"\n"
	io.WriteString(w, result)
}

func main(){
	http.HandleFunc("/", getRoot)
	http.ListenAndServe(":8080", nil)
}
