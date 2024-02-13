package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

const (
	HOST    = "localhost"
	PORT    = "6378"
	TYPE    = "tcp"
	PORTSER = "8080"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create/", CreateURL).Methods("GET")
	r.HandleFunc("/{hash}", Redirect).Methods("GET")
	log.Printf("Запуск веб-сервера на %s", HOST+":"+PORTSER)
	http.ListenAndServe(":"+PORTSER, r)
}

func CreateURL(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		http.Error(w, "Port listening error", http.StatusInternalServerError)
		log.Fatalf("Port listening error: %s", err)
	}
	defer conn.Close()

	rsrv := bufio.NewReader(conn)
	wsrv := bufio.NewWriter(conn)

	params := r.URL.Query()
	longURL := strings.Join(params["url"], "")
	_, err = http.Get(longURL)
	if err != nil {
		http.Error(w, "Bad url!", http.StatusBadRequest)
		log.Print(err.Error())
		return
	} else {
		var response, hash string
		for response != "ok" {
			hash = GetHash()
			fmt.Fprint(wsrv, hash+"\n")
			fmt.Fprint(wsrv, longURL+"\n")
			wsrv.Flush()
			response, _ = rsrv.ReadString('\n')
			response = strings.TrimSpace(response)
		}
		fmt.Fprint(w, "Your url:\nlocalhost:"+PORTSER+"/", hash)
	}
}

func GetHash() string {
	b := make([]byte, 7)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		fmt.Fprint(w, "home page")
		return
	}
	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Port listening error: %s", err)
	}
	defer conn.Close()

	rsrv := bufio.NewReader(conn)
	wsrv := bufio.NewWriter(conn)

	params := mux.Vars(r)
	shortURL := params["hash"]

	fmt.Fprint(wsrv, shortURL+"\n")
	fmt.Fprint(wsrv, "-1\n")
	wsrv.Flush()
	longURL, _ := rsrv.ReadString('\n')
	longURL = strings.TrimSpace(longURL)
	if longURL == "404" {
		fmt.Fprint(w, "baad url!")
	} else {
		fmt.Println("REDIRECT ", longURL)
		http.Redirect(w, r, longURL, http.StatusMovedPermanently)
	}
}
