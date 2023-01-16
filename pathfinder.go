package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func GetRemoteRoot(root string) (string, error) {
  dir, err := os.Open(root)
  if err != nil {
    return "", err
  }
  defer dir.Close()

  fileInfos, err := dir.Readdir(-1)
  if err != nil {
    return "", err
  }

  for _, fileInfo := range fileInfos {
    return filepath.Join(root, fileInfo.Name(), "\\760\\remote\\"), nil
  }

  return "", fmt.Errorf("Could not find remote directory")
}

type NamedPath struct {
  Name string
  Path string
}

func SearchGames(root string) []NamedPath {
  dir, err := os.Open(root)
  if err != nil {
    Die(err)
  }
  defer dir.Close()

  fileInfos, err := dir.Readdir(-1)
  if err != nil {
    Die(err)
  }

  result := make([]NamedPath, 0)

  for _, fileInfo := range fileInfos {
    item := NamedPath {
      Name: fileInfo.Name(),
      Path: filepath.Join(root, fileInfo.Name(), "screenshots"),
    }
    result = append(result, item)
  }

  return result
}


