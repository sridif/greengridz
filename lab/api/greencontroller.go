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

type GreenController struct {
  
  Datas []*IceData
}

func (r *GreenController) Read(brand string, ctx context.Context) error{

  fmt.Println("already interfacing with cassandra")

  //  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")  
  //newsclusters := db.ReadValues("ice", "user_data","1")  

 /* for _, cluster := range newsclusters {
    fmt.Println(cluster)
  }*/
  fmt.Println("rest")
  data := new(IceData)
  data.Output = []string{"weojrign?"}
  data.TimeStamp=[]string{"testing", "timestamp" }
  return goweb.API.RespondWithData(ctx , data)
}
