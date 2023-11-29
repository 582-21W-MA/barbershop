package barbershop

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

// fileExists() checks if a file exists at the given path.
func fileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return true
}

func copyFile(from, to string) error {
	if !fileExists(from) {
		return nil // TODO: Return an error instead (?)
	}
	if err := os.MkdirAll(to, 0777); err != nil {
		return fmt.Errorf("creating parent directories of %q: %w",
			to, err)
	}
	// TODO: Using cp is not portable. Use native Go code instead.
	// TODO: Return proper error if the cp command is unknown.
	cmd := exec.Command("cp", "-r", from, to)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("executing bash's copy command: %w", err)
	}
	return nil
}

// createFile creates an entry at the given path with the given content,
// and creates parent directories if necessary.
func createFile(path, content string) (*os.File, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return nil, fmt.Errorf("creating parent directories of entry %q: %w",
			path, err)
	}
	file, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("creating entry %q: %w", path, err)
	}
	defer file.Close()
	if _, err = file.WriteString(content); err != nil {
		return nil, fmt.Errorf("writing to entry %q: %w", path, err)
	}
	return file, nil
}
