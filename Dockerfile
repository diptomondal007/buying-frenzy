ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /src
COPY ./ ./

RUN CGO_ENABLED=0 go build -mod=mod -o /buying-frenzy .

FROM scratch AS final

COPY --from=builder buying-frenzy buying-frenzy

EXPOSE 8000
ENTRYPOINT ["/buying-frenzy"]