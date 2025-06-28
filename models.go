package geoip2

import "net/netip"

// Names contains localized names for geographic entities.
type Names struct {
	// German localized name
	German string `json:"de,omitzero"    maxminddb:"de"`
	// English localized name
	English string `json:"en,omitzero"    maxminddb:"en"`
	// Spanish localized name
	Spanish string `json:"es,omitzero"    maxminddb:"es"`
	// French localized name
	French string `json:"fr,omitzero"    maxminddb:"fr"`
	// Japanese localized name
	Japanese string `json:"ja,omitzero"    maxminddb:"ja"`
	// BrazilianPortuguese localized name (pt-BR)
	BrazilianPortuguese string `json:"pt-BR,omitzero" maxminddb:"pt-BR"` //nolint:tagliatelle,lll // pt-BR matches MMDB format
	// Russian localized name
	Russian string `json:"ru,omitzero"    maxminddb:"ru"`
	// SimplifiedChinese localized name (zh-CN)
	SimplifiedChinese string `json:"zh-CN,omitzero" maxminddb:"zh-CN"` //nolint:tagliatelle // zh-CN matches MMDB format
}

var zeroNames Names

// HasData returns true if the Names struct has any localized names.
func (n Names) HasData() bool {
	return n != zeroNames
}

// The Enterprise struct corresponds to the data in the GeoIP2 Enterprise
// database.
type Enterprise struct {
	// Continent contains data for the continent record associated with the IP
	// address.
	Continent struct {
		// Names contains localized names for the continent
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// Code is a two character continent code like "NA" (North America) or
		// "OC" (Oceania)
		Code string `json:"code,omitzero" maxminddb:"code"`
		// GeoNameID for the continent
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
	} `json:"continent,omitzero"           maxminddb:"continent"`
	// City contains data for the city record associated with the IP address.
	City struct {
		// Names contains localized names for the city
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// GeoNameID for the city
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the city is correct
		Confidence uint8 `json:"confidence,omitzero" maxminddb:"confidence"`
	} `json:"city,omitzero"                maxminddb:"city"`
	// Postal contains data for the postal record associated with the IP address.
	Postal struct {
		// Code is the postal code of the location. Postal codes are not
		// available for all countries.
		// In some countries, this will only contain part of the postal code.
		Code string `json:"code,omitzero" maxminddb:"code"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the postal code is correct
		Confidence uint8 `json:"confidence,omitzero" maxminddb:"confidence"`
	} `json:"postal,omitzero"              maxminddb:"postal"`
	// Subdivisions contains data for the subdivisions associated with the IP
	// address. The subdivisions array is ordered from largest to smallest. For
	// instance, the response for Oxford in the United Kingdom would have England
	// as the first element and Oxfordshire as the second element.
	Subdivisions []struct {
		// Names contains localized names for the subdivision
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is a string up to three characters long containing the
		// subdivision portion
		// of the ISO 3166-2 code. See https://en.wikipedia.org/wiki/ISO_3166-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// GeoNameID for the subdivision
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the subdivision is correct
		Confidence uint8 `json:"confidence,omitzero" maxminddb:"confidence"`
	} `json:"subdivisions,omitzero"        maxminddb:"subdivisions"`
	// RepresentedCountry contains data for the represented country associated
	// with the IP address. The represented country is the country represented
	// by something like a military base or embassy.
	RepresentedCountry struct {
		// Names contains localized names for the represented country
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the represented
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// Type is a string indicating the type of entity that is representing
		// the country.
		// Currently this is only "military" but may expand in the future.
		Type string `json:"type,omitzero" maxminddb:"type"`
		// GeoNameID for the represented country
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the represented country is a member
		// state of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union,omitzero" maxminddb:"is_in_european_union"`
	} `json:"represented_country,omitzero" maxminddb:"represented_country"`
	// Country contains data for the country record associated with the IP
	// address. This record represents the country where MaxMind believes the IP
	// is located.
	Country struct {
		// Names contains localized names for the country
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// GeoNameID for the country
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the country is correct
		Confidence uint8 `json:"confidence,omitzero" maxminddb:"confidence"`
		// IsInEuropeanUnion is true if the country is a member state of the
		// European Union
		IsInEuropeanUnion bool `json:"is_in_european_union,omitzero" maxminddb:"is_in_european_union"`
	} `json:"country,omitzero"             maxminddb:"country"`
	// RegisteredCountry contains data for the registered country associated
	// with the IP address. This record represents the country where the ISP has
	// registered the IP block and may differ from the user's country.
	RegisteredCountry struct {
		// Names contains localized names for the registered country
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the registered
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// GeoNameID for the registered country
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// Confidence is a value from 0-100 indicating MaxMind's confidence that
		// the registered country is correct
		Confidence uint8 `json:"confidence,omitzero" maxminddb:"confidence"`
		// IsInEuropeanUnion is true if the registered country is a member state
		// of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union,omitzero" maxminddb:"is_in_european_union"`
	} `json:"registered_country,omitzero"  maxminddb:"registered_country"`
	// Traits contains various traits associated with the IP address
	Traits struct {
		// Network is the largest network prefix where all fields besides
		// IPAddress have the same value.
		Network netip.Prefix `json:"network,omitzero"`
		// IPAddress is the IP address used during the lookup
		IPAddress netip.Addr `json:"ip_address,omitzero"`
		// AutonomousSystemOrganization for the registered ASN
		AutonomousSystemOrganization string `json:"autonomous_system_organization,omitzero" maxminddb:"autonomous_system_organization"` //nolint:lll
		// ConnectionType indicates the connection type. May be Dialup,
		// Cable/DSL, Corporate, Cellular, or Satellite
		ConnectionType string `json:"connection_type,omitzero" maxminddb:"connection_type"`
		// Domain is the second level domain associated with the IP address
		// (e.g., "example.com")
		Domain string `json:"domain,omitzero" maxminddb:"domain"`
		// ISP is the name of the ISP associated with the IP address
		ISP string `json:"isp,omitzero" maxminddb:"isp"`
		// MobileCountryCode is the mobile country code (MCC) associated with
		// the IP address and ISP.
		// See https://en.wikipedia.org/wiki/Mobile_country_code
		MobileCountryCode string `json:"mobile_country_code,omitzero" maxminddb:"mobile_country_code"`
		// MobileNetworkCode is the mobile network code (MNC) associated with
		// the IP address and ISP.
		// See https://en.wikipedia.org/wiki/Mobile_network_code
		MobileNetworkCode string `json:"mobile_network_code,omitzero" maxminddb:"mobile_network_code"`
		// Organization is the name of the organization associated with the IP
		// address
		Organization string `json:"organization,omitzero" maxminddb:"organization"`
		// UserType indicates the user type associated with the IP address
		// (business, cafe, cellular, college, etc.)
		UserType string `json:"user_type,omitzero" maxminddb:"user_type"`
		// StaticIPScore is an indicator of how static or dynamic an IP address
		// is, ranging from 0 to 99.99
		StaticIPScore float64 `json:"static_ip_score,omitzero" maxminddb:"static_ip_score"`
		// AutonomousSystemNumber for the IP address
		AutonomousSystemNumber uint `json:"autonomous_system_number,omitzero" maxminddb:"autonomous_system_number"`
		// IsAnycast is true if the IP address belongs to an anycast network.
		// See https://en.wikipedia.org/wiki/Anycast
		IsAnycast bool `json:"is_anycast,omitzero" maxminddb:"is_anycast"`
		// IsLegitimateProxy is true if MaxMind believes this IP address to be a
		// legitimate proxy, such as an internal VPN used by a corporation
		IsLegitimateProxy bool `json:"is_legitimate_proxy,omitzero" maxminddb:"is_legitimate_proxy"`
	} `json:"traits,omitzero"              maxminddb:"traits"`
	// Location contains data for the location record associated with the IP
	// address
	Location struct {
		// TimeZone is the time zone associated with location, as specified by
		// the IANA Time Zone Database (e.g., "America/New_York")
		TimeZone string `json:"time_zone,omitzero" maxminddb:"time_zone"`
		// Latitude is the approximate latitude of the location associated with
		// the IP address. This value is not precise and should not be used to
		// identify a particular address or household.
		Latitude float64 `json:"latitude" maxminddb:"latitude"`
		// Longitude is the approximate longitude of the location associated with
		// the IP address. This value is not precise and should not be used to
		// identify a particular address or household.
		Longitude float64 `json:"longitude" maxminddb:"longitude"`
		// MetroCode is a metro code for targeting advertisements.
		//
		// Deprecated: Metro codes are no longer maintained and should not be used.
		MetroCode uint `json:"metro_code,omitzero" maxminddb:"metro_code"`
		// AccuracyRadius is the approximate accuracy radius in kilometers around
		// the latitude and longitude. This is the radius where we have a 67%
		// confidence that the device using the IP address resides within the
		// circle.
		AccuracyRadius uint16 `json:"accuracy_radius,omitzero" maxminddb:"accuracy_radius"`
	} `json:"location,omitzero"            maxminddb:"location"`
}

// HasData returns true if any GeoIP data was found for the IP in the Enterprise database.
// This excludes the Network and IPAddress fields which are always populated for found IPs.
func (e Enterprise) HasData() bool {
	return e.hasContinentData() || e.hasCityData() || e.hasPostalData() ||
		e.hasSubdivisionsData() || e.hasRepresentedCountryData() ||
		e.hasCountryData() || e.hasRegisteredCountryData() ||
		e.hasTraitsData() || e.hasLocationData()
}

func (e Enterprise) hasContinentData() bool {
	return e.Continent.Names.HasData() || e.Continent.Code != "" || e.Continent.GeoNameID != 0
}

func (e Enterprise) hasCityData() bool {
	return e.City.Names.HasData() || e.City.GeoNameID != 0 || e.City.Confidence != 0
}

func (e Enterprise) hasPostalData() bool {
	return e.Postal.Code != "" || e.Postal.Confidence != 0
}

func (e Enterprise) hasSubdivisionsData() bool {
	for _, sub := range e.Subdivisions {
		if sub.Names.HasData() || sub.ISOCode != "" || sub.GeoNameID != 0 || sub.Confidence != 0 {
			return true
		}
	}
	return false
}

func (e Enterprise) hasRepresentedCountryData() bool {
	return e.RepresentedCountry.Names.HasData() || e.RepresentedCountry.ISOCode != "" ||
		e.RepresentedCountry.Type != "" || e.RepresentedCountry.GeoNameID != 0 ||
		e.RepresentedCountry.IsInEuropeanUnion
}

func (e Enterprise) hasCountryData() bool {
	return e.Country.Names.HasData() || e.Country.ISOCode != "" || e.Country.GeoNameID != 0 ||
		e.Country.Confidence != 0 || e.Country.IsInEuropeanUnion
}

func (e Enterprise) hasRegisteredCountryData() bool {
	return e.RegisteredCountry.Names.HasData() || e.RegisteredCountry.ISOCode != "" ||
		e.RegisteredCountry.GeoNameID != 0 || e.RegisteredCountry.Confidence != 0 ||
		e.RegisteredCountry.IsInEuropeanUnion
}

func (e Enterprise) hasTraitsData() bool {
	return e.Traits.AutonomousSystemOrganization != "" || e.Traits.ConnectionType != "" ||
		e.Traits.Domain != "" || e.Traits.ISP != "" || e.Traits.MobileCountryCode != "" ||
		e.Traits.MobileNetworkCode != "" || e.Traits.Organization != "" ||
		e.Traits.UserType != "" || e.Traits.StaticIPScore != 0 ||
		e.Traits.AutonomousSystemNumber != 0 || e.Traits.IsAnycast ||
		e.Traits.IsLegitimateProxy
}

func (e Enterprise) hasLocationData() bool {
	return e.Location.TimeZone != "" || e.Location.Latitude != 0 || e.Location.Longitude != 0 ||
		e.Location.MetroCode != 0 || e.Location.AccuracyRadius != 0
}

// The City struct corresponds to the data in the GeoIP2/GeoLite2 City
// databases.
type City struct {
	// Traits contains various traits associated with the IP address
	Traits struct {
		// IPAddress is the IP address used during the lookup
		IPAddress netip.Addr `json:"ip_address,omitzero"`
		// Network is the network prefix for this record. This is the largest
		// network where all of the fields besides IPAddress have the same value.
		Network netip.Prefix `json:"network,omitzero"`
		// IsAnycast is true if the IP address belongs to an anycast network.
		// See https://en.wikipedia.org/wiki/Anycast
		IsAnycast bool `json:"is_anycast,omitzero" maxminddb:"is_anycast"`
	} `json:"traits,omitzero"              maxminddb:"traits"`
	// Postal contains data for the postal record associated with the IP address
	Postal struct {
		// Code is the postal code of the location. Postal codes are not
		// available for all countries.
		// In some countries, this will only contain part of the postal code.
		Code string `json:"code,omitzero" maxminddb:"code"`
	} `json:"postal,omitzero"              maxminddb:"postal"`
	// Continent contains data for the continent record associated with the IP address
	Continent struct {
		// Names contains localized names for the continent
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// Code is a two character continent code like "NA" (North America) or
		// "OC" (Oceania)
		Code string `json:"code,omitzero" maxminddb:"code"`
		// GeoNameID for the continent
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
	} `json:"continent,omitzero"           maxminddb:"continent"`
	// City contains data for the city record associated with the IP address
	City struct {
		// Names contains localized names for the city
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// GeoNameID for the city
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
	} `json:"city,omitzero"                maxminddb:"city"`
	// Subdivisions contains data for the subdivisions associated with the IP
	// address. The subdivisions array is ordered from largest to smallest. For
	// instance, the response for Oxford in the United Kingdom would have England
	// as the first element and Oxfordshire as the second element.
	Subdivisions []struct {
		Names     Names  `json:"names,omitzero" maxminddb:"names"`
		ISOCode   string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		GeoNameID uint   `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
	} `json:"subdivisions,omitzero"        maxminddb:"subdivisions"`
	// RepresentedCountry contains data for the represented country associated
	// with the IP address. The represented country is the country represented
	// by something like a military base or embassy.
	RepresentedCountry struct {
		// Names contains localized names for the represented country
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the represented
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// Type is a string indicating the type of entity that is representing
		// the country.
		// Currently this is only "military" but may expand in the future.
		Type string `json:"type,omitzero" maxminddb:"type"`
		// GeoNameID for the represented country
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the represented country is a member
		// state of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union,omitzero" maxminddb:"is_in_european_union"`
	} `json:"represented_country,omitzero" maxminddb:"represented_country"`
	// Country contains data for the country record associated with the IP
	// address. This record represents the country where MaxMind believes the IP
	// is located.
	Country struct {
		// Names contains localized names for the country
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// GeoNameID for the country
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the country is a member state of the
		// European Union
		IsInEuropeanUnion bool `json:"is_in_european_union,omitzero" maxminddb:"is_in_european_union"`
	} `json:"country,omitzero"             maxminddb:"country"`
	// RegisteredCountry contains data for the registered country associated
	// with the IP address. This record represents the country where the ISP has
	// registered the IP block and may differ from the user's country.
	RegisteredCountry struct {
		// Names contains localized names for the registered country
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the registered
		// country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// GeoNameID for the registered country
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the registered country is a member state
		// of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union,omitzero" maxminddb:"is_in_european_union"`
	} `json:"registered_country,omitzero"  maxminddb:"registered_country"`
	// Location contains data for the location record associated with the IP
	// address
	Location struct {
		// TimeZone is the time zone associated with location, as specified by
		// the IANA Time Zone Database (e.g., "America/New_York")
		TimeZone string `json:"time_zone,omitzero" maxminddb:"time_zone"`
		// Latitude is the approximate latitude of the location associated with
		// the IP address. This value is not precise and should not be used to
		// identify a particular address or household.
		Latitude float64 `json:"latitude" maxminddb:"latitude"`
		// Longitude is the approximate longitude of the location associated with
		// the IP address. This value is not precise and should not be used to
		// identify a particular address or household.
		Longitude float64 `json:"longitude" maxminddb:"longitude"`
		// MetroCode is a metro code for targeting advertisements.
		//
		// Deprecated: Metro codes are no longer maintained and should not be used.
		MetroCode uint `json:"metro_code,omitzero" maxminddb:"metro_code"`
		// AccuracyRadius is the approximate accuracy radius in kilometers around
		// the latitude and longitude. This is the radius where we have a 67%
		// confidence that the device using the IP address resides within the
		// circle.
		AccuracyRadius uint16 `json:"accuracy_radius,omitzero" maxminddb:"accuracy_radius"`
	} `json:"location,omitzero"            maxminddb:"location"`
}

// HasData returns true if any GeoIP data was found for the IP in the City database.
// This excludes the Network and IPAddress fields which are always populated for found IPs.
func (c City) HasData() bool {
	return c.hasTraitsData() || c.hasPostalData() || c.hasContinentData() ||
		c.hasCityData() || c.hasSubdivisionsData() || c.hasRepresentedCountryData() ||
		c.hasCountryData() || c.hasRegisteredCountryData() || c.hasLocationData()
}

func (c City) hasTraitsData() bool {
	return c.Traits.IsAnycast
}

func (c City) hasPostalData() bool {
	return c.Postal.Code != ""
}

func (c City) hasContinentData() bool {
	return c.Continent.Names.HasData() || c.Continent.Code != "" || c.Continent.GeoNameID != 0
}

func (c City) hasCityData() bool {
	return c.City.Names.HasData() || c.City.GeoNameID != 0
}

func (c City) hasSubdivisionsData() bool {
	for _, sub := range c.Subdivisions {
		if sub.Names.HasData() || sub.ISOCode != "" || sub.GeoNameID != 0 {
			return true
		}
	}
	return false
}

func (c City) hasRepresentedCountryData() bool {
	return c.RepresentedCountry.Names.HasData() || c.RepresentedCountry.ISOCode != "" ||
		c.RepresentedCountry.Type != "" || c.RepresentedCountry.GeoNameID != 0 ||
		c.RepresentedCountry.IsInEuropeanUnion
}

func (c City) hasCountryData() bool {
	return c.Country.Names.HasData() || c.Country.ISOCode != "" || c.Country.GeoNameID != 0 ||
		c.Country.IsInEuropeanUnion
}

func (c City) hasRegisteredCountryData() bool {
	return c.RegisteredCountry.Names.HasData() || c.RegisteredCountry.ISOCode != "" ||
		c.RegisteredCountry.GeoNameID != 0 || c.RegisteredCountry.IsInEuropeanUnion
}

func (c City) hasLocationData() bool {
	return c.Location.TimeZone != "" || c.Location.Latitude != 0 || c.Location.Longitude != 0 ||
		c.Location.MetroCode != 0 || c.Location.AccuracyRadius != 0
}

// The Country struct corresponds to the data in the GeoIP2/GeoLite2
// Country databases.
type Country struct {
	// Traits contains various traits associated with the IP address
	Traits struct {
		// IPAddress is the IP address used during the lookup
		IPAddress netip.Addr `json:"ip_address,omitzero"`
		// Network is the largest network prefix where all fields besides
		// IPAddress have the same value.
		Network netip.Prefix `json:"network,omitzero"`
		// IsAnycast is true if the IP address belongs to an anycast network.
		// See https://en.wikipedia.org/wiki/Anycast
		IsAnycast bool `json:"is_anycast,omitzero" maxminddb:"is_anycast"`
	} `json:"traits,omitzero"              maxminddb:"traits"`
	// Continent contains data for the continent record associated with the IP address
	Continent struct {
		// Names contains localized names for the continent
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// Code is a two character continent code like "NA" (North America) or
		// "OC" (Oceania)
		Code string `json:"code,omitzero" maxminddb:"code"`
		// GeoNameID for the continent
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
	} `json:"continent,omitzero"           maxminddb:"continent"`
	// RepresentedCountry contains data for the represented country associated
	// with the IP address. The represented country is the country represented
	// by something like a military base or embassy.
	RepresentedCountry struct {
		// Names contains localized names for the represented country
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the represented
		// country. See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// Type is a string indicating the type of entity that is representing
		// the country. Currently this is only "military" but may expand in the future.
		Type string `json:"type,omitzero" maxminddb:"type"`
		// GeoNameID for the represented country
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the represented country is a member
		// state of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union,omitzero" maxminddb:"is_in_european_union"`
	} `json:"represented_country,omitzero" maxminddb:"represented_country"`
	// Country contains data for the country record associated with the IP
	// address. This record represents the country where MaxMind believes the IP
	// is located.
	Country struct {
		// Names contains localized names for the country
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the country.
		// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// GeoNameID for the country
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the country is a member state of the
		// European Union
		IsInEuropeanUnion bool `json:"is_in_european_union,omitzero" maxminddb:"is_in_european_union"`
	} `json:"country,omitzero"             maxminddb:"country"`
	// RegisteredCountry contains data for the registered country associated
	// with the IP address. This record represents the country where the ISP has
	// registered the IP block and may differ from the user's country.
	RegisteredCountry struct {
		// Names contains localized names for the registered country
		Names Names `json:"names,omitzero" maxminddb:"names"`
		// ISOCode is the two-character ISO 3166-1 alpha code for the registered
		// country. See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
		ISOCode string `json:"iso_code,omitzero" maxminddb:"iso_code"`
		// GeoNameID for the registered country
		GeoNameID uint `json:"geoname_id,omitzero" maxminddb:"geoname_id"`
		// IsInEuropeanUnion is true if the registered country is a member state
		// of the European Union
		IsInEuropeanUnion bool `json:"is_in_european_union,omitzero" maxminddb:"is_in_european_union"`
	} `json:"registered_country,omitzero"  maxminddb:"registered_country"`
}

// HasData returns true if any GeoIP data was found for the IP in the Country database.
// This excludes the Network and IPAddress fields which are always populated for found IPs.
func (c Country) HasData() bool {
	return c.Continent.Names.HasData() || c.Continent.Code != "" || c.Continent.GeoNameID != 0 ||
		c.RepresentedCountry.Names.HasData() || c.RepresentedCountry.ISOCode != "" ||
		c.RepresentedCountry.Type != "" || c.RepresentedCountry.GeoNameID != 0 ||
		c.RepresentedCountry.IsInEuropeanUnion ||
		c.Country.Names.HasData() || c.Country.ISOCode != "" || c.Country.GeoNameID != 0 ||
		c.Country.IsInEuropeanUnion ||
		c.RegisteredCountry.Names.HasData() || c.RegisteredCountry.ISOCode != "" ||
		c.RegisteredCountry.GeoNameID != 0 || c.RegisteredCountry.IsInEuropeanUnion ||
		c.Traits.IsAnycast
}

// The AnonymousIP struct corresponds to the data in the GeoIP2
// Anonymous IP database.
type AnonymousIP struct {
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address,omitzero"`
	// Network is the largest network prefix where all fields besides
	// IPAddress have the same value.
	Network netip.Prefix `json:"network,omitzero"`
	// IsAnonymous is true if the IP address belongs to any sort of anonymous network.
	IsAnonymous bool `json:"is_anonymous,omitzero"         maxminddb:"is_anonymous"`
	// IsAnonymousVPN is true if the IP address is registered to an anonymous
	// VPN provider. If a VPN provider does not register subnets under names
	// associated with them, we will likely only flag their IP ranges using the
	// IsHostingProvider attribute.
	IsAnonymousVPN bool `json:"is_anonymous_vpn,omitzero"     maxminddb:"is_anonymous_vpn"`
	// IsHostingProvider is true if the IP address belongs to a hosting or VPN provider
	// (see description of IsAnonymousVPN attribute).
	IsHostingProvider bool `json:"is_hosting_provider,omitzero"  maxminddb:"is_hosting_provider"`
	// IsPublicProxy is true if the IP address belongs to a public proxy.
	IsPublicProxy bool `json:"is_public_proxy,omitzero"      maxminddb:"is_public_proxy"`
	// IsResidentialProxy is true if the IP address is on a suspected
	// anonymizing network and belongs to a residential ISP.
	IsResidentialProxy bool `json:"is_residential_proxy,omitzero" maxminddb:"is_residential_proxy"`
	// IsTorExitNode is true if the IP address is a Tor exit node.
	IsTorExitNode bool `json:"is_tor_exit_node,omitzero"     maxminddb:"is_tor_exit_node"`
}

// HasData returns true if any data was found for the IP in the AnonymousIP database.
// This excludes the Network and IPAddress fields which are always populated for found IPs.
func (a AnonymousIP) HasData() bool {
	return a.IsAnonymous || a.IsAnonymousVPN || a.IsHostingProvider ||
		a.IsPublicProxy || a.IsResidentialProxy || a.IsTorExitNode
}

// The ASN struct corresponds to the data in the GeoLite2 ASN database.
type ASN struct {
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address,omitzero"`
	// Network is the largest network prefix where all fields besides
	// IPAddress have the same value.
	Network netip.Prefix `json:"network,omitzero"`
	// AutonomousSystemOrganization for the registered autonomous system number.
	AutonomousSystemOrganization string `json:"autonomous_system_organization,omitzero" maxminddb:"autonomous_system_organization"` //nolint:lll
	// AutonomousSystemNumber for the IP address.
	AutonomousSystemNumber uint `json:"autonomous_system_number,omitzero"       maxminddb:"autonomous_system_number"` //nolint:lll
}

// HasData returns true if any data was found for the IP in the ASN database.
// This excludes the Network and IPAddress fields which are always populated for found IPs.
func (a ASN) HasData() bool {
	return a.AutonomousSystemNumber != 0 || a.AutonomousSystemOrganization != ""
}

// The ConnectionType struct corresponds to the data in the GeoIP2
// Connection-Type database.
type ConnectionType struct {
	// ConnectionType indicates the connection type. May be Dialup, Cable/DSL,
	// Corporate, Cellular, or Satellite. Additional values may be added in the
	// future.
	ConnectionType string `json:"connection_type,omitzero" maxminddb:"connection_type"`
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address,omitzero"`
	// Network is the largest network prefix where all fields besides
	// IPAddress have the same value.
	Network netip.Prefix `json:"network,omitzero"`
}

// HasData returns true if any data was found for the IP in the ConnectionType database.
// This excludes the Network and IPAddress fields which are always populated for found IPs.
func (c ConnectionType) HasData() bool {
	return c.ConnectionType != ""
}

// The Domain struct corresponds to the data in the GeoIP2 Domain database.
type Domain struct {
	// Domain is the second level domain associated with the IP address
	// (e.g., "example.com")
	Domain string `json:"domain,omitzero"     maxminddb:"domain"`
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address,omitzero"`
	// Network is the largest network prefix where all fields besides
	// IPAddress have the same value.
	Network netip.Prefix `json:"network,omitzero"`
}

// HasData returns true if any data was found for the IP in the Domain database.
// This excludes the Network and IPAddress fields which are always populated for found IPs.
func (d Domain) HasData() bool {
	return d.Domain != ""
}

// The ISP struct corresponds to the data in the GeoIP2 ISP database.
type ISP struct {
	// Network is the largest network prefix where all fields besides
	// IPAddress have the same value.
	Network netip.Prefix `json:"network,omitzero"`
	// IPAddress is the IP address used during the lookup
	IPAddress netip.Addr `json:"ip_address,omitzero"`
	// AutonomousSystemOrganization for the registered ASN
	AutonomousSystemOrganization string `json:"autonomous_system_organization,omitzero" maxminddb:"autonomous_system_organization"` //nolint:lll
	// ISP is the name of the ISP associated with the IP address
	ISP string `json:"isp,omitzero"                            maxminddb:"isp"`
	// MobileCountryCode is the mobile country code (MCC) associated with the IP address and ISP.
	// See https://en.wikipedia.org/wiki/Mobile_country_code
	MobileCountryCode string `json:"mobile_country_code,omitzero"            maxminddb:"mobile_country_code"`
	// MobileNetworkCode is the mobile network code (MNC) associated with the IP address and ISP.
	// See https://en.wikipedia.org/wiki/Mobile_network_code
	MobileNetworkCode string `json:"mobile_network_code,omitzero"            maxminddb:"mobile_network_code"`
	// Organization is the name of the organization associated with the IP address
	Organization string `json:"organization,omitzero"                   maxminddb:"organization"`
	// AutonomousSystemNumber for the IP address
	AutonomousSystemNumber uint `json:"autonomous_system_number,omitzero"       maxminddb:"autonomous_system_number"`
}

// HasData returns true if any data was found for the IP in the ISP database.
// This excludes the Network and IPAddress fields which are always populated for found IPs.
func (i ISP) HasData() bool {
	return i.AutonomousSystemOrganization != "" || i.ISP != "" ||
		i.MobileCountryCode != "" || i.MobileNetworkCode != "" ||
		i.Organization != "" || i.AutonomousSystemNumber != 0
}