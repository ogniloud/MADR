FROM node:22.2.0-alpine as build
WORKDIR /app
COPY package.json ./
COPY package-lock.json package-lock.json
RUN npm install
COPY . ./
RUN npm run build .env

# release step
FROM nginx:1.26.0-alpine as release
COPY --from=build /app/build /usr/share/nginx/html/
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]