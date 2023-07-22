# Приклад використання MaxMind в Go
https://dou.ua/forums/topic/35195/

# Support Ukraine 🇺🇦
- Help Ukraine via [SaveLife](https://savelife.in.ua/en/donate-en/) fund
- Help Ukraine via [National Bank of Ukraine](https://bank.gov.ua/en/news/all/natsionalniy-bank-vidkriv-spetsrahunok-dlya-zboru-koshtiv-na-potrebi-armiyi)
- More info on [war.ukraine.ua](https://war.ukraine.ua/) and [MFA of Ukraine](https://twitter.com/MFA_Ukraine)

# Structure
```bash
tree . -h
```
```text
├── [4.0K]  GeoLite2-City_20230721
│   ├── [  55]  COPYRIGHT.txt
│   ├── [ 68M]  GeoLite2-City.mmdb
│   ├── [ 398]  LICENSE.txt
│   └── [ 116]  README.txt
├── [ 33M]  GeoLite2-City_20230721.tar.gz
├── [4.0K]  GeoLite2-Country_20230721
│   ├── [  55]  COPYRIGHT.txt
│   ├── [5.8M]  GeoLite2-Country.mmdb
│   └── [ 398]  LICENSE.txt
├── [3.0M]  GeoLite2-Country_20230721.tar.gz
├── [ 253]  go.mod
├── [ 881]  go.sum
├── [1.0K]  LICENSE
├── [ 155]  main_test.go
├── [ 525]  Makefile
└── [  91]  README.md
```

# Testing
```bash
go test ./... -bench=. -benchmem -benchtime=1000000x
```
```text
goos: linux
goarch: amd64
pkg: github.com/doutivity/maxminddb-golang-example
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkLookupTurkey-12    	 1000000	      2448 ns/op	      34 B/op	       3 allocs/op
PASS
ok  	github.com/doutivity/maxminddb-golang-example	2.453s
```

# Packages
* https://github.com/oschwald/maxminddb-golang
* https://github.com/ip2location/ip2location-go
