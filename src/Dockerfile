FROM golang:1.17.2-alpine3.14

ENV PS1='\[\033[1;36m\][\u@\h] \[\033[1;34m\]\w \[\033[1;36m\]# \[\033[0m\]'

RUN apk add --no-cache --virtual .persistent-deps \
    util-linux \
    bash \
    vim \
    sudo \
    git

RUN apk update \
    && apk add -U tzdata \
    && cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime \
    && echo "Europe/Moscow" > /etc/timezone \
    && apk del tzdata \
    && rm -rf /var/cache/apk/*

RUN mkfifo -m 0666 /var/log/cron.log \
    && ln -s /var/log/cron.log /var/log/crond.log

ENV WORKDIR /usr/data/app
ENV PATH $WORKDIR/bin:$PATH
ENV GOCACHE /usr/data/.cache/go-build
ENV GOENV /usr/data/.config/go/env

RUN set -xe \
    && mkdir -p "$WORKDIR/src" "$WORKDIR/bin" \
    && mkdir -p "/usr/data/.cache/go-build" "/usr/data/.config/go/env" \
    && chmod -R 777 /usr/data/.cache \
    && chmod -R 777 /usr/data/.config \
    && go env

WORKDIR $WORKDIR

CMD ["crond", "-f", "-d", "8"]
