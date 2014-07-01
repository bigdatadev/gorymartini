// Riemann middleware for martini framework
//
// Copyright (C) 2014 by Christopher Gilbert <christopher.john.gilbert@gmail.com>
package gorymartini

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bigdatadev/goryman"
	"github.com/go-martini/martini"
)

// NewGoryMartini - Factory
func NewGoryMartini(host string) (*goryman.GorymanClient, martini.Handler) {
	riemann := goryman.NewGorymanClient(host)
	err := riemann.Connect()
	if err != nil {
		return nil, nil
	}

	return riemann, func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		start := time.Now()

		rw := res.(martini.ResponseWriter)
		c.Next()

		metric := float64(time.Since(start)) / float64(time.Millisecond)

		err := riemann.SendEvent(&goryman.Event{
			Service:     "http req",
			Metric:      metric,
			Description: fmt.Sprintf("Request took %f seconds.", metric),
			Tags: []string{
				"http",
			},
			Attributes: map[string]string{
				"path":   req.URL.Path,
				"status": strconv.Itoa(rw.Status()),
			},
		})
		if err != nil {
			log.Fatal("Riemann client SendEvent failed!")
		}
	}
}
