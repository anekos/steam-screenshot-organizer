package main

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "github.com/jessevdk/go-flags"
)

func Die(err error) {
  fmt.Println("Error", err)
  os.Exit(1)
}

type AppOptions struct {
  SteamRoot string `short:"s" long:"steam-root" description:"Steam App Root" default:"C:\\Program Files (x86)\\Steam\\userdata\\"`
  Shutdown bool `short:"x" long:"shutdown" description:"Shutdown OS"`
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
    Die(err)
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
    if smart, err := FindScreenshotDirectory(target); err == nil {
      CopyScreenshots(game.Path, smart)
    } else {
      CopyScreenshots(game.Path, target)
    }
  }

  if options.Shutdown {
    if err := exec.Command("shutdown.exe", "/s", "/t", "0").Run(); err != nil {
      Die(err)
    }
  }
}
