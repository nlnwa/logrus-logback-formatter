package logback

import (
	. "github.com/sirupsen/logrus"
	"runtime"
)

func init() {
	SetFormatter(new(Formatter))
}

type LogBack struct {
	*Logger
}

func functionName(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	return runtime.FuncForPC(pc).Name()
}

func (logger *LogBack) withLogger() *Entry {
	return WithField("logger", functionName(2))
}

func (logger *LogBack) WithField(key string, value interface{}) *Entry {
	return logger.withLogger().WithField(key, value)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (logger *LogBack) WithFields(fields Fields) *Entry {
	return logger.withLogger().WithFields(fields)
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func (logger *LogBack) WithError(err error) *Entry {
	return logger.withLogger().WithError(err)
}

func (logger *LogBack) Debugf(format string, args ...interface{}) {
	if logger.Logger.Level >= DebugLevel {
		entry := logger.withLogger().newEntry()
		entry.Debugf(format, args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Infof(format string, args ...interface{}) {
	if logger.withLogger().level() >= InfoLevel {
		entry := logger.withLogger().newEntry()
		entry.Infof(format, args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Printf(format string, args ...interface{}) {
	entry := logger.withLogger().newEntry()
	entry.Printf(format, args...)
	logger.withLogger().releaseEntry(entry)
}

func (logger *LogBack) Warnf(format string, args ...interface{}) {
	if logger.withLogger().level() >= WarnLevel {
		entry := logger.withLogger().newEntry()
		entry.Warnf(format, args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Warningf(format string, args ...interface{}) {
	if logger.withLogger().level() >= WarnLevel {
		entry := logger.withLogger().newEntry()
		entry.Warnf(format, args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Errorf(format string, args ...interface{}) {
	if logger.withLogger().level() >= ErrorLevel {
		entry := logger.withLogger().newEntry()
		entry.Errorf(format, args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Fatalf(format string, args ...interface{}) {
	if logger.withLogger().level() >= FatalLevel {
		entry := logger.withLogger().newEntry()
		entry.Fatalf(format, args...)
		logger.withLogger().releaseEntry(entry)
	}
	Exit(1)
}

func (logger *LogBack) Panicf(format string, args ...interface{}) {
	if logger.withLogger().level() >= PanicLevel {
		entry := logger.withLogger().newEntry()
		entry.Panicf(format, args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Debug(args ...interface{}) {
	if logger.withLogger().level() >= DebugLevel {
		entry := logger.withLogger().newEntry()
		entry.Debug(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Info(args ...interface{}) {
	if logger.withLogger().level() >= InfoLevel {
		entry := logger.withLogger().newEntry()
		entry.Info(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Print(args ...interface{}) {
	entry := logger.withLogger().newEntry()
	entry.Info(args...)
	logger.withLogger().releaseEntry(entry)
}

func (logger *LogBack) Warn(args ...interface{}) {
	if logger.withLogger().level() >= WarnLevel {
		entry := logger.withLogger().newEntry()
		entry.Warn(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Warning(args ...interface{}) {
	if logger.withLogger().level() >= WarnLevel {
		entry := logger.withLogger().newEntry()
		entry.Warn(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Error(args ...interface{}) {
	if logger.withLogger().level() >= ErrorLevel {
		entry := logger.withLogger().newEntry()
		entry.Error(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Fatal(args ...interface{}) {
	if logger.withLogger().level() >= FatalLevel {
		entry := logger.withLogger().newEntry()
		entry.Fatal(args...)
		logger.withLogger().releaseEntry(entry)
	}
	Exit(1)
}

func (logger *LogBack) Panic(args ...interface{}) {
	if logger.withLogger().level() >= PanicLevel {
		entry := logger.withLogger().newEntry()
		entry.Panic(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Debugln(args ...interface{}) {
	if logger.withLogger().level() >= DebugLevel {
		entry := logger.withLogger().newEntry()
		entry.Debugln(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Infoln(args ...interface{}) {
	if logger.withLogger().level() >= InfoLevel {
		entry := logger.withLogger().newEntry()
		entry.Infoln(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Println(args ...interface{}) {
	entry := logger.withLogger().newEntry()
	entry.Println(args...)
	logger.withLogger().releaseEntry(entry)
}

func (logger *LogBack) Warnln(args ...interface{}) {
	if logger.withLogger().level() >= WarnLevel {
		entry := logger.withLogger().newEntry()
		entry.Warnln(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Warningln(args ...interface{}) {
	if logger.withLogger().level() >= WarnLevel {
		entry := logger.withLogger().newEntry()
		entry.Warnln(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Errorln(args ...interface{}) {
	if logger.withLogger().level() >= ErrorLevel {
		entry := logger.withLogger().newEntry()
		entry.Errorln(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

func (logger *LogBack) Fatalln(args ...interface{}) {
	if logger.withLogger().level() >= FatalLevel {
		entry := logger.withLogger().newEntry()
		entry.Fatalln(args...)
		logger.withLogger().releaseEntry(entry)
	}
	Exit(1)
}

func (logger *LogBack) Panicln(args ...interface{}) {
	if logger.withLogger().level() >= PanicLevel {
		entry := logger.withLogger().newEntry()
		entry.Panicln(args...)
		logger.withLogger().releaseEntry(entry)
	}
}

//When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func (logger *LogBack) SetNoLock() {
	logger.withLogger().mu.Disable()
}