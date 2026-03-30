package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const uploadDir = "./uploads"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form tanpa batasan ukuran
	if err := r.ParseMultipartForm(0); err != nil {
		http.Error(w, "Gagal parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Ketik 'file' di key untuk mengupload file atau gambar
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Gagal membaca file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Membuat direktori upload jika belum ada
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, "Gagal membuat direktori: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Buat file tujuan
	dstPath := filepath.Join(uploadDir, filepath.Base(header.Filename))
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Gagal membuat file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Salin isi file ke disk
	written, err := io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Gagal menyimpan file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File '%s' berhasil diupload (%d bytes)\n", header.Filename, written)
}
