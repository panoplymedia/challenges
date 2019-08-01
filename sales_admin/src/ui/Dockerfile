FROM node:lts-alpine

RUN apk update

RUN mkdir /app

COPY ./tslint.json /app/
COPY ./tsconfig.spec.json /app/
COPY ./tsconfig.json /app/
COPY ./tsconfig.app.json /app/
COPY ./package.json /app/
COPY ./package-lock.json /app/
COPY ./karma.conf.js /app/
COPY ./browserslist /app/
COPY ./angular.json /app/

COPY ./e2e /app/e2e
COPY ./src /app/src