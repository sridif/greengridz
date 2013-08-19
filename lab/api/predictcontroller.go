package api

import (
  "github.com/stretchr/goweb"
  "github.com/stretchr/goweb/context"
  //"net/url"
  "fmt"
  //"time"
  //"db"

)

/*type IceData struct {
  
  Output []string
  TimeStamp []string

}*/

type PredictController struct {
  
  Datas []*IceData

}

func (r *PredictController) Read(brand string, ctx context.Context) error{

  fmt.Println("attempting to perform regression")
  fmt.Println("rest")

  data := new(IceData)
 
  //regression data format 
  data.Output = []string{"{{10,10,12,12,32}}"}
  data.TimeStamp=[]string{"testing", "timestamp" }
  return goweb.API.RespondWithData(ctx , data)

}
