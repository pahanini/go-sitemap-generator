# Golang Sitemap Generator

Creates sitemap files

[![Build Status](https://travis-ci.org/pahanini/go-sitemap-generator.svg)](https://travis-ci.org/pahanini/go-sitemap-generator)

## Usage

```go
package main

import (
	"github.com/pahanini/go-sitemap-generator"
)

g := sitemap.New(sitmap.Options{
	Dir:         "sitemap",
	BaseURL:     "http://example.com/",
})
g.Open()
g.Add(sitemap.URL{Location:`http://example.com`, Priority: `0.5`})
g.Add(sitemap.URL{Location:`http://example.com/test`, Priority: `0.5`})
g.Close()
```

## Install

```console
$ go get github.com/pahanini/go-sitemap-generator
```

## Available Options

Sitemap comes with a variety of configuration options. Available Options:

```go
type Options struct {
	// Filename is base file name for sitemap w/o extension
	// - single file <filename>.xml
	// - many files with sitemap index <filename>.xml (index file) and <filename>-<n>.xml (files with urls)
	Filename string
	// Max file size (default 10485760)
	MaxFileSize int
	// Max links in one file (default 50000)
	MaxURLs int
	// Dir keeps directory name for sitemap files
	Dir string
	// BaseURL used for generate sitemap index file
	BaseURL string
}
```

                                                                                                                              Render comes with a variety of configuration options (Note: these are not the default option values. See the defaults below.):