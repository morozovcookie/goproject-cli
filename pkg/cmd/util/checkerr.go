package util

import (
	"fmt"
	"os"
)

func CheckErr(err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
