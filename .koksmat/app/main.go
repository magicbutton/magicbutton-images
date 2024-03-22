package main

import (
	"runtime/debug"
	"strings"

	"github.com/365admin/magicbutton-images/magicapp"
	"github.com/365admin/magicbutton-images/utils"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: magicbutton-images
description: Describe the main purpose of this kitchen
---

# magicbutton-images
`
	magicapp.Setup(".env")
	magicapp.RegisterServeCmd("magicbutton-images", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	magicapp.RootCmd.PersistentFlags().BoolVarP(&utils.Verbose, "verbose", "v", false, "verbose output")

	magicapp.Execute(name, "magicbutton-images", "")
}
