FROM node:alpine

WORKDIR '/npm'

COPY package.json /npm/package.json
RUN npm install

COPY . /npm/.

CMD npm run serve