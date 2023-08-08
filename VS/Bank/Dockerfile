FROM golang:1.17

LABEL maintainer=mohamadkousa

EXPOSE 6543/TCP

RUN mkdir "/app"

COPY . /app
# Update go.mod
RUN go clean -modcache

WORKDIR /app

CMD ["go", "run", "."]