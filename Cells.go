import (
	"os"
	"net/http"
)

os.ReadFile("IMG_0059.png")
package main

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

	http.ListenAndServe(":8080", nil)
}
