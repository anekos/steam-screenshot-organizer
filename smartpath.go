package main

import (
  "fmt"
  "path/filepath"
)


func FindScreenshotDirectory(path string) (string, error) {
  files, err := filepath.Glob(path + "*")
  if err != nil {
    return "", err
  }

  if len(files) < 1 {
    return "", fmt.Errorf("Directory not found")
  }

  return files[0], nil
}
