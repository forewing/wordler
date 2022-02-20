package main

import (
	"embed"
	"html/template"
	"io/fs"
)

const (
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
	t, err := template.New("").Delims("[[", "]]").ParseFS(templates, "*.html")
	if err != nil {
		panic(err)
	}
	return t
}
