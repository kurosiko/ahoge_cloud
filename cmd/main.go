package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type FileInfo struct {
	Name  string
	Size  int64
	IsDir bool
}

func fileListHandler(w http.ResponseWriter, r *http.Request) {
	entries, _ := os.ReadDir("data")
	var files []FileInfo
	tmpl,err := template.ParseFiles("cmd/templates/list.html")
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	for _, entry := range entries {
		info, _ := entry.Info()
		files = append(files, FileInfo{
			Name:  entry.Name(),
			Size:  info.Size(),
			IsDir: entry.IsDir(),
		})
	}
	tmpl.Execute(w, files)
}

func main() {
	fmt.Println("run on http://localhost:8000")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/list", fileListHandler)
	http.ListenAndServe(":8000", nil)
}
