# 
# FROM node:lts-alpine as build-stage
# RUN addgroup app && adduser -S -G app app
# WORKDIR /app
# COPY --chown=app:node package*.json ./
# RUN chown -R app:node /app
# USER app
# RUN npm install
# COPY --chown=app:node . .
# RUN chmod +x node_modules/.bin/*
# RUN npm run build
# production (production-stage)
FROM nginx:stable-alpine as production-stage
# COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY ./index.html /usr/share/nginx/html/
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]