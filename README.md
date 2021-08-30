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
<summary> 👉 gocrawler help menu 👈</summary>

```
Usage of ./src:
  -o string
        Output the results to a file
  -t int
        The number of concurrent requests (default 10)
  -w string
        The wordlist to use against a valid endpoint to traverse
```

</details>

### Examples

Finding Server-Side Path Traversal Files/Directories

```sh
▶ echo "https://example.com/api/v1" | tprox -w wordlist 
```

### License

Tprox is distributed under [MIT License](https://github.com/ethicalhackingplayground/tprox/blob/main/LICENSE)

<h1 align="left">
  <a href="https://discord.gg/MQWCem5b"><img src="static/Join-Discord.png" width="380" alt="Join Discord"></a>
</h1>
