# SlowLoris

Welcome to SlowLoris, the tool that lets you unleash a delightfully sneaky low-bandwidth Denial of Service (DoS) attack! Crafted in the mystical land of Go, this nifty little program simulates the SlowLoris attack—an underhanded method to keep your target’s web server on its toes (or, more accurately, utterly bewildered).

**Disclaimer:** This is all for educational fun! Please, use responsibly, and never attack servers without permission—unless you fancy a friendly chat with the law. 

## What’s the Buzz About?

The SlowLoris attack is like a magician that keeps your server guessing, holding open numerous connections and sending incomplete requests. The result? Your target spends valuable resources waiting for requests that never complete. Voilà! Instant chaos!

## Features
* **Stealth Mode:** Simulates a SlowLoris attack with grace.

* **Customizable Chaos:** Easily adjust the number of open connections to amplify your impact.

* **User-Friendly:** A simple command-line interface that’s as easy as pie.

* **Docker-Ready:** Test your attacks in a safe and contained environment.

## Requirements

* Go 1.13 or higher 
* Docker (Optional, but highly recommended for testing)

## How to Use This Tool

### Running the Attack

Ready to unleash the magic? Here’s how you can run the attack using ```go run```:

```bash
go run main.go -t 127.0.0.1:port -r [optional; default is 1000]
```
* ``-t`` is your target (IP address and port).
* ``-r`` is the number of concurrent requests you wish to conjure up (default is 1000).

### Building the Ultimate Spell
Want to make it official? Build your own executable and let it roam free:

1. Compile your magical creation:

```bash
go build main.go
```

2. Let it loose:

```
./main -t 127.0.0.1:port -r [optional; default is 1000]
```

## Testing with Docker 🐳

Why not test your spells in a controlled environment? Follow these steps:

1. Build your Docker image:

```bash
docker build -t name-of-the-image .
```

2. Summon your Docker container:

```bash
docker run -d --name container-name -p 8800:80 name-of-the-image
```

This will create a test server on port 8800, just waiting to be gently overwhelmed by your SlowLoris attack.

## A Word of Caution
Remember, folks: With great power comes great responsibility! This tool is designed for educational use only. Using it on systems without permission could lead to some serious trouble—like a chat with law enforcement! 🚓💼

Created by [Matei-Stoian](https://blog-app-steel-omega.vercel.app/). Happy attacking (ethically, of course)! 