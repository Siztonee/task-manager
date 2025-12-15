package dto

type CreateTaskRequest struct {
	Title  string `json:"title" binding:"required,max=255"`
	UserID uint   `json:"user_id" binding:"required"`
}

type UpdateTaskRequest struct {
	Title     *string `json:"title" binding:"omitempty,required,max=255"`
	Completed *bool   `json:"completed"`
}
