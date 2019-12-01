package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	gen := &Generator{}

	flag.IntVar(&gen.Day, "day", 0, "the day to generate code for (required)")
	flag.IntVar(&gen.Year, "year", 0, "the year to generate code for (required)")
	flag.StringVar(&gen.PackageName, "package", "", "the name of the package to create")
	flag.StringVar(&gen.ImportPath, "import", "", "the import path of the package to create")
	flag.StringVar(&gen.OutputDir, "dir", "", "the directory to write the files into")
	flag.StringVar(&gen.Cookie, "cookie", "", "your Advent of Code authentication cookie")
	flag.BoolVar(&gen.Overwrite, "overwrite", false, "whether to overwrite existing files")

	flag.Parse()

	err := gen.Complete()
	if err != nil {
		return fmt.Errorf("failed to complete generator values: %w", err)
	}

	err = gen.Validate()
	if err != nil {
		return fmt.Errorf("generator validation failed: %w", err)
	}

	err = gen.Run()
	if err != nil {
		return fmt.Errorf("file generation failed: %w", err)
	}

	return nil
}

// A Generator creates a directory with all contents required to kickstart
// a solution to a puzzle in the Advent of Code calendar.
type Generator struct {
	// The Day and Year to generate code for.
	Day, Year int
	// The name of the package to generate and its import path.
	PackageName, ImportPath string
	// The directory to write the files into.
	OutputDir string
	// Advent of Code authentication Cookie.
	Cookie string
	// Whether to Overwrite existing files.
	Overwrite bool
}

// Complete fills in missing values in g.
func (g *Generator) Complete() error {
	if g.Day == 0 {
		return fmt.Errorf("no day is set")
	}
	if g.Year == 0 {
		return fmt.Errorf("no year is set")
	}

	if g.ImportPath == "" {
		moduleName, err := getModuleName()
		if err != nil {
			return fmt.Errorf("could not get module name: %w", err)
		}
		g.ImportPath = filepath.Join(
			moduleName,
			fmt.Sprintf("%04d", g.Year),
			fmt.Sprintf("day%02d", g.Day),
		)
	}

	if g.PackageName == "" {
		g.PackageName = filepath.Base(g.ImportPath)
	}

	if g.OutputDir == "" {
		moduleName, err := getModuleName()
		if err != nil {
			return fmt.Errorf("could not get module name: %w", err)
		}

		rootDir, err := getRootDir()
		if err != nil {
			return fmt.Errorf("could not get module's root directory: %w", err)
		}

		relImportPath, err := filepath.Rel(moduleName, g.ImportPath)
		if err != nil {
			return fmt.Errorf("failed to infer output directory from module name %q and import path %q", moduleName, g.ImportPath)
		}
		g.OutputDir = filepath.Clean(filepath.Join(rootDir, relImportPath))
	}

	return nil
}

// Validate checks that the values in g make sense.
func (g *Generator) Validate() error {
	if g.Day == 0 {
		return fmt.Errorf("no day is set")
	}
	if g.Year == 0 {
		return fmt.Errorf("no year is set")
	}

	moduleName, err := getModuleName()
	if err != nil {
		return fmt.Errorf("failed to get Go module name: %w", err)
	}
	if !strings.HasPrefix(g.ImportPath, moduleName) {
		return fmt.Errorf("import path %q does not start with module name %q", g.ImportPath, moduleName)
	}

	if filepath.Base(g.ImportPath) != g.PackageName {
		return fmt.Errorf("import path %q and package name %q do not match", g.ImportPath, g.PackageName)
	}

	if filepath.Base(g.OutputDir) != g.PackageName {
		return fmt.Errorf("output directory %q and package name %q do not match", g.OutputDir, g.PackageName)
	}

	if g.Cookie == "" {
		return fmt.Errorf("authentication cookie is not set")
	}

	return nil
}

// Run generates files based on values in g.
func (g *Generator) Run() error {
	err := os.MkdirAll(g.OutputDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory %q and any necessary parents: %w", g.OutputDir, err)
	}

	files := []struct {
		output, template string
	}{
		{fmt.Sprintf("%s.go", g.PackageName), "day.go.tmpl"},
		{fmt.Sprintf("%s_test.go", g.PackageName), "day_test.go.tmpl"},
	}

	for _, f := range files {
		err := g.renderTemplate(f.output, f.template)
		if err != nil {
			return fmt.Errorf("failed to render file %q: %w", f.output, err)
		}
	}

	inputData, err := g.downloadInput()
	if err != nil {
		return fmt.Errorf("failed to download puzzle input: %w", err)
	}

	inputFile := filepath.Join(g.OutputDir, "testdata", "input.txt")
	err = g.writeInput(inputFile, inputData)
	if err != nil {
		return fmt.Errorf("failed to write puzzle input to file %q: %w", inputFile, err)
	}

	return nil
}

// renderTemplate renders templateFile into outputFile, based on values of g.
func (g *Generator) renderTemplate(outputFile, templateFile string) error {
	outputFile = filepath.Join(g.OutputDir, outputFile)
	templateFile = filepath.Join(getBaseDir(), "templates", templateFile)

	if !g.Overwrite && fileExists(outputFile) {
		return fmt.Errorf("file %q already exists and overwriting is disabled", outputFile)
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template file %q: %w", templateFile, err)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create/open file %q: %w", outputFile, err)
	}
	defer f.Close()

	err = tmpl.ExecuteTemplate(f, filepath.Base(templateFile), g)
	if err != nil {
		return fmt.Errorf("failed to execute template %q: %w", templateFile, err)
	}

	return nil
}

func (g *Generator) writeInput(file string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(file), 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory %q and any necessary parents: %w", filepath.Dir(file), err)
	}

	if !g.Overwrite && fileExists(file) {
		return fmt.Errorf("file %q already exists and overwriting is disabled", file)
	}

	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write input to file %q: %w", file, err)
	}

	return nil
}

func (g *Generator) downloadInput() ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", g.Year, g.Day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request for url %q: %w", url, err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: g.Cookie})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query url %q: %w", url, err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read input from url %q: %w", url, err)
	}

	return data, nil
}

// getBaseDir returns the directory containing this source file.
func getBaseDir() string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Dir(f)
}

// getRootDir returns the Go module's root directory.
func getRootDir() (string, error) {
	for dir := getBaseDir(); dir != "/"; dir = filepath.Dir(dir) {
		if fileExists(filepath.Join(dir, "go.mod")) {
			return dir, nil
		}
	}

	return "", fmt.Errorf("could not find go.mod file")
}

// getModuleName returns the name of the Go module this code belongs to.
func getModuleName() (string, error) {
	rootDir, err := getRootDir()
	if err != nil {
		return "", fmt.Errorf("failed to find module's root directory: %w", err)
	}

	modFile := filepath.Join(rootDir, "go.mod")
	f, err := os.Open(modFile)
	if err != nil {
		return "", fmt.Errorf("failed to open file %q: %w", modFile, err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("failed to read contents of file %q: %w", modFile, err)
	}

	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		if strings.HasPrefix(l, "module") {
			words := strings.Split(l, " ")
			return words[1], nil
		}
	}

	return "", fmt.Errorf("no line starting with \"module\" found in file %q", modFile)
}

// getTmplDir returns the directory where file templates are stored.
func getTmplDir() string {
	return filepath.Join(getBaseDir(), "templates")
}

// fileExists checks whether filename exists.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		fmt.Println(filename)
		panic(err)
	}
	return !info.IsDir()
}
