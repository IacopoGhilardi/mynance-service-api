package utils

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func init() {
	textHandler := slog.NewTextHandler(os.Stdout, nil)
	Logger = slog.New(textHandler)
}
