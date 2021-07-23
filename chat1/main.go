package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Type string      `json:"type"`
	User string      `json:"user"`
	Data interface{} `json:"data"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	for {
		m := &Message{}
		err := conn.ReadJSON(m)

		if err != nil {
			log.Println(err)
			return
		}

		data := (*m).Data
		user := (*m).User

		(*m).Data = user + " : " + data.(string)

		err = conn.WriteJSON(m)
		if err != nil {
			log.Panicln(err)
			return
		}
	}
}

func main() {
	mux := pat.New()
	mux.Get("/chat", handler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":8081", n)
}
