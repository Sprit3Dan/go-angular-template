FROM node:lts-alpine3.16 as build

RUN mkdir /app
WORKDIR /app

COPY . .

RUN npm i
RUN npm run build --prod

FROM nginx:latest

COPY --from=build /app/dist/ui /usr/share/nginx/html

EXPOSE 80

CMD [ "nginx", "-g", "daemon off;" ]