FROM golang:1.17
WORKDIR /go/src
RUN apt-get update && apt install build-essential -y
EXPOSE 3000 3000
CMD ["tail", "-f","/dev/null"]