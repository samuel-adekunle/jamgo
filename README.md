
[<img src="https://raw.githubusercontent.com/SamtheSaint/jamgo/master/jamgo.svg" style="text-align: center;"></img>](https://raw.githubusercontent.com/SamtheSaint/jamgo/master/jamgo.svg)

Jamgo is a minimal, superfast golang static site generator.

[![Go Report Card](https://goreportcard.com/badge/github.com/SamtheSaint/jamgo)](https://goreportcard.com/report/github.com/SamtheSaint/jamgo) [![PkgGoDev](https://pkg.go.dev/badge/github.com/SamtheSaint/jamgo)](https://pkg.go.dev/github.com/SamtheSaint/jamgo) ![GitHub](https://img.shields.io/github/license/SamtheSaint/jamgo) ![GitHub Release Date](https://img.shields.io/github/release-date/SamtheSaint/jamgo) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/SamtheSaint/jamgo)

# Table of Contents
- [Overview](#overview)
- [Installing](#installing)
- [Getting Started](#getting-started)
- [Command Documentation](#command-documentation)
  * [Build](#build)
  * [Init](#init)
  * [New](#)
    - [Page](#page)
    - [Template](#template)
  * [Shell Completion](#shell-completion)
  * [Help](#help)
- [License](#license)

# Overview

Jamgo is a superfast, minimal golang static site generator. The command line interface is designed to be quick, lightweight and simple to use.

# Installing

> _Requirements_: Golang 1.10+

Installing and using Jamgo is easy and pain free.
If golang is not already installed, [follow this guide](https://golang.org/doc/install) to install the only requirement of jamgo.

Once it's confirmed that `Golang 1.10+` is installed, run the following command to install jamgo:

```bash
go get -u github.com/SamtheSaint/jamgo
```

> This will download and install the latest version of Jamgo.

To get a specific verison, [look here](https://github.com/SamtheSaint/jamgo/releases)

# Getting Started

Getting started with jamgo is made as easy as possible. The `init command` sets up the recommended folder structure for a minimal barebones jamgo application. To get started run:

```bash
jamgo init first-app
cd first-app
```

The folder structure will look like this:

```
first-app
├── assets
│   ├── css
│   ├── img
│   └── js
├── go.mod
└── pages
    ├── index
    │   ├── index.go
    │   └── index.gohtml
    └── templates
        ├── footer.gohtml
        ├── header.gohtml
        └── head.gohtml
```

To build the website, run:

```bash
jamgo build --dir public
```

# Command Documentation

Jamgo has a very reduced set of commands to minimize complexity and enabling ease of development. A copy of the `cobra` generated markdown documentation for each command can be found in the `docs` directory of this repository.

The root command [`jamgo`](./docs/jamgo.md) doesn't do anything and needs to be chained with one of the commands below

Usage: `jamgo [command] [flags]`

Flags: `-h, --help (help for jamgo)`

## Build

The [`build`](./docs/jamgo_build.md) command generates the html website in the directory of choice. It builds the pages in the `pages` directory of the project. Each page will need to be in its own directory, for example:

```
pages
├── about
│   ├── about.go
│   └── about.gohtml
├── blog
│   ├── blog.go
│   └── blog.gohtml
│   └── blog_multiple.gohtml
├── index
│   ├── index.go
│   └── index.gohtml
└── templates
    ├── footer.gohtml
    ├── header.gohtml
    └── head.gohtml
```

The `{page}.gohtml` describes the structure of the page, and can make use of templates defined in the `templates` directory. The pages built this way, will be found in the root of the build directory as `{page}.html`.

If generating multiple pages from a single template is required, the template is designated as `{page}_multiple.gohtml` and the multiple pages generated will be found at the `{page}` directory of the build directory. Eash file would be designated as `{page} | {Title}.html`. For example:

```
public
├── about.html
├── blog
│   ├── blog | 1.html
│   ├── blog | 2.html
│   └── blog | 3.html
├── blog.html
└── index.html
```

The `{page}.go` file acts as a configuration file for each website, and is executed before each page is generated. Page data is passed in through this file, and all data initialization, fetching or manipulation should be done here. For example, the `blog.go` page used above:

```go
package main

import "github.com/SamtheSaint/jamgo/tools"

// PageData supples data for the page to parse.
// Parses {folderName}.gohtml template and
// is stored in the root directory of the build directory
// should be left as nil if only multiple page needed
var PageData tools.Page

// PageDataCollection is used to generate multiple pages from the same template
// uses template {folderName}_multiple.gohtml and is stored in
// {buildDir}/{folderName}
// should be left as nil if only single page needed
var PageDataCollection []tools.Page

func init() {
	PageData = tools.Page{
		Title: "Enter Title Here",
		Data:  nil,
	}
	PageDataCollection = nil
}
```

> _Note:_ Data in PageData and PageDataCollection must be of type _tools.Page_ in order for page to render properly.

> _Tip:_ The Data field is an empty interface so any and all data can go in there.

Usage:
```
jamgo build [flags]
```

Flags: 
```
-d, --dir (build directory, default "public")
-h, --help (help for build)
```

## Init

The [`init`](./docs/jamgo_init.md) command simply sets up a barebone jamgo app with the basic requirements and structure. The folder structure is shown in the [getting started](#getting-started) section.

Usage:
```
jamgo init [name] [flags]
```

Flags: 
```
-h, --help (help for init)
```

## New

The [`new`](./docs/jamgo_new.md) command is the recommend way for generating new resources when building your site. It needs to be chained together with the type of resource used.

Usage:
```
jamgo new [resource] [flags]
```

Flags: 
```
-h, --help (help for new)
```

## Page

The [`Page`](./docs/jamgo_new_page.md) sub command is used for generating new pages for the website. It generates the directory, gohtml templates and go config file. 

Usage:
```
jamgo new page [name] [flags]
```

Flags: 
```
-m, --multiple (toggles generation of multiple page template)
-h, --help (help for page)
```

## Template

The [`Template`](./docs/jamgo_new_template.md) sub command is used for generating new templates to be used in the pages. It generates the gohtml template in the templates directory.

Usage:
```
jamgo new template [name] [flags]
```

Flags: 
```
-h, --help (help for template)
```

## Shell Completion

The [`completion`](/.docs/jamgo_completion.md) command enables shell completion, which in addition to cobra's nearest command suggestion makes for a very smooth developer experience. To set up the shell completion, follow [this guide](./docs/jamgo_completion.md). Currently supported shells include:

- Bash
- Zsh
- Fish
- Powershell

## Help

The help command can be used to bring up the help information for any command.

Usage: `jamgo help [command]`

# License

Jamgo is released under The MIT License. See [LICENSE](./LICENSE).