package command

import (
	// "fmt"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func gitAdd() error {
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func gitCommit() {
	msg := "Commit from ezgit"
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Run()
}

func GetBranch() (string, error) {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(branch), "\n"), nil
}

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
	if err := MergeGitBranch(sourcebranch); err != nil {
		return err
	}
	return nil
}

func CreateGitBranch(currentBranch, newBranch string) error {
	if err := gitAdd(); err != nil {
		fmt.Println("Error in git add")
		return err
	}
	gitCommit()
	cmd := exec.Command("git", "checkout", "-b", newBranch)
	if err := cmd.Run(); err != nil {
		return err
	}
	if err := MergeGitBranch(currentBranch); err != nil {
		return err
	}
	return nil
}

func MergeGitBranch(branch string) error {
	cmd := exec.Command("git", "merge", branch)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func DeleteGitBranch(branch string) error {
	cmd := exec.Command("git", "branch", "-D", branch)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

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
