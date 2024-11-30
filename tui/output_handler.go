package tui

import (
	"fmt"
	"mainak55512/ezgit/command"
)

func (op Outputs) RunCommands() {
	op.Handler()
}

// TODO: need to implement the commands for each input.
func (op Outputs) Handler() string {
	switch op.text_output {
	case "Push updates to Remote":
		fmt.Println("Action: ", "Push")
		if err := command.GitPushExec(); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		return "Push"
	case "Pull updates from Remote":
		fmt.Println("Action: ", "Pull")
		if err := command.GitPullExec(); err != nil {
			return fmt.Sprintf("%s, %s", "something went wrong", err)
		}
		return "Pull"
	case "Fetch from pull request":
		return "Fetch"
	case "Create new Local Branch":
		return "Create branch"
	default:
		return "Something went wrong"
	}
}
