package tui

import (
	"fmt"
	"mainak55512/ezgit/command"
)

func (op Outputs) SpitSelectedOutput() {
	fmt.Println("Selected Output: ", op.Handler())
}

// TODO: need to implement the commands for each input.
func (op Outputs) Handler() string {
	switch op.text_output {
	case "Push updates to Remote":
		if err := command.GitPush(); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
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
