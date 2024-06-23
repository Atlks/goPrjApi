package lib

import (
	"bufio"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

func iniToMap(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	configMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines and comments
		if len(line) == 0 || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		configMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return configMap, nil
}

func iniToMapEmedMap(filename string) (map[string]map[string]string, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}

	result := make(map[string]map[string]string)
	for _, section := range cfg.Sections() {
		sectionName := section.Name()
		if sectionName == ini.DefaultSection {
			sectionName = "default"
		}
		sectionMap := make(map[string]string)
		for _, key := range section.Keys() {
			sectionMap[key.Name()] = key.String()
		}
		result[sectionName] = sectionMap
	}

	return result, nil
}
