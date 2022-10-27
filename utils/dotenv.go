package utils

import (
	"errors"
	"fmt"
	_ "fmt"
	"os"
	"path/filepath"
	"strings"
)

// init auto load available .env file in project root
// utils is imported to main.go so this would work fine.
func init() {
	load()
}

// tokenize source string into key-value pairs 
func Tokenize(env_source []byte) (map[string]string, error) {
	index := 0
	kMap := make(map[string]string, 5)

	lenght := len(env_source)
	for index < lenght {
		switch env_source[index] {
			// ignore comments
			case '#': {
				fw_index := index + 1
				for fw_index < lenght {
					if env_source[fw_index] == '\n' {
						break
					}
					fw_index++
				}
				index = fw_index
				break
			}
			case '\n', '\t': {
				index++
				break
			}
			default: {
				var key, value strings.Builder
				// parse key 
				for index < lenght && env_source[index] != '=' {
					key.WriteByte(env_source[index])
					index++
				}
				if index >= lenght {
					return kMap, fmt.Errorf("Invalid variable key")
				}
				index++ 
				// parse value
				for index < lenght && env_source[index] != '\n' {
					value.WriteByte(env_source[index])	
					index++
				}
				kMap[key.String()] = value.String()
				key.Reset()
				value.Reset() 
				break
			}
			
 		}
	}
	return kMap, nil 
}

func load() {
	dir, _ := os.Getwd()
	abs_path := filepath.Join(dir, "/.env")
	if _, err := os.Stat(abs_path); errors.Is(err, os.ErrNotExist) {
		return
	}

	source_in_bytes, err := os.ReadFile(abs_path)
	if err != nil {
		return 
	}
	// parse source to key-value pairs
	kv_pairs, err := Tokenize(source_in_bytes)

	for key, value := range kv_pairs {
		os.Setenv(key, value)
	}
}