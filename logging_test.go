package logging

import (
	"testing"
	"time"

	"github.com/rs/zerolog/log"
)

func TestSyncLogger(t *testing.T) {
	Debug(true)

	// Concurrent logging should have no data race.
	for i := 0; i < 20; i++ {
		go log.Info().Msg("hello")
	}

	time.Sleep(3 * time.Second)
}
