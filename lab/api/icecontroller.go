package api

import (
  "github.com/stretchr/goweb"
  "github.com/stretchr/goweb/context"
  //"net/url"
  "fmt"
  //"time"
  "net/url"
  "db"

)

type IceData struct {
  
  Output []string
  TimeStamp []string

}

type IceController struct {
  
  Datas []*IceData
}

func GetQueryValue( name string, query string) (metric string){

  values, _ := url.ParseQuery(query)
  //fmt.Println( values)
  metrics := values[name]
  //fmt.Println( metrics[0])
  if len(metrics) > 0 {
     metric = metrics[0]
  }
 
  return 

}

func (r *IceController) Read(brand string, ctx context.Context) error{

  fmt.Println("already interfacing with cassandra")
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")  
  
  query := ctx.HttpRequest().URL.RawQuery
  id := GetQueryValue("user_id",query)
  column := GetQueryValue("op", query)
  fmt.Println(id)
  newsclusters := db.ReadValue("ice", "user_data", id,column)  
 
  /*for _, cluster := range newsclusters {
    fmt.Println(cluster)
  }*/
  fmt.Println(column)
  data := new(IceData)
  data.Output = newsclusters
  data.TimeStamp=[]string{"testing", "timestamp" }
  return goweb.API.RespondWithData(ctx , data)
}
