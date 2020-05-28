package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
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
	vid := p.ByName("vid")

	vl := VIDEO_DIR + vid

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	file, _, err := r.FormFile("file")

	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "form")
		return
	}

	data, _ := ioutil.ReadAll(file)

	fn := p.ByName("vid")
	err = ioutil.WriteFile(vl+fn, data, 0666)

	if err != nil {
		fmt.Print("writefile")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "success")

}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload/upload.html")

	t.Execute(w, nil)
}
