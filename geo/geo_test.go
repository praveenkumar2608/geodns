package geo

import (
	"net"
	"testing"
)

func TestGeoIPv4(t *testing.T) {
	ip := net.ParseIP("207.171.1.1")

	geoIP := New()
	geoIP.(*GeoIP).DisableV6 = true

	if geoIP.SetupCity() == nil {
		_, _, regionGroup, _, _ := geoIP.GetCountryRegion(ip)
		if regionGroup != "us-west" {
			t.Errorf("Expected regionGroup '%s', got '%s'", "us-west", regionGroup)
		}
	}
	if geoIP.SetupCountry() == nil {
		country, continent, _ := geoIP.GetCountry(ip)
		if country != "us" {
			t.Errorf("Expected country '%s', got '%s'", "us", country)
		}
		if continent != "north-america" {
			t.Errorf("Expected continent '%s', got '%s'", "north-america", continent)
		}
	}
	if geoIP.SetupASN() == nil {
		asn, _ := geoIP.GetASN(ip)
		if asn != "as7012" {
			t.Errorf("Expected ASN '%s', got '%s'", "as7012", asn)
		}
	}

	asn, _ := geoIP.GetASN(ip)
	if asn != "as7012" {
		t.Errorf("Expected ASN '%s', got '%s'", "as7012", asn)
	}
}

func TestGeoIPv6(t *testing.T) {

	// los angeles
	ip := net.ParseIP("2607:f238:3::1")

	// lux
	// ip := net.ParseIP("2001:888:2156::1")

	geoIP := New()

	if false {
		// todo: find a test IP that works for city database
		if geoIP.SetupCity() == nil {
			_, _, regionGroup, _, _ := geoIP.GetCountryRegion(ip)
			if regionGroup != "us-west" {
				t.Errorf("Expected regionGroup '%s', got '%s'", "us-west", regionGroup)
			}
		}
	}

	if geoIP.SetupCountry() == nil {
		country, continent, _ := geoIP.GetCountry(ip)
		if country != "us" {
			t.Errorf("Expected country '%s', got '%s'", "us", country)
		}
		if continent != "north-america" {
			t.Errorf("Expected continent '%s', got '%s'", "north-america", continent)
		}
	}
	if geoIP.SetupASN() == nil {
		asn, _ := geoIP.GetASN(ip)
		if asn != "as7012" {
			t.Errorf("Expected ASN '%s', got '%s'", "as7012", asn)
		}
	}

	asn, _ := geoIP.GetASN(ip)
	if asn != "as7012" {
		t.Errorf("Expected ASN '%s', got '%s'", "as7012", asn)
	}
}

func TestIsV6(t *testing.T) {
	if isv6(net.ParseIP("207.171.1.1")) {
		t.Error("isv6(207.171.1.1) == true, oops!")
	}
	if !isv6(net.ParseIP("2607:f238:3::1")) {
		t.Error("isv6(207.171.1.1) == false, oops!")
	}
}