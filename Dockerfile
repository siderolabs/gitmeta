ARG GOLANG_IMAGE
FROM ${GOLANG_IMAGE}

ENV CGO_ENABLED 0
ENV GO111MODULES on

WORKDIR /gitmeta
COPY ./ ./
RUN go mod download
RUN go mod verify
RUN go mod tidy
RUN go mod vendor

ENV GOOS linux
ENV GOARCH amd64
RUN go build -o /build/gitmeta-${GOOS}-${GOARCH} .

ENV GOOS darwin
ENV GOARCH amd64
RUN go build -o /build/gitmeta-${GOOS}-${GOARCH} .

ENV GOOS linux
ENV GOARCH amd64
COPY ./hack ./hack
RUN chmod +x ./hack/test.sh
RUN ./hack/test.sh --all

FROM scratch
COPY /build/gitmeta-linux-amd64 /gitmeta
ENTRYPOINT [ "/gitmeta" ]
