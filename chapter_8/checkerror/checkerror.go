package checkerror

import (
	"fmt"
	"os"
)

func CheckError(err error)  {
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error()); err != nil {
			os.Exit(1)
			return
		}
		os.Exit(1)
	}
}