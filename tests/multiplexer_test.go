package tests

import (
	"testing"
	"wallet-sandbox/utils"
)


func TestMultiplexer(t *testing.T) {
	mux := utils.NewHTTPMultiplexer()
	mux.GET("/", testHandler)

	matched_node := mux.URLPatternTree.Search("/")

	if matched_node == nil {
		t.Errorf("expected matched, found %v", matched_node)
	}
	
	if matched_node.Handler() == nil {
		t.Errorf("match node (%v) has nil handler", matched_node.Path())
	}

	if matched_node.Method() != "GET" {
		t.Errorf("expected method = GET, found %v", matched_node.Method())
	}
}

func TestHandler(t *testing.T) {
	// mux := utils.NewHTTPMultiplexer()
	// mux.GET("/", testHandler)
}