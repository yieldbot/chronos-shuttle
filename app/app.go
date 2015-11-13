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
	"os"
	"strings"

	"github.com/yieldbot/chronos-client"
	"github.com/yieldbot/gocli"
)

var (
	cli            gocli.Cli
	chronosURL     string
	chronosClient  client.Client
	usageFlag      bool
	versionFlag    bool
	versionExtFlag bool
	prettyPrint    bool
	chronos        string
)

func init() {
	flag.BoolVar(&usageFlag, "h", false, "Display usage")
	flag.BoolVar(&usageFlag, "help", false, "Display usage")
	flag.BoolVar(&versionFlag, "version", false, "Display version information")
	flag.BoolVar(&versionFlag, "v", false, "Display version information")
	flag.BoolVar(&versionExtFlag, "vv", false, "Display extended version information")
	flag.BoolVar(&prettyPrint, "pp", false, "Pretty print for JSON output")
	flag.StringVar(&chronos, "chronos", "", "Chronos url (default \"http://localhost:8080\")")
}

// Run runs the app
func Run() {

	// Init cli
	cli = gocli.Cli{
		AppName:    "chronos-shuttle",
		AppVersion: "1.0.0",
		AppDesc:    "An opinionated CLI for Chronos",
		CommandList: map[string]string{
			"jobs":  "Retrieve jobs",
			"add":   "Add a job",
			"run":   "Run a job",
			"kill":  "Kill tasks of the job",
			"del":   "Delete a job",
			"graph": "Retrieve the dependency graph",
		},
	}
	cli.Init()

	// Run the app

	// Command
	if cli.Command != "" {

		// Init the Chronos client
		if chronos != "" {
			chronosURL = chronos
		} else if os.Getenv("CHRONOS_URL") != "" {
			chronosURL = os.Getenv("CHRONOS_URL")
		} else {
			chronosURL = "http://localhost:8080"
		}
		chronosClient = client.Client{URL: chronosURL}

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
	if err := chronosClient.PrintJobs(prettyPrint); err != nil {
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

	// Run the job
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
