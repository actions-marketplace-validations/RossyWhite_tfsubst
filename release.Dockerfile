FROM golang:1.19.4

COPY tfsubst /usr/local/bin/tfsubst

ENTRYPOINT ["tfsubst"]
