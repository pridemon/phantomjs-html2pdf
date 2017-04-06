# phantomjs-html2pdf
## Description
This is HTML to PDF converter microservice based on phantomjs and rasterize.js

Container is based on https://hub.docker.com/r/wernight/phantomjs/ version and just adds `rasterize.js` script and `server` executable to it.

## Usage
Simply start this container with command
```bash
docker run -d --name phantomjs-html2pdf -p 7777:7777 ontrif/phantomjs-html2pdf
```

Then you can convert any html to pdf by sending it to `http://localhost:7777/pdf` endpoint. Html content of page must be sent as POST body:
```bash
curl -d '@/path/to/file.html' 'http://localhost:7777/pdf' -o book.pdf
```
