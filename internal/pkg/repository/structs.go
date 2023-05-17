package repository

type Student struct {
	ID     uint64 `db:"id" json:"id,omitempty"`
	Name   string `db:"name" json:"name,omitempty"`
	Grades int16  `db:"grades" json:"grades,omitempty"`
	UnivID uint64 `db:"univ_apply_id" json:"univ_id"`
}

type University struct {
	ID       uint64 `db:"id" json:"id,omitempty"`
	Name     string `db:"univ_name" json:"name,omitempty"`
	Facility string `db:"facility" json:"facility,omitempty"`
}
