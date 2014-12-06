package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"os"
	"hash/crc32"
	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

const srcDir = "/home/josh/shittyfs";
const dbString = "root:pass/analyshit";
var db *sql.DB = nil;

var addr = flag.String("addr", ":1718", "http service address")

func main() {
	fmt.Printf("heel\n");
	checkStartup()
	defer db.Close()
	
	flag.Parse()
	http.Handle("/login", http.HandlerFunc(serveLogin))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func serveLogin(w http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	if username == "" || password == "" {
		w.Write([]byte("need both username and password!"));
		return
	}
	passwordHash := getPasswordHash(password)
	log.Printf("%d\n", passwordHash);
	result := login(username, passwordHash)
	if result == "" { // success
		w.Write([]byte("successfully logged in"));
	} else {
		w.Write([]byte(result));
	}
}

func getPasswordHash(pass string) uint32 {
	return crc32.ChecksumIEEE([]byte(pass));
}

func checkStartup() {
	if !fileExists(srcDir) {
		panic(srcDir + " doesn't exist! Exiting..")
	}
	// TODO database stuff
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		panic(err.Error())
	} 
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close();
}

func register(username, passwordHash string) error {
	return nil
}

func login(username string, passwordHash uint32) string {
	return ""
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
