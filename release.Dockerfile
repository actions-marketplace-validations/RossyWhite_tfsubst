FROM scratch

COPY tfsubst /tfsubst

ENTRYPOINT ["/tfsubst"]
