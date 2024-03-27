FROM golang:alpine AS be-builder
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
# ENV GOPROXY https://goproxy.cn,direct
WORKDIR /backend
COPY ./backend .
# RUN apk add --no-cache gcc musl-dev
RUN go build -ldflags="-w -s" -o svc .

FROM node:20-alpine AS fe-builder
WORKDIR /frontend
COPY ./frontend .
# RUN npm config set registry https://registry.npmmirror.com
RUN yarn install --slient
RUN yarn global add vite
RUN vite build

FROM scratch
WORKDIR /app
COPY --from=be-builder /backend/svc .
COPY --from=fe-builder /frontend/dist/ ./dist/
EXPOSE 8080
CMD ["/app/svc"]