package main

import (
	"pricetrack/collector"
	dbinterface "pricetrack/db_interface"
)

func main(){
	con := dbinterface.Createconnection()
	urls := dbinterface.ListURLS(con)
	c := collector.NewColWithConfig(true,2,true,"amz")
	for _,url := range urls{
		c.Visit(url)
	}
	c.Wait()
}