package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// git add
func gitAdd() error {
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// git commit
func gitCommit() {
	msg := "Commit from ezgit"
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Run()
}

// Get the current branch
func GetBranch() (string, error) {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(branch), "\n"), nil
}

// git pull
func gitPull(branch string) error {
	cmd := exec.Command("git", "pull", "origin", branch, "--allow-unrelated-histories")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error in pull", err)
		return err
	}
	return nil
}

// git push
func gitPush(branch string) error {
	cmd := exec.Command("git", "push", "-u", "origin", branch)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error in push", err)
		return err
	}
	return nil
}

// List all available branches on the repo
func ListGitBranch() ([]string, error) {
	cmd := exec.Command("git", "branch")
	result, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var branchNames []string
	stringList := strings.Split(string(result), "\n")
	for _, x := range stringList {
		if x != "" {
			x := strings.Replace(x, "*", "", -1)
			branchNames = append(branchNames, strings.TrimSpace(x))
		}
	}
	// fmt.Println("Available branches: ", branchNames, "length: ", len(branchNames))
	return branchNames, nil
}

// git switch
func SwitchGitBranch(sourcebranch, destinationBranch string) error {
	if err := gitAdd(); err != nil {
		fmt.Println("Error in git add")
		return err
	}
	gitCommit()
	cmd := exec.Command("git", "switch", destinationBranch)
	if err := cmd.Run(); err != nil {
		return err
	}
	if sourcebranch != destinationBranch {
		if err := MergeGitBranch(sourcebranch); err != nil {
			return err
		}
	}
	return nil
}

// git checkout
func CreateGitBranch(baseBranch, newBranch string) error {
	if err := gitAdd(); err != nil {
		fmt.Println("Error in git add")
		return err
	}
	gitCommit()
	cmd := exec.Command("git", "checkout", "-b", newBranch)
	if err := cmd.Run(); err != nil {
		return err
	}
	if err := MergeGitBranch(baseBranch); err != nil {
		return err
	}
	return nil
}

// git merge
func MergeGitBranch(branch string) error {
	cmd := exec.Command("git", "merge", branch)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// git branch -D
func DeleteGitBranch(branch string) error {
	cmd := exec.Command("git", "branch", "-D", branch)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Handle git pull with all necessary commands like add, commit etc.
func GitPullExec(baseBranch string) error {
	if err := gitAdd(); err != nil {
		fmt.Println("Error in git add")
		return err
	}
	gitCommit()
	fmt.Println("Pulling from: ", baseBranch)
	if err := gitPull(baseBranch); err != nil {
		return err
	}
	return nil
}

// Handle git push with all necessary commands like add, commit etc.
func GitPushExec(baseBranch string) error {
	if err := GitPullExec(baseBranch); err != nil {
		return err
	}
	branch, err := GetBranch()
	if err != nil {
		return err
	}
	if err := gitPush(branch); err != nil {
		return err
	}
	return nil
}
