test:
	go test ./... -v -count=1

# make download-maxminddb-country YOUR_LICENSE_KEY=YOUR_LICENSE_KEY
download-maxminddb-country:
	echo "# NOT WORKING!"
	# wget https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=$(YOUR_LICENSE_KEY)&suffix=tar.gz

# make download-maxminddb-city YOUR_LICENSE_KEY=YOUR_LICENSE_KEY
download-maxminddb-city:
	echo "# NOT WORKING!"
	# wget https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=$(YOUR_LICENSE_KEY)&suffix=tar.gz

