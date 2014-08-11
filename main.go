package basicproxy

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"io"
	"net/http"
	"strings"

	"appengine"
	"appengine/urlfetch"
)

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/serve", serveHandler)
	mux.HandleFunc("/", indexHandler)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)
	http.Handle("/", n)
}

func serveHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if len(url) == 0 {
		fmt.Fprintln(w, "<h1>No URL was passed</h1>")
		return
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	resp, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(w, "<h1>Error fetching from %s</h1>\n%s\n", url, err.Error())
		return
	}

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		fmt.Fprintln(w, "<h1>Error occured while copying data streams</h1>")
		return
	}

	resp.Body.Close()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "index.html")
}
