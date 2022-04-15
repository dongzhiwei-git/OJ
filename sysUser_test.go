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

func TestDelProblemById(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	pro := new(services.Problem)
	pro.DelProblemById(3)
}

func TestAddSolution(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	sol := new(services.Solution)
	_, err := sol.AddSolution(34, "!23 ", 1, 1, 34, 1, "3")
	if err != nil {
		fmt.Println()
	}

}

func TestQueryContestByConId(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	con := new(services.Contest)
	conID, err := con.QueryContestByConId(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conID)
}
