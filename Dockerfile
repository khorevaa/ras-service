FROM scratch
COPY ras-service /
ENTRYPOINT ["/ras-service"]