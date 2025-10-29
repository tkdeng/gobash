package bash

import (
	"strings"
)

// If will run a bash if [[ logic ]] method, and return true or false to the output for go to process
func If(logic string, dir string, env []string) bool {
	if out, err := RunRaw(`if [[ `+logic+` ]]; then echo "true"; else echo "false"; fi`, dir, env); err == nil {
		if strings.TrimSpace(string(out)) == "true" {
			return true
		}
	}

	return false
}
