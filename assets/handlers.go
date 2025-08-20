package assets

import (
	"fmt"
	"html/template"
	"net/http"
)

type Asset struct {
	ID             string
	Owner          string
	DateOfCreation string
}
type AssetCollection map[string]Asset

func (ac AssetCollection) AddAsset(asset Asset) {
	ac[asset.ID] = asset
}

func (ac AssetCollection) GetAsset(id string) (Asset, bool) {
	asset, exists := ac[id]
	return asset, exists
}

var assetTestData = []Asset{
	{ID: "A001", Owner: "Peter", DateOfCreation: "2025.04.01"},
	{ID: "A002", Owner: "Lucas", DateOfCreation: "2023.06.02"},
}
var assetCollection = make(AssetCollection)

var tmpl = template.Must(template.ParseFiles("./static/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	assetCollection.AddAsset(assetTestData[0])
	assetCollection.AddAsset(assetTestData[1])

	tmpl.ExecuteTemplate(w, "index", assetCollection)
}

func NewAssetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newAsset := Asset{
		ID:             r.FormValue("id"),
		Owner:          r.FormValue("owner"),
		DateOfCreation: r.FormValue("date-of-creation"),
	}
	assetCollection.AddAsset(newAsset)

	tmpl.ExecuteTemplate(w, "asset-list", assetCollection)
}

func EditAssetHandler(w http.ResponseWriter, r *http.Request) {
	// ID change creates new asset, non-key fields fine
	id := r.URL.Query().Get("id")

	asset, exists := assetCollection.GetAsset(id)
	if !exists {
		fmt.Println("Asset does not exist")
	} else {
		fmt.Println(asset)
	}
	tmpl.ExecuteTemplate(w, "edit-asset", asset)
}
