package main

import (
	fmt "fmt"
	"inherited/internal"
	"inherited/internal/models"
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

func TestAddProblem(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	problem := new(services.Problem)
	data := models.Problem{}
	info := problem.AddProblem(data)
	fmt.Println(info)

}

func TestDelProblemById(t *testing.T){
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	pro := new(services.Problem)
	pro.DelProblemById(3)
}