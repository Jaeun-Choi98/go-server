package model

import "database/sql"

type VMSWithServinfo struct {
	SpotName string         `db:"SPOT_NM" json:"spot_name"`
	VmsId    int            `db:"VMS_ID" json:"vms_id"`
	Seq      int            `db:"SEQ" json:"seq"`
	ServDt   []uint8        `db:"SERV_DT" json:"serv_dt"`
	LibId    sql.NullString `db:"LIB_ID" json:"lib_id"`
	ServLib  sql.NullString `db:"SERV_LIB" json:"serv_liv"`
	ShowTm   sql.NullInt32  `db:"SHOW_TM" json:"show_tm"`
	ShowType sql.NullInt32  `db:"SHOW_TYPE" json:"show_type"`
	ShowDir  sql.NullInt32  `db:"SHOW_DIR" json:"show_dir"`
	ServRst  sql.NullInt32  `db:"SERV_RST" json:"serv_rst"`
}

type GroupAndAdmin struct {
	GroupName sql.NullString `db:"GRP_NM" json:"grp_nm"`
	AdminName sql.NullString `db:"OPRT_NM" json:"oprt_nm"`
}

// mysql, oracle test
type User struct {
	UserID   int    `db:"user_id"`
	UserName string `db:"user_name"`
	Age      int    `db:"age"`
}
