#build a tiny doker image
FROM alpine:latest
RUN mkdir /app
COPY goApp /app
CMD ["/app/goApp"]
