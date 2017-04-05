# phantomjs-html2pdf
## Description
This container is based on https://hub.docker.com/r/wernight/phantomjs/ version and adds to it `rasterize.js` script for converting html to pdf and `server` executable for providing http server for sending html to and receiving resulting pdf back.

## Usage
Simply start this container with command
```bash
docker run -d --name phantomjs-html2pdf -p 7777:7777 ontrif/phantomjs-html2pdf
```

Then you can convert any html to pdf by sending it to `http://localhost:7777/pdf` endpoint. Html content of page must be sent as POST body:
```bash
curl -d '@/path/to/file.html' 'http://localhost:7777/pdf' -o book.pdf
```
