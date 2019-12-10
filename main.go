package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"text/template"
	"time"
)

func isAlive(w http.ResponseWriter, r *http.Request) {

	// //Allow CORS here By * or specific origin
	// // w.Header().Set("Access-Control-Allow-Origin", "*")
	// // w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// // if r.Method == "OPTIONS" || r.Method == "GET" {
	// // 	return
	// // }

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	var status = "down"

	d := map[string]string{
		"status": status,
	}

	var url string
	url = r.URL.Query().Get("ip")
	if url == "" {
		url = "https://172.16.198.14:8888/EBSGateway/isAlive"
		// url = "https://beta.soluspay.net/api/test"
	}

	// url = "https://beta.soluspay.net/api/test"
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			}},
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Printf("The error is: %v", err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Printf("The error is: %v", err)
		tmpl.ExecuteTemplate(w, "index.html", d)
		return
	}

	if res.StatusCode == http.StatusOK {
		d["status"] = "up"
	}

	err = tmpl.ExecuteTemplate(w, "index.html", d)
	if err != nil {
		panic(err)
	}

}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", fs)
	http.HandleFunc("/", isAlive)
	log.Fatal(http.ListenAndServe(":55555", nil))
}
