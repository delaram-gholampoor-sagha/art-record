

package protocol


/*
	Storage Data Structure
*/

// https://en.wikipedia.org/wiki/Uniform_Resource_Identifier
// https://en.wikipedia.org/wiki/Uniform_Resource_Name
// https://en.wikipedia.org/wiki/Electronic_Product_Code
type QuiddityURI interface {
	URI() []byte          // Locale name in the Computer world.
	QuiddityID() [16]byte // quiddity domain
	RequestID() [16]byte  // user-request domain
}

type QuiddityURI_StorageServices interface {
	Save(q QuiddityURI) (err error)

	Count(uri []byte) (numbers uint64, err error)
	Get(uri []byte, versionOffset uint64) (q QuiddityURI, err error)
	Last(uri []byte) (q QuiddityURI, err error)

	ListQuiddityURIs(quiddityID [16]byte, offset, limit uint64) (uris [][]byte, err error)
}
