FROM node:20.9.0-alpine as node-pnpm
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app
COPY styles ./styles
COPY pkg/view ./templates
COPY static ./static
COPY package.json pnpm-lock.yaml *.js ./

FROM node-pnpm as build-css
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm build

FROM golang:1.21.5-alpine as base
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd ./cmd
COPY internal ./internal
COPY --from=build-css /app/templates ./pkg/view
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

FROM base as test
RUN go test ./...

FROM test as build
RUN go build -o main ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build-css /app/static ./static
COPY --from=build /app/main ./
COPY configs ./configs
CMD [ "/app/main" ]
