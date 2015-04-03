package logger

import (
	"bytes"
	"testing"
)

var Data = []struct {
	Level   Level
	Message string
	Info    string
	Debug   string
	Warn    string
	Error   string
}{
	{
		Level:   LevelInfo,
		Message: "at=info",
		Info:    "ns=test at=info fn=info\n",
		Debug:   "ns=test at=info fn=debug\n",
		Warn:    "ns=test at=info fn=warn\n",
		Error:   "ns=test at=info fn=error\n",
	},
	{
		Level:   LevelDebug,
		Message: "at=debug",
		Info:    "",
		Debug:   "ns=test at=debug fn=debug\n",
		Warn:    "ns=test at=debug fn=warn\n",
		Error:   "ns=test at=debug fn=error\n",
	},
	{
		Level:   LevelWarn,
		Message: "at=warn",
		Info:    "",
		Debug:   "",
		Warn:    "ns=test at=warn fn=warn\n",
		Error:   "ns=test at=warn fn=error\n",
	},
	{
		Level:   LevelError,
		Message: "at=error",
		Info:    "",
		Debug:   "",
		Warn:    "",
		Error:   "ns=test at=error fn=error\n",
	},
}

var (
	b = bytes.NewBuffer([]byte{})
	l = New(b, "ns=test")
)

func TestPrefix(t *testing.T) {
	b.Reset()
	l.Print("at=prefix")
	assert(t, "ns=test at=prefix\n", b.String())
}

func TestAll(t *testing.T) {
	for _, d := range Data {
		l.SetLevel(d.Level)

		b.Reset()
		l.Info(d.Message + " fn=info")
		assert(t, d.Info, b.String())

		b.Reset()
		l.Debug(d.Message + " fn=debug")
		assert(t, d.Debug, b.String())

		b.Reset()
		l.Warn(d.Message + " fn=warn")
		assert(t, d.Warn, b.String())

		b.Reset()
		l.Error(d.Message + " fn=error")
		assert(t, d.Error, b.String())
	}
}

func assert(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("expected %q; was %q", a, b)
	}
}
