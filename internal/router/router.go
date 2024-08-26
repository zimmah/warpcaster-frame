package router

import 
(
	"net/http"
	"os"
	"log"
)

func Router() {
	mux := http.NewServeMux()
	mux.Handle("GET /api/frames/{frameID}", logger(http.HandlerFunc(handleGetFrameByID)))

	port := os.Getenv("PORT")
	server := &http.Server{
		Addr: 	":" + port,
		Handler: mux,
	}

	log.Printf("Listening on port: %s", port)
	log.Fatal(server.ListenAndServe())
}