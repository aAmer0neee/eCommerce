package logger

import (
	"log/slog"
	"os"
)

type slogLogger struct {
	log *slog.Logger
}

func (s *slogLogger) Info(msg string, args ...any) {
	s.log.Info(msg, args...)
}
func (s *slogLogger) Debug(msg string, args ...any) {
	s.log.Debug(msg, args...)
}
func (s *slogLogger) Warn(msg string, args ...any) {
	s.log.Warn(msg, args...)
}
func (s *slogLogger) Error(msg string, args ...any) {
	s.log.Error(msg, args...)
}

func newSlog(level string) *slogLogger {
	opts := &slog.HandlerOptions{}
	switch level {
	case "info":
		opts.AddSource = false
		opts.Level = slog.LevelInfo
	case "debug":
		opts.AddSource = true
		opts.Level = slog.LevelDebug
	}

	opts.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			t := a.Value.Time()
			return slog.String(slog.TimeKey, t.Format("2006-01-02 15:04:05"))
		}
		return a
	}

	return &slogLogger{
		log: slog.New(slog.NewTextHandler(os.Stdout, opts))}
}
