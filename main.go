package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/rfraposa/hadoopbeat/beater"
)

func main() {
	err := beat.Run("hadoopbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
