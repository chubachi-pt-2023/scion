FROM golang:1.23

WORKDIR .

# Air のインストール
RUN go install github.com/cosmtrek/air@v1.29.0

ENV PATH="/go/bin:${PATH}"

COPY . .

# COPY go.mod go.sum ./

RUN go mod download

# RUN go build -o main ./cmd/server

EXPOSE 4322

COPY .air.toml .

CMD ["air", "-c", ".air.toml"]