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
        req.URL.Host = "localhost:9922"
        fmt.Println("overhere")
     
      case "api.greengridz.com" :
        req.URL.Host = "localhost:9923"
        fmt.Println("testing")

      /* default :
        req.URL.Host = "loaclhost:9922" */

    }
  
  }}

 http.ListenAndServe(*addr, &proxy)

}

