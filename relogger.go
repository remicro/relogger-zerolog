package reZerologer

import (
	"github.com/remicro/api/logging"
	"github.com/rs/zerolog"
)

func New(l *zerolog.Logger) logging.Logger {
	return &implZeroLogger{
		log: l,
	}
}

type implZeroLogger struct {
	log *zerolog.Logger
}

func (log *implZeroLogger) Info() logging.Entry {
	return &implEntry{
		event: log.log.Info(),
	}
}

func (log *implZeroLogger) Error() logging.Entry {
	return &implEntry{
		event: log.log.Error(),
	}
}

func (log *implZeroLogger) Debug() logging.Entry {
	return &implEntry{
		event: log.log.Debug(),
	}
}

func (log *implZeroLogger) Warn() logging.Entry {
	return &implEntry{
		event: log.log.Warn(),
	}
}

func (log *implZeroLogger) Critical() logging.Entry {
	return &implEntry{
		event: log.log.Error(),
	}
}
