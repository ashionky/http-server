/**
 * @Author pibing
 */

package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)
var store_arr []string   //全局
var lock sync.Mutex

func init() {
	store_arr =make([]string,0)   //每次启动归0
}
func main() {

	http.HandleFunc("/api",testHandler)

	//err :=http.ListenAndServeTLS(":11111","server.crt","server_no_password.key", nil)
	err :=http.ListenAndServe(":11111",nil)
	if err !=nil{
		fmt.Println("ListenAndServeTLS err:",err)
		return
	}
}
//响应结构体
type ResData struct {
	ErrMsg string  `json:"err_msg"`
	Barr   []bool  `json:"bool_arr"`
}

//请求结构体
type reqStruct struct {
	ReqArr   []string   `json:"req_arr"`
}

func testHandler(res http.ResponseWriter, req *http.Request)  {
	reqdata := reqStruct{}
	res_arr :=make([]bool,0)   //响应的布尔数组
	data :=ResData{}          //响应结构体
	var err error
	var arr  []string     //请求时的字符串数组
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		data.ErrMsg="参数错误"
		resSend(res,data)
		return
	}

	err =json.Unmarshal(body,&reqdata)
	if err != nil {
		data.ErrMsg="参数错误"
		resSend(res,data)
		return
	}
	arr = reqdata.ReqArr

	if  len(arr) == 0 {
		data.ErrMsg="参数错误"
		resSend(res,data)
		return
	}

    lock.Lock()      //加锁
	for i:=0;i<len(arr);i++ {
		str :=arr[i]
		isbool := CheckStoreArr(str)
		res_arr=append(res_arr,isbool)
		//不存在则放入store_arr
		if !isbool {
			store_arr=append(store_arr,str)
		}
	}
	lock.Unlock()   //释放锁

	data.Barr=res_arr
	resSend(res,data)
	return
}
//响应
func resSend(res http.ResponseWriter, data interface{}){
	res.Header().Set("Content-Type","application/json")
	bytese, err := json.Marshal(data)
	if err !=nil {
		fmt.Println("Marshal data err",err)

		res.Write([]byte("500"))
		return
	}
	res.Write(bytese)
}

func CheckStoreArr(str string) bool  {
	for i:=0;i<len(store_arr);i++ {
		if store_arr[i]==str {
			return true
		}
	}
	return false

}
