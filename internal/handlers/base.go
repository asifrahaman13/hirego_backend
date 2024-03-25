package handlers

import "github.com/asifrahaman13/hirego/config"

var Base *base

type base struct {
	conf *config.Config
}

func (b *base) Initialize(conf *config.Config) {
	Base = &base{
		conf: conf,
	}
}
