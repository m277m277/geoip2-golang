package geoip2

import (
	"math/rand"
	"net"
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoIP2-City-Test.mmdb")
	require.NoError(t, err)

	defer reader.Close()

	testAddr := netip.MustParseAddr("81.2.69.160")
	record, err := reader.City(testAddr)
	require.NoError(t, err)

	m := reader.Metadata()
	assert.Equal(t, uint(2), m.BinaryFormatMajorVersion)
	assert.Equal(t, uint(0), m.BinaryFormatMinorVersion)
	assert.NotZero(t, m.BuildEpoch)
	assert.Equal(t, "GeoIP2-City", m.DatabaseType)
	assert.Equal(t,
		map[string]string{
			"en": "GeoIP2 City Test Database (fake GeoIP2 data, for example purposes only)",
			"zh": "小型数据库",
		},
		m.Description,
	)
	assert.Equal(t, uint(6), m.IPVersion)
	assert.Equal(t, []string{"en", "zh"}, m.Languages)
	assert.NotZero(t, m.NodeCount)
	assert.Equal(t, uint(28), m.RecordSize)

	assert.Equal(t, uint(2643743), record.City.GeoNameID)
	expectedNames := Names{
		German:              "London",
		English:             "London",
		Spanish:             "Londres",
		French:              "Londres",
		Japanese:            "ロンドン",
		BrazilianPortuguese: "Londres",
		Russian:             "Лондон",
	}
	assert.Equal(t, expectedNames, record.City.Names)

	assert.Equal(t, uint(6255148), record.Continent.GeoNameID)
	assert.Equal(t, "EU", record.Continent.Code)
	expectedContinentNames := Names{
		German:              "Europa",
		English:             "Europe",
		Spanish:             "Europa",
		French:              "Europe",
		Japanese:            "ヨーロッパ",
		BrazilianPortuguese: "Europa",
		Russian:             "Европа",
		SimplifiedChinese:   "欧洲",
	}
	assert.Equal(t, expectedContinentNames, record.Continent.Names)

	assert.Equal(t, uint(2635167), record.Country.GeoNameID)
	assert.False(t, record.Country.IsInEuropeanUnion)
	assert.Equal(t, "GB", record.Country.ISOCode)
	expectedCountryNames := Names{
		German:              "Vereinigtes Königreich",
		English:             "United Kingdom",
		Spanish:             "Reino Unido",
		French:              "Royaume-Uni",
		Japanese:            "イギリス",
		BrazilianPortuguese: "Reino Unido",
		Russian:             "Великобритания",
		SimplifiedChinese:   "英国",
	}
	assert.Equal(t, expectedCountryNames, record.Country.Names)

	assert.Equal(t, uint16(100), record.Location.AccuracyRadius)
	assert.InEpsilon(t, 51.5142, record.Location.Latitude, 1e-10)
	assert.InEpsilon(t, -0.0931, record.Location.Longitude, 1e-10)
	assert.Equal(t, "Europe/London", record.Location.TimeZone)

	assert.Equal(t, uint(6269131), record.Subdivisions[0].GeoNameID)
	assert.Equal(t, "ENG", record.Subdivisions[0].ISOCode)
	expectedSubdivisionNames := Names{
		English:             "England",
		BrazilianPortuguese: "Inglaterra",
		French:              "Angleterre",
		Spanish:             "Inglaterra",
	}
	assert.Equal(t, expectedSubdivisionNames, record.Subdivisions[0].Names)

	assert.Equal(t, uint(6252001), record.RegisteredCountry.GeoNameID)
	assert.False(t, record.RegisteredCountry.IsInEuropeanUnion)
	assert.Equal(t, "US", record.RegisteredCountry.ISOCode)
	expectedRegisteredCountryNames := Names{
		German:              "USA",
		English:             "United States",
		Spanish:             "Estados Unidos",
		French:              "États-Unis",
		Japanese:            "アメリカ合衆国",
		BrazilianPortuguese: "Estados Unidos",
		Russian:             "США",
		SimplifiedChinese:   "美国",
	}
	assert.Equal(t, expectedRegisteredCountryNames, record.RegisteredCountry.Names)

	assert.False(t, record.RepresentedCountry.IsInEuropeanUnion)

	// Test Network and IPAddress fields
	assert.Equal(t, testAddr, record.Traits.IPAddress)
	assert.True(t, record.Traits.Network.IsValid())
	assert.True(t, record.Traits.Network.Contains(testAddr))
}

func TestIsAnycast(t *testing.T) {
	for _, test := range []string{"Country", "City", "Enterprise"} {
		t.Run(test, func(t *testing.T) {
			reader, err := Open("test-data/test-data/GeoIP2-" + test + "-Test.mmdb")
			require.NoError(t, err)
			defer reader.Close()

			record, err := reader.City(netip.MustParseAddr("214.1.1.0"))
			require.NoError(t, err)

			assert.True(t, record.Traits.IsAnycast)
		})
	}
}

func TestMetroCode(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoIP2-City-Test.mmdb")
	require.NoError(t, err)
	defer reader.Close()

	record, err := reader.City(netip.MustParseAddr("216.160.83.56"))
	require.NoError(t, err)

	assert.Equal(t, uint(819), record.Location.MetroCode)
}

func TestAnonymousIP(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoIP2-Anonymous-IP-Test.mmdb")
	require.NoError(t, err)
	defer reader.Close()

	testAddr := netip.MustParseAddr("1.2.0.0")
	record, err := reader.AnonymousIP(testAddr)
	require.NoError(t, err)

	assert.True(t, record.IsAnonymous)

	assert.True(t, record.IsAnonymousVPN)
	assert.False(t, record.IsHostingProvider)
	assert.False(t, record.IsPublicProxy)
	assert.False(t, record.IsTorExitNode)
	assert.False(t, record.IsResidentialProxy)

	// Test Network and IPAddress fields
	assert.Equal(t, testAddr, record.IPAddress)
	assert.True(t, record.Network.IsValid())
	assert.True(t, record.Network.Contains(testAddr))
}

func TestASN(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoLite2-ASN-Test.mmdb")
	require.NoError(t, err)
	defer reader.Close()

	testAddr := netip.MustParseAddr("1.128.0.0")
	record, err := reader.ASN(testAddr)
	require.NoError(t, err)

	assert.Equal(t, uint(1221), record.AutonomousSystemNumber)

	assert.Equal(t, "Telstra Pty Ltd", record.AutonomousSystemOrganization)

	// Test Network and IPAddress fields
	assert.Equal(t, testAddr, record.IPAddress)
	assert.True(t, record.Network.IsValid())
	assert.True(t, record.Network.Contains(testAddr))
}

func TestConnectionType(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoIP2-Connection-Type-Test.mmdb")
	require.NoError(t, err)

	defer reader.Close()

	record, err := reader.ConnectionType(netip.MustParseAddr("1.0.1.0"))
	require.NoError(t, err)

	assert.Equal(t, "Cellular", record.ConnectionType)
}

func TestCountry(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoIP2-Country-Test.mmdb")
	require.NoError(t, err)

	defer reader.Close()

	record, err := reader.Country(netip.MustParseAddr("81.2.69.160"))
	require.NoError(t, err)

	assert.False(t, record.Country.IsInEuropeanUnion)
	assert.False(t, record.RegisteredCountry.IsInEuropeanUnion)
	assert.False(t, record.RepresentedCountry.IsInEuropeanUnion)
}

func TestDomain(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoIP2-Domain-Test.mmdb")
	require.NoError(t, err)
	defer reader.Close()

	record, err := reader.Domain(netip.MustParseAddr("1.2.0.0"))
	require.NoError(t, err)
	assert.Equal(t, "maxmind.com", record.Domain)
}

func TestEnterprise(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoIP2-Enterprise-Test.mmdb")
	require.NoError(t, err)

	defer reader.Close()

	testAddr1 := netip.MustParseAddr("74.209.24.0")
	record, err := reader.Enterprise(testAddr1)
	require.NoError(t, err)

	assert.Equal(t, uint8(11), record.City.Confidence)

	assert.Equal(t, uint(14671), record.Traits.AutonomousSystemNumber)
	assert.Equal(t, "FairPoint Communications", record.Traits.AutonomousSystemOrganization)
	assert.Equal(t, "Cable/DSL", record.Traits.ConnectionType)
	assert.Equal(t, "frpt.net", record.Traits.Domain)
	assert.InEpsilon(t, float64(0.34), record.Traits.StaticIPScore, 1e-10)

	testAddr2 := netip.MustParseAddr("149.101.100.0")
	record, err = reader.Enterprise(testAddr2)
	require.NoError(t, err)

	assert.Equal(t, uint(6167), record.Traits.AutonomousSystemNumber)

	assert.Equal(t, "CELLCO-PART", record.Traits.AutonomousSystemOrganization)
	assert.Equal(t, "Verizon Wireless", record.Traits.ISP)
	assert.Equal(t, "310", record.Traits.MobileCountryCode)
	assert.Equal(t, "004", record.Traits.MobileNetworkCode)

	// Test Network and IPAddress fields (for the second lookup)
	assert.Equal(t, testAddr2, record.Traits.IPAddress)
	assert.True(t, record.Traits.Network.IsValid())
	assert.True(t, record.Traits.Network.Contains(testAddr2))
}

func TestISP(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoIP2-ISP-Test.mmdb")
	require.NoError(t, err)
	defer reader.Close()

	record, err := reader.ISP(netip.MustParseAddr("149.101.100.0"))
	require.NoError(t, err)

	assert.Equal(t, uint(6167), record.AutonomousSystemNumber)

	assert.Equal(t, "CELLCO-PART", record.AutonomousSystemOrganization)
	assert.Equal(t, "Verizon Wireless", record.ISP)
	assert.Equal(t, "310", record.MobileCountryCode)
	assert.Equal(t, "004", record.MobileNetworkCode)
	assert.Equal(t, "Verizon Wireless", record.Organization)
}

// This ensures the compiler does not optimize away the function call.
var cityResult *City

func BenchmarkCity(b *testing.B) {
	db, err := Open("GeoLite2-City.mmdb")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	//nolint:gosec // this is just a benchmark
	r := rand.New(rand.NewSource(0))

	var city *City

	ip := make(net.IP, 4)
	for range b.N {
		randomIPv4Address(r, ip)
		addr, _ := netip.AddrFromSlice(ip)
		city, err = db.City(addr)
		if err != nil {
			b.Fatal(err)
		}
	}
	cityResult = city
}

// This ensures the compiler does not optimize away the function call.
var asnResult *ASN

func BenchmarkASN(b *testing.B) {
	db, err := Open("GeoLite2-ASN.mmdb")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	//nolint:gosec // this is just a benchmark
	r := rand.New(rand.NewSource(0))

	var asn *ASN

	ip := make(net.IP, 4)
	for range b.N {
		randomIPv4Address(r, ip)
		addr, _ := netip.AddrFromSlice(ip)
		asn, err = db.ASN(addr)
		if err != nil {
			b.Fatal(err)
		}
	}
	asnResult = asn
}

func randomIPv4Address(r *rand.Rand, ip net.IP) {
	num := r.Uint32()
	ip[0] = byte(num >> 24)
	ip[1] = byte(num >> 16)
	ip[2] = byte(num >> 8)
	ip[3] = byte(num)
}

func TestIsZero(t *testing.T) {
	reader, err := Open("test-data/test-data/GeoIP2-City-Test.mmdb")
	require.NoError(t, err)
	defer reader.Close()

	// Test with an IP that has data
	ipWithData := netip.MustParseAddr("81.2.69.160")
	record, err := reader.City(ipWithData)
	require.NoError(t, err)
	assert.False(t, record.IsZero(), "Record with data should not be zero")

	// Test with an IP that has no data (private IP)
	ipWithoutData := netip.MustParseAddr("192.168.1.1")
	emptyRecord, err := reader.City(ipWithoutData)
	require.NoError(t, err)
	assert.True(t, emptyRecord.IsZero(), "Record without data should be zero")

	// Test Names IsZero
	var emptyNames Names
	assert.True(t, emptyNames.IsZero(), "Empty Names should be zero")

	nonEmptyNames := Names{English: "Test"}
	assert.False(t, nonEmptyNames.IsZero(), "Names with data should not be zero")

	// Test other struct types
	var emptyASN ASN
	assert.True(t, emptyASN.IsZero(), "Empty ASN should be zero")

	nonEmptyASN := ASN{AutonomousSystemNumber: 123}
	assert.False(t, nonEmptyASN.IsZero(), "ASN with data should not be zero")
}

func TestAllStructsHaveIsZero(t *testing.T) {
	// Ensure all result structs have IsZero methods
	var city City
	var country Country
	var enterprise Enterprise
	var anonymousIP AnonymousIP
	var asn ASN
	var connectionType ConnectionType
	var domain Domain
	var isp ISP
	var names Names

	// These should all compile and return true for zero values
	assert.True(t, city.IsZero())
	assert.True(t, country.IsZero())
	assert.True(t, enterprise.IsZero())
	assert.True(t, anonymousIP.IsZero())
	assert.True(t, asn.IsZero())
	assert.True(t, connectionType.IsZero())
	assert.True(t, domain.IsZero())
	assert.True(t, isp.IsZero())
	assert.True(t, names.IsZero())
}
