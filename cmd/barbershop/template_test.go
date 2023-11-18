package barbershop

import (
	"os"
	"reflect"
	"testing"
)

func TestParseDataFromFile(t *testing.T) {
	tests := []struct {
		name     string
		pathname string
		want     parsedJson
	}{
		{
			"base",
			"testdata/parsedata/data.json",
			parsedJson{"title": "Test"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := parseDataFromFile(test.pathname)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

// TODO: Test for slice (JSON with first level array)
func TestOverrideData(t *testing.T) {
	tests := []struct {
		name                  string
		localData, globalData parsedJson
		want                  parsedJson
	}{
		{
			"noData",
			parsedJson{},
			parsedJson{},
			parsedJson{},
		},
		{
			"local",
			parsedJson{"page_title": "page"},
			parsedJson{},
			parsedJson{"page_title": "page"},
		},
		{
			"localAndGlobal",
			parsedJson{"page_title": "page"},
			parsedJson{"site_title": "site"},
			parsedJson{"page_title": "page", "site_title": "site"},
		},
		{
			"global",
			parsedJson{},
			parsedJson{"site_title": "site"},
			parsedJson{"site_title": "site"},
		},
		{
			"globalOverrrideLocal",
			parsedJson{"page_title": "page", "color": "blue"},
			parsedJson{"site_title": "site", "color": "red"},
			parsedJson{
				"page_title": "page",
				"site_title": "site",
				"color":      "red",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := overrideData(test.localData, test.globalData)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestIsPartial(t *testing.T) {
	tests := []struct {
		name        string
		tmpl        template
		partialsDir string
		want        bool
	}{
		{
			"tmplOutsidePartials",
			template{path: "src/index.mustache"},
			"src/_partials",
			false,
		}, {
			"tmplInsidePartials",
			template{path: "src/_partials/index.mustache"},
			"src/_partials",
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.tmpl.isPartial(test.partialsDir)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestCreateHTMLFile(t *testing.T) {
	tests := []struct {
		name                string
		tmpl                template
		inputDir, outputDir string
		want                string
	}{
		{
			"base",
			template{path: "testdata/createhtml/index.mustache"},
			"testdata/createhtml",
			"testdata/createhtml/site",
			"testdata/createhtml/site/index.html",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.tmpl.createHTMLFile(test.inputDir, test.outputDir)
			if got.Name() != test.want || !fileExists(got.Name()) {
				t.Errorf("got %v, want %v: %v", got, test.want, err)
			}
		})
		t.Cleanup(func() {
			err := os.RemoveAll(test.outputDir)
			if err != nil {
				panic(err)
			}
		})
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		name string
		tmpl template
		want string
	}{
		{
			"simplePartial",
			template{path: "testdata/render/index.mustache"},
			"header",
		},
		{
			"nestedPartial",
			template{path: "testdata/render/index2.mustache"},
			"header",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.tmpl.render("testdata/render/_partials")
			got := test.tmpl.html
			if got != test.want {
				t.Errorf("got %q, want %q: %v", got, test.want, err)
			}
		})
	}
}
