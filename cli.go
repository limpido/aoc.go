package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"
	"time"

	"aoc.go/shared"
)

func initFile(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: aoc init <year> <day>")
		os.Exit(1)
	}

	yearStr := args[1]
	dayStr := args[2]
	day, _ := strconv.Atoi(dayStr)

	solutionFilePath := filepath.Join("solutions", yearStr, fmt.Sprintf("%02d.go", day))
	if _, err := os.Stat(solutionFilePath); err == nil {
		fmt.Printf("Solution file already exists: %s\n", solutionFilePath)
		os.Exit(1)
	}

	template := `package main

import (
	"strings"
	
	"aoc.go/shared"
)

var Instance shared.Solver = &Solver{}

type Solver struct{}

func parseInput(input string) any {
	lines := strings.Split(input, "\n")
}

func (s *Solver) Part1(input string) any {
	return 0
}

func (s *Solver) Part2(input string) any {
	return 0
}

func main() {}
`

	err := os.MkdirAll(filepath.Dir(solutionFilePath), os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating directories for %s: %v", solutionFilePath, err)
	}

	err = os.WriteFile(solutionFilePath, []byte(template), 0644)
	if err != nil {
		log.Fatalf("Error writing solution file %s: %v", solutionFilePath, err)
	}

	fmt.Printf("Created solution file: %s\n", solutionFilePath)
}

func main() {
	var partFlag int
	var testFlag bool

	flag.IntVar(&partFlag, "part", 0, "Run part 1 or 2. Default to run both.")
	flag.IntVar(&partFlag, "p", 0, "Run part 1 or 2. Default to run both.")
	flag.BoolVar(&testFlag, "test", false, "Run solution against test input file instead of real input.")
	flag.BoolVar(&testFlag, "t", false, "Run solution against test input file instead of real input.")

	flag.Parse()
	args := flag.Args()

	if len(args) > 0 && args[0] == "init" {
		initFile(args)
		return
	}

	if len(args) < 2 {
		fmt.Println("Usage: aoc [options] <year> <day>")
		os.Exit(1)
	}

	yearStr := args[0]
	dayStr := args[1]
	year, _ := strconv.Atoi(yearStr)
	day, _ := strconv.Atoi(dayStr)

	// compile plugin
	srcPath := filepath.Join("solutions", yearStr, fmt.Sprintf("%02d.go", day))
	pluginPath, cleanup := compilePlugin(srcPath)
	defer cleanup()

	// load plugin
	p, err := plugin.Open(pluginPath)
	if err != nil {
		log.Fatalf("Error loading plugin from %s: %v", srcPath, err)
	}

	sym, err := p.Lookup("Instance")
	if err != nil {
		log.Fatalf("Error: Solution file %s must export a variable named 'Instance': %v", srcPath, err)
	}
	solver, ok := sym.(*shared.Solver)
	if !ok {
		log.Fatalf("Error: 'Instance' in %s does not implement shared.Solver interface", srcPath)
	}

	// read input
	var inputPath string
	if testFlag {
		inputPath = filepath.Join("input", yearStr, fmt.Sprintf("%02d-test.txt", day))
	} else {
		inputPath = filepath.Join("input", yearStr, fmt.Sprintf("%02d.txt", day))
	}
	content, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %v", inputPath, err)
	}
	inputStr := string(content)
	inputStr = strings.TrimSpace(inputStr)
	inputStr = strings.ReplaceAll(inputStr, "\r\n", "\n")

	// run solution
	fmt.Printf("--- AoC %d Day %d ---\n", year, day)
	if partFlag == 1 || partFlag == 0 {
		runPart("Part 1", (*solver).Part1, inputStr)
	}
	if partFlag == 2 || partFlag == 0 {
		runPart("Part 2", (*solver).Part2, inputStr)
	}
}

func runPart(name string, fn func(string) any, input string) {
	start := time.Now()
	res := fn(input)
	duration := time.Since(start)

	fmt.Printf("%s: %v (%v)\n", name, res, duration)
}

func compilePlugin(srcPath string) (string, func()) {
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		log.Fatalf("Solution file not found: %s", srcPath)
	}

	// Create a temp file for the binary
	f, err := os.CreateTemp("", "aoc-plugin-*.so")
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	pluginPath := f.Name()

	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginPath, srcPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Compilation failed for %s", srcPath)
	}

	return pluginPath, func() {
		os.Remove(pluginPath)
	}
}
