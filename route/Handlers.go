package route

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	fmt.Print(r.RemoteAddr + "\n")
	if err != nil {
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
		http.Error(w, err.Error(), 500)
	}

	fmt.Print(string(body[:]))
}
