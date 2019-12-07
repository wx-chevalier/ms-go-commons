package errutils

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"syscall"

	"golang.org/x/xerrors"
)

func Wrapf(err error, format string, args ...interface{}) error {
	errString := fmt.Sprintf(format, args...)
	return xerrors.Errorf("%s: %w", errString, err)
}

func Errorf(format string, args ...interface{}) error {
	return xerrors.Errorf(format, args...)
}

func Errors(msgs []string) error {
	return xerrors.Errorf(strings.Join(msgs, "\n"))
}

func New(text string) error {
	return xerrors.New(text)
}

func Is(err, target error) bool {
	return xerrors.Is(err, target)
}

// IsExitStatusError ...
func IsExitStatusError(err error) bool {
	return IsExitStatusErrorStr(err.Error())
}

// IsExitStatusErrorStr ...
func IsExitStatusErrorStr(errString string) bool {
	// example exit status error string: exit status 1
	var rex = regexp.MustCompile(`^exit status [0-9]{1,3}$`)
	return rex.MatchString(errString)
}

// CmdExitCodeFromError ...
func CmdExitCodeFromError(err error) (int, error) {
	cmdExitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus, ok := exitError.Sys().(syscall.WaitStatus)
			if !ok {
				return 1, errors.New("Failed to cast exit status")
			}
			cmdExitCode = waitStatus.ExitStatus()
		}
		return cmdExitCode, nil
	}
	return 0, nil
}
