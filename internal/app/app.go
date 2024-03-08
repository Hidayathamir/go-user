// Package app contains the application starter.
package app

import (
	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/controller/http"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
	err := config.Init()
	if err != nil {
		logrus.Fatalf("config.Init: %v", err)
	}

	err = http.RunServer()
	if err != nil {
		logrus.Fatalf("http.RunServer: %v", err)
	}
}
