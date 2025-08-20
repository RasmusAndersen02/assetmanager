package main

import (
	"net/http"
	// "fmt"
	"github.com/Rasmusandersen02/assetmanager/assets"
)

func main() {
	http.HandleFunc("/", assets.IndexHandler)
	http.HandleFunc("/uploadAsset", assets.NewAssetHandler)
	http.HandleFunc("/updateAsset", assets.EditAssetHandler)

	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3333", nil)
}
