package services

type TaskService struct {
	taskRepo TaskRepo
}

func CreateTaskService(taskRepo TaskRepo) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

func (s *TaskService) GetTasks() *[]GetTaskDto {
	tasks := s.taskRepo.GetTasks()
	return tasks
}

func (s *TaskService) GetTask(id int) *GetTaskDto {
	task := s.taskRepo.GetTask(id)

	return task
}

func (s *TaskService) CreateTask(taskDto *CreateTaskDto) *GetTaskDto {
	// save to db
	newTask := s.taskRepo.CreateTask(taskDto)

	return newTask
}
