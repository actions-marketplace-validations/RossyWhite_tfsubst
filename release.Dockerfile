FROM bash

COPY tfsubst /usr/local/bintfsubst

ENTRYPOINT ["tfsubst"]
