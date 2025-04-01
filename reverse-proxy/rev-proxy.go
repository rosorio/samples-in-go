/*
 * Code submitted on Github gist
 * https://gist.github.com/JalfResi/6287706
 */

package main

import(
    "flag"
    "log"
    "net/url"
    "net/http"
    "net/http/httputil"
)


func main() {
    var rurl = flag.String("remote", "", "Remote URL")
    var lbind = flag.String("local", "", "Bind address")
    flag.Parse()

    remote, err := url.Parse(*rurl)
    if err != nil {
            panic(err)
    }

    handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
        return func(w http.ResponseWriter, r *http.Request) {
                log.Println(r.URL)
                r.Host = remote.Host
                w.Header().Set("X-Ben", "Rad")
                p.ServeHTTP(w, r)
            }
    }
        
    proxy := httputil.NewSingleHostReverseProxy(remote)
    http.HandleFunc("/", handler(proxy))
    err = http.ListenAndServe(*lbind, nil)
    if err != nil {
        panic(err)
    }
}
