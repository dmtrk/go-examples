package docker

import (
	"path/filepath"
	"os"
	"io"
	"bufio"
	"strings"
	"log"
)

const (
	SWARM_CONFIG_FILE_NAME_PATTERN = "*.cfg"
	SWARM_SECRETS_DIR = "/run/secrets/"
	SWARM_CONFIGS_DIR = "/"
)

func FindAndParseProperties(args []string) (map[string]string, error) {
	var err error
	var files []string
	properties := make(map[string]string, 10)
	if len(args) > 1 && len(args[1]) > 0 {
		loadPropertiesFromFile(&properties, filepath.Clean(args[1]))
	} else {
		// swarm 'configs' - swarm mounts config files under '/'
		files, err = filepath.Glob(filepath.Join(SWARM_CONFIGS_DIR, SWARM_CONFIG_FILE_NAME_PATTERN))
		if err == nil {
			for _, config := range files {
				err = loadPropertiesFromFile(&properties, config)
				//log.Printf("loadconfigs(), %v: ", properties)
			}
		}
		// swarm 'secrets' - swarm mounts secrets files under '/run/secrets/'
		files, err = filepath.Glob(filepath.Join(SWARM_SECRETS_DIR, SWARM_CONFIG_FILE_NAME_PATTERN))
		if err == nil {
			for _, secret := range files {
				err = loadPropertiesFromFile(&properties, secret)
				//log.Printf("loadsecrets(), %v: ", properties)
			}
		}
	}
	return properties, err
}

func loadPropertiesFromFile(properties *map[string]string, filename string) (error) {
	log.Printf("loadPropertiesFromFile(%v)", filename)
	file, err := os.Open(filename)
	if err == nil {
		loadProperties(properties, file)
	}
	defer file.Close()
	return err
}

func loadProperties(properties *map[string]string, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if (len(line) > 0 && strings.Contains(line, "=")) {
			keyVals := strings.Split(line, "=")
			if (len(keyVals) > 1) {
				(*properties)[strings.TrimSpace(keyVals[0])] = strings.TrimSpace(keyVals[1])
			}
		}
	}
}
