package exo

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        "exo",
	Summary:     "exo CLI",
	Usage:       "",
	Version:     "0.0.1",
	Description: "CLI helper for exocortex",
	Commands:    []*Z.Cmd{help.Cmd, pageCmd, listPagesCmd, dayCmd, todayCmd, yesterdayCmd, syncCmd},
}

var pageCmd = &Z.Cmd{
	Name:     "page",
	Summary:  "open a page",
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(z *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return fmt.Errorf("no page specified")
		}

		page := args[0]
		markdown := fmt.Sprintf("%s.md", page)
		filePath := filepath.Join(os.Getenv("HOME"), "ruby", "exo", "pages", markdown)

		openInVim(filePath)

		return nil
	},
}

var listPagesCmd = &Z.Cmd{
	Name:     "list",
	Summary:  "list pages",
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(z *Z.Cmd, _ ...string) error {
		filePath := filepath.Join(os.Getenv("HOME"), "ruby", "exo", "pages")
		cmd := exec.Command("ls", filePath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println("Error listing pages:", err)
		}

		return nil
	},
}

var dayCmd = &Z.Cmd{
	Name:     "day",
	Summary:  "open a daily file",
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(z *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return fmt.Errorf("no page specified")
		}

		day := args[0]
		openDay(day)

		return nil
	},
}

var todayCmd = &Z.Cmd{
	Name:     "today",
	Summary:  "open today's daily file",
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(z *Z.Cmd, _ ...string) error {
		today := time.Now().Format("20060102")
		createToday()
		openDay(today)

		return nil
	},
}

var yesterdayCmd = &Z.Cmd{
	Name:     "yesterday",
	Summary:  "open yesterday's daily file",
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(z *Z.Cmd, _ ...string) error {
		yesterday := time.Now().AddDate(0, 0, -1).Format("20060102")
		openDay(yesterday)

		return nil
	},
}

var syncCmd = &Z.Cmd{
	Name:     "sync",
	Summary:  "git sync exo",
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(z *Z.Cmd, _ ...string) error {
		repoDir := filepath.Join(os.Getenv("HOME"), "ruby", "exo")

		// Define git commands
		gitCommands := [][]string{
			{"git", "add", "."},
			{"git", "commit", "-m", "update"},
			{"git", "push", "kicomp", "master"},
		}

		// Execute the git commands in the specified directory
		for _, cmd := range gitCommands {
			if err := runCommand(repoDir, cmd[0], cmd[1:]...); err != nil {
				return fmt.Errorf("failed to run command %v: %w", cmd, err)
			}
		}

		// Define rsync command
		rsyncArgs := []string{
			"-auhv", "--delete", "--no-perms", "--no-times", "--exclude=.*",
			"-e", "ssh -p 920",
			filepath.Join(os.Getenv("HOME"), "ruby", "exo") + "/",
			"emil@kicomp.xyz:/mnt/md0/exo/",
		}

		// Execute rsync command
		if err := runCommand("", "rsync", rsyncArgs...); err != nil {
			return fmt.Errorf("failed to run rsync command: %w", err)
		}

		return nil
	},
}

func createToday() {
	// Define the template file path
	templateFilePath := filepath.Join(os.Getenv("HOME"), "ruby", "exo", "daily", "daily-template.md")

	// Get the current date
	currentTime := time.Now()
	dateString := currentTime.Format("20060102")
	dayName := currentTime.Weekday().String()

	// Create the new filename
	newFileName := fmt.Sprintf("%s-daily.md", dateString)
	newFilePath := filepath.Join(os.Getenv("HOME"), "ruby", "exo", "daily", newFileName)

	// Check if the new file already exists
	if _, err := os.Stat(newFilePath); err == nil {
		return
	}

	// Read the template file
	templateContent, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	// Create the new content
	newContent := fmt.Sprintf("# %s %s\n\n%s", dateString, dayName, templateContent)

	// Write the new content to the new file
	err = ioutil.WriteFile(newFilePath, []byte(newContent), 0644)
	if err != nil {
		fmt.Println("Error writing new file:", err)
		return
	}

	fmt.Println("New daily file created.")
}

func openDay(date string) {
	filename := fmt.Sprintf("%s-daily.md", date)
	filePath := filepath.Join(os.Getenv("HOME"), "ruby", "exo", "daily", filename)

	openInVim(filePath)
}

func openInVim(filePath string) {
	fmt.Println("exo:", filePath)
	cmd := exec.Command("vim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error opening file in Vim:", err)
	}
}

func runCommand(dir string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if dir != "" {
		cmd.Dir = dir
	}
	return cmd.Run()
}
