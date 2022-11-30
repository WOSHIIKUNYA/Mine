package main

import (
	"test/api"
	"test/dao"
)

func main() {
	dao.Open()
	api.SetApi()
}
