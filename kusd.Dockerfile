FROM golang:1.9.2-alpine as builder
RUN apk update && apk add --update git make gcc musl-dev linux-headers

WORKDIR /kusd/
ADD . .
RUN make kusd

FROM alpine:3.7
WORKDIR /kusd/
COPY --from=builder /kusd/build/bin/kusd .
EXPOSE 22334
EXPOSE 22334/udp
ADD release/kusd_with_new_account.sh .
ENTRYPOINT ["./kusd_with_new_account.sh"]
