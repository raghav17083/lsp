package main

import (
	"fmt"
	"sort"

	c "github.com/mitchellh/colorstring"
)

const (
	briefcaseRune     = '💼'
	gitRune           = '😻'
	musicRune         = '🎼'
	pythonRune        = '🐍'
	javaRune          = '🍵'
	documentRune      = '📄'
	commonPrefix      = "[blue]"
	descriptionIndent = "                "
)

func render() {
	SetColumnSize()
	Traverse()
	renderSummary()
}

func renderSummary() {
	printHR()
	printCentered(fmt.Sprintf(c.Color("[white]lsp \"[red]%s[white]\""), presentPath(mode.targetPath)) + fmt.Sprintf(c.Color(", [red]%v[white] files, [red]%v[white] directories\n\n"), len(FileList), len(Trie.Ch["dirs"].Fls)))
}

func renderFiles(fls []*FileInfo) {
	switch {
	case mode.size:
		sort.Sort(sizeSort(fls))
	case mode.time:
		sort.Sort(timeSort(fls))
	default:
		sort.Sort(alphabeticSort(fls))
	}
	for _, fl := range fls {
		PrintColumns(fl.f.Name(), fl.Description())
	}
}
