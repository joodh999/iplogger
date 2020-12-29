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

type IP struct {
	gorm.Model
	id int
	ip string
}

// Newid ...
func Newid(w http.ResponseWriter, r *http.Request) {
	id := rand.Intn(1000)
	db.Create(&IP{id: id})
	defer fmt.Fprintf(w, "%d id:", id)
}

// Logip ...
func Logip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey")
}

// Stats ...
func Stats(w http.ResponseWriter, r *http.Request) {

	ID, err := r.URL.Query()["id"]
	if err {
		fmt.Fprintf(w, "Url Param 'id' is missing")
		return
	}

	var ip IP
	result := db.First(&ip, ID)
	if result.Error != nil {
		fmt.Print(&result.Error)
		fmt.Println("err")
	}
	fmt.Print(*result)
}

func main() {

	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&IP{})

	http.HandleFunc("/newid", Newid)
	http.HandleFunc("/", Logip)
	http.HandleFunc("/Stats/", Stats)

	http.ListenAndServe(":8000", nil)
}
