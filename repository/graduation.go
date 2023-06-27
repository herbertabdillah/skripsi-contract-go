package repository

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
)

func (r Repository) GetTranscript(id string) (*state.Transcript, error) {
	res, err := r.context.State().Get("Transcript."+id, &state.Transcript{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.Transcript)

	return &obj, nil
}

func (r Repository) UpdateTranscript(obj *state.Transcript) (*state.Transcript, error) {
	err := r.context.State().Put("Transcript."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) InsertTranscript(obj *state.Transcript) (*state.Transcript, error) {
	err := r.context.State().Insert("Transcript."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
