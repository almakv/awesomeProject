package logger

import (
	"log"
	"log/slog"
)

func LogLoad() *log.Logger {
	Mlogg := slog.NewLogLogger(slog.Default().Handler(), slog.LevelInfo)
	Mlogg.Println("LoggerMain was created")
	return Mlogg
}
