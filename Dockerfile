FROM golang:1.17-alpine3.14

# Ignore APT warnings about not having a TTY
ENV DEBIAN_FRONTEND noninteractive

RUN apk update
RUN apk add wget gcc build-base libwebp-dev pkgconfig

# Install ImageMagick deps
RUN apk add libjpeg-turbo-dev libpng-dev libx11-dev

ENV IMAGEMAGICK_VERSION=6.9.10-11

RUN cd && \
	wget https://github.com/ImageMagick/ImageMagick6/archive/${IMAGEMAGICK_VERSION}.tar.gz && \
	tar xvzf ${IMAGEMAGICK_VERSION}.tar.gz && \
	cd ImageMagick* && \
	./configure \
	    --without-magick-plus-plus \
	    --without-perl \
	    --disable-openmp \
	    --with-gvc=no \
	    --disable-docs && \
	make -j$(nproc) && make install && \
	ldconfig /usr/local/lib

WORKDIR /go/projects/resizer

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o build/app

CMD ["build/app"]