package main

import (
	"html/template"
	"net/http"
	"path"

	log "github.com/sirupsen/logrus"
)

var liveness bool
var readiness bool

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

	liveness = true
	readiness = true

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

	if err = tmpl.Execute(w, Application{Liveness: liveness, Readiness: readiness}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Liveness ...")
	if !liveness {
		log.Info("Imitate processing problem")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Readiness ...")
	if !readiness {
		log.Info("Imitate processing problem")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleHealthToggle(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Toggle ...")
	liveness = !liveness

	w.WriteHeader(http.StatusOK)
}

func handleReadinessToggle(w http.ResponseWriter, r *http.Request) {
	log.Info("Handle Toggle ...")
	readiness = !readiness

	w.WriteHeader(http.StatusOK)
}
