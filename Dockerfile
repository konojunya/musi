FROM mhart/alpine-node:10.10.0

WORKDIR /go/src/github.com/konojunya/musi
ENV GO_ENV production
ENV GIN_MODE release

RUN set -x \
  && apk upgrade --no-cache \
  && apk --update add tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  && rm -rf /var/cache/apk/*
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN yarn install

ADD ./public/css/common.css ./public/css/common.css
ADD ./public/css/first.css ./public/css/first.css

ADD ./public/images/dj.jpg ./public/images/dj.jpg
ADD ./public/images/musi_logo_white.png ./public/images/musi_logo_white.png
ADD ./public/images/musi_logo_black.png ./public/images/musi_logo_black.png
ADD ./public/images/musi_logo_shape_black.png ./public/images/musi_logo_shape_black.png
ADD ./public/images/musi_logo_shape_white.png ./public/images/musi_logo_shape_white.png
ADD ./public/images/musi_logo_text_black.png ./public/images/musi_logo_text_black.png
ADD ./public/images/musi_logo_text_white.png ./public/images/musi_logo_text_white.png

ADD ./public/js/bundle.js ./public/js/bundle.js

ADD ./views/first.html ./views/first.html
ADD ./views/index.html ./views/index.html
ADD ./.env .
ADD ./cmd/main .

EXPOSE 3000

ENTRYPOINT ["./main"]