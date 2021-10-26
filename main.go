package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)


//1 遍历文件中的所有文件 并获取到文件名字 和文件数量
//2 判断是正反面 然后进行名字修改
func main() {
	FileSlices := []string{}
	newfileslice :=[]string{}
	files, err := ioutil.ReadDir("./file")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		FileSlices = append(FileSlices, file.Name())
	}
	SlicesLen := len(FileSlices) / 2

	for i := 0; i <SlicesLen; i++{
		zhengmian :=i
		fanmian :=i+SlicesLen
		newfileslice=append(newfileslice,FileSlices[zhengmian])
		newfileslice=append(newfileslice,FileSlices[fanmian])
		zhengmian +=1
	}
	for k,v :=range newfileslice{
		startfilenum :=k+1
		filename :=fmt.Sprintf("%s%s","./file/",v)
		//fmt.Println(filename)
		f,_ :=os.Open(filename)
		ret :=make([]byte,0)
		buf :=make([]byte,1024*10) //按照1024的大小读取文件
		n,_ :=f.Read(buf)
			//如果n==0 就说明读完了
			if n==0{
				break
			}
			ret =append(ret,buf...)
			savename :=fmt.Sprintf("./newfiles/合同(%d).jpg",startfilenum)
		    ioutil.WriteFile(savename,ret,0600)
	}
}