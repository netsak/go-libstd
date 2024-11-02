package libError

import (
	"log/slog"
)

// FailOnError exit the program with code 1 when the error is not nil
// and logs the given message with default slog.
func FailOnError(err error, msg string) {
	if err != nil {
		slog.Error(msg, "err", err)
		panic(err)
	}
}
