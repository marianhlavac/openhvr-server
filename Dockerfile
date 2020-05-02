FROM node:alpine AS clientbuild

RUN mkdir -p /tmp/clientbuild
WORKDIR /tmp/clientbuild

ADD ./configurator-app/package*.json /tmp/clientbuild/
RUN npm install

ADD ./configurator-app /tmp/clientbuild/
RUN npm run build

# ---

FROM library/golang AS serverbuild
ENV APP_DIR $GOPATH/src/github.com/mmajko/openhvr-server

RUN mkdir -p $APP_DIR
WORKDIR $APP_DIR

COPY --from=clientbuild /tmp/clientbuild/ $APP_DIR/configurator-app
ADD . $APP_DIR

RUN go build

ENTRYPOINT ./openhvr-server
EXPOSE 47023
