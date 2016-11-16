package endpoints

import (
	"bitbucket.org/llg/vcard"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"net/http"
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
