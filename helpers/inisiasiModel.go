package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

//InisiasiModel digunakan untuk membuat tabel di dalam database sesuai dengan model yang di buat
func InisiasiModel(con *sql.DB, namaTable string, model interface{}) error {
	ref := reflect.TypeOf(model)
	sqlFields := []string{}
	for i := 0; i < ref.NumField(); i++ {
		sqlFields = append(sqlFields, ref.Field(i).Tag.Get("json")+" "+ref.Field(i).Tag.Get("sql"))
	}
	fields := strings.Join(sqlFields, ", ")
	tableStr := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", namaTable, fields)
	_, err := con.Query(tableStr)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
