FROM node:alpine

WORKDIR '/npm'

COPY ./admin/package.json /npm/package.json
RUN npm install

COPY ./admin /npm/.

CMD npm run serve