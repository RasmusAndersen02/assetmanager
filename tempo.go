



import "text/template"

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


