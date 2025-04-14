package dao

import (
	"database/sql"
	"log"
	"root/model"

	_ "github.com/go-mysql-org/go-mysql/driver"
)

type MariaDB struct {
	db *sql.DB
}

func NewMariaDB(dbCon string) DaoInterface {
	db, err := sql.Open("mysql", dbCon)
	if err != nil {
		log.Println("failed to connect db")
		return nil
	}
	return &MariaDB{
		db: db,
	}
}

func (m *MariaDB) Close() {
	m.db.Close()
}

func (m *MariaDB) GetGroupAndAdmin() []model.GroupAndAdmin {
	query := `SELECT og.GRP_NM AS '그룹 이름', o.OPRT_NM AS '운영자 이름' FROM oprt_grp og
LEFT JOIN oprt o ON og.GRP_ID = o.GRP_ID;`
	rows, err := m.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}

	ret := make([]model.GroupAndAdmin, 0)
	for rows.Next() {
		instance := model.GroupAndAdmin{}
		if err := rows.Scan(&instance.AdminName, &instance.GroupName); err != nil {
			log.Println(err)
			return nil
		}
		ret = append(ret, instance)
	}
	return ret
}

func (m *MariaDB) GetVMSWithServeInfo() []model.VMSWithServinfo {
	query := `SELECT v.SPOT_NM, s.* FROM vms v
INNER JOIN vms_servinfo_stt s ON v.VMS_ID = s.VMS_ID
WHERE date(s.SERV_DT) >= '2024-04-01';`
	rows, err := m.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}

	ret := make([]model.VMSWithServinfo, 0)
	for rows.Next() {
		instance := model.VMSWithServinfo{}
		if err := rows.Scan(&instance.SpotName, &instance.VmsId,
			&instance.Seq, &instance.ServDt, &instance.LibId, &instance.ServLib,
			&instance.ShowTm, &instance.ShowType, &instance.ShowDir, &instance.ServRst); err != nil {
			log.Println(err)
			return nil
		}
		ret = append(ret, instance)
	}

	return ret
}

func (m *MariaDB) Exec(query string, params ...any) bool {
	rst, err := m.db.Exec(query, params...)
	if err != nil {
		log.Println(err)
		return false
	}
	cnt, err := rst.RowsAffected()
	if err != nil {
		log.Println(err)
		return false
	}
	if cnt == 0 {
		log.Println("nothin exchange")
		return false
	}
	return true
}
