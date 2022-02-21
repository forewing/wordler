package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/url"

	"github.com/forewing/wordler"
)

const (
	queryHashLength = 7

	staticsPath   = "statics"
	templatesPath = "templates"
)

var (
	//go:embed statics/*
	staticsEmbed embed.FS
	statics      fs.FS

	//go:embed templates/*
	templatesEmbed embed.FS
	templates      fs.FS
)

func init() {
	statics = mustStripFSPrefix(staticsEmbed, staticsPath)
	templates = mustStripFSPrefix(templatesEmbed, templatesPath)
}

func mustStripFSPrefix(sfs fs.FS, prefix string) fs.FS {
	dfs, err := fs.Sub(sfs, prefix)
	if err != nil {
		panic(err)
	}
	return dfs
}

func mustLoadTemplate() *template.Template {
	t, err := template.New("").Delims("[[", "]]").Funcs(template.FuncMap{
		"generateStaticURL": func(origin string) string {
			u, err := url.Parse(origin)
			if err != nil {
				panic(err)
			}

			hash := wordler.Hash
			if len(hash) == 0 {
				hash = wordler.HashDefault
			}
			if len(hash) > queryHashLength {
				hash = hash[0:queryHashLength]
			}

			q := u.Query()
			q.Set("v", hash)
			u.RawQuery = q.Encode()

			return u.String()
		},
	}).ParseFS(templates, "*.html")
	if err != nil {
		panic(err)
	}
	return t
}
