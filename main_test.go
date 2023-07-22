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

func TestLookupTurkey(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	db, err := maxminddb.Open("./GeoLite2-Country_20230721/GeoLite2-Country.mmdb")
	require.NoError(t, err)
	defer db.Close()

	ip := net.ParseIP("95.0.0.0")

	var record Country

	err = db.Lookup(ip, &record)
	require.NoError(t, err)

	require.Equal(t, "TR", record.Country.ISOCode)
}

func BenchmarkLookupTurkey(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	db, err := maxminddb.Open("./GeoLite2-Country_20230721/GeoLite2-Country.mmdb")
	require.NoError(b, err)
	defer db.Close()

	ipRoundRobin := NewIPV4RoundRobin(net.ParseIP("95.0.0.0"), net.ParseIP("95.7.255.255"))

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
