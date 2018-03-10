package controllers

import "net/http"

type HomeController struct{}

func (c *HomeController) Get() (string, int) {
	return "Welcome to Burj Center", http.StatusOK
}

func (c *HomeController) GetHeathCheck() (string, int) {
	return "pong", http.StatusOK
}
