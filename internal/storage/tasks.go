package storage

import (
	"github.com/LigeronAhill/planify/internal/enums/status"
	"github.com/LigeronAhill/planify/internal/models"
)

func (r *Repository) InsertTask(task *models.Task) (uint, error) {
	res := r.db.Create(&task)
	if res.Error != nil {
		return 0, res.Error
	}
	return task.ID, nil
}

func (r *Repository) UpdateTask(update *models.Task) (*models.Task, error) {
	res := r.db.Save(&update)
	if res.Error != nil {
		return nil, res.Error
	}
	return update, nil
}

func (r *Repository) GetTask(taskID uint) (*models.Task, error) {
	var result *models.Task
	err := r.db.First(&result, taskID).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) GetNotCompletedTasksByAuthor(authorID uint) ([]*models.Task, error) {
	var result []*models.Task
	res := r.db.Where("NOT status = ? AND author_id = ?", status.DONE, authorID).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

func (r *Repository) GetNotCompletedTasksByExecutor(executorID uint) ([]*models.Task, error) {
	var result []*models.Task
	res := r.db.Where("NOT status = ? AND executor_id = ?", status.DONE, executorID).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

func (r *Repository) GetCompletedTasksByAuthor(authorID uint) ([]*models.Task, error) {
	var result []*models.Task
	res := r.db.Where("status = ? AND author_id = ?", status.DONE, authorID).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

func (r *Repository) GetCompletedTasksByExecutor(executorID uint) ([]*models.Task, error) {
	var result []*models.Task
	res := r.db.Where("status = ? AND executor_id = ?", status.DONE, executorID).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

func (r *Repository) DeleteTask(taskID uint) error {
	err := r.db.Delete(&models.Task{}, taskID).Error
	if err != nil {
		return err
	}
	return nil
}
