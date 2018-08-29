package reZerologer

import (
	"github.com/remicro/api/logging"
	"github.com/rs/zerolog"
	"time"
)

type implEntry struct {
	event *zerolog.Event
}

func (e *implEntry) String(key, value string) logging.Entry {
	e.event.Str(key, value)
	return e
}

func (e *implEntry) Int(key string, value int) logging.Entry {
	e.event.Int(key, value)
	return e
}

func (e *implEntry) Err(err error) logging.Entry {
	e.event.Err(err)
	return e
}

func (e *implEntry) Bool(key string, value bool) logging.Entry {
	e.event.Bool(key, value)
	return e
}

func (e *implEntry) Time(key string, value time.Time) logging.Entry {
	e.Time(key, value)
	return e
}

func (e *implEntry) Duration(key string, duration time.Duration) logging.Entry {
	e.event.Dur(key, duration)
	return e
}

func (e *implEntry) Float64(key string, value float64) logging.Entry {
	e.event.Float64(key, value)
	return e
}

func (e *implEntry) Uint64(key string, value uint64) logging.Entry {
	e.event.Uint64(key, value)
	return e
}

func (e *implEntry) Logf(message string, args ...interface{}) {
	e.event.Msgf(message, args...)
}

func (e *implEntry) Log(message string) {
	e.event.Msg(message)
}
