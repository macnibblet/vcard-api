package endpoints

import (
	"bitbucket.org/llg/vcard"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"net/http"
)

func handleUpload(resp http.ResponseWriter, req *http.Request) {

	uploadedFile, _, err := req.FormFile("card")
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(RESP_MALFORMED_REQUEST)
		return
	}

	card := vcard.VCard{}
	card.ReadFrom(vcard.NewDirectoryInfoReader(uploadedFile))

	name := fmt.Sprintf("./static/%s.vcf", uuid.NewV4())
	file, err := os.Create(name)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write(RESP_INTERNAL_SERVER_ERROR)
		return
	}

	defer file.Close()

	card.WriteTo(vcard.NewDirectoryInfoWriter(file))
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(name[1:]))
}
