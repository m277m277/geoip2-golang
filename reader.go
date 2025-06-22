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
	// German localized name
	German string `json:"de"    maxminddb:"de"`
	// English localized name
	English string `json:"en"    maxminddb:"en"`
	// Spanish localized name
	Spanish string `json:"es"    maxminddb:"es"`
	// French localized name
	French string `json:"fr"    maxminddb:"fr"`
	// Japanese localized name
	Japanese string `json:"ja"    maxminddb:"ja"`
	// BrazilianPortuguese localized name (pt-BR)
	BrazilianPortuguese string `json:"pt-BR" maxminddb:"pt-BR"` //nolint:tagliatelle // pt-BR matches MMDB format
	// Russian localized name
	Russian string `json:"ru"    maxminddb:"ru"`
	// SimplifiedChinese localized name (zh-CN)
	SimplifiedChinese string `json:"zh-CN" maxminddb:"zh-CN"` //nolint:tagliatelle // zh-CN matches MMDB format
}

var zeroNames Names

// IsZero returns true if the Names struct has no localized names.
func (n Names) IsZero() bool {
	return n == zeroNames
}

// The Enterprise struct corresponds to the data in the GeoIP2 Enterprise
// database.
type Enterprise struct {
	// Continent contains data for the continent record associated with the IP
	// address.
	Continent struct {
		// Names contains localized names for the continent
		Names Names `json:"names" maxminddb:"names"`
		// Code is a two character continent code like "NA" (North America) or
		// "OC" (Oceania)
		Code string `json:"code" maxminddb:"code"`
		// GeoNameID is the GeoName ID for the continent
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
	} `json:"continent"           maxminddb:"continent"`
	// City contains data for the city record associated with the IP address.
	City struct {
		// Names contains localized names for the city
		Names Names `json:"names" maxminddb:"names"`
		// GeoNameID is the GeoName ID for the city
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the city is correct
		Confidence uint8 `json:"confidence" maxminddb:"confidence"`
	} `json:"city"                maxminddb:"city"`
	// Postal contains data for the postal record associated with the IP address.
	Postal struct {
		// Code is the postal code of the location. Postal codes are not
		// available for all countries.
		// In some countries, this will only contain part of the postal code.
		Code string `json:"code" maxminddb:"code"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the postal code is correct
		Confidence uint8 `json:"confidence" maxminddb:"confidence"`
	} `json:"postal"              maxminddb:"postal"`
	// Subdivisions contains data for the subdivisions associated with the IP
	// address.
	// The subdivisions array is ordered from largest to smallest. For instance,
	// the response
	// for Oxford in the United Kingdom would have England as the first element
	// and
	// Oxfordshire as the second element.
	Subdivisions []struct {
		// Names contains localized names for the subdivision
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is a string up to three characters long containing the
		// subdivision portion
		// of the ISO 3166-2 code. See https://en.wikipedia.org/wiki/ISO_3166-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// GeoNameID is the GeoName ID for the subdivision
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the subdivision is correct
		Confidence uint8 `json:"confidence" maxminddb:"confidence"`
	} `json:"subdivisions"        maxminddb:"subdivisions"`
	// RepresentedCountry contains data for the represented country associated
	// with the IP address.
	// The represented country is the country represented by something like a
	// military base or embassy.
	RepresentedCountry struct {
		// Names contains localized names for the represented country
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the represented
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// Type is a string indicating the type of entity that is representing
		// the country.
		// Currently this is only "military" but may expand in the future.
		Type string `json:"type" maxminddb:"type"`
		// GeoNameID is the GeoName ID for the represented country
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the represented country is a member
		// state of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union" maxminddb:"is_in_european_union"`
	} `json:"represented_country" maxminddb:"represented_country"`
	// Country contains data for the country record associated with the IP
	// address.
	// This record represents the country where MaxMind believes the IP is
	// located.
	Country struct {
		// Names contains localized names for the country
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// GeoNameID is the GeoName ID for the country
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the country is correct
		Confidence uint8 `json:"confidence" maxminddb:"confidence"`
		// IsInEuropeanUnion is true if the country is a member state of the
		// European Union
		IsInEuropeanUnion bool `json:"is_in_european_union" maxminddb:"is_in_european_union"`
	} `json:"country"             maxminddb:"country"`
	// RegisteredCountry contains data for the registered country associated
	// with the IP address.
	// This record represents the country where the ISP has registered the IP
	// block and may differ from the user's country.
	RegisteredCountry struct {
		// Names contains localized names for the registered country
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the registered
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// GeoNameID is the GeoName ID for the registered country
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the registered country is correct
		Confidence uint8 `json:"confidence" maxminddb:"confidence"`
		// IsInEuropeanUnion is true if the registered country is a member state
		// of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union" maxminddb:"is_in_european_union"`
	} `json:"registered_country"  maxminddb:"registered_country"`
	// Traits contains various traits associated with the IP address
	Traits struct {
		// Network is the network prefix for this record. This is the largest
		// network where all
		// of the fields besides IPAddress have the same value.
		Network netip.Prefix `json:"network"`
		// IPAddress is the IP address used during the lookup
		IPAddress netip.Addr `json:"ip_address"`
		// AutonomousSystemOrganization is the organization associated with the
		// registered ASN for the IP address
		AutonomousSystemOrganization string `json:"autonomous_system_organization" maxminddb:"autonomous_system_organization"` //nolint:lll // long struct tag //nolint:lll // long struct tag
		// ConnectionType indicates the connection type. May be Dialup,
		// Cable/DSL, Corporate, Cellular, or Satellite
		ConnectionType string `json:"connection_type" maxminddb:"connection_type"`
		// Domain is the second level domain associated with the IP address
		// (e.g., "example.com")
		Domain string `json:"domain" maxminddb:"domain"`
		// ISP is the name of the ISP associated with the IP address
		ISP string `json:"isp" maxminddb:"isp"`
		// MobileCountryCode is the mobile country code (MCC) associated with
		// the IP address and ISP.
		// See https://en.wikipedia.org/wiki/Mobile_country_code
		MobileCountryCode string `json:"mobile_country_code" maxminddb:"mobile_country_code"`
		// MobileNetworkCode is the mobile network code (MNC) associated with
		// the IP address and ISP.
		// See https://en.wikipedia.org/wiki/Mobile_network_code
		MobileNetworkCode string `json:"mobile_network_code" maxminddb:"mobile_network_code"`
		// Organization is the name of the organization associated with the IP
		// address
		Organization string `json:"organization" maxminddb:"organization"`
		// UserType indicates the user type associated with the IP address
		// (business, cafe, cellular, college, etc.)
		UserType string `json:"user_type" maxminddb:"user_type"`
		// StaticIPScore is an indicator of how static or dynamic an IP address is, ranging from 0 to 99.99
		StaticIPScore float64 `json:"static_ip_score" maxminddb:"static_ip_score"`
		// AutonomousSystemNumber is the autonomous system number associated with the IP address
		AutonomousSystemNumber uint `json:"autonomous_system_number" maxminddb:"autonomous_system_number"`
		// IsAnonymousProxy is true if the IP is an anonymous proxy.
		//
		// Deprecated: Use the GeoIP2 Anonymous IP database instead.
		IsAnonymousProxy bool `json:"is_anonymous_proxy" maxminddb:"is_anonymous_proxy"`
		// IsAnycast is true if the IP address belongs to an anycast network.
		// See https://en.wikipedia.org/wiki/Anycast
		IsAnycast bool `json:"is_anycast" maxminddb:"is_anycast"`
		// IsLegitimateProxy is true if MaxMind believes this IP address to be a legitimate proxy,
		// such as an internal VPN used by a corporation
		IsLegitimateProxy bool `json:"is_legitimate_proxy" maxminddb:"is_legitimate_proxy"`
		// IsSatelliteProvider is true if the IP address is from a satellite
		// provider that provides service to multiple countries.
		//
		// Deprecated: Due to increased coverage by mobile carriers, very few
		// satellite providers now serve multiple countries.
		IsSatelliteProvider bool `json:"is_satellite_provider" maxminddb:"is_satellite_provider"`
	} `json:"traits"              maxminddb:"traits"`
	// Location contains data for the location record associated with the IP address
	Location struct {
		// TimeZone is the time zone associated with location, as specified by
		// the IANA Time Zone Database (e.g., "America/New_York")
		TimeZone string `json:"time_zone" maxminddb:"time_zone"`
		// Latitude is the approximate latitude of the location associated with the IP address.
		// This value is not precise and should not be used to identify a particular address or household.
		Latitude float64 `json:"latitude" maxminddb:"latitude"`
		// Longitude is the approximate longitude of the location associated with the IP address.
		// This value is not precise and should not be used to identify a particular address or household.
		Longitude float64 `json:"longitude" maxminddb:"longitude"`
		// MetroCode is a metro code for targeting advertisements.
		//
		// Deprecated: Metro codes are no longer maintained and should not be used.
		MetroCode uint `json:"metro_code" maxminddb:"metro_code"`
		// AccuracyRadius is the approximate accuracy radius in kilometers around the latitude and longitude.
		// This is the radius where we have a 67% confidence that the device
		// using the IP address resides within the circle.
		AccuracyRadius uint16 `json:"accuracy_radius" maxminddb:"accuracy_radius"`
	} `json:"location"            maxminddb:"location"`
}

var zeroEnterprise Enterprise

// IsZero returns true if no data was found for the IP in the Enterprise database.
func (e Enterprise) IsZero() bool {
	return reflect.DeepEqual(e, zeroEnterprise)
}

// The City struct corresponds to the data in the GeoIP2/GeoLite2 City
// databases.
type City struct {
	// City contains data for the city record associated with the IP address
	City struct {
		// Names contains localized names for the city
		Names Names `json:"names" maxminddb:"names"`
		// GeoNameID is the GeoName ID for the city
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
	} `json:"city"                maxminddb:"city"`
	// Postal contains data for the postal record associated with the IP address
	Postal struct {
		// Code is the postal code of the location. Postal codes are not
		// available for all countries.
		// In some countries, this will only contain part of the postal code.
		Code string `json:"code" maxminddb:"code"`
	} `json:"postal"              maxminddb:"postal"`
	// Continent contains data for the continent record associated with the IP address
	Continent struct {
		// Names contains localized names for the continent
		Names Names `json:"names" maxminddb:"names"`
		// Code is a two character continent code like "NA" (North America) or
		// "OC" (Oceania)
		Code string `json:"code" maxminddb:"code"`
		// GeoNameID is the GeoName ID for the continent
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
	} `json:"continent"           maxminddb:"continent"`
	// Subdivisions contains data for the subdivisions associated with the IP
	// address.
	// The subdivisions array is ordered from largest to smallest. For instance,
	// the response
	// for Oxford in the United Kingdom would have England as the first element
	// and
	// Oxfordshire as the second element.
	Subdivisions []struct {
		// Names contains localized names for the subdivision
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is a string up to three characters long containing the
		// subdivision portion
		// of the ISO 3166-2 code. See https://en.wikipedia.org/wiki/ISO_3166-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// GeoNameID is the GeoName ID for the subdivision
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
	} `json:"subdivisions"        maxminddb:"subdivisions"`
	// RepresentedCountry contains data for the represented country associated
	// with the IP address.
	// The represented country is the country represented by something like a
	// military base or embassy.
	RepresentedCountry struct {
		// Names contains localized names for the represented country
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the represented
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// Type is a string indicating the type of entity that is representing
		// the country.
		// Currently this is only "military" but may expand in the future.
		Type string `json:"type" maxminddb:"type"`
		// GeoNameID is the GeoName ID for the represented country
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the represented country is a member
		// state of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union" maxminddb:"is_in_european_union"`
	} `json:"represented_country" maxminddb:"represented_country"`
	// Country contains data for the country record associated with the IP
	// address.
	// This record represents the country where MaxMind believes the IP is
	// located.
	Country struct {
		// Names contains localized names for the country
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// GeoNameID is the GeoName ID for the country
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the country is a member state of the
		// European Union
		IsInEuropeanUnion bool `json:"is_in_european_union" maxminddb:"is_in_european_union"`
	} `json:"country"             maxminddb:"country"`
	// RegisteredCountry contains data for the registered country associated
	// with the IP address.
	// This record represents the country where the ISP has registered the IP
	// block and may differ from the user's country.
	RegisteredCountry struct {
		// Names contains localized names for the registered country
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the registered
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// GeoNameID is the GeoName ID for the registered country
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the registered country is a member state
		// of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union" maxminddb:"is_in_european_union"`
	} `json:"registered_country"  maxminddb:"registered_country"`
	// Location contains data for the location record associated with the IP address
	Location struct {
		// TimeZone is the time zone associated with location, as specified by
		// the IANA Time Zone Database (e.g., "America/New_York")
		TimeZone string `json:"time_zone" maxminddb:"time_zone"`
		// Latitude is the approximate latitude of the location associated with the IP address.
		// This value is not precise and should not be used to identify a particular address or household.
		Latitude float64 `json:"latitude" maxminddb:"latitude"`
		// Longitude is the approximate longitude of the location associated with the IP address.
		// This value is not precise and should not be used to identify a particular address or household.
		Longitude float64 `json:"longitude" maxminddb:"longitude"`
		// MetroCode is a metro code for targeting advertisements.
		//
		// Deprecated: Metro codes are no longer maintained and should not be used.
		MetroCode uint `json:"metro_code" maxminddb:"metro_code"`
		// AccuracyRadius is the approximate accuracy radius in kilometers around the latitude and longitude.
		// This is the radius where we have a 67% confidence that the device
		// using the IP address resides within the circle.
		AccuracyRadius uint16 `json:"accuracy_radius" maxminddb:"accuracy_radius"`
	} `json:"location"            maxminddb:"location"`
	// Traits contains various traits associated with the IP address
	Traits struct {
		// IPAddress is the IP address used during the lookup
		IPAddress netip.Addr `json:"ip_address"`
		// IsAnonymousProxy is true if the IP is an anonymous proxy.
		//
		// Deprecated: Use the GeoIP2 Anonymous IP database instead.
		IsAnonymousProxy bool `json:"is_anonymous_proxy" maxminddb:"is_anonymous_proxy"`
		// IsAnycast is true if the IP address belongs to an anycast network.
		// See https://en.wikipedia.org/wiki/Anycast
		IsAnycast bool `json:"is_anycast" maxminddb:"is_anycast"`
		// IsSatelliteProvider is true if the IP address is from a satellite
		// provider that provides service to multiple countries.
		//
		// Deprecated: Due to increased coverage by mobile carriers, very few
		// satellite providers now serve multiple countries.
		IsSatelliteProvider bool `json:"is_satellite_provider" maxminddb:"is_satellite_provider"`
		// Network is the network prefix for this record. This is the largest
		// network where all
		// of the fields besides IPAddress have the same value.
		Network netip.Prefix `json:"network"`
	} `json:"traits"              maxminddb:"traits"`
}

var zeroCity City

// IsZero returns true if no data was found for the IP in the City database.
func (c City) IsZero() bool {
	return reflect.DeepEqual(c, zeroCity)
}

// The Country struct corresponds to the data in the GeoIP2/GeoLite2
// Country databases.
type Country struct {
	// Continent contains data for the continent record associated with the IP address
	Continent struct {
		// Names contains localized names for the continent
		Names Names `json:"names" maxminddb:"names"`
		// Code is a two character continent code like "NA" (North America) or
		// "OC" (Oceania)
		Code string `json:"code" maxminddb:"code"`
		// GeoNameID is the GeoName ID for the continent
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
	} `json:"continent"           maxminddb:"continent"`
	// Country contains data for the country record associated with the IP
	// address.
	// This record represents the country where MaxMind believes the IP is
	// located.
	Country struct {
		// Names contains localized names for the country
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// GeoNameID is the GeoName ID for the country
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the country is a member state of the
		// European Union
		IsInEuropeanUnion bool `json:"is_in_european_union" maxminddb:"is_in_european_union"`
	} `json:"country"             maxminddb:"country"`
	// RegisteredCountry contains data for the registered country associated
	// with the IP address.
	// This record represents the country where the ISP has registered the IP
	// block and may differ from the user's country.
	RegisteredCountry struct {
		// Names contains localized names for the registered country
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the registered
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// GeoNameID is the GeoName ID for the registered country
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the registered country is a member state
		// of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union" maxminddb:"is_in_european_union"`
	} `json:"registered_country"  maxminddb:"registered_country"`
	// RepresentedCountry contains data for the represented country associated
	// with the IP address.
	// The represented country is the country represented by something like a
	// military base or embassy.
	RepresentedCountry struct {
		// Names contains localized names for the represented country
		Names Names `json:"names" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the represented
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code" maxminddb:"iso_code"`
		// Type is a string indicating the type of entity that is representing
		// the country.
		// Currently this is only "military" but may expand in the future.
		Type string `json:"type" maxminddb:"type"`
		// GeoNameID is the GeoName ID for the represented country
		GeoNameID uint `json:"geoname_id" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the represented country is a member
		// state of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union" maxminddb:"is_in_european_union"`
	} `json:"represented_country" maxminddb:"represented_country"`
	// Traits contains various traits associated with the IP address
	Traits struct {
		// IPAddress is the IP address used during the lookup
		IPAddress netip.Addr `json:"ip_address"`
		// IsAnonymousProxy is true if the IP is an anonymous proxy.
		//
		// Deprecated: Use the GeoIP2 Anonymous IP database instead.
		IsAnonymousProxy bool `json:"is_anonymous_proxy" maxminddb:"is_anonymous_proxy"`
		// IsAnycast is true if the IP address belongs to an anycast network.
		// See https://en.wikipedia.org/wiki/Anycast
		IsAnycast bool `json:"is_anycast" maxminddb:"is_anycast"`
		// IsSatelliteProvider is true if the IP address is from a satellite
		// provider that provides service to multiple countries.
		//
		// Deprecated: Due to increased coverage by mobile carriers, very few
		// satellite providers now serve multiple countries.
		IsSatelliteProvider bool `json:"is_satellite_provider" maxminddb:"is_satellite_provider"`
		// Network is the network prefix for this record. This is the largest
		// network where all
		// of the fields besides IPAddress have the same value.
		Network netip.Prefix `json:"network"`
	} `json:"traits"              maxminddb:"traits"`
}

var zeroCountry Country

// IsZero returns true if no data was found for the IP in the Country database.
func (c Country) IsZero() bool {
	return c == zeroCountry
}

// The AnonymousIP struct corresponds to the data in the GeoIP2
// Anonymous IP database.
type AnonymousIP struct {
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address"`
	// IsAnonymous is true if the IP address belongs to any sort of anonymous network
	IsAnonymous bool `json:"is_anonymous"         maxminddb:"is_anonymous"`
	// IsAnonymousVPN is true if the IP address is registered to an anonymous VPN provider.
	// If a VPN provider does not register subnets under names associated with them, we will
	// likely only flag their IP ranges using the IsHostingProvider attribute.
	IsAnonymousVPN bool `json:"is_anonymous_vpn"     maxminddb:"is_anonymous_vpn"`
	// IsHostingProvider is true if the IP address belongs to a hosting or VPN provider
	IsHostingProvider bool `json:"is_hosting_provider"  maxminddb:"is_hosting_provider"`
	// IsPublicProxy is true if the IP address belongs to a public proxy
	IsPublicProxy bool `json:"is_public_proxy"      maxminddb:"is_public_proxy"`
	// IsResidentialProxy is true if the IP address is on a suspected anonymizing network
	// and belongs to a residential ISP
	IsResidentialProxy bool `json:"is_residential_proxy" maxminddb:"is_residential_proxy"`
	// IsTorExitNode is true if the IP address is a Tor exit node
	IsTorExitNode bool `json:"is_tor_exit_node"     maxminddb:"is_tor_exit_node"`
	// Network is the network prefix for this record. This is the largest network where all
	// of the fields besides IPAddress have the same value.
	Network netip.Prefix `json:"network"`
}

var zeroAnonymousIP AnonymousIP

// IsZero returns true if no data was found for the IP in the AnonymousIP database.
func (a AnonymousIP) IsZero() bool {
	return a == zeroAnonymousIP
}

// The ASN struct corresponds to the data in the GeoLite2 ASN database.
type ASN struct {
	// AutonomousSystemNumber is the autonomous system number associated with the IP address
	AutonomousSystemNumber uint `json:"autonomous_system_number"       maxminddb:"autonomous_system_number"`
	// AutonomousSystemOrganization is the organization associated with the registered ASN for the IP address
	AutonomousSystemOrganization string `json:"autonomous_system_organization" maxminddb:"autonomous_system_organization"` //nolint:lll // long struct tag
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address"`
	// Network is the network prefix for this record. This is the largest network where all
	// of the fields besides IPAddress have the same value.
	Network netip.Prefix `json:"network"`
}

var zeroASN ASN

// IsZero returns true if no data was found for the IP in the ASN database.
func (a ASN) IsZero() bool {
	return a == zeroASN
}

// The ConnectionType struct corresponds to the data in the GeoIP2
// Connection-Type database.
type ConnectionType struct {
	// ConnectionType indicates the connection type. May be Dialup, Cable/DSL, Corporate, Cellular, or Satellite.
	// Additional values may be added in the future.
	ConnectionType string `json:"connection_type" maxminddb:"connection_type"`
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address"`
	// Network is the network prefix for this record. This is the largest network where all
	// of the fields besides IPAddress have the same value.
	Network netip.Prefix `json:"network"`
}

var zeroConnectionType ConnectionType

// IsZero returns true if no data was found for the IP in the ConnectionType database.
func (c ConnectionType) IsZero() bool {
	return c == zeroConnectionType
}

// The Domain struct corresponds to the data in the GeoIP2 Domain database.
type Domain struct {
	// Domain is the second level domain associated with the IP address (e.g., "example.com")
	Domain string `json:"domain"     maxminddb:"domain"`
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address"`
	// Network is the network prefix for this record. This is the largest network where all
	// of the fields besides IPAddress have the same value.
	Network netip.Prefix `json:"network"`
}

var zeroDomain Domain

// IsZero returns true if no data was found for the IP in the Domain database.
func (d Domain) IsZero() bool {
	return d == zeroDomain
}

// The ISP struct corresponds to the data in the GeoIP2 ISP database.
type ISP struct {
	// Network is the network prefix for this record. This is the largest network where all
	// of the fields besides IPAddress have the same value.
	Network netip.Prefix `json:"network"`
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address"`
	// AutonomousSystemOrganization is the organization associated with the registered ASN for the IP address
	AutonomousSystemOrganization string `json:"autonomous_system_organization" maxminddb:"autonomous_system_organization"` //nolint:lll // long struct tag
	// ISP is the name of the ISP associated with the IP address
	ISP string `json:"isp"                            maxminddb:"isp"`
	// MobileCountryCode is the mobile country code (MCC) associated with the IP address and ISP.
	// See https://en.wikipedia.org/wiki/Mobile_country_code
	MobileCountryCode string `json:"mobile_country_code"            maxminddb:"mobile_country_code"`
	// MobileNetworkCode is the mobile network code (MNC) associated with the IP address and ISP.
	// See https://en.wikipedia.org/wiki/Mobile_network_code
	MobileNetworkCode string `json:"mobile_network_code"            maxminddb:"mobile_network_code"`
	// Organization is the name of the organization associated with the IP address
	Organization string `json:"organization"                   maxminddb:"organization"`
	// AutonomousSystemNumber is the autonomous system number associated with the IP address
	AutonomousSystemNumber uint `json:"autonomous_system_number"       maxminddb:"autonomous_system_number"`
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
