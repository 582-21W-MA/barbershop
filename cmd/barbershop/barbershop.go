package barbershop

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	OutputDir string
	Quiet     bool
)

func Run(inputDir string) error {
	log.SetFlags(0)

	assetsDir := filepath.Join(inputDir, "assets")
	globalDataPath := filepath.Join(inputDir, "global.json")
	partialsDir := filepath.Join(inputDir, "_partials")
	inputDirname := filepath.Base(inputDir)
	OutputDir = strings.Replace(inputDir, inputDirname, "site", -1)

	// globalData is declared here to avoid making multiple system calls.
	// There is only one global data file.
	globalData, err := parseDataFromFile(globalDataPath)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("parsing global data file: %w", err)
	}

	pageCount := pageCounter{0}

	err = filepath.WalkDir(inputDir, func(entryPath string, entry fs.DirEntry, err error) error {
		// filepath.WalkDir's error is passed to this function.
		// errors.Is(err, fs.ErrNotExist) is used to verify inputDir
		// exists without having to make two system calls (one to verify
		// it exists, and one to walk its content).
		if errors.Is(err, fs.ErrNotExist) {
			pwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("getting path of working directory")
			}
			log.Fatalf("Input directory %q not found in %q.\n", inputDir, pwd)
		} else if filepath.Ext(entryPath) != ".mustache" {
			return nil
		}

		tmpl := newTemplate(entryPath)
		if tmpl.isPartial(partialsDir) {
			return nil
		}
		if err = tmpl.transform(inputDir, partialsDir, OutputDir, globalData); err != nil {
			return err
		}

		pageCount.incr()

		return nil
	})
	if err != nil {
		return fmt.Errorf("walking directory %q: %v", inputDir, err)
	}

	if err = copyFile(assetsDir, OutputDir); err != nil {
		return fmt.Errorf("copying assets directory %q to %q: %w",
			assetsDir, OutputDir, err)
	}

	if Quiet {
		return nil
	}

	fmt.Printf("All mustaches shaved. %s\n", pageCount.print())

	return nil
}
