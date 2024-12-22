package models

import "Go-Blog/config"

type PigeonholeRes struct {
	config.Viewer
	config.SystemConfig
	Categorys []Category
	Lines     map[string][]Post
}
