/*
 * Code submitted on Github gist
 * https://gist.github.com/JalfResi/6287706
 */

package main

import(
    "flag"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "net/http/httputil"
)

func main() {
    rurl := flag.String("remote", "", "Remote URL")
    lbind := flag.String("local", "", "Bind address")
    flag.Parse()
    remote, err := url.Parse(*rurl)

    handler := func(w http.ResponseWriter,r *http.Request) {

        r.Header["X-Forwarded-For"] = nil
        r.Host = remote.Host

        log.Println("Passing request to remote")

        dump, err := httputil.DumpRequest(r, true)
        if err == nil {
            fmt.Printf("Request:\n%q\n", dump)
        }
    
        proxy := httputil.NewSingleHostReverseProxy(remote)
        proxy.ServeHTTP(w, r)
    }
        
    http.HandleFunc("/", handler)

    log.Println("Starting")
    err = http.ListenAndServe(*lbind, nil)
    if err != nil {
        panic(err)
    }
}
