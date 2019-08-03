FROM node:lts-alpine AS builder

RUN apk update

COPY package.json package-lock.json ./

RUN npm set progress=false && npm config set depth 0 && npm cache clean --force

RUN npm i && mkdir /ng-app && cp -R ./node_modules ./ng-app

WORKDIR /ng-app

COPY . .

RUN $(npm bin)/ng build --prod --build-optimizer

FROM nginx:stable-alpine

COPY nginx/default.conf /etc/nginx/conf.d/

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /ng-app/dist/sales-portal/ /usr/share/nginx/html/

ENTRYPOINT ["nginx", "-g", "daemon off;"]