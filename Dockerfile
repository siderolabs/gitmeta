ARG GOLANG_IMAGE
FROM ${GOLANG_IMAGE} AS build

ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org

WORKDIR /gitmeta
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
RUN go mod verify
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./main.go ./
RUN go list -mod=readonly all >/dev/null
RUN ! go mod tidy -v 2>&1 | grep .

ENV GOOS linux
ENV GOARCH amd64
RUN go build -o /gitmeta-${GOOS}-${GOARCH} .

ENV GOOS darwin
ENV GOARCH amd64
RUN go build -o /gitmeta-${GOOS}-${GOARCH} .

ENV GOOS linux
ENV GOARCH amd64
COPY ./hack ./hack
RUN chmod +x ./hack/test.sh
RUN ./hack/test.sh --all

FROM scratch AS image
COPY --from=build /gitmeta-linux-amd64 /gitmeta
ENTRYPOINT [ "/gitmeta" ]
