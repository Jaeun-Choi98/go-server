package model

/**
 * if value is null, needs sql.NullString or NullInt32
 * also, datetime type needs []uint8
 * e.g.
 	SpotName string         `db:"SPOT_NM" json:"spot_name"`
	VmsId    int            `db:"VMS_ID" json:"vms_id"`
	Seq      int            `db:"SEQ" json:"seq"`
	ServDt   []uint8        `db:"SERV_DT" json:"serv_dt"`
	LibId    sql.NullString `db:"LIB_ID" json:"lib_id"`
	ServLib  sql.NullString `db:"SERV_LIB" json:"serv_liv"`
	ShowTm   sql.NullInt32  `db:"SHOW_TM" json:"show_tm"`
*/

// mysql, oracle, maria test
type User struct {
	UserID   int    `db:"user_id"`
	UserName string `db:"user_name"`
	Age      string `db:"age"`
}
