package main

import (
	"fmt"
)

func main()  {
	fmt.Println("测试")
	mgr := InitModules()
	
	test_md, err := mgr.GetModule("test_module")
	if err != nil {
		fmt.Println("错误")
	}

	test_md.module.Init()
}