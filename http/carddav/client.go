package carddav

//Request holds methods necessary for requesting cardDAV HTTP servers.
type Request interface {
	Get() (string, error)
}


//TODO decorreleer basic auth de carddav ?



