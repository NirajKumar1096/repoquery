package util

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/user", GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/user/{"+Userparm+"}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/user/{"+Userparm+"}", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{"+Userparm+"}", DeleteUserEndpoint).Methods("DELETE")

	router.HandleFunc("/user/{"+Userparm+"}/fav/{"+Favidparm+"}", CreateUserFavEndpoint).Methods("POST")
	router.HandleFunc("/user/{"+Userparm+"}/fav/{"+Favidparm+"}", DeleteUserFavEndpoint).Methods("DELETE")
	router.HandleFunc("/user/{"+Userparm+"}/like/{"+Likeidparm+"}", CreateUserLikeEndpoint).Methods("POST")

	router.HandleFunc("/github/repo", GetReposEndpoint).Methods("GET")
	router.HandleFunc("/github/{query}", QueryReposEndpoint).Methods("GET")

	return router
}
