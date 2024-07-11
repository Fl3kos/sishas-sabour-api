package main

import (
	"ShishaSabourAPI/consts"
	restHandlers "ShishaSabourAPI/internal/handlers"
	"ShishaSabourAPI/internal/logger"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	logger.InitLogger()

	router := mux.NewRouter()

	r := restHandlers.NewTobbacoHandler()

	router.HandleFunc("/helloWorld", restHandlers.HelloWorld).Methods("GET")
	router.HandleFunc(consts.UrlPath+"/getTobaccos", r.GetTobbacos).Methods("GET")
	// TODO: create handler to post tabaccos

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"}) // Reemplaza con la URL de tu aplicación Vue.js
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Configura los encabezados CORS utilizando el middleware Cors
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
