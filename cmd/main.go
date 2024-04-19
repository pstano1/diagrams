package main

import (
    "go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
    defer logger.Sync()

    logger.Info("Hello World!")
}