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

type RegController struct {
  
  Datas []*IceData

}

func (r *RegController) Read(brand string, ctx context.Context) error{

  fmt.Println("attempting to perform regression")
  fmt.Println("rest")

  data := new(IceData)
 
  //regression data format 
  data.Output = []string{"{10, 20, 30, 231}, {1,2,3,5}"}
  data.TimeStamp=[]string{"testing", "timestamp" }
  return goweb.API.RespondWithData(ctx , data)

}
