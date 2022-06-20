// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovnnb

type (
	BFDStatus = string
)

var (
	BFDStatusDown      BFDStatus = "down"
	BFDStatusInit      BFDStatus = "init"
	BFDStatusUp        BFDStatus = "up"
	BFDStatusAdminDown BFDStatus = "admin_down"
)

// BFD defines an object in BFD table
type BFD struct {
	UUID        string            `ovsdb:"_uuid"`
	DetectMult  *int              `ovsdb:"detect_mult"`
	DstIP       string            `ovsdb:"dst_ip"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	LogicalPort string            `ovsdb:"logical_port"`
	MinRx       *int              `ovsdb:"min_rx"`
	MinTx       *int              `ovsdb:"min_tx"`
	Options     map[string]string `ovsdb:"options"`
	Status      *BFDStatus        `ovsdb:"status"`
}
