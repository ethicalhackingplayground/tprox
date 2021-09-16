<h1 align="center">
  <br>
<img src="static/icon.png" width="200px" alt="TProx">
</h1>

<h4 align="center">TProx is a fast reverse proxy path traversal detector and directory bruteforcer</h4>

<p align="center">
<a href="https://goreportcard.com/report/github.com/ethicalhackingplayground/tprox"><img src="https://goreportcard.com/badge/github.com/ethicalhackingplayground/tprox"></a>
<a href="https://github.com/ethicalhackingplayground/tprox/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://github.com/ethicalhackingplayground/tprox/releases"><img src="https://img.shields.io/github/release/ethicalhackingplayground/tprox"></a>
<a href="https://twitter.com/z0idsec"><img src="https://img.shields.io/twitter/follow/z0idsec.svg?logo=twitter"></a>
<a href="https://discord.gg/MQWCem5b"><img src="https://img.shields.io/discord/862900124740616192.svg?logo=discord"></a>
</p>

<p align="center">
  <a href="#install">Install</a> •
  <a href="#usage">Usage</a> •
  <a href="#examples">Examples</a> •
  <a href="https://discord.gg/MQWCem5b">Join Discord</a> 
</p>

---

### Install Options

#### From Source

```sh
▶  GO111MODULE=on go get -v  github.com/ethicalhackingplayground/tprox/tprox
```

#### Docker

```sh
▶  git clone https://github.com/ethicalhackingplayground/tprox && cd tprox && docker build -t tprox .
```

---

### Usage

```sh
▶ tprox -h
```

```sh
▶  docker run tprox -h
```



This will display help for the tool. Here are all the switches it supports.

<details>
<summary> 👉 tprox help menu 👈</summary>

```
Usage of ./tprox:
  -c int
        The number of concurrent requests (default 10)
  -check
        Check if a path/folder/file is internal
  -crawl
        crawl the resolved domain while testing for proxy misconfigs
  -depth int
        The crawl depth (default 5)
  -discover
        Discover path/folder/file with already found traversal
  -o string
        Output the results to a file
  -progress
        This flag will allow you to turn on the progress bar
  -regex string
        Filter crawl with regex pattern
  -scope string
        Specify a scope to crawl with in using regexs
  -silent
        Show Silent output
  -test
        Enable/Disable test mode only
  -traverse
        This flag will allow you to turn on traversing
  -w string
        The wordlist to use against a valid endpoint to traverse
```

</details>

### Examples

#### Traversal with Brute

```sh
▶ echo "https://example.com/api/v1" | tprox -w wordlist -traverse
```

#### Traversal with Crawling & Brute

```sh
▶ echo "https://example.com" | tprox -w wordlist -crawl -traverse
```

#### Traversal with Crawling, Regex Match & Brute

```sh
▶ echo "https://example.com" | tprox -w wordlist -crawl -traverse -regex "/api/"
```

#### Traversal With Crawling InScope & Brute

```sh
▶ echo "https://example.com" | tprox -w wordlist -crawl -traverse -regex "/api/" -scope ".*.\.example.com"
```

#### Traversal with Test Only

```sh
▶ echo "https://example.com/api" | tprox -test -traverse
```

#### Check if File is Internal

```sh
▶ echo "https://example.com/api/internalfile.html" | tprox -check
```

#### Discover Content 

```sh
▶ echo "https://example.com/api/..%2f" | tprox -discover -progress -w wordlist
```


<h1 align="center">
  <br>
<img src="static/example.png" alt="example">
</h1>

--- 

### Changes

- Added some additional flags to help aid finding traversal misconfigurations
- Optimised the crawler
- Added a flag to disable/enable the progress bar
- Fixed the silent flag
- Added check,test & discover flags

### Fixes

- Fixed a crawling bug.
- Fixed a traversal bug, it now only prints internal files & endpoints very low % of false positives.
- Made some optimization fixes.
- Discover content fix, it was not finding content.
- Optimisation fixes.

### Known Fixes

if for some reason the program fails to install or update run:

```sh
sudo rm -r /home/<user-name>/go/pkg/mod/github.com/ethicalhackingplayground/tprox
go clean --modcache
go clean
```

Then try and install it again.

### License

Tprox is distributed under [MIT License](https://github.com/ethicalhackingplayground/tprox/blob/main/LICENSE)

<h1 align="left">
  <a href="https://discord.gg/MQWCem5b"><img src="static/Join-Discord.png" width="380" alt="Join Discord"></a>
</h1>

