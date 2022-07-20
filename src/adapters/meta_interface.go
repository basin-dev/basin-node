/* The metaadaptere just takes a URL and makes the correct call to the adapter that lives there.
Basically this is only responsible for deciding whether the URL is Basin, Local, HTTP, or other and then routing to the correct adapter.
The meta adapter lives on the node.
It is only a NICE TO HAVE feature right now.
*/
package adapters

type Adapter interface {
	Read(url string) []byte
	Write(url string, value []byte) bool
}

type MetaAdapter struct{}

func (m MetaAdapter) Read() {

}

func (m MetaAdapter) Write() {

}
