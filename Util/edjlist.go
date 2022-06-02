package main

type Info struct {
	v   interface{}
	ebj []interface{}
}

type EdjList []Info

func (el *EdjList) InsertVertex(v interface{}) {

}

func (el *EdjList) DleteVertex(v interface{}) (Info, error) {

	return Info{}, nil
}
