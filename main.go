package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("build")))
	http.HandleFunc("/api/Persegi", Persegi)
	http.HandleFunc("/api/Lingkaran", Lingkaran)
	http.HandleFunc("/api/Segitiga", Segitiga)
	http.HandleFunc("/api/Persegi_Panjang", Persegi_Panjang)

	fmt.Println("Web Berjalan dengan port 8000")
	http.ListenAndServe(":8000", nil)
}

// ==============================
type RequestBody_Persegi struct {
	Value string `json:"value"`
}
type ResponseBody_Persegi struct {
	Result float64 `json:"result"`
}

func Persegi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RequestBody_Persegi
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	val, _ := strconv.ParseFloat(requestBody.Value, 64)
	var result float64 = val * val

	responseBody := ResponseBody_Persegi{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

//===============================

type RequestBody_Lingkaran struct {
	Value string `json:"value"`
}
type ResponseBody_Lingkaran struct {
	Result float64 `json:"result"`
}

func Lingkaran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RequestBody_Lingkaran
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	val, _ := strconv.ParseFloat(requestBody.Value, 64)
	var result float64 = 3.14 * val * val

	responseBody := ResponseBody_Lingkaran{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

//===============================

type RequestBody_Segitiga struct {
	Value  string `json:"value"`
	Value2 string `json:"value2"`
}
type ResponseBody_Segitiga struct {
	Result float64 `json:"result"`
}

func Segitiga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RequestBody_Segitiga
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	val, _ := strconv.ParseFloat(requestBody.Value, 64)
	val2, _ := strconv.ParseFloat(requestBody.Value2, 64)
	var result float64 = val * val2 * 0.5

	responseBody := ResponseBody_Segitiga{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

//===============================

type RequestBody_Persegi_Panjang struct {
	Value  string `json:"value"`
	Value2 string `json:"value2"`
}
type ResponseBody_Persegi_Panjang struct {
	Result float64 `json:"result"`
}

func Persegi_Panjang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RequestBody_Persegi_Panjang
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	val, _ := strconv.ParseFloat(requestBody.Value, 64)
	val2, _ := strconv.ParseFloat(requestBody.Value2, 64)
	var result float64 = val * val2

	responseBody := ResponseBody_Persegi_Panjang{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
