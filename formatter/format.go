package formatter

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func Format(path string) error {
	var lines []string
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	defer file.Close()

	if err != nil {
		return errors.New("unable to read file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)

	line := 1
	isInsideImport := false
	var imports []string

	var startLine int
	var endLine int

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if strings.Contains(scanner.Text(), ")") && !isInsideImport {
			endLine = line
		}

		if isInsideImport {
			imports = append(imports, strings.TrimSpace(scanner.Text()))
		}

		if strings.Contains(scanner.Text(), "import") && strings.Contains(scanner.Text(), "(") {
			isInsideImport = true
			startLine = line
		}

		line++
	}

	sort.Strings(imports)
	j := 0

	for i, _ := range lines {
		if i > startLine && i < endLine {
			lines[i] = "	" + imports[j]
			j++
		}
		i++
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}
