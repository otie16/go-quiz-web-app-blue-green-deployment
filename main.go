package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    // "os"
    // "time"
)

type QuizQuestion struct {
    ID       int      `json:"id"`
    Question string   `json:"question"`
    Options  []string `json:"options"`
}

type StatusResponse struct {
    Message   string `json:"message"`
    Version   string `json:"version"`
    Timestamp string `json:"timestamp"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "OK")
}

func quizHandler(w http.ResponseWriter, r *http.Request) {
    quiz := []QuizQuestion{
        {1, "What does Kubernetes use to manage container networking?", []string{"Calico", "Ingress", "Istio", "Etcd"}},
        {2, "What command is used to build a Docker image?", []string{"docker make", "docker run", "docker build", "docker image"}},
        {3, "Which HTTP status code means 'Not Found'?", []string{"500", "301", "404", "200"}},
        {4, "What language is Kubernetes primarily written in?", []string{"Python", "Java", "Go", "Rust"}},
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(quiz)
}

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/quiz", quizHandler)
    http.HandleFunc("/health", healthHandler)

    port := "8080"
    log.Printf("Starting server on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

