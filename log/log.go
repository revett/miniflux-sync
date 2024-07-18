package log

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	zerologlog "github.com/rs/zerolog/log"
)

// Metadata is a map of string to any, which is used to add metadata fields to a log event.
type Metadata map[string]interface{}

// New returns a new zerolog.Logger, using a ConsoleWriter.
func New() zerolog.Logger {
	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02T15:04:05.0000",
	}

	writer.FormatTimestamp = func(i interface{}) string {
		return time.Now().Format(writer.TimeFormat)
	}

	return zerologlog.Output(writer)
}

// Info logs a message at the info level.
func Info(ctx context.Context, msg string, meta ...Metadata) {
	logger := getLogger(ctx)
	logEvent := logger.Info()

	if len(meta) > 0 {
		logEvent = logEvent.Fields(mergeMetadata(meta...))
	}

	logEvent.Msg(msg)
}

// Warn logs a message at the warn level.
func Warn(ctx context.Context, msg string, meta ...Metadata) {
	logger := getLogger(ctx)
	logEvent := logger.Warn()

	if len(meta) > 0 {
		logEvent = logEvent.Fields(mergeMetadata(meta...))
	}

	logEvent.Msg(msg)
}

// Error logs a message at the error level.
func Error(ctx context.Context, err error, meta ...Metadata) {
	if err == nil {
		return
	}

	logger := getLogger(ctx)
	logEvent := logger.Error()

	if len(meta) > 0 {
		logEvent = logEvent.Fields(mergeMetadata(meta...))
	}

	logEvent.Msg(err.Error())
	fmt.Printf("%+v\n", err) //nolint:forbidigo
}

// Fatal logs a message at the fatal level, which ends by calling os.Exit(1).
func Fatal(ctx context.Context, err error, meta ...Metadata) {
	if err == nil {
		return
	}

	logger := getLogger(ctx)
	logEvent := logger.Fatal()

	if len(meta) > 0 {
		logEvent = logEvent.Fields(mergeMetadata(meta...))
	}

	// This has to happen before the call to .Msg(), as that will call os.Exit(1).
	fmt.Printf("%+v\n", err) //nolint:forbidigo
	logEvent.Msg(err.Error())
}

// WithMetadata returns a new context with the given metadata.
func WithMetadata(ctx context.Context, meta Metadata) context.Context {
	if len(meta) == 0 {
		return ctx
	}

	logger := getLogger(ctx)
	return logger.With().Fields(mergeMetadata(meta)).Logger().WithContext(ctx)
}

func getLogger(ctx context.Context) *zerolog.Logger {
	logger := zerologlog.Ctx(ctx)

	if logger == nil {
		panic("nil logger returned from context, this should never happen")
	}

	if logger.GetLevel() == zerolog.Disabled {
		panic("disable logger returned from context, this should never happen")
	}

	return logger
}

func mergeMetadata(meta ...Metadata) map[string]interface{} {
	merged := map[string]interface{}{}

	for _, m := range meta {
		for k, v := range m {
			merged[k] = v
		}
	}

	return merged
}
