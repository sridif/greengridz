/* 

Access point code

Todo: 
1: migrate gron moln and persanal site away from this aws instance.
2: 


*/

package main
import (
  "flag"
  "fmt"
  "net/http"
  "net/http/httputil"
)


var addr = flag.String("addr", "demo.greengridz.com:9921", "address of the service ")

func main() {

  proxy := httputil.ReverseProxy{Director : func(req *http.Request) {
    req.URL.Scheme = "http"
    switch req.Host {

      case "demo.greengridz.com" :
        req.URL.Host = "www.gronmoln.com:9925"
        fmt.Println("overhere")
     
      case "moderator.kanybal.com" :
        req.URL.Host = "moderator.kanybal.com:9927"

      case "metrics.kanybal.com" :
        req.URL.Host = "metrics.kanybal.com:9926"
      case "www.kanybal.com" :
        req.URL.Host = "www.kanybal.com:9922"
      case "demo.kanybal.com":
        req.URL.Host = "demo.kanybal.com:9923"
      case "api.kanybal.com" :
        req.URL.Host = "api.kanybal.com:9924"

      default :
        req.URL.Host = "www.kanybal.com:9922"

    }
  
  }}

 http.ListenAndServe(*addr, &proxy)

}

