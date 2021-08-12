package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"

	log "github.com/sirupsen/logrus"
)

func init() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestProcessFileChange(t *testing.T) {
	RunArgs.SourceFile, _ = filepath.Abs("testdata/acme_v1.golden")
	RunArgs.TargetDir, _ = filepath.Abs("testdata")
	RunArgs.ProducePEM = true

	processFileChange(RunArgs.SourceFile, "all")
	assert.FileExists(t, filepath.Join("testdata", "baz.foo.bar.key"))
	assert.FileExists(t, filepath.Join("testdata", "baz.foo.bar.pem"))
	assert.FileExists(t, filepath.Join("testdata", "foo.bar.key"))
	assert.FileExists(t, filepath.Join("testdata", "foo.bar.pem"))

	RunArgs.SourceFile, _ = filepath.Abs("testdata/acme_v2.golden")
	RunArgs.ProducePKCS = true
	RunArgs.Domains = "all"

	processFileChange(RunArgs.SourceFile, RunArgs.Domains)
	assert.FileExists(t, filepath.Join("testdata", "baz.foo.bar.pkcs12"))
	assert.FileExists(t, filepath.Join("testdata", "foo.bar.pkcs12"))
}
