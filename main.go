package main

import (
	"time"

	"github.com/Ohimma/beegoSre/models"
	_ "github.com/Ohimma/beegoSre/routers"

	"github.com/Ohimma/beegoSre/utils"
	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
)

func main() {
	models.Init()
	utils.Che = cache.New(60*time.Minute, 120*time.Minute)
	beego.Run()
}
