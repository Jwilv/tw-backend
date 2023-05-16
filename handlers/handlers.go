package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Jwilv/tw-backend/middlewares"
	"github.com/Jwilv/tw-backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Drivers manejo de seteo y ListenServe
func Drivers() {

	//Router objet
	router := mux.NewRouter()
	//route register, crea user en la base de datos
	router.HandleFunc("/register", middlewares.CheckDb(routers.Register)).Methods("POST")
	//route login, se encarga de logear la data contra la base de datos
	router.HandleFunc("/login", middlewares.CheckDb(routers.Login)).Methods("POST")
	//obtiene un id y devuelve un modelo user que se encuentre en la base de datos. que contenga la id otorgada
	router.HandleFunc("/getprofile", middlewares.CheckDb(middlewares.ValidateJwt(routers.GetProfile))).Methods("GET")
	//obtiene data y la pasa a un modelo de user, para despues buscarlo en la base de datos y modificarlo
	router.HandleFunc("/changeProfile", middlewares.CheckDb(middlewares.ValidateJwt(routers.ChangeProfile))).Methods("PUT")
	//obtiene la data y la pasa a un modelo de note para despues subirlo a la base de datos
	router.HandleFunc("/saveNote", middlewares.CheckDb(middlewares.ValidateJwt(routers.SaveNote))).Methods("POST")
	//se obtiene las notas de la base de datos y se le pasa a un arreglo, se obtiene todas las notas del user.
	// ya que se envia el id de dicho user
	router.HandleFunc("/getNotes", middlewares.CheckDb(middlewares.ValidateJwt(routers.GetNotes))).Methods("GET")
	//route removeNote, obtiene el id del user por el jwt y se le envia el id de la nota a eliminar.
	//y la remueve de la base de datos
	router.HandleFunc("/removeNote", middlewares.CheckDb(middlewares.ValidateJwt(routers.RemoveNote))).Methods("DELETE")
	//route updateAvatar, se encarga de recibir un archivo, crea un archivo en local y modifica o agrega
	//el path a la base de datos en el campo avatar
	router.HandleFunc("/updateAvatar", middlewares.CheckDb(middlewares.ValidateJwt(routers.SaveAvatar))).Methods("POST")
	//route updateBanner, se encarga de recibir un archivo, crea un archivo en local y modifica o agrega
	//el path a la base de datos en el campo banner
	router.HandleFunc("/updateBanner", middlewares.CheckDb(middlewares.ValidateJwt(routers.SaveBanner))).Methods("POST")
	//obtiene el path de la base de datos del campo banner y busca el archivo local para copiarlo
	//en la respuesta y asi enviar el archivo
	router.HandleFunc("/getBanner", middlewares.CheckDb(routers.GetBanner)).Methods("GET")
	//obtiene el path de la base de datos del campo  avatar y busca el archivo local para copiarlo
	//en la respuesta y asi enviar el archivo
	router.HandleFunc("/getAvatar", middlewares.CheckDb(routers.GetAvatar)).Methods("GET")

	//route relationUsers, inserta la relacion entre 2 usuarios y la guarda en la base de datos
	router.HandleFunc("/relationUsers", middlewares.CheckDb(middlewares.ValidateJwt(routers.Relation))).Methods("POST")

	//elimina la relacion con el user enviado mediante el id
	router.HandleFunc("/deleteRelation", middlewares.CheckDb(middlewares.ValidateJwt(routers.DeleteRelation))).Methods("DELETE")

	//check de relacion existente
	router.HandleFunc("/checkRelation", middlewares.CheckDb(middlewares.ValidateJwt(routers.GetRelation))).Methods("GET")

	//retorna un listado de usuarios, dependiendo del tipo, los new o los follow
	router.HandleFunc("/getUsers", middlewares.CheckDb(middlewares.ValidateJwt(routers.UsersList))).Methods("GET")

	//se le envia el numero de la pagina y retorna los notas de los usuarios que el usuario de la cuenta sigue
	router.HandleFunc("/notesFollow", middlewares.CheckDb(middlewares.ValidateJwt(routers.GetNotesFollow))).Methods("GET")

	//renueva el token
	router.HandleFunc("/renew", middlewares.CheckDb(middlewares.ValidateJwt(routers.RenewToken))).Methods("GET")

	router.HandleFunc("/notesrandom", middlewares.CheckDb(routers.GetRandomNotes)).Methods("GET")

	router.HandleFunc("/inadmin", middlewares.CheckDb(routers.LoginAdminUser)).Methods("POST")

	router.HandleFunc("/upadmin", middlewares.CheckDb(routers.RegisterAdminUser)).Methods("POST")

	router.HandleFunc("/msgadm", middlewares.CheckDb(middlewares.ValidateJwt(routers.SaveMsgAdmin))).Methods("POST")

	router.HandleFunc("/removeMsg", middlewares.CheckDb(middlewares.ValidateJwt(routers.RemoveMsg))).Methods("DELETE")

	router.HandleFunc("/getMsg", middlewares.CheckDb(middlewares.ValidateJwt(routers.GetMsg))).Methods("GET")



	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
