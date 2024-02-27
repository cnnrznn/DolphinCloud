package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cnnrznn/DolphinCloud/server/service"
	"github.com/cnnrznn/DolphinCloud/types"
)

const (
	PORT = 36574
)

func main() {
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleUpload(w, r)
	case http.MethodGet:
		//handleDownload()
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	gci := &types.GCI{}

	if err := json.NewDecoder(r.Body).Decode(gci); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	if err := service.Upload(gci); err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
}

func writeError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)

	bs, err := json.Marshal(struct {
		Error error
	}{
		Error: err,
	})
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(bs)
}
