// Package geoip2 provides an easy-to-use API for the MaxMind GeoIP2 and
// GeoLite2 databases; this package does not support GeoIP Legacy databases.
//
// The structs provided by this package match the internal structure of
// the data in the MaxMind databases.
//
// See github.com/oschwald/maxminddb-golang for more advanced used cases.
package geoip2

import (
	"fmt"
	"net/netip"
	"reflect"

	"github.com/oschwald/maxminddb-golang/v2"
)

// Names contains localized names for geographic entities.
// This replaces map[string]string to eliminate map allocation overhead.
// Uses reflection-based decoding which is 69% faster than custom unmarshalers.
type Names struct {
	German              string `maxminddb:"de"`    // German
	English             string `maxminddb:"en"`    // English
	Spanish             string `maxminddb:"es"`    // Spanish
	French              string `maxminddb:"fr"`    // French
	Japanese            string `maxminddb:"ja"`    // Japanese
	BrazilianPortuguese string `maxminddb:"pt-BR"` // Portuguese (Brazil)
	Russian             string `maxminddb:"ru"`    // Russian
	SimplifiedChinese   string `maxminddb:"zh-CN"` // Chinese (Simplified)
}

var zeroNames Names

// IsZero returns true if the Names struct has no localized names.
func (n Names) IsZero() bool {
	return n == zeroNames
}

// The Enterprise struct corresponds to the data in the GeoIP2 Enterprise
// database.
type Enterprise struct {
	Continent struct {
		Names     Names  `maxminddb:"names"`
		Code      string `maxminddb:"code"`
		GeoNameID uint   `maxminddb:"geoname_id"`
	} `maxminddb:"continent"`
	City struct {
		Names      Names `maxminddb:"names"`
		GeoNameID  uint  `maxminddb:"geoname_id"`
		Confidence uint8 `maxminddb:"confidence"`
	} `maxminddb:"city"`
	Postal struct {
		Code       string `maxminddb:"code"`
		Confidence uint8  `maxminddb:"confidence"`
	} `maxminddb:"postal"`
	Subdivisions []struct {
		Names      Names  `maxminddb:"names"`
		ISOCode    string `maxminddb:"iso_code"`
		GeoNameID  uint   `maxminddb:"geoname_id"`
		Confidence uint8  `maxminddb:"confidence"`
	} `maxminddb:"subdivisions"`
	RepresentedCountry struct {
		Names             Names  `maxminddb:"names"`
		ISOCode           string `maxminddb:"iso_code"`
		Type              string `maxminddb:"type"`
		GeoNameID         uint   `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	} `maxminddb:"represented_country"`
	Country struct {
		Names             Names  `maxminddb:"names"`
		ISOCode           string `maxminddb:"iso_code"`
		GeoNameID         uint   `maxminddb:"geoname_id"`
		Confidence        uint8  `maxminddb:"confidence"`
		IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	} `maxminddb:"country"`
	RegisteredCountry struct {
		Names             Names  `maxminddb:"names"`
		ISOCode           string `maxminddb:"iso_code"`
		GeoNameID         uint   `maxminddb:"geoname_id"`
		Confidence        uint8  `maxminddb:"confidence"`
		IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	} `maxminddb:"registered_country"`
	Traits struct {
		Network                      netip.Prefix  // The network prefix for this record
		IPAddress                    netip.Addr    // The IP address used during the lookup
		AutonomousSystemOrganization string        `maxminddb:"autonomous_system_organization"`
		ConnectionType               string        `maxminddb:"connection_type"`
		Domain                       string        `maxminddb:"domain"`
		ISP                          string        `maxminddb:"isp"`
		MobileCountryCode            string        `maxminddb:"mobile_country_code"`
		MobileNetworkCode            string        `maxminddb:"mobile_network_code"`
		Organization                 string        `maxminddb:"organization"`
		UserType                     string        `maxminddb:"user_type"`
		StaticIPScore                float64       `maxminddb:"static_ip_score"`
		AutonomousSystemNumber       uint          `maxminddb:"autonomous_system_number"`
		IsAnonymousProxy             bool          `maxminddb:"is_anonymous_proxy"`
		IsAnycast                    bool          `maxminddb:"is_anycast"`
		IsLegitimateProxy            bool          `maxminddb:"is_legitimate_proxy"`
		IsSatelliteProvider          bool          `maxminddb:"is_satellite_provider"`
	} `maxminddb:"traits"`
	Location struct {
		TimeZone       string  `maxminddb:"time_zone"`
		Latitude       float64 `maxminddb:"latitude"`
		Longitude      float64 `maxminddb:"longitude"`
		MetroCode      uint    `maxminddb:"metro_code"`
		AccuracyRadius uint16  `maxminddb:"accuracy_radius"`
	} `maxminddb:"location"`
}

var zeroEnterprise Enterprise

// IsZero returns true if no data was found for the IP in the Enterprise database.
func (e Enterprise) IsZero() bool {
	return reflect.DeepEqual(e, zeroEnterprise)
}

// The City struct corresponds to the data in the GeoIP2/GeoLite2 City
// databases.
type City struct {
	City struct {
		Names     Names `maxminddb:"names"`
		GeoNameID uint  `maxminddb:"geoname_id"`
	} `maxminddb:"city"`
	Postal struct {
		Code string `maxminddb:"code"`
	} `maxminddb:"postal"`
	Continent struct {
		Names     Names  `maxminddb:"names"`
		Code      string `maxminddb:"code"`
		GeoNameID uint   `maxminddb:"geoname_id"`
	} `maxminddb:"continent"`
	Subdivisions []struct {
		Names     Names  `maxminddb:"names"`
		ISOCode   string `maxminddb:"iso_code"`
		GeoNameID uint   `maxminddb:"geoname_id"`
	} `maxminddb:"subdivisions"`
	RepresentedCountry struct {
		Names             Names  `maxminddb:"names"`
		ISOCode           string `maxminddb:"iso_code"`
		Type              string `maxminddb:"type"`
		GeoNameID         uint   `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	} `maxminddb:"represented_country"`
	Country struct {
		Names             Names  `maxminddb:"names"`
		ISOCode           string `maxminddb:"iso_code"`
		GeoNameID         uint   `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	} `maxminddb:"country"`
	RegisteredCountry struct {
		Names             Names  `maxminddb:"names"`
		ISOCode           string `maxminddb:"iso_code"`
		GeoNameID         uint   `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	} `maxminddb:"registered_country"`
	Location struct {
		TimeZone       string  `maxminddb:"time_zone"`
		Latitude       float64 `maxminddb:"latitude"`
		Longitude      float64 `maxminddb:"longitude"`
		MetroCode      uint    `maxminddb:"metro_code"`
		AccuracyRadius uint16  `maxminddb:"accuracy_radius"`
	} `maxminddb:"location"`
	Traits struct {
		IPAddress           netip.Addr    // The IP address used during the lookup
		IsAnonymousProxy    bool          `maxminddb:"is_anonymous_proxy"`
		IsAnycast           bool          `maxminddb:"is_anycast"`
		IsSatelliteProvider bool          `maxminddb:"is_satellite_provider"`
		Network             netip.Prefix  // The network prefix for this record
	} `maxminddb:"traits"`
}

var zeroCity City

// IsZero returns true if no data was found for the IP in the City database.
func (c City) IsZero() bool {
	return reflect.DeepEqual(c, zeroCity)
}

// The Country struct corresponds to the data in the GeoIP2/GeoLite2
// Country databases.
type Country struct {
	Continent struct {
		Names     Names  `maxminddb:"names"`
		Code      string `maxminddb:"code"`
		GeoNameID uint   `maxminddb:"geoname_id"`
	} `maxminddb:"continent"`
	Country struct {
		Names             Names  `maxminddb:"names"`
		ISOCode           string `maxminddb:"iso_code"`
		GeoNameID         uint   `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	} `maxminddb:"country"`
	RegisteredCountry struct {
		Names             Names  `maxminddb:"names"`
		ISOCode           string `maxminddb:"iso_code"`
		GeoNameID         uint   `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	} `maxminddb:"registered_country"`
	RepresentedCountry struct {
		Names             Names  `maxminddb:"names"`
		ISOCode           string `maxminddb:"iso_code"`
		Type              string `maxminddb:"type"`
		GeoNameID         uint   `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	} `maxminddb:"represented_country"`
	Traits struct {
		IPAddress           netip.Addr    // The IP address used during the lookup
		IsAnonymousProxy    bool          `maxminddb:"is_anonymous_proxy"`
		IsAnycast           bool          `maxminddb:"is_anycast"`
		IsSatelliteProvider bool          `maxminddb:"is_satellite_provider"`
		Network             netip.Prefix  // The network prefix for this record
	} `maxminddb:"traits"`
}

var zeroCountry Country

// IsZero returns true if no data was found for the IP in the Country database.
func (c Country) IsZero() bool {
	return c == zeroCountry
}

// The AnonymousIP struct corresponds to the data in the GeoIP2
// Anonymous IP database.
type AnonymousIP struct {
	IPAddress          netip.Addr    // The IP address used during the lookup
	IsAnonymous        bool          `maxminddb:"is_anonymous"`
	IsAnonymousVPN     bool          `maxminddb:"is_anonymous_vpn"`
	IsHostingProvider  bool          `maxminddb:"is_hosting_provider"`
	IsPublicProxy      bool          `maxminddb:"is_public_proxy"`
	IsResidentialProxy bool          `maxminddb:"is_residential_proxy"`
	IsTorExitNode      bool          `maxminddb:"is_tor_exit_node"`
	Network            netip.Prefix  // The network prefix for this record
}

var zeroAnonymousIP AnonymousIP

// IsZero returns true if no data was found for the IP in the AnonymousIP database.
func (a AnonymousIP) IsZero() bool {
	return a == zeroAnonymousIP
}

// The ASN struct corresponds to the data in the GeoLite2 ASN database.
type ASN struct {
	AutonomousSystemNumber       uint          `maxminddb:"autonomous_system_number"`
	AutonomousSystemOrganization string        `maxminddb:"autonomous_system_organization"`
	IPAddress                    netip.Addr    // The IP address used during the lookup
	Network                      netip.Prefix  // The network prefix for this record
}

var zeroASN ASN

// IsZero returns true if no data was found for the IP in the ASN database.
func (a ASN) IsZero() bool {
	return a == zeroASN
}

// The ConnectionType struct corresponds to the data in the GeoIP2
// Connection-Type database.
type ConnectionType struct {
	ConnectionType string        `maxminddb:"connection_type"`
	IPAddress      netip.Addr    // The IP address used during the lookup
	Network        netip.Prefix  // The network prefix for this record
}

var zeroConnectionType ConnectionType

// IsZero returns true if no data was found for the IP in the ConnectionType database.
func (c ConnectionType) IsZero() bool {
	return c == zeroConnectionType
}

// The Domain struct corresponds to the data in the GeoIP2 Domain database.
type Domain struct {
	Domain    string        `maxminddb:"domain"`
	IPAddress netip.Addr    // The IP address used during the lookup
	Network   netip.Prefix  // The network prefix for this record
}

var zeroDomain Domain

// IsZero returns true if no data was found for the IP in the Domain database.
func (d Domain) IsZero() bool {
	return d == zeroDomain
}

// The ISP struct corresponds to the data in the GeoIP2 ISP database.
type ISP struct {
	Network                      netip.Prefix  // The network prefix for this record
	IPAddress                    netip.Addr    // The IP address used during the lookup
	AutonomousSystemOrganization string        `maxminddb:"autonomous_system_organization"`
	ISP                          string        `maxminddb:"isp"`
	MobileCountryCode            string        `maxminddb:"mobile_country_code"`
	MobileNetworkCode            string        `maxminddb:"mobile_network_code"`
	Organization                 string        `maxminddb:"organization"`
	AutonomousSystemNumber       uint          `maxminddb:"autonomous_system_number"`
}

var zeroISP ISP

// IsZero returns true if no data was found for the IP in the ISP database.
func (i ISP) IsZero() bool {
	return i == zeroISP
}

type databaseType int

const (
	isAnonymousIP = 1 << iota
	isASN
	isCity
	isConnectionType
	isCountry
	isDomain
	isEnterprise
	isISP
)

// Reader holds the maxminddb.Reader struct. It can be created using the
// Open and FromBytes functions.
type Reader struct {
	mmdbReader   *maxminddb.Reader
	databaseType databaseType
}

// InvalidMethodError is returned when a lookup method is called on a
// database that it does not support. For instance, calling the ISP method
// on a City database.
type InvalidMethodError struct {
	Method       string
	DatabaseType string
}

func (e InvalidMethodError) Error() string {
	return fmt.Sprintf(`geoip2: the %s method does not support the %s database`,
		e.Method, e.DatabaseType)
}

// UnknownDatabaseTypeError is returned when an unknown database type is
// opened.
type UnknownDatabaseTypeError struct {
	DatabaseType string
}

func (e UnknownDatabaseTypeError) Error() string {
	return fmt.Sprintf(`geoip2: reader does not support the %q database type`,
		e.DatabaseType)
}

// Open takes a string path to a file and returns a Reader struct or an error.
// The database file is opened using a memory map. Use the Close method on the
// Reader object to return the resources to the system.
func Open(file string) (*Reader, error) {
	reader, err := maxminddb.Open(file)
	if err != nil {
		return nil, err
	}
	dbType, err := getDBType(reader)
	return &Reader{reader, dbType}, err
}

// FromBytes takes a byte slice corresponding to a GeoIP2/GeoLite2 database
// file and returns a Reader struct or an error. Note that the byte slice is
// used directly; any modification of it after opening the database will result
// in errors while reading from the database.
func FromBytes(bytes []byte) (*Reader, error) {
	reader, err := maxminddb.FromBytes(bytes)
	if err != nil {
		return nil, err
	}
	dbType, err := getDBType(reader)
	return &Reader{reader, dbType}, err
}

func getDBType(reader *maxminddb.Reader) (databaseType, error) {
	switch reader.Metadata.DatabaseType {
	case "GeoIP2-Anonymous-IP":
		return isAnonymousIP, nil
	case "DBIP-ASN-Lite (compat=GeoLite2-ASN)",
		"GeoLite2-ASN":
		return isASN, nil
	// We allow City lookups on Country for back compat
	case "DBIP-City-Lite",
		"DBIP-Country-Lite",
		"DBIP-Country",
		"DBIP-Location (compat=City)",
		"GeoLite2-City",
		"GeoIP2-City",
		"GeoIP2-City-Africa",
		"GeoIP2-City-Asia-Pacific",
		"GeoIP2-City-Europe",
		"GeoIP2-City-North-America",
		"GeoIP2-City-South-America",
		"GeoIP2-Precision-City",
		"GeoLite2-Country",
		"GeoIP2-Country":
		return isCity | isCountry, nil
	case "GeoIP2-Connection-Type":
		return isConnectionType, nil
	case "GeoIP2-Domain":
		return isDomain, nil
	case "DBIP-ISP (compat=Enterprise)",
		"DBIP-Location-ISP (compat=Enterprise)",
		"GeoIP2-Enterprise":
		return isEnterprise | isCity | isCountry, nil
	case "GeoIP2-ISP", "GeoIP2-Precision-ISP":
		return isISP | isASN, nil
	default:
		return 0, UnknownDatabaseTypeError{reader.Metadata.DatabaseType}
	}
}

// Enterprise takes an IP address as a netip.Addr and returns an Enterprise
// struct and/or an error. This is intended to be used with the GeoIP2
// Enterprise database.
func (r *Reader) Enterprise(ipAddress netip.Addr) (*Enterprise, error) {
	if isEnterprise&r.databaseType == 0 {
		return nil, InvalidMethodError{"Enterprise", r.Metadata().DatabaseType}
	}
	result := r.mmdbReader.Lookup(ipAddress)
	var enterprise Enterprise
	err := result.Decode(&enterprise)
	if err != nil {
		return &enterprise, err
	}
	if result.Found() {
		enterprise.Traits.IPAddress = ipAddress
		enterprise.Traits.Network = result.Prefix()
	}
	return &enterprise, nil
}

// City takes an IP address as a netip.Addr and returns a City struct
// and/or an error. Although this can be used with other databases, this
// method generally should be used with the GeoIP2 or GeoLite2 City databases.
func (r *Reader) City(ipAddress netip.Addr) (*City, error) {
	if isCity&r.databaseType == 0 {
		return nil, InvalidMethodError{"City", r.Metadata().DatabaseType}
	}
	result := r.mmdbReader.Lookup(ipAddress)
	var city City
	err := result.Decode(&city)
	if err != nil {
		return &city, err
	}
	if result.Found() {
		city.Traits.IPAddress = ipAddress
		city.Traits.Network = result.Prefix()
	}
	return &city, nil
}

// Country takes an IP address as a netip.Addr and returns a Country struct
// and/or an error. Although this can be used with other databases, this
// method generally should be used with the GeoIP2 or GeoLite2 Country
// databases.
func (r *Reader) Country(ipAddress netip.Addr) (*Country, error) {
	if isCountry&r.databaseType == 0 {
		return nil, InvalidMethodError{"Country", r.Metadata().DatabaseType}
	}
	result := r.mmdbReader.Lookup(ipAddress)
	var country Country
	err := result.Decode(&country)
	if err != nil {
		return &country, err
	}
	if result.Found() {
		country.Traits.IPAddress = ipAddress
		country.Traits.Network = result.Prefix()
	}
	return &country, nil
}

// AnonymousIP takes an IP address as a netip.Addr and returns a
// AnonymousIP struct and/or an error.
func (r *Reader) AnonymousIP(ipAddress netip.Addr) (*AnonymousIP, error) {
	if isAnonymousIP&r.databaseType == 0 {
		return nil, InvalidMethodError{"AnonymousIP", r.Metadata().DatabaseType}
	}
	result := r.mmdbReader.Lookup(ipAddress)
	var anonIP AnonymousIP
	err := result.Decode(&anonIP)
	if err != nil {
		return &anonIP, err
	}
	if result.Found() {
		anonIP.IPAddress = ipAddress
		anonIP.Network = result.Prefix()
	}
	return &anonIP, nil
}

// ASN takes an IP address as a netip.Addr and returns a ASN struct and/or
// an error.
func (r *Reader) ASN(ipAddress netip.Addr) (*ASN, error) {
	if isASN&r.databaseType == 0 {
		return nil, InvalidMethodError{"ASN", r.Metadata().DatabaseType}
	}
	result := r.mmdbReader.Lookup(ipAddress)
	var val ASN
	err := result.Decode(&val)
	if err != nil {
		return &val, err
	}
	if result.Found() {
		val.IPAddress = ipAddress
		val.Network = result.Prefix()
	}
	return &val, nil
}

// ConnectionType takes an IP address as a netip.Addr and returns a
// ConnectionType struct and/or an error.
func (r *Reader) ConnectionType(ipAddress netip.Addr) (*ConnectionType, error) {
	if isConnectionType&r.databaseType == 0 {
		return nil, InvalidMethodError{"ConnectionType", r.Metadata().DatabaseType}
	}
	result := r.mmdbReader.Lookup(ipAddress)
	var val ConnectionType
	err := result.Decode(&val)
	if err != nil {
		return &val, err
	}
	if result.Found() {
		val.IPAddress = ipAddress
		val.Network = result.Prefix()
	}
	return &val, nil
}

// Domain takes an IP address as a netip.Addr and returns a
// Domain struct and/or an error.
func (r *Reader) Domain(ipAddress netip.Addr) (*Domain, error) {
	if isDomain&r.databaseType == 0 {
		return nil, InvalidMethodError{"Domain", r.Metadata().DatabaseType}
	}
	result := r.mmdbReader.Lookup(ipAddress)
	var val Domain
	err := result.Decode(&val)
	if err != nil {
		return &val, err
	}
	if result.Found() {
		val.IPAddress = ipAddress
		val.Network = result.Prefix()
	}
	return &val, nil
}

// ISP takes an IP address as a netip.Addr and returns a ISP struct and/or
// an error.
func (r *Reader) ISP(ipAddress netip.Addr) (*ISP, error) {
	if isISP&r.databaseType == 0 {
		return nil, InvalidMethodError{"ISP", r.Metadata().DatabaseType}
	}
	result := r.mmdbReader.Lookup(ipAddress)
	var val ISP
	err := result.Decode(&val)
	if err != nil {
		return &val, err
	}
	if result.Found() {
		val.IPAddress = ipAddress
		val.Network = result.Prefix()
	}
	return &val, nil
}

// Metadata takes no arguments and returns a struct containing metadata about
// the MaxMind database in use by the Reader.
func (r *Reader) Metadata() maxminddb.Metadata {
	return r.mmdbReader.Metadata
}

// Close unmaps the database file from virtual memory and returns the
// resources to the system.
func (r *Reader) Close() error {
	return r.mmdbReader.Close()
}
