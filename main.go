package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/api/binary", func(w http.ResponseWriter, r *http.Request) {
		// Header CORS Lengkap
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 1. Baca data binary dari client
		clientData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Gagal membaca data", 500)
			return
		}

		fmt.Printf("Client mengirim: %s\n", string(clientData))

		// 2. Buat balasan: "Ini Dari Server " + data client
		prefix := []byte("Ini Dari Server ")
		balasan := append(prefix, clientData...)

		// 3. Kirim balik sebagai binary
		w.Write(balasan)
	})

	fmt.Println("Server aktif di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
