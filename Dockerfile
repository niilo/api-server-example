FROM centurylink/ca-certs

WORKDIR /app

COPY app /app/

EXPOSE 1323

ENTRYPOINT ["./app"]
