package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func isAlive(w http.ResponseWriter, r *http.Request) {
	var url string
	url = r.URL.Query().Get("ip")
	if url == "" {
		url = "https://172.16.198.14:8888/EBSGateway/isAlive"
		url = "https://beta.soluspay.net/api/test"
	}

	url = "https://beta.soluspay.net/api/test"
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			}},
		Timeout: 5 * time.Second,
	}

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Printf("The error is: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := httpClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("The error is: %v", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func main() {
	http.HandleFunc("/pinger", isAlive)
	log.Fatal(http.ListenAndServe(":55555", nil))
}
