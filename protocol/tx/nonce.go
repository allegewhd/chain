package tx

import "chain/protocol/bc"

type nonce struct {
	body struct {
		Program   bc.Program
		TimeRange EntryRef
		ExtHash   extHash
	}
}

func (nonce) Type() string         { return "nonce1" }
func (n *nonce) Body() interface{} { return n.body }

func newNonce(p bc.Program, tr EntryRef) *nonce {
	n := new(nonce)
	n.body.Program = p
	n.body.TimeRange = tr
	return n
}
