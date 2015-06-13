package faroo

import ( 
	"appengine"
)

 

type Error string

func (e Error) Error() string {
	return string(e)
}

func init() {
	 
}
 