<h1 align="center">
  <br>
<img src="static/icon.png" width="200px" alt="TProx">
</h1>

<h4 align="center">TProx is a fast reverse proxy path traversal detector and directory bruteforcer</h4>

<p align="center">
  <a href="#install">Install</a> •
  <a href="#usage">Usage</a> •
  <a href="#examples">Examples</a> •
  <a href="https://discord.gg/MQWCem5b">Join Discord</a> 
</p>

---

### Install From Source

```sh
▶  GO111MODULE=on go get -v  github.com/ethicalhackingplayground/tprox/tprox
```

### Install With Docker

```sh
▶  git clone https://github.com/ethicalhackingplayground/tprox && cd tprox && docker build -t tprox .
```

```sh
▶  docker run tprox -h
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

```sh
▶ echo "https://example.com/api/v1" | tprox -w wordlist
```

```sh
▶ echo "https://example.com" | tprox -w wordlist -c
```

```sh
▶ echo "https://example.com" | tprox -w wordlist -c -r "/api/"
```

```sh
▶ cat urls.txt | tprox -w wordlist
```

### License

Tprox is distributed under [MIT License](https://github.com/ethicalhackingplayground/tprox/blob/main/LICENSE)

<h1 align="left">
  <a href="https://discord.gg/MQWCem5b"><img src="static/Join-Discord.png" width="380" alt="Join Discord"></a>
</h1>
