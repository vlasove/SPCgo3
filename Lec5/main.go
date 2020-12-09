package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//UploadPageFormHandler ...
func UploadPageFormHandler(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/upload.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error parsing template:", err)
		return
	}
}

//FileUploaderHandler ...
func FileUploaderHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file") //По умолчанию файл будет открыт
	if err != nil {
		log.Println("error getting a file from form:", err)
		return
	}
	defer file.Close()

	//Куда сохраним этот файл?
	outFile, pathError := os.Create("uploadedFile")
	if pathError != nil {
		log.Println("error creating a file for writing:", pathError)
		return
	}
	defer outFile.Close()

	//Копируем все из входного файла в выходной
	_, copyFileError := io.Copy(outFile, file)
	if copyFileError != nil {
		log.Println("error while copy file from to:", copyFileError)
		return
	}
	fmt.Fprintf(w, "File uploaded successfully : "+header.Filename)

}

func main() {
	http.HandleFunc("/", UploadPageFormHandler)
	http.HandleFunc("/upload", FileUploaderHandler)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
