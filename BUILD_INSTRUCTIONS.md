## Building and running

> Personal note: Void Linux doesn't have ffmpeg7 (and no libavcodec7) so I had
> to create a Debian trixie distrobox. The dockerfile uses Debian so I used
> that.

First, create a Debian distrobox:

```sh
distrobox create -i docker.io/debian:trixie filestash-debian
distrobox enter filestash-debian

sudo apt update && sudo apt install -y git make
```

Here you might need to add trixie-backports to
`/etc/apt/sources/list.d/debian.sources` in order to install golang-1.27. Once
that's done, run `sudo ln -s /usr/lib/go-1.27/bin/go /usr/bin/go`.

Then, install all the dependencies:

```sh
sudo apt install \
    libjpeg-dev \
    libtiff-dev \
    libpng-dev \
    libwebp-dev \
    libraw-dev \
    libheif-dev \
    libgif-dev \
    libvips-dev \
    libavcodec-dev \
    libavdevice-dev \
    libavfilter-dev \
    libavformat-dev \
    libswresample-dev \
    libswscale-dev \
    libavutil-dev

make init
make build
```

## Config

If you're looking for `config.json` to have it regenerated, it's in `dist/data/state/config/config.json`.
