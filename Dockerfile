##
## NOTE: to build this image you must be in the root of this repository
##

FROM wernight/phantomjs:latest

COPY rasterize.js rasterize.js
COPY server server

EXPOSE 7777

ENTRYPOINT ["./server"]
