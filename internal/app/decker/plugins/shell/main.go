package main

import (
	"fmt"
	"os"
	"os/exec"
)

type plugin string

// --- wants input: ---
// inputsMap{
//   "host": "example.com",
// }
//
// --- gives output: ---
// resultsMap{
//  "raw_output": "...",
// }
func (p plugin) Run(inputsMap, resultsMap *map[string]string, resultsListMap *map[string][]string) {
	var (
		cmdOut []byte
		err    error
	)

	cmdName := "sh"
	cmdArgs := []string{"-c"}
	cmdArgs = append(cmdArgs, (*inputsMap)["command"])

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running shell command: ", err)
		return
	}

	output := string(cmdOut)

	(*resultsMap)["raw_output"] = output
}

// Plugin is an implementation of github.com/stevenaldinger/decker/pkg/plugins.Plugin
// All this includes is a single function, "Run(*map[string]string, *map[string]string)"
// which takes a map of inputs and an empty map of outputs that the Plugin
// is expected to populate
var Plugin plugin
