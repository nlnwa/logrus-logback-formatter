package logback

import (
	"bytes"
	"fmt"
	"sort"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultTimestampFormat = time.RFC3339
	defaultThread          = "main"
)

type Formatter struct {
	TimestampFormat  string
	DisableSorting   bool
	QuoteEmptyFields bool
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b = &bytes.Buffer{}
	keys := make([]string, 0, len(entry.Data))
	for key := range entry.Data {
		if key == "logger" || key == "thread" {
			continue
		}
		keys = append(keys, key)
	}

	if !f.DisableSorting {
		sort.Strings(keys)
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	// timestamp
	b.WriteString(entry.Time.Format(timestampFormat))

	// space
	b.WriteByte(' ')

	// thread
	b.WriteByte('[')
	if thread, ok := entry.Data["thread"]; ok {
		f.appendValue(b, thread)
	} else {
		f.appendValue(b, defaultThread)
	}
	b.WriteByte(']')

	// space
	b.WriteByte(' ')

	// level
	level := ""
	switch entry.Level {
	case logrus.DebugLevel:
		level = "TRACE"
	case logrus.InfoLevel:
		level = "INFO"
	case logrus.WarnLevel:
		level = "WARN"
	case logrus.FatalLevel:
		fallthrough
	case logrus.PanicLevel:
		fallthrough
	case logrus.ErrorLevel:
		level = "ERROR"
	}
	b.WriteString(level)

	// space
	b.WriteByte(' ')

	// logger
	if logger, ok := entry.Data["logger"]; ok {
		f.appendValue(b, logger)
	} else {
		b.WriteString("unknown")
	}

	// space - space
	b.WriteString(" - ")

	// fields
	b.WriteByte('{')
	for i, key := range keys {
		f.appendKeyValue(b, key, entry.Data[key], len(keys)-1 != i)
	}
	b.WriteByte('}')

	// message
	if entry.Message != "" {
		b.WriteByte(' ')
		b.WriteString(entry.Message)
	}

	// entry buffer
	if entry.Buffer != nil {
		b.WriteByte(' ')
		b.Write(entry.Buffer.Bytes())
	}

	// new line
	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (f *Formatter) needsQuoting(text string) bool {
	if f.QuoteEmptyFields && len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}
	return false
}

func (f *Formatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}, appendComma bool) {
	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)

	if appendComma {
		b.WriteString(", ")
	}
}

func (f *Formatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	if !f.needsQuoting(stringVal) {
		b.WriteString(stringVal)
	} else {
		b.WriteString(fmt.Sprintf("%q", stringVal))
	}
}
