package main

import (
	"fmt"
	"net/http"

	"github.com/alienimnida/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API service")
	fmt.Println(`
   ____   ____    _    ____  _ 
  / ___| /  _ \  / \  |  _ \| |
 | |  _  | | | |/ _ \ | |_) | |
 | |_| | | |_| / ___ \|  __/| |
  \____| |____/_/   \_\_|   |_|
	`)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
