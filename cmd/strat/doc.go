// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// The command start contains various subcommands related to Stratumn.
//
// Usage
//
// 	The Stratumn CLI provides various commands to generate and work with Stratumn's technology.
//
// 	Usage:
// 	  strat [command]
//
// 	Available Commands:
// 	  build       Build project
// 	  deploy      Deploy project
// 	  down        Stop project services
// 	  generate    Generate a project
// 	  generators  List available generators
// 	  info        Display Stratumn CLI info
// 	  pull        Pull project Docker images
// 	  push        Push project Docker images
// 	  run         Run project command
// 	  test        Run project test suite
// 	  up          Start project services
// 	  update      Update Stratumn CLI or generators
// 	  version     Display version info
//
// 	Flags:
// 	  -c, --config string             Location of Stratumn configuration files (default "/Users/stephan/.stratumn")
// 	      --generators-owner string   Github owner of generators repository (default "stratumn")
// 	  -p, --generators-path string    Location where generators are stored locally (default "/Users/stephan/.stratumn/generators")
// 	      --generators-ref string     Git branch, tag, or commit of generators repository (default "master")
// 	      --generators-repo string    Name of generators Git repository (default "generators")
// 	      --gh-token string           Github API token
// 	      --stdin                     Attach stdin to process when executing project commands (default true)
//
// 	Use "strat [command] --help" for more information about a command.
package main
