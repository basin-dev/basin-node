/* The meta interface just takes a URL and makes the correct call to the interface that lives there.
Basically this is only responsible for deciding whether the URL is Basin or HTTP and then making the corresponding call.
The meta interface lives on the node.
*/
package interfaces

import (
	. "github.com/sestinj/basin-node/util"
)

func callRead(url string) {
	scheme := ParseUrl(url).scheme

	if scheme == "basin" {
		// For now, just use leveldb_local. Eventually might need to get some metadata from other nodes?
		read()
	} else if scheme == "https" || scheme == "http" {
		// Make HTTP request

	} else {

	}
}

func callWrite(url string, value []byte) {
	// Might just reuse everything from the callRead function.
}
