# 🍢 Oden
Oden is a cross-platform desktop GUI toolkit for Gophers.

## Introduction
The goal of Oden is to provide a single fine framework for gophers easy to create desktop GUI applications for several OS platforms. The idea of Oden is to successfully achevie the goal by combining browser technology available on most platforms with the easy cross-compilation that is an important feature of Go.

## Features
- Oden is a pure Go library without cgo.
- Oden utilizes a browser that is already installed on the desktop where an Oden app is installed.
- Oden doesn't require Gophers to have any knowledge of HTML/CSS/JS.
- Oden provides easy to use API. The `widget` module supports decralative UI style programming.

## A Quick Look
This example is a very simple counter app using Oden.

### Source (main.go)
```go
package main

import (
    core "github.com/i2y/oden/core"
    . "github.com/i2y/oden/widget"
)

func main() {
    app := core.NewApp(
        "Counter",
        300, 300,
        counter(0),
    )
    app.Run()
}

func counter(value int) Widget {
    opBtn := func(label string) Widget {
        return Button(label).
            BorderRadius(0).
            FontSize(TwoXLarge)
    }

    // When Go 1.18 is released and generics-related syntax is available,
    // this will be a State, not an IntState. Maybe.
    count := IntState(0)
    return Column(
        Text(count).
            FgColor(White).
            BgColor(Orange).
            BorderColor(Gray).
            FontSize(TwoXLarge),
        Row(
            opBtn("+").OnClick(func(ev core.Event) {
                count.Increment()
            }),
            opBtn("-").OnClick(func(_ core.Event) {
                count.Decrement()
            }),
        ),
    )
}
```

### Build and Run
```sh
$ go build main.go
$ ./main
```

### Screen Shot
<img width="215" alt="Counter" src="https://user-images.githubusercontent.com/6240399/144830545-46059e3c-0c00-41e6-a4c1-dd0f173e4431.png">

## A little detailed explanation
### Source code repository Layout
Oden source code repository is a multi-module repository.
Oden is composed of two modules: `core` module and `widget` module.
- The `core` module provides the interaction with a browser, widget interface required by core module, etc.
- The `widget` module provides a standard set of widgets for Oden.

Note: You do not necessarily need to use the `widget` module. If you want, you can define and use your own widget set module. For this purpose, Oden provides `core` and `widget` as independent modules, to make the boundary between them clear. This also reduces the size of the generated binary when combining core module with your own widget module, without having to include the standard `widget` module.

### Supported Browsers
- WebView2
- Chrome
- Edge
- Chromium
- Firefox

When an Oden application starts, Oden will try to detect a browser installed on the host machine in the order of the list above.
In other words, the order of priority is as follows.
`WebView2 > Chrome > Edge > Chromium > Firefox`

### Limitations
- If you use a browser other than WebView2, the app window will be opened as the browser's one.
- If you use Firefox, address/tool bar won't be hidden.

### Dependencies
Oden currently depends on the following packages.
Thanks to the creators and contributors of each package.
- [go-webview2](https://github.com/jchv/go-webview2)
- [Shoelace](https://shoelace.style/)
- [Hotwired Turbo](https://turbo.hotwired.dev/)
- [EventBus](https://github.com/asaskevich/eventbus)
- and libraries that the above libraries depend on

These dependencies may be changed for internal implementation reasons.

### Distribution/Packaging
Basically, you just need to run `go build` with appropriate `GOOS` and `GOARCH`, but there are a few tips for each platform.

#### Mac OS X
You can create a application bundle (`.app` package) using [appify](https://github.com/machinebox/appify) or [macappshell](https://github.com/Xeoncross/macappshell).

#### Windows
You can run your app in background by specifying the `-ldflags "-H windowsgui"` option in `go build`.
For example:
```
> go build -ldflags "-H windowsgui" main.go
```

You can also use [go-winres](https://github.com/tc-hib/go-winres), [rsrc](https://github.com/akavel/rsrc), or [GoVersionInfo](https://github.com/josephspurrier/goversioninfo) to embed resources (such as the application icon) into your windows app.

## Near Future Plan
- Add documentation
- Add more widgets
- Refine Oden's API a little bit after Generics related features is introduced in Go 1.18

## License
[MIT License](https://github.com/i2y/oden/blob/main/LICENSE)
