package rcfm18

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrNoHandler = errors.New("no handler installed")

	addrRe = regexp.MustCompile(`^/(\d{2})/(\d{2})/(..)_(\d{3})$`)
)

func boolToFloat(b bool) float32 {
	if b {
		return 1
	}
	return 0
}

func boolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func floatToBool(f float32) (bool, error) {
	switch f {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("invalid bool value %v", f)
	}
}

func stringToBool(s string) (bool, error) {
	switch s {
	case "0":
		return false, nil
	case "1":
		return true, nil
	default:
		return false, fmt.Errorf("invalid bool value %q", s)
	}
}
