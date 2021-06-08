package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.GET("/user/:uid", getuser)
	router.POST("/adduser/:uid", adduser)
	router.PUT("/moduser/:uid", modifyuser)
	router.DELETE("/deluser/:uid", deleteuser)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if _, err := fmt.Fprintf(w, "Welcome!\n"); err != nil {
		return
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if _, err := fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name")); err != nil {
		return
	}
}

func getuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if _, err := fmt.Fprintf(w, "You are get user %s\n", uid); err != nil {
		return
	}
}

func modifyuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if _, err := fmt.Fprintf(w, "You are modify user %s\n", uid); err != nil {
		return
	}
}

func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if _, err := fmt.Fprintf(w, "You are delete user %s\n", uid); err != nil {
		return
	}
}

func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")
	if _, err := fmt.Fprintf(w, "You are add user %s\n", uid); err != nil {
		return
	}
}
