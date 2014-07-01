// Riemann middleware for martini framework
//
// Copyright (C) 2014 by Christopher Gilbert <christopher.john.gilbert@gmail.com>
package gorymartini

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bigdatadev/goryman"
	"github.com/go-martini/martini"
)

// NewGoryMartini - Factory
func NewGoryMartini(host string) (*GorymanClient, martini.Handler) {
	riemann := goryman.NewGorymanClient(host)
	err := riemann.Connect()
	if err != nil {
		return nil, nil
	}

	return riemann, func(res http.ResponseWriter, req *http.Request, c Context, log *log.Logger) {
		start := time.Now()

		rw := res.(ResponseWriter)
		c.Next()

		metric := float64(time.Since(start)) / float64(time.Milliseconds)

		err := riemann.SendEvent(&goryman.Event{
			Service:     "http req",
			Metric:      metric,
			Description: fmt.Printf("Request took %f seconds.", metric),
			Tags: []string{
				"http",
			},
			Attributes: map[string]interface{}{
				"path":   req.URL.Path,
				"status": strconf.Itoa(rw.Status()),
			},
		})
		if err != nil {
			log.Fatal("Riemann client SendEvent failed!")
		}
	}
}
