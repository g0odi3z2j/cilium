package config

import (
	"bytes"
	"fmt"
	"os"

	ginkgo "github.com/cilium/cilium/test/ginkgo-ext"

	"github.com/sirupsen/logrus"
)

var (
	// Formatter is the format configuration to write logs into text
	Formatter = logrus.TextFormatter{
		DisableTimestamp: false,
	}

	// TestLogWriter is a buffer in which all logs generated by a test are
	// stored
	TestLogWriter bytes.Buffer
	// TestLogFileName is the file name to dump `TestLogWriter` content when
	// test finish
	TestLogFileName = "test-output.log"
)

// TestLogWriterReset resets the current buffer
func TestLogWriterReset() {
	TestLogWriter.Reset()
}

// LogHook to send logs via `ginkgo.GinkgoWriter`.
type LogHook struct{}

// Levels defined levels to send logs to FireAction
func (h *LogHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

// Fire is a callback function used by logrus to write logs that match in
// the given by `Levels` method
func (h *LogHook) Fire(entry *logrus.Entry) (err error) {
	line, err := Formatter.Format(entry)
	if err == nil {
		fmt.Fprintf(ginkgo.GinkgoWriter, string(line))
	} else {
		fmt.Fprintf(os.Stderr, "LogHook.Fire: unable to format log entry (%v)", err)
	}
	return
}
