module github.com/0Delta/yudenbot_devel

go 1.12

require (
	github.com/ChimeraCoder/anaconda v2.0.0+incompatible
	github.com/ChimeraCoder/tokenbucket v0.0.0-20131201223612-c5a927568de7 // indirect
	github.com/azr/backoff v0.0.0-20160115115103-53511d3c7330 // indirect
	github.com/dustin/go-jsonpointer v0.0.0-20160814072949-ba0abeacc3dc // indirect
	github.com/dustin/gojson v0.0.0-20160307161227-2e71ec9dd5ad // indirect
	github.com/garyburd/go-oauth v0.0.0-20180319155456-bca2e7f09a17 // indirect
	golang.org/x/net v0.0.0-20190327091125-710a502c58a2 // indirect
	gopkg.in/yaml.v2 v2.2.2
	github.com/0Delta/yudenbot_devel/eventdata v0.0.0
	github.com/0Delta/yudenbot_devel/twitter v0.0.0
	github.com/0Delta/yudenbot_devel/discord v0.0.0
)

	replace github.com/0Delta/yudenbot_devel/eventdata => ./eventdata
	replace github.com/0Delta/yudenbot_devel/twitter => ./twitter
	replace github.com/0Delta/yudenbot_devel/discord => ./discord