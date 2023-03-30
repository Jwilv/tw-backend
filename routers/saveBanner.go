package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"
)

//modifica el user en la base de datos
// y crea un archivo en la carpeta updates y el path va aser guardado en la base de datos

func SaveBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, errFile := r.FormFile("banner")

	if errFile != nil {
		http.Error(w, "Error al obtener el archivo"+errFile.Error(), http.StatusBadRequest)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]
	var document string = "uploads/banners/" + IDUser + "." + extension

	fileOs, errOs := os.OpenFile(document, os.O_WRONLY|os.O_CREATE, 0666)

	if errOs != nil {
		http.Error(w, "Error al subir la imagen ! "+errOs.Error(), http.StatusBadRequest)
		return
	}

	_, errCopy := io.Copy(fileOs, file)

	if errCopy != nil {
		http.Error(w, "Error al copiar la img"+errCopy.Error(), http.StatusBadRequest)
		return
	}

	var user models.User

	user.Banner = IDUser + "." + extension

	status, errDb := db.ModifyRegister(user, IDUser)

	if errDb != nil || !status {
		http.Error(w, "Error al grabar en la base de datos"+errDb.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
