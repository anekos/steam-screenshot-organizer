package main

import (
  "fmt"
  "os"
  "path/filepath"
  "github.com/n-marshall/go-cp"
)

func CopyScreenshots(source string, target string) {
  fmt.Println("  Directory:", target)

  dir, err := os.Open(source)
  if err != nil {
    Die(err)
  }
  defer dir.Close()

  fileInfos, err := dir.Readdir(-1)
  if err != nil {
    Die(err)
  }

  for _, fileInfo := range fileInfos {
    if fileInfo.IsDir() {
      continue
    }
    left := filepath.Join(source, fileInfo.Name())
    right := filepath.Join(target, fileInfo.Name())
    if _, err := os.Stat(right); os.IsNotExist(err) {
      fmt.Println("  Copying", left, "to", right)
      if err := cp.CopyFile(left, right); err != nil {
        Die(err)
      }
    }
  }
}

