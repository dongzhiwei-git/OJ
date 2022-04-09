package main

import (
	fmt "fmt"
	"inherited/internal"
	"inherited/internal/services"
	"log"
	"testing"
)

func TestCreateSysUser(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	SysUser := new(services.SysUser)
	SysUser.CreateSysUser("rot", "2355")
	fmt.Println("er")

}


func TestQueryAllProblem(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	problem := new(services.Problem)
	info, _ := problem.QueryAllProblem()
	fmt.Println(info)

}


