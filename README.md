# Harbor: API server

[![Build Status](http://github.dronehippie.de/api/badges/harborapp/harbor-api/status.svg)](http://github.dronehippie.de/harborapp/harbor-api)
[![Coverage Status](http://coverage.dronehippie.de/badges/harborapp/harbor-api/coverage.svg)](http://coverage.dronehippie.de/harborapp/harbor-api)
[![Go Doc](https://godoc.org/github.com/harborapp/harbor-api?status.svg)](http://godoc.org/github.com/harborapp/harbor-api)
[![Go Report](http://goreportcard.com/badge/harborapp/harbor-api)](http://goreportcard.com/report/harborapp/harbor-api)
[![Join the chat at https://gitter.im/harborapp/harbor-api](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/harborapp/harbor-api)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)

**This project is under heavy development, it's not in a working state yet!**

Harbor will be a pretty simple web interface and authentication method for the
new docker distribution, the new opensource private docker registry. I thought
it's time to implement a shiny application with Go for the API and with React
for the UI.

The structure of the code base is heavily inspired by Drone, so those credits
are getting to [bradrydzewski](https://github.com/bradrydzewski), thank you for
this awesome project!


## Install

You can download prebuilt binaries from the GitHub releases or from our
[download site](http://dl.webhippie.de/harbor-api). You are a Mac user? Just take
a look at our [homebrew formula](https://github.com/harborapp/homebrew-harbor).
If you are missing an architecture just write us on our nice
[Gitter](https://gitter.im/harborapp/harbor-api) chat. Take a look at the help
output, you can enable auto updates to the binary to avoid bugs related to old
versions. If you find a security issue please contact thomas@webhippie.de first.


## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`

```bash
go get -d github.com/harborapp/harbor-api
cd $GOPATH/src/github.com/harborapp/harbor-api
make deps build

bin/harbor-api -h
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2016 Thomas Boerger <http://www.webhippie.de>
```
