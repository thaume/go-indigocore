## Release notes

This is an early version of Stratumn Go packages and tools for experimental testing.

[Change log](CHANGE_LOG.md)

### Installing the CLI tools

#### Prerequisite

The only requirements are Docker and Docker Compose.

#### MacOS 64bit

Download `strat-darwin-amd64.zip`.

After downloading, unzip the archive, and move the `strat` binary somewhere in your executables `$PATH`.

```bash
$ unzip strat-darwin-amd64.zip
$ mv strat/strat /usr/local/bin/strat
```

#### Linux 64bit

Download `strat-linux-amd64.zip`.

After downloading, unzip the archive, and move the `strat` binary somewhere in your executables `$PATH`.

```bash
$ unzip strat-linux-amd64.zip
$ mv strat/strat /usr/local/bin/strat
```

#### Windows 64bit

Download `strat-windows-amd64.zip`.

After downloading, unzip the archive, and move `strat.exe` somewhere in your `%PATH%`.

### CLI quickstart

To generate a project in a directory named `demo`, run:

```bash
$ strat generate demo
$ cd demo
```

To launch all services locally, run within the project folder: 

```bash
$ strat up
```

To launch tests, run within the project folder: 

```bash
$ strat test
```

To update the CLI to the latest stable release, run: 

```bash
$ strat update
```

To update the CLI to the latest prerelease, run: 

```bash
$ strat update -prerelease
```

To update the generators, run: 

```bash
$ strat update -generators
```

## PGP signatures

Each zip file contains a cryptographic signature of the binary.

The signature can be verified using this public key:

https://keybase.io/stephan83/key.asc

This key is linked to the CTO of Stratumn on Keybase:

https://keybase.io/stephan83

If you have Keybase installed, you can verify the binary using
(replace `BINARY` with the name of the binary):

```
$ keybase verify -d BINARY.sig -i BINARY
```
