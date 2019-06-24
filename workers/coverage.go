package workers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/GaneshaThirumurthi/metrics-dashboard/consts"
)

const (
	fileName         = "coverage.out"
	coverageFileName = "coveragefunc.out"
	offset           = 19
)

func New(path, executable string) *Coverage {
	if path == "" {
		path = consts.AksRepoPath
	}
	if executable == "" {
		executable = consts.GoExecutable
	}
	return &Coverage{
		Path:       path,
		Executable: executable,
	}
}

// Coverage has core information required to determine the % coverage of a given repo
type Coverage struct {
	Path       string
	Executable string
}

// GenerateCoverage creates a general go test coverage file
func (c *Coverage) GenerateCoverage(fileName string) error {
	fmt.Println("generating code coverage...")
	args := []string{"test",
		"./...",
		"-coverprofile=" + fileName,
	}
	cmd := exec.Command(c.Executable, args...)
	cmd.Dir = c.Path

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error occured", err)
		return err
	}
	return nil
}

// GenerateFuncCoverage generates a coverage report based on functional coverage
func (c *Coverage) GenerateFuncCoverage(coverageFileName, fileName string) error {
	fmt.Println("generating code function coverage...")
	args := []string{"tool",
		"cover",
		"-func=" + coverageFileName,
		"-o",
		fileName}
	cmd := exec.Command(c.Executable, args...)
	cmd.Dir = c.Path

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error occured", err)
		return err
	}
	return nil
}

// ParseCoverageFile extracts a coverage total
func (c *Coverage) ParseCoverageFile(fileName string) (float64, error) {
	fmt.Println("parsing coverage file...")
	filePath := c.Path + fileName
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		return 0.0, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var coverage float64
	prefix := "(statements)"

	for scanner.Scan() {
		line := scanner.Text()
		if index := strings.Index(line, prefix); index > -1 {
			c, err := strconv.ParseFloat(line[index+offset:index+offset+4], 1)
			if err != nil {
				fmt.Println("Unable to convert number", err)
				return 0.0, err
			}
			coverage = c
		}
	}
	return coverage, nil
}
