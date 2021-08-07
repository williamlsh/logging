package logging

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func init() {
	// Use the RFC 3339 as the standard for the date-time format.
	zerolog.TimeFieldFormat = time.RFC3339Nano

	// Set default level to be info, suitable for production environment.
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Error Logging with Stacktrace
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

// Debug sets logging at debug mod determined by input if true,
// debug mod suites for development environments.
// Logger is concurrency safe.
func Debug(debug bool) {
	if debug {
		setDebugLevel()

		log.Logger = pretty()
	} else {
		log.Logger = standard()
	}
}

// setDebugLevel sets zerolog's global level to be debug which overides the default info level.
// It suits for development environment.
func setDebugLevel() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

// pretty returns a console prettier.
// It prettifies logs by outputing them to console with caller information.
func pretty() zerolog.Logger {
	logger := log.Output(zerolog.ConsoleWriter{
		Out:        syncWriter(),
		TimeFormat: time.RFC3339Nano,
	})

	return logger.With().Caller().Logger()
}

// standard is standard JSON output logger. Suitable for production environment.
func standard() zerolog.Logger {
	logger := log.Output(syncWriter())
	return logger.With().Caller().Logger()
}

// syncWriter is concurrent safe writer.
func syncWriter() io.Writer {
	return diode.NewWriter(os.Stderr, 1000, 0, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})
}
