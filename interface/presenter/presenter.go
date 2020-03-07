package presenter

import (
	"encoding/json"
	"io"
)

type Presenter struct {
	w io.Writer
}

func NewPresenter(w io.Writer) *Presenter {
	return &Presenter{w: w}
}

func (presenter *Presenter) outputJSON(i interface{}) error {
	bs, err := json.Marshal(i)
	if err != nil {
		return err
	}
	_, err = presenter.w.Write(bs)
	if err != nil {
		return err
	}
	return nil
}
