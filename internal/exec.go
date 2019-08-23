package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// ProgRunner is the interface that runs child programmes with the provided environment variables
type ProgRunner interface {
	Execute() error
}

// ProgRunnerImpl is a type that implements the ProgRunner interface
type ProgRunnerImpl struct {
	ClearEnv  bool
	EnvPath   string
	ChildProg string
	envVars   []string
}

var progRunner ProgRunner

// SetProgRunner is a convenience function.
// It sets ProgRunner instance to be used during testing.
func SetProgRunner(pr ProgRunner) {
	progRunner = pr
}

// GetProgRunner is a convenience function that is used primarily for testing.
func GetProgRunner() ProgRunner {
	return progRunner
}

// NewProgRunnerImpl returns an instance of the
func NewProgRunnerImpl() *ProgRunnerImpl {
	return &ProgRunnerImpl{envVars: make([]string, 0, 2)}
}

// Execute runs a child programme with a set of environment variables.
func (pr *ProgRunnerImpl) Execute() error {
	if err := pr.prepareEnvironment(); err != nil {
		return err
	}
	// Run the child programme and return its status
	cmd := exec.Command(pr.ChildProg)

	// Configure the command's behaviour
	cmd.Env = pr.envVars
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (pr *ProgRunnerImpl) prepareEnvironment() error {
	// Check whether it's a file or a directory
	src, err := os.Stat(pr.EnvPath)
	if err != nil {
		return err
	}
	// Clear environment variables, if it was requested
	if pr.ClearEnv {
		os.Clearenv()
	} else {
		pr.envVars = append(pr.envVars, os.Environ()...)
	}

	// Populate the environment variables
	if src.Mode().IsDir() {
		return pr.readEnvVarsFromDir()
	}
	return pr.readEnvVarsFromFile("")
}

func (pr *ProgRunnerImpl) readEnvVarsFromFile(path string) error {
	// This is done in order to make this method reusable
	pathToFile := pr.EnvPath
	if path != "" {
		pathToFile = path
	}
	// Open file
	f, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}
	// Read data from file
	s := bufio.NewScanner(f)
	for s.Scan() {
		pr.envVars = append(pr.envVars, fmt.Sprintf("%s=%s", info.Name(), s.Text()))
	}
	return s.Err()
}

func (pr *ProgRunnerImpl) readEnvVarsFromDir() error {
	files, err := ioutil.ReadDir(pr.EnvPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			// For now, let's skip nested directories
			continue
		}
		path := fmt.Sprintf("%s%c%s", pr.EnvPath, os.PathSeparator, file.Name())
		if err := pr.readEnvVarsFromFile(path); err != nil {
			return err
		}
	}

	return nil
}
