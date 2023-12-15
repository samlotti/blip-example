FROM alpine
WORKDIR /app
COPY blipServer .


EXPOSE 8181
CMD ["/app/blipServer"]
