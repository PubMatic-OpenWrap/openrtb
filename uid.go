package openrtb

import "encoding/json"

// Uid defines the contract for bidrequest.user.eids[i].uids[j]
type Uid struct {
	ID  string          `json:"id"`
	Ext json.RawMessage `json:"ext,omitempty"`
}
