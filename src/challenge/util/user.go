package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User is main entity
type User struct {
	UserID string `json:"userid,omitempty"`
	Name   string `json:"name,omitempty"`
	Favs   []int  `json:"favorites,omitempty"`
	Likes  []int  `json:"likes,omitempty"`
}

const (
	Userparm   = "user"
	Favidparm  = "favid"
	Likeidparm = "likeid"
)

// This variable servers as memory database container, holds users
var users []User

func Init() {
	// favs := []
	users = append(users, User{UserID: "1", Name: "Nic", Favs: []int{126195231, 93446075}})
	users = append(users, User{UserID: "2", Name: "Mary", Favs: []int{126195231, 120045621}, Likes: []int{93446075, 120045621}})
}

// GetUserEndpoint is
func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range users {
		if item.UserID == params[Userparm] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// GetusersEndpoint is
func GetUsersEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// CreateUserEndpoint is
func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var User User
	_ = json.NewDecoder(req.Body).Decode(&User)
	User.UserID = params[Userparm]
	users = append(users, User)
	json.NewEncoder(w).Encode(users)
}

// DeleteUserEndpoint is
func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range users {
		if item.UserID == params[Userparm] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

// CreateUserFavorite adds favs
func CreateUserFavEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range users {
		if item.UserID == params[Userparm] {
			newFav, err := strconv.Atoi(params[Favidparm])
			if err != nil {
				err = fmt.Errorf("CreateUserFavEndpoint failed with non numeric favid, favid=%s, err='%s'", params[Favidparm], err)
				return
			}
			item.Favs = append(item.Favs, newFav)
			users[index] = item
		}
	}
	json.NewEncoder(w).Encode(users)
}

// DeleteUserFavEndpoint removes favs for user
func DeleteUserFavEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range users {
		if item.UserID == params[Userparm] {
			newFav, err := strconv.Atoi(params[Favidparm])
			if err != nil {
				err = fmt.Errorf("DeleteUserFavEndpoint failed with non numeric favid, favid=%s, err='%s'", params[Favidparm], err)
				return
			}
			for i, fav := range item.Favs {
				if fav == newFav {
					item.Favs = append(item.Favs[:i], item.Favs[i+1:]...)
					break
				}
			}
			users[index] = item
		}
	}
	json.NewEncoder(w).Encode(users)
}

// CreateUserFavorite adds favs
func CreateUserLikeEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range users {
		if item.UserID == params[Userparm] {
			newLike, err := strconv.Atoi(params[Likeidparm])
			if err != nil {
				err = fmt.Errorf("CreateUserLikeEndpoint failed with non numeric favid, favid=%s, err='%s'", params[Likeidparm], err)
				return
			}
			item.Likes = append(item.Likes, newLike)
			users[index] = item
		}
	}
	json.NewEncoder(w).Encode(users)
}
