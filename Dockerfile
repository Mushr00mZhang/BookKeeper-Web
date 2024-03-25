FROM golang:alpine AS be-builder
ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /backend
COPY ./backend /backend
RUN apk add --no-cache gcc musl-dev
RUN go build -ldflags -w -o svc .

FROM node:20-alpine AS fe-builder
WORKDIR /frontend
COPY ./frontend /frontend
RUN npm config set registry https://registry.npmmirror.com
RUN yarn install --slient
RUN yarn global add vite
RUN vite build

FROM scratch
COPY --from=be-builder /backend/svc /app/svc
# COPY --from=fe-builder /frontend/dist /app/dist
EXPOSE 8080
CMD ["/app/svc"]