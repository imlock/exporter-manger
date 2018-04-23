package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ximply/exporter-manger/models"
)


// node exporter
type NodeController struct {
	beego.Controller
}

func (c *NodeController) NodeMetrics() {
	r := models.NodeMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// php-fpm exporter
type PhpfpmController struct {
	beego.Controller
}

func (c *PhpfpmController) PhpfpmMetrics() {
	r := models.PhpfpmMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// redis exporter
type RedisController struct {
	beego.Controller
}

func (c *RedisController) RedisMetrics() {
	r := models.RedisMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// memcached exporter
type MemchachedController struct {
	beego.Controller
}

func (c *MemchachedController) MemcachedMetrics() {
	r := models.MemcachedMetrics()
	c.Ctx.Output.Body([]byte(r))
}