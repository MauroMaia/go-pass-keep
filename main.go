package main

import "go-pass-keeper/src/cmd"

const VERSION = ""
const COMMIT_ID = ""
const BUILD_DATE = ""

func main() {
	cmd.Execute(VERSION, COMMIT_ID, BUILD_DATE)
}
