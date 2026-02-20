package postgres

import "go-blog/internal/model"

func (s *storage) GetArticles() ([]*model.Article, error) {
	var articles []*model.Article
	if err := s.db.Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *storage) GetArticle(id int) (*model.Article, error) {
	var article model.Article
	if err := s.db.First(&article, id).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

func (s *storage) CreateArticle(newArticle *model.Article) error {
	return s.db.Create(&newArticle).Error
}

func (s *storage) UpdateArticle(id int, updateArticle *model.Article) error {
	var article model.Article
	if err := s.db.First(&article, id).Error; err != nil {
		return err
	}

	if err := s.db.Model(&article).Updates(updateArticle).Error; err != nil {
		return err
	}

	return nil
}

func (s *storage) DeleteArticle(id int) error {
	var article model.Article
	if err := s.db.First(&article, id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&article).Error; err != nil {
		return err
	}

	return nil
}
