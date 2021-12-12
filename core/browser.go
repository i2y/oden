package core

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type browser interface {
	open(title string, port, width, height int)
}

func detectBrowser() browser {
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


type chrome struct {
	execPath string
	cmd      *exec.Cmd
}

func (c *chrome) open(_ string, port, width, height int) {
	cmd := exec.Command(
		c.execPath,
		fmt.Sprintf("--app=http://localhost:%d", port),
		fmt.Sprintf("--window-size=%d,%d", width, height),
	)
	cmd.Start()
	c.cmd = cmd
}

func detectChrome() *chrome {
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
		return &chrome{
			execPath: path,
		}
	}
	return nil
}

type edge struct {
	execPath string
	cmd      *exec.Cmd
}

func (e *edge) open(_ string, port, width, height int) {
	cmd := exec.Command(
		e.execPath,
		fmt.Sprintf("--app=http://localhost:%d", port),
		fmt.Sprintf("--window-size=%d,%d", width, height),
	)
	cmd.Start()
	e.cmd = cmd
}

func detectEdge() *edge {
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
		return &edge{
			execPath: path,
		}
	}
	return nil
}

type chromium struct {
	execPath string
	cmd      *exec.Cmd
}

func (c *chromium) Open(_ string, port, width, height int) {
	cmd := exec.Command(
		c.execPath,
		fmt.Sprintf("--app=http://localhost:%d", port),
		fmt.Sprintf("--window-size=%d,%d", width, height),
	)
	cmd.Start()
	c.cmd = cmd
}

func detectChromium() *chrome {
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
		return &chrome{
			execPath: path,
		}
	}
	return nil
}

type firefox struct {
	execPath string
	cmd      *exec.Cmd
}

func (f *firefox) open(_ string, port, width, height int) {
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

func detectFirefox() *firefox {
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
		return &firefox{
			execPath: path,
		}
	}
	return nil
}
