
FROM golang:1.26.0 AS build

WORKDIR /app


COPY . .


RUN CGO_ENABLED=0 go build -o server .


FROM ubuntu:22.04
WORKDIR /root/

COPY --from=build /app/pages ./pages
COPY --from=build /app/server /root/server




RUN chmod +x ./server

EXPOSE 80

CMD ["./server"]