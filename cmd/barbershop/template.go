package barbershop

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cbroglie/mustache"
)

// parsedJson contains data unmarshalled from JSON files.
type parsedJson map[string]any

// overrideData override a given dataset with the next,
// and returns the result.
func overrideData(datasets ...parsedJson) parsedJson {
	var r = make(parsedJson)
	for _, d := range datasets {
		for k, v := range d {
			r[k] = v
		}
	}
	return r
}

// parseDataFromFile parses JSON from a given file path,
// and returns the data. Data is nil if no data file is found
// at the given path.
func parseDataFromFile(pathname string) (parsedJson, error) {
	b, err := os.ReadFile(pathname)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}
	var d parsedJson // Unmarshal returns only an error.
	if err := json.Unmarshal(b, &d); err != nil {
		// We do not exit because it would stop the watch subcommand.
		log.Printf("Error parsing %q: %v. Data will be nil.\n", pathname, err)
		return nil, nil
	}
	return d, nil
}

// A template is a mustache file that contains HTML and mustache tags.
type template struct {
	path string
	data parsedJson
	raw  []byte
	html string
}

func newTemplate(path string) *template {
	t := template{path: path}
	return &t
}

// isPartials checks if the received template t
// is in the given partials directory.
func (t template) isPartial(partialsDir string) bool {
	return strings.Contains(t.path, partialsDir)
}

// getData looks for a data file alongside a given template and parses the data.
// If globalData is not nil, it overrides the local page data with global data.
func (t *template) getData(globalData parsedJson) error {
	dir := filepath.Dir(t.path)
	dataFilePath := filepath.Join(dir, "data.json")

	var err error
	t.data, err = parseDataFromFile(dataFilePath)
	if err != nil {
		return fmt.Errorf("parsing data from %q: %w", dataFilePath, err)
	}

	if globalData == nil {
		return nil
	}
	t.data = overrideData(t.data, globalData)

	return nil
}

func (t *template) transform(inputDir string, partialsDir string, outputDir string, globalData parsedJson) error {
	if err := t.getData(globalData); err != nil {
		return err
	}
	if err := t.render(partialsDir); err != nil {
		return err
	}
	if _, err := t.createHTMLFile(inputDir, outputDir); err != nil {
		return err
	}
	return nil
}

// render parses the received template and puts the resulting
// HTML in t.html.
func (t *template) render(partialsDir string) error {
	var err error
	t.raw, err = os.ReadFile(t.path)
	if err != nil {
		return fmt.Errorf("reading file %q: %w", t.path, err)
	}
	fp := &mustache.FileProvider{Paths: []string{partialsDir}}
	t.html, err = mustache.RenderPartials(string(t.raw), fp, t.data)
	if err != nil {
		return fmt.Errorf("parsing template file %q: %w", t, err)
	}
	return nil
}

// createHTMLFile creates an HTML file containing the parsed template
// and saves it in the output directory.
func (t template) createHTMLFile(inputDir, outputDir string) (*os.File, error) {
	outputFilePath := strings.Replace(t.path, inputDir, outputDir, -1)
	outputFilePath = strings.Replace(outputFilePath, ".mustache", ".html", -1)
	outputFile, err := createFile(outputFilePath, t.html)
	if err != nil {
		return outputFile, fmt.Errorf("creating HTML file %q: %w",
			outputFile.Name(), err)
	}
	return outputFile, nil
}

type pageCounter struct {
	count int
}

func (c *pageCounter) incr() {
	c.count++
}

func (c pageCounter) print() string {
	if c.count == 1 {
		return "1 page created."
	}
	return fmt.Sprintf("%d pages created.", c.count)
}
