package logging

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func ExampleProd() {
	Debug(false)

	log.Info().Str("for", "bar").Msg("Logging in production.")
	log.Debug().Str("for", "bar").Msg("Logging in development.")

	// Output: {"level":"info","for":"bar","time":"2021-04-01T15:39:20.939053648+08:00","caller":"/home/william/go/src/github.com/SB-IM/logging/logging_example_test.go:10","message":"Logging in production."}
}

func ExampleDebug() {
	Debug(true)

	log.Debug().Str("for", "bar").Msg("Logging in development.")

	// Output: 2021-03-15T16:31:34.399023951+08:00 DBG logging_example_test.go:14 > Logging in development. for=bar
}

func ExampleErrorStack() {
	Debug(true)

	err := outer()
	log.Error().Stack().Err(err).Send()
	//Output: 2021-04-06T12:58:53.63045935+08:00 ERR logging_example_test.go:29 >  error="seems we have an error here" stack=[{"func":"inner","line":"35","source":"logging_example_test.go"},{"func":"middle","line":"39","source":"logging_example_test.go"},{"func":"outer","line":"47","source":"logging_example_test.go"},{"func":"ExampleErrorStack","line":"28","source":"logging_example_test.go"},{"func":"runExample","line":"63","source":"run_example.go"},{"func":"runExamples","line":"44","source":"example.go"},{"func":"(*M).Run","line":"1418","source":"testing.go"},{"func":"main","line":"49","source":"_testmain.go"},{"func":"main","line":"225","source":"proc.go"},{"func":"goexit","line":"1371","source":"asm_amd64.s"}]
}

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}

	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}

	return nil
}
