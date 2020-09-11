<h1 align="center">youtube-tui 

![](https://img.shields.io/badge/YouTube-Terminal%20App-black?logoColor=fc2803&style=for-the-badge&logo=youtube)
</h1>

<div align="center">

[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fbelikesayantan%2Fyoutube-tui%2Fbadge%3Fref%3Dmaster&style=for-the-badge&color=green)](https://actions-badge.atrox.dev/belikesayantan/youtube-tui/goto?ref=master)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/belikesayantan/youtube-tui/master?style=for-the-badge&logo=github&color=orange)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/belikesayantan/youtube-tui?style=for-the-badge&logo=go)
![Codacy grade](https://img.shields.io/codacy/grade/661007379bcc4bca841a5447155f02b3?style=for-the-badge&logo=codacy)
</div>

A Terminal Youtube App | Play, Listen &amp; Enjoy, right from your terminal !!
    

## Requirements

* [Youtube-dl](https://youtube-dl.org/downloads/)

* VLC SDK
    - For Unix: https://github.com/adrg/libvlc-go/wiki/Install-on-Linux
    - For Windows: https://github.com/adrg/libvlc-go/wiki/Install-on-Windows

## Installation

```
$ git clone https://github.com/belikesayantan/youtube-tui.git

$ cd youtube-tui

$ CGO_CFLAGS="-w" go build 

$ ./youtube-tui
```
## Docker 

```
docker build -t youtube-tui
docker run -it --device /dev/snd:/dev/snd youtube-tui:latest 
```

![initial-screen](https://user-images.githubusercontent.com/53504602/92327302-486afe00-f076-11ea-92f6-65db95a990f2.png)
