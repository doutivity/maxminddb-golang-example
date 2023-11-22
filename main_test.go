package maxminddb_golang_example

import (
	"net"
	"testing"

	"github.com/oschwald/maxminddb-golang"
	"github.com/stretchr/testify/require"
)

type (
	Country struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	}
)

const (
	countryFile = "./GeoLite2-Country_20230721/GeoLite2-Country.mmdb"
	cityFile    = "./GeoLite2-City_20230721/GeoLite2-City.mmdb"
)

func TestGeoLite2CountryLookupTurkey(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testLookupTurkey(t, countryFile)
}

func TestGeoLite2CityLookupTurkey(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testLookupTurkey(t, cityFile)
}

func testLookupTurkey(t *testing.T, file string) {
	t.Helper()

	db, err := maxminddb.Open(file)
	require.NoError(t, err)
	defer db.Close()

	ip := net.ParseIP("95.0.0.0")

	var record Country

	err = db.Lookup(ip, &record)
	require.NoError(t, err)

	require.Equal(t, "TR", record.Country.ISOCode)
}

func BenchmarkGeoLite2CountryLookupTurkey(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	benchmarkLookupTurkey(b, countryFile)
}

func BenchmarkGeoLite2CityLookupTurkey(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	benchmarkLookupTurkey(b, cityFile)
}

func benchmarkLookupTurkey(b *testing.B, file string) {
	b.Helper()

	db, err := maxminddb.Open(file)
	require.NoError(b, err)
	defer db.Close()

	ipRoundRobin := NewIPV4RoundRobin(net.ParseIP("95.0.0.0"), net.ParseIP("95.7.255.255"))

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var record Country

			ip := ipRoundRobin.Next()
			err = db.Lookup(ip, &record)
			require.NoError(b, err)

			require.Equal(b, "TR", record.Country.ISOCode)
		}
	})
}
