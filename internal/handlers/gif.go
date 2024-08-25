package handlers

import (
	"fmt"
	"net/http"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GifHandler(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	gif := query.Get("gif")

	// there is *probably* a path traversal vulnerability hidden here
	text := query.Get("text")

	input := ffmpeg.Input(fmt.Sprintf("gifs/%s.gif", gif))
	input = input.Drawtext(text, 100, 200, false)
	input = input.Output("pipe:1", ffmpeg.KwArgs{"format": "gif"})
	input = input.WithOutput(writer)

	input.Run()
}
