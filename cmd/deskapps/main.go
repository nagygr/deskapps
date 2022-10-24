package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/rkoesters/xdg/desktop"
)

const (
	uncategorized = "UNCATEGORIZED"
)

var (
	appsByCat = make(map[string][]*desktop.Entry)
)

func main() {
	desktopDirectoryDefault := []string{"/usr/share/applications"}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("WARNING: Couldn't obtain user home directory\n")
	} else {
		homeDeskDir := filepath.Join(homeDir, ".local/share/applications/")
		if _, err := os.Stat(homeDeskDir); errors.Is(err, os.ErrNotExist) {
			fmt.Printf(
				"WARNING: The desktop directory (%s) in the current user's home doesn't exist\n",
				homeDeskDir,
			)
		} else {
			desktopDirectoryDefault = append(
				desktopDirectoryDefault,
				homeDeskDir,
			)
		}
	}

	parser := argparse.NewParser(
		os.Args[0],
		"Lists the currently installed desktop applications",
	)

	deskDirs := parser.StringList(
		"d", "deskDir",
		&argparse.Options{
			Required: false,
			Help:     "Directory that shall be searched for desktop files (multiple entries can appear on the command line)",
			Default:  desktopDirectoryDefault,
		},
	)

	err = parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	desktopFiles := []string{}
	for _, dir := range *deskDirs {
		dfs, err := GetDesktopFiles(dir)

		if err != nil {
			ErrorMsg("Error getting desktop files for %s: %s", dir, err.Error())
		}

		desktopFiles = append(desktopFiles, dfs...)
	}

	entries := []*desktop.Entry{}
	for _, df := range desktopFiles {
		dFile, err := os.Open(df)

		if err != nil {
			ErrorMsg("Error opening desktop file (%s): %s", df, err.Error())
		}

		entry, err := desktop.New(dFile)

		if err != nil {
			ErrorMsg("Error parsing dekstop file (%s): %s", df, err.Error())
		}

		entries = append(entries, entry)
	}

	for _, entry := range entries {
		categories := entry.Categories

		if len(categories) == 0 {
			appsByCat[uncategorized] = append(appsByCat[uncategorized], entry)
		} else {
			for _, category := range categories {
				appsByCat[category] = append(appsByCat[category], entry)
			}
		}
	}

	for category, appList := range appsByCat {
		fmt.Printf("[%s]\n", category)

		for _, app := range appList {
			fmt.Printf("\t%s (%s)\n", app.Name, app.Exec)

			if comment := app.Comment; comment != "" {
				fmt.Printf("\t\t%s\n", comment)
			}
		}

		fmt.Println("")
	}
}

func ErrorMsg(fmtString string, args ...any) {
	fmt.Fprintf(os.Stderr, fmtString+"\n", args...)
	os.Exit(1)
}

func GetDesktopFiles(dirName string) ([]string, error) {
	dir, err := os.Open(dirName)

	if err != nil {
		return nil, fmt.Errorf("Error opening directory: %w", err)
	}

	files, err := dir.ReadDir(-1)

	if err != nil {
		return nil, fmt.Errorf("Error reading files from directory: %w", err)
	}

	desktopFiles := []string{}
	for _, file := range files {
		if fileName := file.Name(); strings.HasSuffix(fileName, ".desktop") {
			desktopFiles = append(desktopFiles, filepath.Join(dirName, fileName))
		}
	}

	return desktopFiles, nil
}
