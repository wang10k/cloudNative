package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {

	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		ret, err := json.Marshal(v)
		if err != nil {
			logrus.Errorf("json.Marshal failed %v", err)
			continue
		}
		w.Header().Add(k, string(ret))
	}
	version := os.Getenv("VERSION")
	w.Header().Add("version", version)
	statusCode := 200
	ip := ""
	addrArr := strings.Split(r.RemoteAddr, ":")
	if len(addrArr) > 1 {
		ip = addrArr[0]
	}
	logrus.WithField("ip", ip).WithField("status_code", statusCode).Infof("request record")
	w.WriteHeader(statusCode)
	io.WriteString(w, "ok")
}
