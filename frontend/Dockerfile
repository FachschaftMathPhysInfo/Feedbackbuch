# build stage
FROM node:lts-alpine as feedbackbuch-frontend-build-stage
WORKDIR /app
COPY package*.json *.config.js ./
RUN npm install
COPY ./src ./src
RUN npm run build

# production stage
FROM nginx:stable-alpine as feedbackbuch-frontend-production-stage
COPY --from=feedbackbuch-frontend-build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
