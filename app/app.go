/*
 * chronos-shuttle
 * Copyright (c) 2015 Yieldbot, Inc. (http://github.com/yieldbot/chronos-shuttle)
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

// Package app package provides the app information
package app

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/yieldbot/chronos-client"
	"github.com/yieldbot/gocli"
)

var (
	cli             gocli.Cli
	chronosURL      string
	proxyURL        string
	chronosClient   client.Client
	usageFlag       bool
	versionFlag     bool
	versionExtFlag  bool
	prettyPrintFlag bool
	chronosFlag     string
	proxyFlag       string
)

func init() {
	flag.BoolVar(&usageFlag, "h", false, "Display usage")
	flag.BoolVar(&usageFlag, "help", false, "Display usage")
	flag.BoolVar(&versionFlag, "version", false, "Display version information")
	flag.BoolVar(&versionFlag, "v", false, "Display version information")
	flag.BoolVar(&versionExtFlag, "vv", false, "Display extended version information")
	flag.BoolVar(&prettyPrintFlag, "pp", false, "Pretty print for JSON output")
	flag.StringVar(&chronosFlag, "chronos", "", "Chronos url (default \"http://localhost:8080\")")
	flag.StringVar(&proxyFlag, "proxy", "", "Proxy url")
}

// Run runs the app
func Run() {

	// Init cli
	cli = gocli.Cli{
		AppName:    "chronos-shuttle",
		AppVersion: "1.1.1",
		AppDesc:    "An opinionated CLI for Chronos",
		CommandList: map[string]string{
			"jobs":  "Retrieve jobs",
			"add":   "Add a job",
			"run":   "Run a job",
			"kill":  "Kill tasks of the job",
			"del":   "Delete a job",
			"graph": "Retrieve the dependency graph",
			"sync":  "Sync jobs via a file or directory",
		},
	}
	cli.Init()

	// Run the app

	// Command
	if cli.Command != "" {

		// Init the Chronos client
		if chronosFlag != "" {
			chronosURL = chronosFlag
		} else if os.Getenv("CHRONOS_URL") != "" {
			chronosURL = os.Getenv("CHRONOS_URL")
		} else {
			chronosURL = "http://localhost:8080"
		}

		if proxyFlag != "" {
			proxyURL = proxyFlag
		} else if os.Getenv("SHUTTLE_PROXY_URL") != "" {
			proxyURL = os.Getenv("SHUTTLE_PROXY_URL")
		}

		if proxyURL != "" {
			p, err := url.Parse(proxyURL)
			if err != nil {
				cli.LogErr.Fatal("invalid proxy value due to " + err.Error())
			}
			chronosClient = client.Client{URL: chronosURL, ProxyURL: p}
		} else {
			chronosClient = client.Client{URL: chronosURL}
		}

		// Run the command
		if cli.Command == "jobs" {
			// Get the jobs
			runJobsCmd()
		} else if cli.Command == "add" {
			// Add a job
			runAddCmd()
		} else if cli.Command == "run" {
			// Run a job
			runRunCmd()
		} else if cli.Command == "kill" {
			// Kill the job tasks
			runKillCmd()
		} else if cli.Command == "del" {
			// Delete a job
			runDelCmd()
		} else if cli.Command == "graph" {
			// Get the dependency graph
			runGraphCmd()
		} else if cli.Command == "sync" {
			// Sync jobs
			runSyncCmd()
		}
	} else if versionFlag || versionExtFlag {
		// Version
		cli.PrintVersion(versionExtFlag)
	} else {
		// Default
		cli.PrintUsage()
	}
}

// runJobsCmd runs the jobs command
func runJobsCmd() {
	if err := chronosClient.PrintJobs(prettyPrintFlag); err != nil {
		cli.LogErr.Fatal(err)
	}
}

// runAddCmd runs the add command
func runAddCmd() {
	// Get the job name
	var jobj string
	if len(cli.CommandArgs) > 0 {
		jobj = cli.CommandArgs[0]
	}

	// Add the job
	if ok, err := chronosClient.AddJob(jobj); !ok && err != nil {
		cli.LogErr.Fatal(err) // fatal error
	} else if err != nil {
		cli.LogErr.Println(err) // print error
	} else {
		cli.LogOut.Printf("The job is added\n")
	}
}

// runRunCmd runs the run command
func runRunCmd() {
	// Get the job name
	var job, ja string
	if len(cli.CommandArgs) > 0 {
		job = cli.CommandArgs[0]
	}
	if len(cli.CommandArgs) > 1 {
		ja = strings.Join(cli.CommandArgs[1:], "")
	}

	// Run the job
	if ok, err := chronosClient.RunJob(job, ja); !ok && err != nil {
		cli.LogErr.Fatal(err) // fatal error
	} else if err != nil {
		cli.LogErr.Println(err) // print error
	} else {
		cli.LogOut.Printf("%s job is running\n", job)
	}
}

// runKillCmd runs the kill command
func runKillCmd() {
	// Get the job name
	var job string
	if len(cli.CommandArgs) > 0 {
		job = cli.CommandArgs[0]
	}

	// Kill the job tasks
	// If it is not ok and there is an error then
	if ok, err := chronosClient.KillJobTasks(job); !ok && err != nil {
		cli.LogErr.Fatal(err) // fatal error
	} else if err != nil {
		cli.LogErr.Println(err) // print error
	} else {
		cli.LogOut.Printf("%s job tasks are killed\n", job)
	}
}

// runDelCmd runs the remove command
func runDelCmd() {
	// Get the job name
	var job string
	if len(cli.CommandArgs) > 0 {
		job = cli.CommandArgs[0]
	}

	// Delete the job
	if ok, err := chronosClient.DeleteJob(job); !ok && err != nil {
		cli.LogErr.Fatal(err) // fatal error
	} else if err != nil {
		cli.LogErr.Println(err) // print error
	} else {
		cli.LogOut.Printf("%s job is removed\n", job)
	}
}

// runGraphCmd runs the graph command
func runGraphCmd() {
	if res, err := chronosClient.DepGraph(); err != nil {
		cli.LogErr.Fatal(err) // fatal error
	} else {
		fmt.Print(res)
	}
}

// syncFile syncs the given file
func syncFile(path string) {
	// Read file
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		cli.LogErr.Fatal(err)
	}

	// Add the job
	if ok, err := chronosClient.AddJob(string(buf)); !ok && err != nil {
		cli.LogErr.Fatal(err) // fatal error
	} else if err != nil {
		cli.LogErr.Println(err) // print error
	} else {
		cli.LogOut.Printf("%s is synced\n", path)
	}
}

// walkFn called for each directory during walk function execution
func walkFn(path string, info os.FileInfo, err error) error {
	// If it is not a directory then
	if !info.IsDir() {
		// Sync the file
		syncFile(path)
	}
	return nil
}

// runSyncCmd runs the sync command
func runSyncCmd() {
	// Get the file or directory path
	var path string
	if len(cli.CommandArgs) > 0 {
		path = cli.CommandArgs[0]
	}

	// Check file
	var fi os.FileInfo
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		cli.LogErr.Fatal("no such file or directory: " + path) // fatal error
	}

	// If it is a file than
	if !fi.IsDir() {
		// Sync the file
		syncFile(path)
	} else {
		// Otherwise recursively sync files
		if err := filepath.Walk(path, walkFn); err != nil {
			cli.LogErr.Fatal(err)
		}
	}
}
