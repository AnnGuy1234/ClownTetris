import (
	"os"
	"net/http"
)

os.ReadFile("IMG_0059.png")
package main

func serveImage(w http.ResponseWriter, r *http.Request) {
	img, err := os.ReadFile("IMG_0059.png")
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(img)
	if err != nil {
		http.Error(w, "Failed to write image", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/get-go-file", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("Cells.go")
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(content)
	})

	http.ListenAndServe(":8080", serveImage())
}
