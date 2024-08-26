package router

import (
	"net/http"
	"strconv"
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
      content="https://images.unsplash.com/photo-1469474968028-56623f02e42e?q=80&w=1474&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
    />
    <meta
      property="fc:frame:image:aspect_ratio"
      name="fc:frame:image:aspect_ratio"
      content="1:1"
    />
    <meta
      property="og:image"
      name="og:image"
      content="https://images.unsplash.com/photo-1469474968028-56623f02e42e?q=80&w=1474&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
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
    <h1>this is a different frame, with a landscape</h1>
    <img
      src="https://images.unsplash.com/photo-1469474968028-56623f02e42e?q=80&w=1474&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(placeholderFrame))
}