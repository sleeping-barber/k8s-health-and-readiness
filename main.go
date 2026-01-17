package main

import (
	"html/template"
	"net/http"
	"path"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

var liveness atomic.Bool
var readiness atomic.Bool

type Application struct {
	Liveness  bool
	Readiness bool
}

func main() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	log.Info("Starting up process")

	liveness.Store(true)
	readiness.Store(true)

	http.HandleFunc("/", handleStatus)
	http.HandleFunc("/healthz", handleHealthz)
	http.HandleFunc("/readiness", handleReadiness)
	http.HandleFunc("/toggle/healthz", handleHealthToggle)
	http.HandleFunc("/toggle/readiness", handleReadinessToggle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Process terminated: %v", err)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Status ...")
	fp := path.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, Application{Liveness: liveness.Load(), Readiness: readiness.Load()}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Liveness ...")
	if !liveness.Load() {
		log.Info("Imitate processing problem")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Readiness ...")
	if !readiness.Load() {
		log.Info("Imitate processing problem")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleHealthToggle(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Toggle ...")
	current := liveness.Load()
	liveness.Store(!current)

	w.WriteHeader(http.StatusOK)
}

func handleReadinessToggle(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Toggle ...")
	current := readiness.Load()
	readiness.Store(!current)

	w.WriteHeader(http.StatusOK)
}
