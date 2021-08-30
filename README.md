<h1 align="center">
  <br>
<img src="static/icon.png" width="200px" alt="TProx">
</h1>

<h4 align="center">TProx is a fast reverse proxy path traversal detector and directory bruteforcer</h4>

<p align="center">
  <a href="#install">Install</a> •
  <a href="#usage">Usage</a> •
  <a href="#examples">Usage</a> •
  <a href="https://discord.gg/MQWCem5b">Join Discord</a> 
</p>

---

### Install

```sh
▶  GO111MODULE=off go get -v -u github.com/ethicalhackingplayground/tprox/src
```

### Usage

```sh
tprox -h
```

This will display help for the tool. Here are all the switches it supports.

<details>
<summary> 👉 tprox help menu 👈</summary>

```
Usage of ./tprox:
  -c    crawl the resolved domain while testing for proxy misconfigs
  -d int
        The crawl depth (default 5)
  -o string
        Output the results to a file
  -r string
        Filter crawl with regex pattern
  -s    Show Silent output
  -t int
        The number of concurrent requests (default 10)
  -w string
        The wordlist to use against a valid endpoint to traverse
```

</details>

### Examples

Finding Path Traversal Files/Directories

```sh
▶ echo "https://example.com/api/v1" | tprox -w wordlist
```

Finding Path Traversal Files/Directories Through Crawling

```sh
▶ echo "https://example.com/api/v1" | tprox -w wordlist -c
```

Finding Path Traversal Files/Directories Through Crawling And Grepping

```sh
▶ echo "https://example.com/api/v1" | tprox -w wordlist -c -r "/api/"
```

Another alternitive to `echo` would be to cat out a list of resolved hosts

```sh
▶ cat urls.txt | tprox -w wordlist
```

### License

Tprox is distributed under [MIT License](https://github.com/ethicalhackingplayground/tprox/blob/main/LICENSE)

<h1 align="left">
  <a href="https://discord.gg/MQWCem5b"><img src="static/Join-Discord.png" width="380" alt="Join Discord"></a>
</h1>
