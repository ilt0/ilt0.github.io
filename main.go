package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", home)
	http.ListenAndServe(":80", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page")

	//  html dosyası gösterme metodları

	/* örnek statik dosya gösterimi
	view, _ := template.ParseFiles(("index.html"))
	view.Execute(w, nil)
	*/

	// dinamik veri göndermek
	view, _ := template.ParseFiles(("index.html"))
	// data := "USER"
	data := make(map[string]interface{})
	data["veri"] = "user"
	data["veri2"] = "user2"
	data["sayilar"] = []int{1,2,3,4,5}
	data["is_admin"] = true
	data["sayi"] = 10



	view.Execute(w, data)
	
}
