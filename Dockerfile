# build base image from official for Go with 1.18 pre-installed
FROM golang:1.18-bullseye

# creates sandbox folder and set it as the workding directory
WORKDIR /go/src/wallet-sandbox
COPY . /go/src/wallet-sandbox/


#  download dependencies
RUN go mod tidy 

# tell Docker to accept port 
EXPOSE 8088

# set entry point command i.e docker run container
CMD ["make", "run-server"]
