package main

import (
	"fmt"
	"os"
)

func main()  {
	n,err:=os.Create("plan.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}

	str:="I’m not afraid of difficulties and insist on learning programming"
	f,err:=n.Write([]byte(str))
	if err!=nil{
		fmt.Println(str)
		return
	}
	fmt.Println("写入字节数",f)
	p:=make([]byte,f)
	m,err:=n.ReadAt(p,0)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("读出字节数",m)
	fmt.Println(string(p))
	n.Close()


}