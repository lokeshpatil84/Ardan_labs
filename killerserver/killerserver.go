package main

import (
	"fmt"
	"os"
)


 func main(){
	fmt.Println("server.pid")
 }

 func Killerserver(Pidfile string) error{
      file, err:=os.Open(Pidfile)
      if err != nil{	
		return err
	  }
	  defer file.Close()

	  var pid int
	if _,err:= fmt.Fscanf(file,"%d",&pid);err!=nil{
	}
	return nil
 }