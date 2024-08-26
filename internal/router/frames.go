package router

import (
	"net/http"
	"strconv"
	"fmt"
)

const placeholderFrame = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <title>Apple frame</title>
    <meta property="fc:frame" name="fc:frame" content="vNext" />
    <meta
      property="fc:frame:image"
      name="fc:frame:image"
      content="https://fastly.picsum.photos/id/237/300/300.jpg"
    />
    <meta
      property="fc:frame:image:aspect_ratio"
      name="fc:frame:image:aspect_ratio"
      content="1:1"
    />
    <meta
      property="og:image"
      name="og:image"
      content="https://fastly.picsum.photos/id/237/300/300.jpg"
    />
    <meta
      name="fc:frame:post_url"
      content="https://frames-bevm.onrender.com/"
    />
    <meta name="fc:frame:button:1" content="Go to Tribally" />
    <meta name="fc:frame:button:1:action" content="link" />
    <meta name="fc:frame:button:1:target" content="https://tribally.games" />
    <meta name="fc:frame:button:2" content="Show next frame" />
    <meta name="fc:frame:button:2:action" content="post" />
    <meta
      name="fc:frame:button:2:target"
      content="https://frames-bevm.onrender.com/api/frames/2"
    />
    <link href="style.css" rel="stylesheet" />
  </head>
  <body>
    <h1>this is a different frame, with an dog</h1>
    <img
      src="https://fastly.picsum.photos/id/237/300/300.jpg"
    />
  </body>
</html>
`

func handleGetFrameByID(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.Atoi(r.PathValue("frameID"))
	if err != nil { 
		respondWithError(w, http.StatusBadRequest, "Invalid frame ID")
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(placeholderFrame))
	w.WriteHeader(http.StatusOK)
}