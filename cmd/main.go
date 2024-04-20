package main

import (
	"github.com/pstano1/diagrams.git/pkg/diagrams"
	"go.uber.org/zap"
)

const (
	diagramsDirectory = "./pkg/diagrams"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	controller := diagrams.New(logger.Named("diagrams"))
	controller.GenerateDiagrams()
}
