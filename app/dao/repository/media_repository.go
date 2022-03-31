package repository

import (
	"database/sql"
	"fun.tvapi/app/dao/model/media"
	"fun.tvapi/app/pkg/database"
	"fun.tvapi/app/pkg/logger"
	"fun.tvapi/app/pkg/util/convert"
	"strings"
)

type MediaRepository struct {
}

func (r *MediaRepository) GetReader() *sql.DB {
	return database.NewMysqlPool().Read("media")
}

func (r *MediaRepository) GetWriter() *sql.DB {
	return database.NewMysqlPool().Write("media")
}

func (r *MediaRepository) GetByIds(idArr []int, fields []string) []media.MediaModel {
	var querySql string = "select " + strings.Join(fields, ",") + " from fm_media where media_id in (" + strings.Join(convert.ToStringSlice(idArr), ",") + ")"
	rows, err := r.GetReader().Query(querySql)
	if err != nil {
		return []media.MediaModel{}
	}
	var result = make([]media.MediaModel, 0, len(idArr))
	var m media.MediaModel
	for rows.Next() {
		er := rows.Scan(m.ToScanField(fields)...)
		if er != nil {
			logger.Fatal("repository: media scan failed!", map[string]interface{}{"err": er})
			continue
		}
		result = append(result, m)
	}
	defer rows.Close()
	return result
}
