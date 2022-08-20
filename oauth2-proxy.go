package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func handleOauth2(router *mux.Router) {
	router.HandleFunc("/-/user", func(w http.ResponseWriter, r *http.Request) {
		fullNameMap, _ := url.ParseQuery(r.Header.Get("X-Forwarded-User"))
		var fullName string
		for k := range fullNameMap {
			fullName = k
			break
		}
		user := &UserInfo{
			Email:    r.Header.Get("X-Forwarded-Email"),
			Name:     fullName,
			NickName: r.Header.Get("X-Forwarded-Preferred-Username"),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		data, _ := json.Marshal(user)
		w.Write(data)
	})
}
