package common

import (
	"context"
	"os"
	"runtime"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/job"
	"github.com/opsidian/basil/logger"
	"github.com/rs/zerolog"
)

func Main(ctx context.Context, parseCtx *basil.ParseContext) {
	level := zerolog.InfoLevel
	if os.Getenv("DEBUG") == "1" {
		level = zerolog.DebugLevel
	}

	zerolog.TimeFieldFormat = basil.LogTimeFormat
	zl := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05.000"}).With().
		Timestamp().
		Logger().
		Level(level)

	scheduler := job.NewScheduler(runtime.NumCPU()*2, 100)
	scheduler.Start()
	defer scheduler.Stop()

	if _, err := basil.Evaluate(
		parseCtx,
		basil.NewBlockContext(ctx, nil, logger.NewZeroLogLogger(zl)),
		scheduler,
		"main",
	); err != nil {
		panic(err)
	}
}