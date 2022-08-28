package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

type Handler = func(http.ResponseWriter, *http.Request)

type HTTPMultiplexer struct {
	URLPatternTree *Trie
}

func (mux *HTTPMultiplexer) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	
	handler := mux.handler(r)
	handler(w, r)
}


func (mux *HTTPMultiplexer) GET(url_pattern string, handler Handler) {
	mux.URLPatternTree.AddHandler(url_pattern, "GET", handler)
}

func (mux *HTTPMultiplexer) POST(url_pattern string, handler Handler) {
	mux.URLPatternTree.AddHandler(url_pattern, "POST", handler)
}


func (mux *HTTPMultiplexer) handler(req *http.Request) (hdlr Handler) {
	escaped := req.URL.EscapedPath()
	matched_node := mux.URLPatternTree.Search(escaped)

	if  matched_node != nil {
		if req.Method != matched_node.method {
			hdlr = func(w http.ResponseWriter, r *http.Request) {
				body, _ := json.Marshal(ErrorMessage{Error: "Method not allowed."})
				w.WriteHeader(http.StatusForbidden)
				w.Write(body)
			}
			return
		}
		hdlr = matched_node.handler
		return 
	}

	hdlr = func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(ErrorMessage{Error: "Resource not found."})
		w.WriteHeader(http.StatusNotFound)
		w.Write(body)
	}
	return
}

func NewHTTPMultiplexer() *HTTPMultiplexer {
	return &HTTPMultiplexer{URLPatternTree: NewTrie()}
}