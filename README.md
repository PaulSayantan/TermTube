<h1 align="center">TermTube

![](https://img.shields.io/badge/YouTube-Terminal%20App-black?logoColor=fc2803&style=for-the-badge&logo=youtube)

<img src="https://user-images.githubusercontent.com/53504602/92931573-454e8400-f461-11ea-904f-8fad59bb0d2a.gif" width="150" height="100"/>

</h1>

<div align="center">

[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2FPaulSayantan%2Fyoutube-tui%2Fbadge%3Fref%3Dmaster&style=for-the-badge&color=green)](https://actions-badge.atrox.dev/PaulSayantan/youtube-tui/goto?ref=master)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/PaulSayantan/youtube-tui/master?style=for-the-badge&logo=github&color=orange)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/PaulSayantan/youtube-tui?style=for-the-badge&logo=go)
![Docker Build](https://img.shields.io/docker/cloud/build/belikesayantan/termtube?logo=docker&style=for-the-badge)
</div>

A Terminal Youtube App | Play, Listen &amp; Enjoy, right from your terminal !!


## Requirements

* [Youtube-dl](https://youtube-dl.org/downloads/)

* VLC SDK
    - For Unix: https://github.com/adrg/libvlc-go/wiki/Install-on-Linux
    - For Windows: https://github.com/adrg/libvlc-go/wiki/Install-on-Windows

## Installation

```
$ git clone https://github.com/PaulSayantan/TermTube.git

$ cd TermTube

$ CGO_CFLAGS="-w" go build 

$ ./TermTube
```
## Docker 
Build
```
docker build -t termtube
```
or download pre-build image
```
docker pull belikesayantan/termtube
```
then run it
```
docker run -it --device /dev/snd:/dev/snd termtube:latest 
```

## Codebase Map
A feature map to visualize the code-base of the current project. Click [here](https://app.codesee.io/maps/9033a700-1d0c-11ec-a981-5f80e2c33b31)

## Show your support

Give a ⭐ if you liked this project !!

---
