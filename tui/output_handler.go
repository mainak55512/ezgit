package tui

import (
	"fmt"
)

func (op Outputs) SpitSelectedOutput() {
	fmt.Println("Selected Output: ", op.Handler())
}

// TODO: need to implement the commands for each input.
func (op Outputs) Handler() string {
	switch op.text_output {
	case "Push updates to Remote":
		return "Push"
	case "Pull updates from Remote":
		return "Pull"
	case "Fetch from pull request":
		return "Fetch"
	case "Create new Local Branch":
		return "Create branch"
	default:
		return "Something went wrong"
	}
}
