package main

type Category struct {
	ID                 int    `json:"id"`
	ParentID           int    `json:"parent_id"`
	CategoryName       string `json:"category_name"`
	ParentCategoryName string `json:"parent_category_name"`
}

const cjson = `
{"1": {"id": 1, "parent_id": 0, "category_name": "ソファー"}, "2": {"id": 2, "parent_id": 1, "category_name": "一人掛けソファー", "parent_category_name": "ソファー"}, "3": {"id": 3, "parent_id": 1, "category_name": "二人掛けソファー", "parent_category_name": "ソファー"}, "4": {"id": 4, "parent_id": 1, "category_name": "コーナーソファー", "parent_category_name": "ソファー"}, "5": {"id": 5, "parent_id": 1, "category_name": "二段ソファー", "parent_category_name": "ソファー"}, "6": {"id": 6, "parent_id": 1, "category_name": "ソファーベッド", "parent_category_name": "ソファー"}, "10": {"id": 10, "parent_id": 0, "category_name": "家庭用チェア"}, "11": {"id": 11, "parent_id": 10, "category_name": "スツール", "parent_category_name": "家庭用チェア"}, "12": {"id": 12, "parent_id": 10, "category_name": "クッションスツール", "parent_category_name": "家庭用チェア"}, "13": {"id": 13, "parent_id": 10, "category_name": "ダイニングチェア", "parent_category_name": "家庭用チェア"}, "14": {"id": 14, "parent_id": 10, "category_name": "リビングチェア", "parent_category_name": "家庭用チェア"}, "15": {"id": 15, "parent_id": 10, "category_name": "カウンターチェア", "parent_category_name": "家庭用チェア"}, "20": {"id": 20, "parent_id": 0, "category_name": "キッズチェア"}, "21": {"id": 21, "parent_id": 20, "category_name": "学習チェア", "parent_category_name": "キッズチェア"}, "22": {"id": 22, "parent_id": 20, "category_name": "ベビーソファ", "parent_category_name": "キッズチェア"}, "23": {"id": 23, "parent_id": 20, "category_name": "キッズハイチェア", "parent_category_name": "キッズチェア"}, "24": {"id": 24, "parent_id": 20, "category_name": "テーブルチェア", "parent_category_name": "キッズチェア"}, "30": {"id": 30, "parent_id": 0, "category_name": "オフィスチェア"}, "31": {"id": 31, "parent_id": 30, "category_name": "デスクチェア", "parent_category_name": "オフィスチェア"}, "32": {"id": 32, "parent_id": 30, "category_name": "ビジネスチェア", "parent_category_name": "オフィスチェア"}, "33": {"id": 33, "parent_id": 30, "category_name": "回転チェア", "parent_category_name": "オフィスチェア"}, "34": {"id": 34, "parent_id": 30, "category_name": "リクライニングチェア", "parent_category_name": "オフィスチェア"}, "35": {"id": 35, "parent_id": 30, "category_name": "投擲用椅子", "parent_category_name": "オフィスチェア"}, "40": {"id": 40, "parent_id": 0, "category_name": "折りたたみ椅子"}, "41": {"id": 41, "parent_id": 40, "category_name": "パイプ椅子", "parent_category_name": "折りたたみ椅子"}, "42": {"id": 42, "parent_id": 40, "category_name": "木製折りたたみ椅子", "parent_category_name": "折りたたみ椅子"}, "43": {"id": 43, "parent_id": 40, "category_name": "キッチンチェア", "parent_category_name": "折りたたみ椅子"}, "44": {"id": 44, "parent_id": 40, "category_name": "アウトドアチェア", "parent_category_name": "折りたたみ椅子"}, "45": {"id": 45, "parent_id": 40, "category_name": "作業椅子", "parent_category_name": "折りたたみ椅子"}, "50": {"id": 50, "parent_id": 0, "category_name": "ベンチ"}, "51": {"id": 51, "parent_id": 50, "category_name": "一人掛けベンチ", "parent_category_name": "ベンチ"}, "52": {"id": 52, "parent_id": 50, "category_name": "二人掛けベンチ", "parent_category_name": "ベンチ"}, "53": {"id": 53, "parent_id": 50, "category_name": "アウトドア用ベンチ", "parent_category_name": "ベンチ"}, "54": {"id": 54, "parent_id": 50, "category_name": "収納付きベンチ", "parent_category_name": "ベンチ"}, "55": {"id": 55, "parent_id": 50, "category_name": "背もたれ付きベンチ", "parent_category_name": "ベンチ"}, "56": {"id": 56, "parent_id": 50, "category_name": "ベンチマーク", "parent_category_name": "ベンチ"}, "60": {"id": 60, "parent_id": 0, "category_name": "座椅子"}, "61": {"id": 61, "parent_id": 60, "category_name": "和風座椅子", "parent_category_name": "座椅子"}, "62": {"id": 62, "parent_id": 60, "category_name": "高座椅子", "parent_category_name": "座椅子"}, "63": {"id": 63, "parent_id": 60, "category_name": "ゲーミング座椅子", "parent_category_name": "座椅子"}, "64": {"id": 64, "parent_id": 60, "category_name": "ロッキングチェア", "parent_category_name": "座椅子"}, "65": {"id": 65, "parent_id": 60, "category_name": "座布団", "parent_category_name": "座椅子"}, "66": {"id": 66, "parent_id": 60, "category_name": "空気椅子", "parent_category_name": "座椅子"}}
`

