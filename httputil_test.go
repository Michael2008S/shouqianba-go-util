package shouqianba

import "testing"

const (
	vendor_sn  = ""
	vendor_key = ""
	code       = ""
)

func TestActivate(t *testing.T) {
	Activate(vendor_sn, vendor_key, code)
}
