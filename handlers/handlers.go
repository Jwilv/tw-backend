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
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDb(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDb(routers.Login)).Methods("POST")
	router.HandleFunc("/getprofile", middlewares.CheckDb(middlewares.ValidateJwt(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/changeProfile", middlewares.CheckDb(middlewares.ValidateJwt(routers.ChangeProfile))).Methods("PUT")
	router.HandleFunc("/saveNote", middlewares.CheckDb(middlewares.ValidateJwt(routers.SaveNote))).Methods("POST")
	router.HandleFunc("/getNotes", middlewares.CheckDb(middlewares.ValidateJwt(routers.GetNotes))).Methods("GET")
	router.HandleFunc("/removeNote", middlewares.CheckDb(middlewares.ValidateJwt(routers.RemoveNote))).Methods("DELETE")
	router.HandleFunc("/updateAvatar", middlewares.CheckDb(middlewares.ValidateJwt(routers.SaveAvatar))).Methods("POST")
	router.HandleFunc("/updateBanner", middlewares.CheckDb(middlewares.ValidateJwt(routers.SaveBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlewares.CheckDb(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/getAvatar", middlewares.CheckDb(routers.GetAvatar)).Methods("GET")



	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
