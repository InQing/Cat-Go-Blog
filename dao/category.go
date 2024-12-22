package dao

import (
	"log"
)

func GetCategoryNameById(cId int) string{
	row := DB.QueryRow("select name from blog_category where cid=?",cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}