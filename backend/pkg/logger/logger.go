package logger

//
//import (
//	"delivery-microservice-goods/backend/config"
//	"log/slog"
//	"os"
//)
//
//const (
//	envLocal = "local"
//	envDev   = "dev"
//	envProd  = "prod"
//)
//
//func New() *slog.Logger {
//	var logger *slog.Logger
//	cfg := config.New()
//
//	switch cfg.Log {
//	case envLocal:
//		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
//	case envDev:
//		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
//	case envProd:
//		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
//	}
//
//	return logger
//}
