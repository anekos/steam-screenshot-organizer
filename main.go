package main

import (
  "fmt"
  "os"
  "path/filepath"
  "github.com/jessevdk/go-flags"
  "github.com/n-marshall/go-cp"
)

func Die(err error) {
  fmt.Println("Error", err)
  os.Exit(1)
}

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

func CopyScreenshots(source string, target string) {
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

type AppOptions struct {
  SteamRoot string `short:"s" long:"steam-root" description:"Steam App Root" default:"C:\\Program Files (x86)\\Steam\\userdata\\"`
}

func Usage() {
  fmt.Println("Usage: steam-screenshots <target>")
  os.Exit(1)
}

func main() {
  options := AppOptions{}
  parser := flags.NewParser(&options, flags.Default)
  parser.Name = "steam-screenshot-organizer"
  parser.Usage = "[OPTIONS] <target>"

  args, err := parser.Parse();
  if err != nil {
    os.Exit(1)
  }

  if len(args) != 1 {
    parser.WriteHelp(os.Stdout)
    os.Exit(1)
  }

  destination := args[0]

  remoteRoot, err := GetRemoteRoot(options.SteamRoot)
  if err != nil {
    Die(err)
  }

  games := SearchGames(remoteRoot)

  for _, game := range games {
    fmt.Println("Game:", game.Name)
    target := filepath.Join(destination, game.Name)
    CopyScreenshots(game.Path, target)
  }
}
