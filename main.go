package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	m := r.MultipartForm
	files := m.File["myFiles"]
	for i := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		}
		f, err := os.Create("./uploads/" + files[i].Filename)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		}
	}
	fmt.Fprintf(w, "%s", "Upload feito com sucesso")

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", upload).Methods("POST")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
