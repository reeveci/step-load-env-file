FROM golang AS builder

WORKDIR /app
COPY . .

ENV GOFLAGS="-buildvcs=false"
ENV CGO_ENABLED=0
RUN go build -o /usr/local/bin/reeve-step .

FROM alpine

RUN apk add jq
COPY --chmod=755 --from=builder /usr/local/bin/reeve-step /usr/local/bin/

# FILE: name of the env file
ENV FILE=.env
# LOAD_ALL=true|false
ENV LOAD_ALL=false
# ENV_<name>: Variables to be loaded from the file and their corresponding runtime variable names to be using in Reeve

ENTRYPOINT ["reeve-step"]
