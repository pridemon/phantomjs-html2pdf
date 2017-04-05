package main

import (
    "os"
    "os/exec"
    "io"
    "fmt"
    "log"
    "bytes"
    "net/http"
    "github.com/satori/go.uuid"
)

func pdfHandler(rw http.ResponseWriter, req *http.Request) {
    // create temp file
    tmpHtml := "/tmp/" + uuid.NewV4().String() + ".html"
    tmpPdf  := "/tmp/" + uuid.NewV4().String() + ".pdf"

    tmpHtmlFile, err := os.Create(tmpHtml)
    if err != nil {
        log.Fatal(err)
    }
    defer os.Remove(tmpHtml)   // delete it after function exit
    defer os.Remove(tmpPdf)    // ...

    // save content from request body to the file
    io.Copy(tmpHtmlFile, req.Body)
    tmpHtmlFile.Close()

    size:= req.URL.Query().Get("size")
    if size == "" {
        size = "25cm*25cm"
    }

    // run phantomjs
    cmd := exec.Command("phantomjs", "rasterize.js", tmpHtml, tmpPdf, size)

    var stderr, stdout bytes.Buffer
    cmd.Stderr = &stderr
    cmd.Stderr = &stdout

    err = cmd.Run()
    if err != nil {
        log.Fatal(fmt.Sprint(err) + ": " + stderr.String() + "; " + stdout.String())
    }

    // send generated pdf in http response
    tmpPdfFile, err := os.Open(tmpPdf)
    if err != nil {
        log.Fatal(err)
    }

    io.Copy(rw, tmpPdfFile)
}

func main() {
    http.HandleFunc("/pdf", pdfHandler)
    log.Fatal(http.ListenAndServe(":7777", nil))
}
