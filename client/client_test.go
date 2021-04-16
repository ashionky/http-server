/**
 * @Author pibing
 */

package client

import (
	"fmt"
	"testing"
	"time"
)

func TestBcjClient_arr1(t *testing.T) {


	//go Start()   //服务端启动
	time.Sleep(2*time.Second)       //也可以用chan阻塞功能

	arr0 :=[]string{}
	_, err0 := BcjClient(arr0)    //参数不能为空
	data0:="参数错误"              //预期结果
	if err0.Error()!=data0 {
		t.Error("请求数组：",arr0," 预期结果与实际结果不同，测试失败")
	}else {
		t.Log("请求数组：",arr0," 测试通过")
	}

	fmt.Println("")
	arr1:=[]string{"a","b"}
	bool1, err1 := BcjClient(arr1) //实际结果
	data1:=[]bool{false,false}     //预期结果
	if err1 !=nil {
		t.Error(err1)
	}
	if !checkResult(data1,bool1) {
		t.Error("请求数组：",arr1," 预期结果与实际结果不同，测试失败")
	}else {
		t.Log("请求数组：",arr1," 测试通过")
	}


	fmt.Println("")
	arr2:=[]string{"a","b","c"}
	bool2, err2 := BcjClient(arr2)      //实际结果
	data2:=[]bool{true,true,false}     //预期结果
	if err2 !=nil {
		t.Error(err2)
	}
	if !checkResult(data2,bool2) {
		t.Error("请求数组：",arr2," 预期结果与实际结果不同，测试失败")
	}else {
		t.Log("请求数组：",arr2," 测试通过")
	}

	fmt.Println("")
	arr3:=[]string{"d","d","c"}        //重复字符串情况
	bool3, err3 := BcjClient(arr3)      //实际结果
	data3:=[]bool{false,true,true}     //预期结果
	if err3 !=nil {
		t.Error(err3)
	}
	if !checkResult(data3,bool3) {
		t.Error("请求数组：",arr3," 预期结果与实际结果不同，测试失败")
	}else {
		t.Log("请求数组：",arr3," 测试通过")
	}

}


//arr 预期结果
//result_arr 实际结果
func checkResult(arr,result_arr []bool) bool {
	if len(arr)==len(result_arr) {
		for i:=0;i<len(arr);i++ {
			if arr[i] != result_arr[i] {
				return false
			}
		}
	}else {
		return false
	}
	return true

}

