package main

import (
	"bitbucket.org/llg/vcard"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"os"
)

func handleUpload(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	uploadedFile, _, err := req.FormFile("card")
	if err != nil {
		panic(err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	card := vcard.VCard{}
	card.ReadFrom(vcard.NewDirectoryInfoReader(uploadedFile))

	name := fmt.Sprintf("./cards/%s.vcf", uuid.NewV4())
	file, err := os.Create(name)
	if err != nil {
		panic(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer file.Close()

	card.WriteTo(vcard.NewDirectoryInfoWriter(file))
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(name[1:]))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	static := http.FileServer(http.Dir("./cards"))

	http.HandleFunc("/upload", handleUpload)
	http.Handle("/cards/", http.StripPrefix("/cards/", static))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
