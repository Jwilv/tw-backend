package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Jwilv/tw-backend/db"

)

//nos permite enviar el archivo 
func GetBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "error, no se encontro un id en el url", http.StatusBadRequest)
		return
	}

	profile, errProfile := db.SearchProfile(ID)

	if errProfile != nil {
		http.Error(w, "usuario no encontrado"+errProfile.Error(), http.StatusBadRequest)
		return
	}

	OpenFile, errFile := os.Open("uploads/banners/" + profile.Banner)

	if errFile != nil {
		http.Error(w, "imagen no encontrada"+errFile.Error(), http.StatusBadRequest)
		return
	}

	_, errCopy := io.Copy(w, OpenFile)

	if errCopy != nil{
	http.Error(w,"no se a podido enviar la imagen " + errCopy.Error(), http.StatusBadRequest)
	return
	}

}
