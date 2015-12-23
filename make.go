package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	version = "0.0.1-alpha0"
)

var steps = map[string]step{
	"deps":     executeDeps,
	"vendor":   executeVendor,
	"lint":     executeLint,
	"fmt":      executeFmt,
	"vet":      executeVet,
	"generate": executeGenerate,
	"scripts":  executeScripts,
	"styles":   executeStyles,
	"test":     executeTest,
	"build":    executeBuild,
	"install":  executeInstall,
	"bindata":  executeBindata,
}

type step func() error

func init() {
	os.Setenv("GO15VENDOREXPERIMENT", "1")
}

func main() {
	for _, arg := range os.Args[1:] {
		step, ok := steps[arg]

		if !ok {
			fmt.Println("Error: Invalid step", arg)
			os.Exit(1)
		}

		err := step()

		if err != nil {
			fmt.Println("Error: Failed step", arg)
			os.Exit(1)
		}
	}
}

func executeDeps() error {
	deps := []string{
		"github.com/jteeuwen/go-bindata/...",
		"github.com/tdewolff/minify/cmd/minify",
	}

	for _, dep := range deps {
		err := run(
			"go",
			"get",
			"-u",
			dep)

		if err != nil {
			return err
		}
	}

	return nil
}

func executeVendor() error {
	err := run(
		"go",
		"get",
		"-u",
		"github.com/Masterminds/glide")

	if err != nil {
		return err
	}

	return run(
		gopath("glide"),
		"update")
}

func executeLint() error {
	err := run(
		"go",
		"get",
		"-u",
		"github.com/golang/lint")

	if err != nil {
		return err
	}

	return run(
		gopath("golint"),
		"./...")
}

func executeFmt() error {
	return run(
		"go",
		"fmt",
		"./...")
}

func executeVet() error {
	return run(
		"go",
		"vet",
		"./...")
}

func executeGenerate() error {
	return run(
		"go",
		"generate",
		"./...")
}

func executeScripts() error {
	var paths = []struct {
		inputs   []string
		output   string
		minified string
	}{
		{
			[]string{
			// "cmd/harbord/assets/scripts/jquery/jquery.js",
			},
			"cmd/harbord/static/scripts/vendor.js",
			"cmd/harbord/static/scripts/vendor.min.js",
		},
		{
			[]string{
			// "cmd/harbord/assets/scripts/harbor/basics.js",
			},
			"cmd/harbord/static/scripts/harbor.js",
			"cmd/harbord/static/scripts/harbor.min.js",
		},
	}

	for _, path := range paths {
		dirname := filepath.Dir(path.output)

		if _, err := os.Stat(dirname); os.IsNotExist(err) {
			fmt.Printf("Directory %s doesn't exist\n", dirname)
			return err
		}

		f, err := os.OpenFile(
			path.output,
			os.O_CREATE|os.O_RDWR|os.O_TRUNC,
			0660)

		defer f.Close()

		if err != nil {
			fmt.Printf("Failed to open %s file\n", path.output)
			return err
		}

		for _, input := range path.inputs {
			content, err := ioutil.ReadFile(input)

			if err != nil {
				return err
			}

			f.Write(content)
		}

		minifyErr := run(
			gopath("minify"),
			"-o",
			path.minified,
			path.output)

		if minifyErr != nil {
			return minifyErr
		}
	}

	return nil
}

func executeStyles() error {
	var paths = []struct {
		inputs   []string
		output   string
		minified string
	}{
		{
			[]string{
			// "cmd/harbord/assets/styles/jquery/jquery.css",
			},
			"cmd/harbord/static/styles/vendor.css",
			"cmd/harbord/static/styles/vendor.min.css",
		},
		{
			[]string{
			// "cmd/harbord/assets/styles/harbor/basics.css",
			},
			"cmd/harbord/static/styles/harbor.css",
			"cmd/harbord/static/styles/harbor.min.css",
		},
	}

	for _, path := range paths {
		dirname := filepath.Dir(path.output)

		if _, err := os.Stat(dirname); os.IsNotExist(err) {
			fmt.Printf("Directory %s doesn't exist\n", dirname)
			return err
		}

		f, err := os.OpenFile(
			path.output,
			os.O_CREATE|os.O_RDWR|os.O_TRUNC,
			0660)

		defer f.Close()

		if err != nil {
			fmt.Printf("Failed to open %s file\n", path.output)
			return err
		}

		for _, input := range path.inputs {
			content, err := ioutil.ReadFile(input)

			if err != nil {
				return err
			}

			f.Write(content)
		}

		minifyErr := run(
			gopath("minify"),
			"-o",
			path.minified,
			path.output)

		if minifyErr != nil {
			return minifyErr
		}
	}

	return nil
}

func executeTest() error {
	ldf := fmt.Sprintf(
		"-X main.version=%s",
		version)

	return run(
		"go",
		"test",
		"-ldflags",
		ldf,
		"github.com/webhippie/harbor/...")
}

func executeBuild() error {
	var bins = []struct {
		input  string
		output string
	}{
		{
			"github.com/webhippie/harbor",
			"harbor",
		},
	}

	for _, bin := range bins {
		ldf := fmt.Sprintf(
			"-X main.version=%s",
			version)

		err := run(
			"go",
			"build",
			"-o",
			bin.output,
			"-ldflags",
			ldf,
			bin.input)

		if err != nil {
			return err
		}
	}

	return nil
}

func executeInstall() error {
	var bins = []struct {
		input string
	}{
		{
			"github.com/webhippie/harbor",
		},
	}

	for _, bin := range bins {
		ldf := fmt.Sprintf(
			"-X main.version=%s",
			version)

		err := run(
			"go",
			"install",
			"-ldflags",
			ldf,
			bin.input)

		if err != nil {
			return err
		}
	}

	return nil
}

func executeBindata() error {
	var paths = []struct {
		input  string
		output string
	}{
		{
			"static/...",
			"bindata.go",
		},
	}

	for _, path := range paths {
		err := run(
			"go-bindata",
			"-o",
			path.output,
			path.input)

		if err != nil {
			return err
		}
	}

	return nil
}

func gopath(exe string) string {
	return strings.Join(
		[]string{
			os.Getenv("GOPATH"),
			"bin",
			exe,
		},
		"/")
}

func run(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	trace(cmd.Args)
	return cmd.Run()
}

func trace(args []string) {
	print("+ ")
	println(strings.Join(args, " "))
}
