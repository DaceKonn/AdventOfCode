package helpers

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogFile() (file *os.File) {
	crnTime := fmt.Sprint(time.Now().Format("2006-01-02-15-04-05"))
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("Failed to create directory")
		return
	}
	file, err := os.OpenFile(strings.Join([]string{"logs/logfile-", crnTime, "-.log"}, ""), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
		return
	}
	return file
}

func ConfigureLogger(file *os.File, stdout bool) {
	zerolog.TimeFieldFormat = time.RFC3339
	var w io.Writer
	switch stdout {
	case true:
		cw := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
		w = zerolog.MultiLevelWriter(cw, file)
	default:
		w = file
	}
	log.Logger = zerolog.New(w).With().Timestamp().Logger()
	log.Info().Msg("Logger configured")
}

func CloseLogFile(file *os.File) {
	if file == nil {
		log.Error().Msg("Can't close log file - file nil")
		return
	}
	if err := file.Close(); err != nil {
		log.Error().Err(err).Msg("Failed to close log file")
	}
}
