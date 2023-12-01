package httpcore

import (
	"net/http"
	"os"
	"path/filepath"
)

func ServeStaticFiles(staticDir http.Dir, staticHandler http.Handler, fallbackPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := staticDir.Open(r.URL.Path)
		if err != nil {
			http.ServeFile(w, r, fallbackPath)
			return
		}
		defer f.Close()
		s, err := f.Stat()
		if err != nil {
			http.ServeFile(w, r, fallbackPath)
			return
		}
		if s.IsDir() {
			index := filepath.Join(r.URL.Path, "index.html")
			if _, err := staticDir.Open(index); os.IsNotExist(err) {
				http.ServeFile(w, r, fallbackPath)
				return
			}
		}

		staticHandler.ServeHTTP(w, r)
	}
}
