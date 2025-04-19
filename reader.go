package goconfiglib

import (
	"bufio"
	"os"
	"strings"
)

func LoadConfigs(filePath string, settings Settings) (*Configs, error) {
	if settings.UseXDGConfigHome {
		filePath = prependXDGConfigPath(filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rootSection *Section

	rootSection, err = parseConfigFile(scanner)

	if err != nil {
		return nil, err
	}
	result := Configs{
		FilePath: filePath,
		Root:     *rootSection,
	}
	return &result, nil
}

func parseConfigFile(scanner *bufio.Scanner) (*Section, error) {
	rootSection := Section{
		Parent:      nil,
		RawValue:    "",
		Name:        "global",
		Values:      make(map[string]string),
		Subsections: make([]Section, 0),
	}
	currentSection := &rootSection

	for scanner.Scan() {
		line := scanner.Text()
		currentSection.RawValue += line

		trimmedLine := strings.TrimSpace(line)

		if isEmptyTrimmedLine(trimmedLine) || isCommentTrimmedLine(trimmedLine) {
			continue
		}

		if isSectionTrimmedLine(trimmedLine) {
			newSection := Section{
				Parent:      nil,
				Name:        trimmedLine[1 : len(trimmedLine)-1],
				RawValue:    line,
				Subsections: make([]Section, 0),
				Values:      make(map[string]string),
			}
			newSection.Parent = &rootSection
			rootSection.Subsections = append(rootSection.Subsections, newSection)
			currentSection = &newSection
			continue
		}
		values := strings.Split(trimmedLine, "=")
		if len(values) < 2 {
			continue
		}
		key := strings.TrimSpace(values[0])
		value := strings.TrimSpace(strings.Join(values[1:], "="))
		for line[len(line)-1] == '\\' {
			scanner.Scan()
			line = scanner.Text()
			currentSection.RawValue += line

			trimmedLine := strings.TrimSpace(line)
			value = value[:len(value)-1] + "\n" + trimmedLine

		}
		currentSection.Values[key] = value
	}

	return &rootSection, nil
}

func isCommentTrimmedLine(trimmedLine string) bool {
	return trimmedLine[0] == ';' || trimmedLine[0] == '#'
}

func isEmptyTrimmedLine(trimmedLine string) bool {
	return len(trimmedLine) == 0
}

func isSectionTrimmedLine(trimmedLine string) bool {
	return trimmedLine[0] == '[' && trimmedLine[len(trimmedLine)-1] == ']'
}
