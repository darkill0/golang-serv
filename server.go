package main 

import (
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author *Author `json:"author"`
	
}

type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var books[]Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type". "application/json")
	json.NewEncoder(w).Encode(books)
}