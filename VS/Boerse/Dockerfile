FROM golang:1.17

LABEL maintainer=mohamadkousa

RUN mkdir "/app"

COPY . /app

WORKDIR /app

EXPOSE 6543/udp


CMD ["go", "run", "."]




