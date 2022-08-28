package tests

import (
	"fmt"
	"net/http"
	"testing"
	"wallet-sandbox/utils"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test handler")
}

func TestAddHandler(t *testing.T) {
	tree := utils.NewTrie()
	tree.AddHandler("/", "GET", testHandler)
	tree.AddHandler("/path", "GET", testHandler)
	tree.AddHandler("/path/*", "GET", testHandler)
	tree.AddHandler("/path/*/test", "GET", testHandler)
	matched_node := tree.Search("/path/handler")

	if matched_node == nil {
		t.Errorf("expected matched, found %v", matched_node)
	}
	if matched_node.Handler() == nil {
		t.Errorf("match node (%v) has nil handler", matched_node.Path())
	}
}