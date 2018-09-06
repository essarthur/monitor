/****************************************************************************************
 *	★ ESSENTIA CLIENT MONITOR
      Copyright ESSENTIA  (2018)
 *    Savchenko Arthur
 ****************************************************************************************/
package main


import (
	"net/http"
	"fmt"
	"os"
)


var(
    ipserv="http://111.222.111.222:5898"
)


/****************************************************************************************
 * DATETIME         : 06-09-2018 17:44
 * NOTES            : Send log information to server
 ****************************************************************************************/
func main() {
     
     // Args
     arg    :=os.Args
     Project:=arg[1]
     Module :=arg[2]
     Text   :=arg[3]
     Status :=arg[4]
  
     Send_Info(Project,Module, Text, Status)
}


// *************************************************************
// Send to log server information
// *************************************************************
func Send_Info(Proj,Module,Text,Status string ){

    url    := ipserv+"/api/add/"+Proj+"*"+Module+"*"+Text+"*"+Status
	re,errr:= http.NewRequest("GET", url, nil)
	
	if errr!=nil{
       fmt.Println(errr.Error()) 
       return
	}

	res, erd := http.DefaultClient.Do(re)
	if erd!=nil{
       fmt.Println(erd.Error()) 
       return
	}

	defer res.Body.Close()
}
