# The Golang build container.
FROM golang:stretch AS go-build

WORKDIR /usr/src/app
COPY . .
RUN make compile-go install

# The Node build container
FROM node:stretch AS node-build

WORKDIR /usr/src/app
RUN mkdir lib
COPY lib/*.cc lib/
COPY package*.json *.gyp ./
COPY --from=go-build /usr/src/app/lib/*.so /usr/src/app/lib/*.h ./lib/
RUN npm install && npm run build

# The install container
FROM node:stretch AS install

WORKDIR /usr/src/app
RUN mkdir -p lib build/Release
COPY lib/*.js ./lib/
COPY cossd package*.json ./
RUN npm install --production
COPY --from=node-build /usr/src/app/build/Release/*.node build/Release/
COPY --from=go-build /usr/src/app/lib/*.so /usr/src/app/lib/
COPY --from=go-build /go/bin/cosscli /usr/src/app/

# By default, run the daemon with specified arguments.
ENTRYPOINT [ "./cossd" ]