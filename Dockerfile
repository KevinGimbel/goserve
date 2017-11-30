FROM alpine:3.6
LABEL maintainer="Kevin Gimbel <docker@kevin.codes>"

ENV version "1.2.1"

RUN mkdir -p /var/www/
WORKDIR /var/www

# Update certificates
RUN apk update \
 && apk add ca-certificates wget \
 && update-ca-certificates

# Install goserve
RUN mkdir -p /tmp/install/
RUN wget https://github.com/kevingimbel/goserve/releases/download/v"$version"/goserve_"$version"_linux_64-bit.tar.gz -O /tmp/install/goserve.tar.gz
RUN tar -xzf /tmp/install/goserve.tar.gz -C /tmp/install/
RUN ls -ahl /tmp/install/
RUN chmod +x /tmp/install/goserve
RUN mv /tmp/install/goserve /bin/
RUN rm -r /tmp/install

CMD ["goserve"]
