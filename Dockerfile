FROM golang:1.25-bookworm AS build

WORKDIR /app

RUN apt-get update && apt-get install -y nodejs npm && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY package.json package-lock.json* ./
RUN npm install --include=optional
RUN npm install -D @tailwindcss/oxide-linux-x64-gnu @tailwindcss/oxide-linux-arm64-gnu

RUN go install github.com/a-h/templ/cmd/templ@v0.3.1020

COPY . .

RUN bash bobby.sh
RUN templ generate
RUN npx @tailwindcss/cli -i ./public/input.css -o ./public/styles.css
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=build /server /app/server
COPY --from=build /app/public /app/public

ENV PORT=8080

EXPOSE 8080

CMD ["/app/server"]