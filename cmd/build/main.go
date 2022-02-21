package main

import (
	"flag"
	"fmt"

	"github.com/forewing/gobuild"
)

const (
	nameCLI = "wordler"
	nameWeb = "wordler-web"

	sourceCLI = "./cmd/" + nameCLI
	sourceWeb = "./cmd/" + nameWeb

	module = "github.com/forewing/wordler"
)

var (
	flagAll = flag.Bool("all", false, "build for all platforms")
	flagCLI = flag.Bool("cli", false, "build CLI")
	flagWeb = flag.Bool("web", false, "build Web")

	target = gobuild.Target{
		Source:      sourceCLI,
		OutputName:  nameCLI,
		OutputPath:  "./output",
		CleanOutput: false,

		ExtraFlags:   []string{"-trimpath"},
		ExtraLdFlags: "-s -w",

		VersionPath: module + ".Version",
		HashPath:    module + ".Hash",

		Compress:  gobuild.CompressRaw,
		Platforms: []gobuild.Platform{{}},
	}
)

func init() {
	flag.Parse()
}

func main() {
	if !*flagCLI && !*flagWeb {
		*flagCLI = true
		*flagWeb = true
	}

	if *flagCLI {
		build(nameCLI, sourceCLI, *flagAll)
	}
	if *flagWeb {
		build(nameWeb, sourceWeb, *flagAll)
	}
}

func build(name string, source string, all bool) {
	t := target
	t.OutputName = name
	t.Source = source

	if all {
		t.OutputName = fmt.Sprintf("%s-%s-%s-%s",
			name,
			gobuild.PlaceholderVersion,
			gobuild.PlaceholderOS,
			gobuild.PlaceholderArch)
		t.Compress = gobuild.CompressZip
		t.Platforms = gobuild.PlatformCommon
	}

	err := t.Build()
	if err != nil {
		panic(err)
	}
}
