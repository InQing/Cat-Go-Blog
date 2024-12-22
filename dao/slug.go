package dao

import "Go-Blog/models"

func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where slug = ? limit ?,?", slug, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CountGetAllPostBySlug(slug string) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where slug=?", slug)
	_ = rows.Scan(&count)
	return
}