ifeq ($(shell id -u), 0)
	SUDO_USER=
else
	SUDO_USER=sudo
endif

NAME=RnD_Jetson_optimization
BASE_DIR=$(shell pwd)

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20

#.SILENT:

define colored
	@echo '${GREEN}$1${RESET}'
endef


test:
	${call colored,test is running...}
	go test ./... -v -count=1

bench:
	${call colored,bechmarks is running...}
	go test ./... -v -bench=. -benchmem -benchtime=1000000x -count=10 > bench.txt
	benchstat bench.txt

# make download-maxminddb-country YOUR_LICENSE_KEY=YOUR_LICENSE_KEY
download-maxmind-db-country:
	${call colored,download-maxmind-db-country is running...}
	 wget 'https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=$(YOUR_LICENSE_KEY)&suffix=tar.gz'

# make download-maxminddb-city YOUR_LICENSE_KEY=YOUR_LICENSE_KEY
download-maxminddb-city:
	${call colored,download-maxminddb-city is running...}
	wget 'https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=$(YOUR_LICENSE_KEY)&suffix=tar.gz'

## Formats the code.
format:
	${call colored,formatting is running...}
	go vet ./...
	go fmt ./...

## Fix-imports order.
fix-imports:
	${call colored,fixing imports...}
	./scripts/fix-imports-order.sh