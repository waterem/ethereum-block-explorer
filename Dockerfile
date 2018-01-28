FROM node:latest

RUN npm install mysql
RUN npm install jayson
RUN npm install getopts
RUN npm install log4js
RUN npm install web3@0.20