# build stage
FROM node:lts-alpine as feedbackbuch-frontend-build-stage
WORKDIR /app
COPY package*.json *.config.js ./
RUN npm install
COPY ./src ./src

ARG SERVER_WS
ENV SERVER_WS "$SERVER_WS"
ARG SERVER_HTTP
ENV SERVER_HTTP "$SERVER_HTTP"
ENV NODE_OPTIONS "--openssl-legacy-provider"

RUN sed -i "s,SERVER\_WS,\"$SERVER_WS\",g" ./src/vue-apollo.js
RUN sed -i "s,SERVER\_HTTP,\"$SERVER_HTTP\",g"  ./src/vue-apollo.js
RUN npm run build

# production stage
FROM nginx:stable-alpine as feedbackbuch-frontend-production-stage
COPY --from=feedbackbuch-frontend-build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
