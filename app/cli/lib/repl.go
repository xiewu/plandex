package lib

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"plandex/fs"
	"plandex/term"
)

var ReplSettingsDir string

type ReplMode string

const (
	ReplModeTell ReplMode = "tell"
	ReplModeChat ReplMode = "chat"
)

type ReplState struct {
	Mode    ReplMode
	IsMulti bool
}

var CurrentReplState = ReplState{
	Mode:    ReplModeChat,
	IsMulti: false,
}

type ReplSettings struct {
	State   ReplState
	History []string
}

var ReplCmdAliases = map[string]string{
	"chat":  "ch",
	"tell":  "t",
	"multi": "m",
	"quit":  "q",
	"help":  "h",
	"run":   "r",
	"send":  "s",
}

func init() {
	ReplSettingsDir = filepath.Join(fs.HomePlandexDir, "repl_settings")
}

func EnsureReplSettingsFile() {
	if err := os.MkdirAll(ReplSettingsDir, os.ModePerm); err != nil {
		term.OutputErrorAndExit("Error creating repl history directory: %v", err)
	}

	settingsFile := filepath.Join(ReplSettingsDir, CurrentProjectId+".json")
	if _, err := os.Stat(settingsFile); os.IsNotExist(err) {
		file, err := os.Create(settingsFile)
		if err != nil {
			term.OutputErrorAndExit("Error creating history file: %v", err)
		}
		defer file.Close()

		// Write empty settings object
		var settings ReplSettings

		data, err := json.Marshal(settings)
		if err != nil {
			term.OutputErrorAndExit("Error converting settings to JSON: %v", err)
		}

		if _, err := file.Write(data); err != nil {
			term.OutputErrorAndExit("Error writing to history file: %v", err)
		}
	}
}

func writeSettings(settings *ReplSettings) {
	settingsFile := filepath.Join(ReplSettingsDir, CurrentProjectId+".json")
	data, err := json.Marshal(settings)
	if err != nil {
		term.OutputErrorAndExit("Error converting settings to JSON: %v", err)
	}

	if err := os.WriteFile(settingsFile, data, 0644); err != nil {
		term.OutputErrorAndExit("Error writing settings file: %v", err)
	}
}

func getSettings() *ReplSettings {
	EnsureReplSettingsFile()

	settingsFile := filepath.Join(ReplSettingsDir, CurrentProjectId+".json")

	// Read existing settings
	data, err := os.ReadFile(settingsFile)
	if err != nil {
		term.OutputErrorAndExit("Error reading history file: %v", err)
	}

	// Parse JSON
	var settings ReplSettings
	if err := json.Unmarshal(data, &settings); err != nil {
		term.OutputErrorAndExit("Error parsing history file: %v", err)
	}

	return &settings
}

func LoadState() {
	settings := getSettings()

	if settings.State.Mode != "" {
		CurrentReplState = settings.State
	} else {
		// Write default state
		WriteState()
	}
}

func WriteState() {
	settings := getSettings()
	settings.State = CurrentReplState
	writeSettings(settings)
}

func WriteHistory(input string) {
	settings := getSettings()
	// Add new input
	settings.History = append(settings.History, input)
	writeSettings(settings)
}

func GetHistory() []string {
	settings := getSettings()
	return settings.History
}

// ExecPlandexCommand spawns the same binary, wiring std streams directly so you
// don't have to capture output. Any os.Exit calls in the child won't kill your REPL.
func ExecPlandexCommand(args []string) (string, error) {
	// Create temp file
	tmpFile, err := os.CreateTemp("", "plandex-output-*")
	if err != nil {
		return "", err
	}
	tmpPath := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpPath)

	// Set env vars
	env := append(os.Environ(),
		"PLANDEX_REPL=1",
		"PLANDEX_REPL_OUTPUT_FILE="+tmpPath,
	)

	// Run command
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = env

	// Connect stdin directly
	cmd.Stdin = os.Stdin
	// Use MultiWriter to send output to both buffer and os.Stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return "", nil
		}
		return "", err
	}

	// Read output from temp file
	output, err := os.ReadFile(tmpPath)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
