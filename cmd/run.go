package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var executeSnippet bool

var runCmd = &cobra.Command{
	Use:   "run <id>",
	Short: "Preview or execute a saved snippet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			return fmt.Errorf("snippet ID must be a positive number")
		}

		item, err := snippet.Get(id)
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("snippet %d was not found", id)
		}
		if err != nil {
			return err
		}

		if !executeSnippet {
			fmt.Fprintf(cmd.OutOrStdout(), "Preview (not executed): %s\n", item.Command)
			return nil
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Execute this command? [y/N]: %s\n", item.Command)
		var answer string
		if _, err := fmt.Fscanln(cmd.InOrStdin(), &answer); err != nil {
			return fmt.Errorf("reading response: %w", err)
		}
		if !strings.EqualFold(answer, "y") && !strings.EqualFold(answer, "yes") {
			fmt.Fprintln(cmd.OutOrStdout(), "Execution cancelled.")
			return nil
		}

		return runCommand(item.Command)
	},
}

func init() {
	runCmd.Flags().BoolVarP(&executeSnippet, "execute", "x", false,
		"execute the snippet after confirmation")
	rootCmd.AddCommand(runCmd)
}

func runCommand(command string) error {
	var shellCommand *exec.Cmd
	if runtime.GOOS == "windows" {
		shellCommand = exec.Command("cmd", "/C", command)
	} else {
		shellCommand = exec.Command("sh", "-c", command)
	}
	shellCommand.Stdout = os.Stdout
	shellCommand.Stderr = os.Stderr
	shellCommand.Stdin = os.Stdin
	return shellCommand.Run()
}
