# Go Overview

This is an overview of the go programming language for developers familiar with either Python, Java or C.

## Run

To view the slides you'll need go [installed](https://golang.org/dl/), and then you'll need the `present` command. This can be obtained by running `go install golang.org/x/tools/cmd/present`.

To launch the slides run `./run.sh` in a checkout of this repository and visit: http://localhost:3999/class.slide.

## Resources
[Go Standard Library](https://golang.org/pkg/#stdlib)
[DevDocs](https://devdocs.io/go/)
[Effective Go](https://golang.org/doc/effective_go.html)
[Less is exponentially more](https://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html)
[Visualizing Concurrency in Go](https://www.youtube.com/watch?v=KyuFeiG3Y60 Visualizing)

## Setting up on Mac
Make sure you have git installed via Xcode.

Download the binary .tar.gz (not the .pkg) from golang.org/dl: https://dl.google.com/go/go1.12.5.darwin-amd64.tar.gz

Extract into ~/goroot so GOROOT and GOPATH don't conflict:
`
tar xzvf go1.12.5.darwin-amd64.tar.gz
mv go ~/goroot
`

Add to your PATH in `~/.bash_profile`:
`
PATH="${HOME}/goroot/bin:${HOME}/go/bin:${PATH}"
export PATH
`

Reload your `~/.bash_profile`:
`
source ~/.bash_profile
`

Make sure you're running the expected version of Go:
`
which go
go version
`
