# syntax=docker/dockerfile:1

# Step 1: build binary
FROM golang:1.18 as builder

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and 
# only redownloading them in subsequent builds if they change
COPY Makefile ./
RUN make cli
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build OS="linux"


# Step 2: copy binary from step 1
FROM loads/alpine:3.8

ENV GF_GERROR_BRIEF=true

WORKDIR /app

COPY --from=builder /src/bin/linux_amd64/gf2-demo-api .

EXPOSE 9000

ENTRYPOINT [ "./gf2-demo-api" ]
