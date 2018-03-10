package net

import (
	"net/http"
	"net/url"
)

type NodeInfo struct {
	Node string `json:"node"`
}

func (n *NodeInfo) toUrlForm() url.Values {
	form := url.Values{}
	form.Add("node", n.Node)
	return form
}

func RigisterServer(url string, info NodeInfo) (err error) {
	_, err = http.PostForm(url, info.toUrlForm())
	return
}
