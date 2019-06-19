package main

import (
	"log"
	"net/http"
	"os"
)

func sendCommand(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("ParseForm() err: %v", err)
		return
	}
	controlID := r.FormValue("control_id")
	log.Printf("control_id = %s\n", controlID)
	req, err := http.NewRequest("GET", "http://10.0.0.10:5000/?gpio="+controlID, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println("response Status:", resp.Status)

}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/command", sendCommand)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3030", nil))
}
