package godaemon

import (
	"fmt"
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
	fmt.Println("call window lib function .....")
	return nil, nil, nil
}
