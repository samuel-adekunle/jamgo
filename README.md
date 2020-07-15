
[<img src="./jamgo.svg" style="text-align: center;"></img>](./jamgo.svg)

Jamgo is a minimal, superfast golang static site generator.

[![Go Report Card](https://goreportcard.com/badge/github.com/SamtheSaint/jamgo)](https://goreportcard.com/report/github.com/SamtheSaint/jamgo) [![PkgGoDev](https://pkg.go.dev/badge/github.com/SamtheSaint/jamgo)](https://pkg.go.dev/github.com/SamtheSaint/jamgo) ![GitHub](https://img.shields.io/github/license/SamtheSaint/jamgo) ![GitHub Release Date](https://img.shields.io/github/release-date/SamtheSaint/jamgo)

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
<!-- - [Contributing](#contributing) -->
- [License](#license)

# Overview

Jamgo is a superfast, minimal golang static site generator
The command line interface is designed to be quick, lightweight and simple to use.
Jamgo doesn't impose any structure on the website and allows complete developer freedom.

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

Getting started with jamgo is made as easy as possible. The [init command](./docs/jamgo_init.md) sets up the recommended folder structure for a minimal barebones jamgo application. To get started run:

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

The root command [`jamgo`](./docs/jamgo.md) doesn't do anything and will need to be chained with one of the commands below

Usage: `jamgo [command] [flags]`

Flags: `-h, --help (help for jamgo)`

## Build

The build command generates the html website in the directory of choice. It builds the pages in the `pages` directory of the project. Each page will need to be in its own directory, for example:

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
INSERT PUBLIC DIR TREE
```

## Init

## New

### Page

### Template

## Shell Completion

# License

Jamgo is released under The MIT License. See [LICENSE](./LICENSE).