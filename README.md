## chronos-shuttle

[![Build Status][travis-image]][travis-url] [![GoDoc][godoc-image]][godoc-url] [![Release][release-image]][release-url]

An opinionated CLI for Chronos

### Installation

```
go get github.com/yieldbot/chronos-shuttle
cd $GOPATH/src/github.com/yieldbot/chronos-shuttle
go build
```

### Usage

#### Help

```bash
./chronos-shuttle -h
```
```
Usage: chronos-shuttle [OPTIONS] COMMAND [arg...]

An opinionated CLI for Chronos

Options:
  --chronos     : Chronos url (default "http://localhost:8080")
  -h, --help    : Display usage
  -pp           : Pretty print for JSON output
  -v, --version : Display version information
  -vv           : Display extended version information

Commands:
  add           : Add a job
  del           : Delete a job
  graph         : Retrieve the dependency graph
  jobs          : Retrieve jobs
  kill          : Kill tasks of the job
  run           : Run a job
```

#### Setting Chronos Url

Default Chronos url is `http://localhost:8080`. But also you can use `--chronos` argument on each
command or set ENV variable with following command

```bash
export CHRONOS_URL=http://localhost
```

#### Getting jobs

```bash
./chronos-shuttle jobs
```

#### Adding a job

```bash
./chronos-shuttle add '{"schedule": "R/2015-11-09T00:00:00Z/PT24H", "name": "test-1", "epsilon": "PT30M", "command": "echo test-1 && sleep 60", "owner": "localhost@localhsot", "async": false}'
```

#### Running a job

```bash
./chronos-shuttle run test-1
```

#### Killing job tasks

```bash
./chronos-shuttle kill test-1
```

#### Retrieving the dependency graph

```bash
./chronos-shuttle graph test-1
```

#### Deleting a job

```bash
./chronos-shuttle del test-1
```

### TODO

- [ ] Auto binary release
- [ ] Add tests
- [ ] Proxy support

### License

Licensed under The MIT License (MIT)  
For the full copyright and license information, please view the LICENSE.txt file.

[travis-url]: https://travis-ci.org/yieldbot/chronos-shuttle
[travis-image]: https://travis-ci.org/yieldbot/chronos-shuttle.svg?branch=master

[godoc-url]: https://godoc.org/github.com/yieldbot/chronos-shuttle
[godoc-image]: https://godoc.org/github.com/yieldbot/chronos-shuttle?status.svg

[release-url]: https://github.com/yieldbot/chronos-shuttle/releases/tag/v1.0.0
[release-image]: https://img.shields.io/badge/release-v1.0.0-blue.svg

[coverage-url]: https://coveralls.io/github/yieldbot/chronos-shuttle?branch=master
[coverage-image]: https://coveralls.io/repos/yieldbot/chronos-shuttle/badge.svg?branch=master&service=github)