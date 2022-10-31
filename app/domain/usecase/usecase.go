package usecase

type CourseUsecaseInterface interface {
	// AddCourse()
	// EditCourse()
	// EdotCourseByID(id int64) error
	// DeleteCourse() error
	// DeleteCoursebyID(id int64) error
}

type ProjectUseCaseInterface interface {
	// AddNewProject()
	// UpdateProject()
	// GetProjects()
	// GetProjectByID()
	// DeleteProject()
}

type OperationLogUseCaseInterface interface {
	// Get()
	// Post()
	// Delete()
}

type OperationUseCaseInterface interface {
	// GetAll()
	// GetByID()
	// Post()
	// Put()
	// Delete()
}

type OperationTypeUseCaseInterface interface {
	// GetAll()
	// GetByID()
	// Post()
	// Put()
	// Delete()
}

type ContactTypeUseCaseInterface interface {
	// GetAll()
	// GetByID()
	// Post()
	// Put()
	// Delete()
}
