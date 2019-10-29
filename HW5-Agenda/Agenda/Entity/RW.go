package Entity

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"os"
)

var filePath string = "./Entity/Users.json"

func GetUsers() (users []User){
	josnStr,err := ioutil.ReadFile(filePath)
	str := strings.Replace(string(josnStr),"\n","",1)
	if str == ""{
		return 
	}
	if err == nil{
	}
	json.Unmarshal(josnStr,&users)
	return
}

func UpdateUsers(userList []User){
	os.Truncate(filePath,0)
	josnStr,err := json.Marshal(userList)
	if err == nil{
	}
	ioutil.WriteFile(filePath,josnStr,os.ModeAppend)
	os.Chmod(filePath,0777)
}


