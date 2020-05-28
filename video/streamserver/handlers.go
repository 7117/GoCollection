package main

import (
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)

	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "internal")
		return
	}

	w.Header().Set("content-type", "video/mp4")

	//二进制流
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()

}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
