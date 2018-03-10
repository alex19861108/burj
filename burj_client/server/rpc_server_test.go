package server

import "testing"

func TestDeCompress(t *testing.T) {
	err := packDirectory("/usr/local/Cellar/go/src/github.com/alex19861108/burj/burj_client/config",
		"/usr/local/Cellar/go/src/github.com/alex19861108/burj/burj_client/main/test.zip")
	if err != nil {
		t.Fatal(err)
	}
}
