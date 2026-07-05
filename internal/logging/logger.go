package logging

import (
	"log/slog"
	"os"
)

var Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Group(name string, attrs ...any) slog.Attr {
	return slog.Group(name, attrs...)
}
