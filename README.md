# Umschlag: API server

[![Build Status](https://cloud.drone.io/api/badges/umschlag/umschlag-api/status.svg)](https://cloud.drone.io/umschlag/umschlag-api)
[![Stories in Ready](https://badge.waffle.io/umschlag/umschlag-api.svg?label=ready&title=Ready)](http://waffle.io/umschlag/umschlag-api)
[![Join the Matrix chat at https://matrix.to/#/#umschlag:matrix.org](https://img.shields.io/badge/matrix-%23umschlag-7bc9a4.svg)](https://matrix.to/#/#umschlag:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/cbe28cf646c34c98b58967079e9ae990)](https://www.codacy.com/app/umschlag/umschlag-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=umschlag/umschlag-api&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/umschlag/umschlag-api?status.svg)](http://godoc.org/github.com/umschlag/umschlag-api)
[![Go Report](http://goreportcard.com/badge/github.com/umschlag/umschlag-api)](http://goreportcard.com/report/github.com/umschlag/umschlag-api)
[![](https://images.microbadger.com/badges/image/umschlag/umschlag-api.svg)](http://microbadger.com/images/umschlag/umschlag-api "Get your own image badge on microbadger.com")

**This project is under heavy development, it's not in a working state yet!**

Umschlag will be a pretty simple web interface and authentication method for the new docker distribution, the new open-source private docker registry. I thought it's time to implement a shiny application with Go for the API and with VueJS for the UI.

*Where does this name come from or what does it mean? It's quite simple, it's one german word for transshipment, I thought it's a good match as it is related to containers and a harbor.*


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.umschlag.tech/api). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/umschlag/homebrew-umschlag).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.11.

```bash
git clone https://github.com/umschlag/umschlag-api.git
cd umschlag-api

make sync generate build

./bin/umschlag-api -h
```


## Security

If you find a security issue please contact umschlag@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
