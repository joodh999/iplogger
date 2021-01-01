package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Schema ...
type Schema struct {
	gorm.Model
	ID int
	IP string
}

// Newid ...
func Newid(w http.ResponseWriter, r *http.Request) {
	newid := rand.Intn(10000)
	db.Create(&Schema{ID: newid, IP: "NULL"})
	defer fmt.Fprintf(w, "id: %v\n", newid)
}

// Logip ...
func Logip(w http.ResponseWriter, r *http.Request) {

	param, err := r.URL.Query()["id"]
	if !err {
		fmt.Fprintf(w, "Url Param 'id' is missing")
		return
	}

	ip := r.RemoteAddr
	db.Model(&Schema{}).Where("ID = ?", param).Update("IP", ip)

	fmt.Fprintf(w, "hahahahahahahaha got your ip")
}

// Stats ...
func Stats(w http.ResponseWriter, r *http.Request) {

	param, err := r.URL.Query()["id"]
	if !err {
		fmt.Fprintf(w, "Url Param 'id' is missing")
		return
	}

	var result Schema
	db.First(&result, "ID = ?", param)
	fmt.Fprintf(w, "ip: %s\n ", result.IP)

}

func main() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db.AutoMigrate(&Schema{})

	http.HandleFunc("/newid", Newid)
	http.HandleFunc("/", Logip)
	http.HandleFunc("/Stats", Stats)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
