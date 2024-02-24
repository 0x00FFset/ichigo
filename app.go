package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

type AppMetadata struct {
	Info struct {
		ProductVersion string `json:"productVersion"`
	} `json:"info"`
}

type GitHubRelease struct {
    TagName string `json:"tag_name"`
    Name    string `json:"name"`
    HtmlUrl string `json:"html_url"`
}

type CustomResponseVersion struct {
    Version string `json:"version"`
    Url     string `json:"url"`
    Message string `json:"message"`
	Color   string `json:"color"`
	DisableButton string `json:"disableButton"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// check for updates
func (a *App) CheckForUpdates() string {
    // The GitHub repository you're interested in (change "owner" and "repo" to the actual values)
    owner := "0x00FFset"
    repo := "ichigo"
	upToDate := "You're currently using the latest version of Ichigo."
	needUpdate := "There's a new version of Ichigo available. Would you like to download it?"
	successColor := "text-green-400"
	errorColor := "text-red-400"
	disableButton := "text-white opacity-20 cursor-not-allowed"
	
	release, err := getLatestRelease(owner, repo)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    response := CustomResponseVersion{
        Version: release.TagName,
        Url:     release.HtmlUrl,
    }
	installedVersion := CheckVersion();

	if release.TagName == installedVersion {
		response.Message = upToDate
		response.Color = successColor
		response.DisableButton = disableButton
	} else {
		response.Message = needUpdate
		response.Color = errorColor
		response.DisableButton = ""
	}

		jsonResponse, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			os.Exit(1)
		}

		return fmt.Sprint(string(jsonResponse))
}

func getLatestRelease(owner, repo string) (*GitHubRelease, error) {
    url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var release GitHubRelease
    if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
        return nil, err
    }

    return &release, nil
}

func CheckVersion() string {
    // Assuming `wails.json` is in the current directory
    content, err := os.ReadFile("wails.json")
    if err != nil {
        fmt.Printf("Error reading wails.json: %v\n", err)
        return ""
    }

    var metadata AppMetadata
    err = json.Unmarshal(content, &metadata)
    if err != nil {
        fmt.Printf("Error parsing JSON: %v\n", err)
        return ""
    }

    return metadata.Info.ProductVersion
}