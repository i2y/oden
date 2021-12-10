package oden

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Browser interface {
	Open(title string, port, width, height int)
}

func DetectBrowser() Browser {
	if runtime.GOOS == "windows" {
		w := detectWebview2()
		if w != nil {
			return w
		}
	}

	chrome := detectChrome()
	if chrome != nil {
		return chrome
	}

	edge := detectEdge()
	if edge != nil {
		return edge
	}

	chromium := detectChromium()
	if chromium != nil {
		return chromium
	}

	firefox := detectFirefox()
	if firefox != nil {
		return firefox
	}

	return nil
}


type Chrome struct {
	execPath string
	cmd      *exec.Cmd
}

func (c *Chrome) Open(_ string, port, width, height int) {
	cmd := exec.Command(
		c.execPath,
		fmt.Sprintf("--app=http://localhost:%d", port),
		fmt.Sprintf("--window-size=%d,%d", width, height),
	)
	cmd.Start()
	c.cmd = cmd
}

func detectChrome() *Chrome {
	var paths []string
	switch runtime.GOOS {
	case "darwin":
		paths = []string{
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
			"/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary",
			"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
			"/usr/bin/google-chrome-stable",
			"/usr/bin/google-chrome",
		}
	case "windows":
		paths = []string{
			os.Getenv("LocalAppData") + "/Google/Chrome/Application/chrome.exe",
			os.Getenv("ProgramFiles") + "/Google/Chrome/Application/chrome.exe",
			os.Getenv("ProgramFiles(x86)") + "/Google/Chrome/Application/chrome.exe",
			os.Getenv("ProgramFiles(x86)") + "/Microsoft/Edge/Application/msedge.exe",
			os.Getenv("ProgramFiles") + "/Microsoft/Edge/Application/msedge.exe",
		}
	default:
		paths = []string{
			"/usr/bin/google-chrome-stable",
			"/usr/bin/google-chrome",
		}
	}

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return &Chrome{
			execPath: path,
		}
	}
	return nil
}

type Edge struct {
	execPath string
	cmd      *exec.Cmd
}

func (e *Edge) Open(_ string, port, width, height int) {
	cmd := exec.Command(
		e.execPath,
		fmt.Sprintf("--app=http://localhost:%d", port),
		fmt.Sprintf("--window-size=%d,%d", width, height),
	)
	cmd.Start()
	e.cmd = cmd
}

func detectEdge() *Edge {
	var paths []string
	switch runtime.GOOS {
	case "darwin":
		paths = []string{
			"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
		}
	case "windows":
		paths = []string{
			os.Getenv("ProgramFiles(x86)") + "/Microsoft/Edge/Application/msedge.exe",
			os.Getenv("ProgramFiles") + "/Microsoft/Edge/Application/msedge.exe",
		}
	default:
		paths = []string{}
	}

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return &Edge{
			execPath: path,
		}
	}
	return nil
}

type Chromium struct {
	execPath string
	cmd      *exec.Cmd
}

func (c *Chromium) Open(_ string, port, width, height int) {
	cmd := exec.Command(
		c.execPath,
		fmt.Sprintf("--app=http://localhost:%d", port),
		fmt.Sprintf("--window-size=%d,%d", width, height),
	)
	cmd.Start()
	c.cmd = cmd
}

func detectChromium() *Chrome {
	var paths []string
	switch runtime.GOOS {
	case "darwin":
		paths = []string{
			"/Applications/Chromium.app/Contents/MacOS/Chromium",
			"/usr/bin/chromium",
			"/usr/bin/chromium-browser",
		}
	case "windows":
		paths = []string{
			os.Getenv("LocalAppData") + "/Chromium/Application/chrome.exe",
			os.Getenv("ProgramFiles") + "/Chromium/Application/chrome.exe",
			os.Getenv("ProgramFiles(x86)") + "/Chromium/Application/chrome.exe",
		}
	default:
		paths = []string{
			"/usr/bin/chromium",
			"/usr/bin/chromium-browser",
			"/snap/bin/chromium",
		}
	}

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return &Chrome{
			execPath: path,
		}
	}
	return nil
}

type Firefox struct {
	execPath string
	cmd      *exec.Cmd
}

func (f *Firefox) Open(_ string, port, width, height int) {
	cmd := exec.Command(
		f.execPath,
		"-no-remote",
		"-private-window",
		fmt.Sprintf("http://localhost:%d", port),
		"-foreground",
	)
	cmd.Start()
	f.cmd = cmd
}

func detectFirefox() *Firefox {
	var paths []string
	switch runtime.GOOS {
	case "darwin":
		paths = []string{
			"/Applications/Chromium.app/Contents/MacOS/Firefox",
			"/usr/bin/firefox",
		}
	case "windows":
		paths = []string{
			os.Getenv("LocalAppData") + "/Mozilla Firefox/firefox.exe",
			os.Getenv("ProgramFiles") + "/Mozilla Firefox/firefox.exe",
			os.Getenv("ProgramFiles(x86)") + "/Mozilla Firefox/firefox.exe",
		}
	default:
		paths = []string{
			"/usr/bin/firefox",
			"/snap/bin/firefox",
		}
	}

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return &Firefox{
			execPath: path,
		}
	}
	return nil
}
