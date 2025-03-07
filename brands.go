package main

var brands = map[string]bool{
	// Fashion
	"adidas":         true,
	"nike":           true,
	"puma":           true,
	"reebok":         true,
	"under armour":   true,
	"new balance":    true,
	"vans":           true,
	"converse":       true,
	"skechers":       true,
	"timberland":     true,
	"clarks":         true,
	"zara":           true,
	"h&m":            true,
	"uniqlo":         true,
	"gucci":          true,
	"prada":          true,
	"hmns":           true,
	"zwitsal":        true,
	"chanel":         true,
	"louis vuitton":  true,
	"hermes":         true,
	"implora":        true,
	"saff & co ":     true,
	"giorgio armani": true,
	"aigner":         true,
	"one million":    true,
	"mont blanc":     true,
	"docmart":        true,
	"elizabeth":      true,
	"hoka":           true,
	"matahari":       true,
	"asics":          true,
	"aerostreet":     true,
	"warrior":        true,
	"ortuseight":     true,
	"specs":          true,
	"onitsuka":       true,
	"all star":       true,
	"kodachi":        true,
	"ventela":        true,
	"johnson":        true,
	"ortus":          true,
	"nineten":        true,
	"pro att":        true,
	"christian dior": true,
	"kanky":          true,
	"rayban":         true,
	"burberry":       true,
	"ralph lauren":   true,
	"apple":          true,
	"samsung":        true,
	"xiaomi":         true,
	"huawei":         true,
	"oppo":           true,
	"vivo":           true,
	"levis":          true,
	"sony":           true,
	"lg":             true,
	"rucas":          true,
	"lea":            true,
	"dickies":        true,
	"cardinal":       true,
	"sorex":          true,
	"eiger":          true,
	"gt man":         true,
	"calvin klein":   true,
	"wrangler":       true,
	"lois":           true,
	"giordano":       true,
	"nyam nyam":      true,
	"bebelove":       true,
	"bebelac":        true,
	"vidoran":        true,
	"kimia farma":    true,
	"antam":          true,
	"ubs":            true,
	"stone island":   true,

	"gap":       true,
	"miniso":    true,
	"maspion":   true,
	"cosmos":    true,
	"sekai":     true,
	"arashi":    true,
	"panasonic": true,
	"mitochiba": true,
	"krisbow":   true,
	"erigos":    true,

	"philips":         true,
	"polytron":        true,
	"sharp":           true,
	"asus":            true,
	"lenovo":          true,
	"hp":              true,
	"acer":            true,
	"dell":            true,
	"toshiba":         true,
	"erspo":           true,
	"league":          true,
	"mils":            true,
	"dorskin":         true,
	"lining":          true,
	"3 desember":      true,
	"nevada":          true,
	"hush puppies":    true,
	"evangeline":      true,
	"kahf":            true,
	"victoria secret": true,
	"jo malone":       true,
	"bvlgari":         true,
	"casablanca":      true,
	"dunhill":         true,
	"dior":            true,
	"axe":             true,
	"vitalis":         true,
	"bellagio":        true,
	"ysl":             true,
	"sauvage":         true,
	"versace ":        true,
	"scarlett":        true,

	// Elektronik
	"oneplus":      true,
	"bosch":        true,
	"dyson":        true,
	"breville":     true,
	"ninja":        true,
	"delonghi":     true,
	"whirlpool":    true,
	"kenmore":      true,
	"kitchenaid":   true,
	"black+decker": true,
	"dewalt":       true,
	"makita":       true,
	"hilti":        true,
	"bosch tools":  true,
	"stanley":      true,
	"milwaukee":    true,
	"fitbit":       true,
	"garmin":       true,
	"realme":       true,

	// Kecantikan
	"l'oreal":           true,
	"loreal":            true,
	"maybelline":        true,
	"nyx":               true,
	"sephora":           true,
	"estee lauder":      true,
	"clinique":          true,
	"lancome":           true,
	"nars":              true,
	"mac":               true,
	"urban decay":       true,
	"tarte":             true,
	"too faced":         true,
	"benefit":           true,
	"shiseido":          true,
	"sk-ii":             true,
	"kiehl's":           true,
	"the body shop":     true,
	"bath & body works": true,
	"scarlet":           true,
	"wardah":            true,
	"make over":         true,
	"pixy":              true,
	"eminence":          true,
	"somethinc":         true,
	"facetologi":        true,
	"yonex":             true,

	// Produk Rumah Tangga
	"ikea":              true,
	"pottery barn":      true,
	"west elm":          true,
	"crate & barrel":    true,
	"home depot":        true,
	"lowe's":            true,
	"wayfair":           true,
	"bed bath & beyond": true,
	"williams-sonoma":   true,
	"target":            true,
	"walmart":           true,
	"ace hardware":      true,
	"informa":           true,
	"olympic":           true,
	"modena":            true,
	"napolly":           true,
	"yongma":            true,
	"miyako":            true,
	"pigeon":            true,
	"sanken":            true,
	"oxy":               true,
	"wellcomm":          true,
	"diamond":           true,
}
