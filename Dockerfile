FROM alpine:3.2
ADD bench-srv /bench-srv
ENTRYPOINT [ "/bench-srv" ]
