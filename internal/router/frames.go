package router

import (
	"net/http"
	"strconv"
	"fmt"
)

func handleGetFrameByID(w http.ResponseWriter, r *http.Request) {
	frameID, err := strconv.Atoi(r.PathValue("frameID"))
	if err != nil { 
		respondWithError(w, http.StatusBadRequest, "Invalid frame ID")
		return
	}

	fmt.Println(frameID)
	// got to serve a new frame by ID
}