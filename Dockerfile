FROM node:20-alpine AS builder
WORKDIR /
COPY . .
RUN npm config set registry https://registry.npmmirror.com
RUN yarn install --slient
RUN yarn build

FROM nginx:alpine
WORKDIR /app
COPY --from=builder /default.conf /etc/nginx/conf.d/
COPY --from=builder /dist ./dist/
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]