FROM golang:1.15
ENV LANG=C.UTF-8
RUN apt-get update && apt-get install -y wget libvlc-dev vlc-plugin-base vlc-plugin-video-output vlc
RUN wget https://yt-dl.org/downloads/latest/youtube-dl -O /usr/local/bin/youtube-dl && chmod a+rx /usr/local/bin/youtube-dl
RUN mkdir /youtube
ADD . /youtube
WORKDIR /youtube
RUN CGO_FLAGS="-w" go build 
CMD ["/youtube/youtube-tui"]