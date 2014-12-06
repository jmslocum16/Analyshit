package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	
	"database/sql"

)

const srcDir = "/home/josh/shittyfs"
const dbUsername = "root";
const dbPassord = "password"; // TODO DO NOT KEEP AS ACTUAL PASSWORD

var addr = flag.String("addr", ":1718", "http service address")

func main() {
	checkStartup()
	flag.Parse()
	//http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func checkStartup() {
	if !fileExists(srcDir) {
		panic(srcDir + " doesn't exist! Exiting..")
	}
	// TODO database stuff
	
}

func register(username, passwordHash string) {
	
}

func login(username, passwordHash string) error {
	
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	}
	return true
}
