// +build window

package godaemon

import (
	"io"
	"os"
)

// DaemonAttr describes the options that apply to daemonization
type DaemonAttr struct {
	ProgramName   string      // child's os.Args[0]; copied from parent if empty
	CaptureOutput bool        // whether to capture stdout/stderr
	Files         []**os.File // files to keep open in the daemon
}

func MakeDaemon(attrs *DaemonAttr) (io.Reader, io.Reader, error) {
	return nil, nil, nil
}
