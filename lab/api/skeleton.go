package api

import (
  "github.com/stretchr/goweb"
  //"github.com/stretchrcom/goweb/context"
  "net/http"
  "log"  
  "time"
  "net"
  "fmt"

)

const (
  Address string = ":9923"
)


func StartService() {

  s := &http.Server{
    Addr:           Address,
    Handler:        goweb.DefaultHttpHandler(),
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  listener, _ := net.Listen("tcp", Address)

  iceController := new(IceController)
  greenController := new(GreenController)
  regController := new(RegController)
  preController := new(PredictController)

  goweb.MapController(regController)
  goweb.MapController(preController)
  goweb.MapController(greenController)
  goweb.MapController(iceController )
 
  fmt.Println("wats up") 
  log.Fatalf("Error in Serve: %s", s.Serve(listener))
}

func Documentation(){
/*


*/
}
