// +build mage

package main

import (
	"fmt"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const name string = "boltchat"
const buildDir string = "build"

const serverPrefix string = "server"
const clientPrefix string = "client"

const serverEntry string = "cmd/server/server.go"
const clientEntry string = "cmd/client/client.go"

type Build mg.Namespace
type Docker mg.Namespace

type BuildOptions struct {
	Static    bool
	Extension string
	Prefix    string
}

func build(os string, arch string, entry string, opts BuildOptions) error {
	env := map[string]string{
		"GOOS":   os,
		"GOARCH": arch,
	}

	// Build static binary
	if opts.Static {
		env["CGO_ENABLED"] = "0"
	}

	outputName := fmt.Sprintf(
		"%s-%s-%s-%s", name, opts.Prefix, os, arch,
	)

	outputPath := path.Join(
		buildDir,
		outputName,
	)

	if opts.Extension != "" {
		outputPath += fmt.Sprintf(".%s", opts.Extension)
	}

	args := []string{
		"build",
		"-o",
		outputPath,
		"-ldflags",
		"-s -w",
		entry,
	}

	fmt.Println(args)

	return sh.RunWith(
		env, "go", args...,
	)
}

/*
Build
*/

// Builds all binaries
func (Build) All() {
	mg.Deps(
		Build.ServerDarwinAmd64,
		Build.ServerLinuxAmd64,
		Build.ServerWindowsAmd64,

		Build.ClientDarwinAmd64,
		Build.ClientLinuxAmd64,
		Build.ClientWindowsAmd64,
	)
}

// Builds the server binary for Linux (amd64)
func (Build) ServerLinuxAmd64() error {
	return build("linux", "amd64", serverEntry, BuildOptions{Prefix: serverPrefix})
}

// Builds the server binary for Windows (amd64)
func (Build) ServerWindowsAmd64() error {
	return build("windows", "amd64", serverEntry, BuildOptions{
		Extension: "exe",
		Prefix:    serverPrefix,
	})
}

// Builds the server binary for Darwin/macOS (amd64)
func (Build) ServerDarwinAmd64() error {
	return build("darwin", "amd64", serverEntry, BuildOptions{Prefix: serverPrefix})
}

// Builds the server binary for Darwin/macOS (arm64, M1)
// func (Build) ServerDarwinArm64() error {
// 	return build("darwin", "arm64", serverEntry, false)
// }

// Builds the server binary for use in a Docker container
func (Build) ServerContainer() error {
	return build("linux", "amd64", serverEntry, BuildOptions{
		Static: true,
		Prefix: serverPrefix,
	})
}

// Builds the client binary for Linux (amd64)
func (Build) ClientLinuxAmd64() error {
	return build("linux", "amd64", clientEntry, BuildOptions{Prefix: clientPrefix})
}

// Builds the client binary for Windows (amd64)
func (Build) ClientWindowsAmd64() error {
	return build("windows", "amd64", clientEntry, BuildOptions{
		Extension: "exe",
		Prefix:    clientPrefix,
	})
}

// Builds the client binary for Darwin/macOS (amd64)
func (Build) ClientDarwinAmd64() error {
	return build("darwin", "amd64", clientEntry, BuildOptions{Prefix: clientPrefix})
}

/*
Docker
*/

// Builds a Docker image for the server
func (Docker) Build() error {
	return sh.RunV("docker", "build", ".", "-t", name)
}

/*
Misc
*/

// Cleans up build directories
func Clean() {
	sh.Rm("build")
}
