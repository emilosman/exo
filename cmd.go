package exo

import (
	"fmt"
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
	Description: "exo CLI",
	Commands:    []*Z.Cmd{help.Cmd, pageCmd, todayCmd, yesterdayCmd},
}

var pageCmd = &Z.Cmd{
	Name:     "page",
	Summary:  "Opens page",
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

var todayCmd = &Z.Cmd{
	Name:     "today",
	Summary:  "Opens today's dialy file",
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(z *Z.Cmd, _ ...string) error {
		today := time.Now().Format("20060102")
		openDay(today)

		return nil
	},
}

var yesterdayCmd = &Z.Cmd{
	Name:     "yesterday",
	Summary:  "Opens yesterday's dialy file",
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(z *Z.Cmd, _ ...string) error {
		yesterday := time.Now().AddDate(0, 0, -1).Format("20060102")
		openDay(yesterday)

		return nil
	},
}

func openDay(date string) {
	filename := fmt.Sprintf("%s-daily.md", date)
	filePath := filepath.Join(os.Getenv("HOME"), "ruby", "exo", "daily", filename)

	openInVim(filePath)
}

func openInVim(filePath string) {
	fmt.Println("Opening file in Vim:", filePath)
	cmd := exec.Command("vim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error opening file in Vim:", err)
	}
}
