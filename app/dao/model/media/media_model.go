package media

type MediaModel struct {
	MediaId         int     `json:"media_id"`
	OriginalId      int64   `json:"original_id"`
	ChannelId       int     `json:"channel_id"`
	AlternativeId   int     `json:"alternative_id"`
	Name            string  `json:"name"`
	Pinyin          string  `json:"pinyin"`
	Alias           string  `json:"alias"`
	Aword           string  `json:"aword"`
	Poster          string  `json:"poster"`
	Still           string  `json:"still"`
	NameImg         string  `json:"name_img"`
	BgImg           string  `json:"bg_img"`
	Director        string  `json:"director"`
	Actor           string  `json:"actor"`
	StaffPinyin     string  `json:"staff_pinyin"`
	Description     string  `json:"description"`
	Score           float32 `json:"score"`
	TotalVv         int     `json:"total_vv"`
	Heat            int     `json:"heat"`
	ReleaseDate     string  `json:"release_date"`
	ReleaseYear     int     `json:"release_year"`
	Upinfo          int     `json:"upinfo"`
	TotalNum        int     `json:"total_num"`
	VipType         string  `json:"vip_type"`
	IsEnd           int     `json:"isend"`
	Disable         int8    `json:"disable"`
	Source          int     `json:"source"`
	BillboardId     int     `json:"billboard_id"`
	BillboardRank   int     `json:"billboard_rank"`
	CornerType      string  `json:"corner_type"`
	DefinitionId    int     `json:"definition_id"`
	LanguageMediaId int     `json:"language_media_id"`
	Language        string  `json:"language"`
	BestvId         string  `json:"bestv_id"`
	BestvStatus     int8    `json:"bestv_status"`
	SelectionMode   int8    `json:"selection_mode"`
	Extend          string  `json:"extend"`
	RecordNum       string  `json:"record_num"`
	Creater         int     `json:"creater"`
	UpdateTime      int     `json:"update_time"`
	CreateTime      int     `json:"create_time"`
}

func (m *MediaModel) ToScanField(fields []string) []interface{} {
	var items []interface{}
	if len(fields) == 1 && fields[0] == "*" {
		items = []interface{}{
			&m.MediaId,
			&m.OriginalId,
			&m.ChannelId,
			&m.AlternativeId,
			&m.Name,
			&m.Pinyin,
			&m.Alias,
			&m.Aword,
			&m.Poster,
			&m.Still,
			&m.NameImg,
			&m.BgImg,
			&m.Director,
			&m.Actor,
			&m.StaffPinyin,
			&m.Description,
			&m.Score,
			&m.TotalVv,
			&m.Heat,
			&m.ReleaseDate,
			&m.ReleaseYear,
			&m.Upinfo,
			&m.TotalNum,
			&m.VipType,
			&m.IsEnd,
			&m.Disable,
			&m.Source,
			&m.BillboardId,
			&m.BillboardRank,
			&m.CornerType,
			&m.DefinitionId,
			&m.LanguageMediaId,
			&m.Language,
			&m.BestvId,
			&m.BestvStatus,
			&m.SelectionMode,
			&m.Extend,
			&m.RecordNum,
			&m.Creater,
			&m.UpdateTime,
			&m.CreateTime,
		}
		return items
	}
	//todo 字段映射
	return []interface{}{}
}
