package log

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rotisserie/eris"
	"strings"
)

// GetTrace returns a formatted stacktrace for err
func GetTrace(err error) string {

	if cause := errors.Cause(err); cause != nil {
		err = cause
	}

	result := eris.ToCustomString(eris.Wrap(err, err.Error()), eris.NewDefaultStringFormat(eris.FormatOptions{
		WithTrace:    true,
		WithExternal: false,
		InvertTrace:  true,
	}))

	resultLines := strings.Split(result, "\n")

	result = ""
	for idx, line := range resultLines {
		if strings.Contains(line, "go-log.") {
			continue
		}
		if idx == 0 {
			result += fmt.Sprintf("%s\n", line)
			continue
		}
		lineParts := strings.SplitN(line, ":", 2)
		if len(lineParts) == 2 {
			result += fmt.Sprintf("%-50s %s\n", lineParts[0], lineParts[1])
		}
	}

	return strings.TrimSuffix(result, "\n")
}
