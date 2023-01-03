FROM bash

COPY tfsubst /usr/local/bin/tfsubst

ENTRYPOINT ["tfsubst"]
