package main

import (
	"embed"
)

//go:embed static/gui/*
var HTML embed.FS
