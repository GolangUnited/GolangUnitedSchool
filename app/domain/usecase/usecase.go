package usecase

type CourseUseCaseInterface interface {
	// AddCourse()
	// EditCourse()
	// EditCourseById(id int64) error
	// DeleteCourse() error
	// DeleteCourseById(id int64) error
}
type PersonUseCaseInterface interface {
	// GetPersonByField() (Person, error)
	// GetPersonById(int64) (Person, error)
	// AddNewPerson() error
	// UpdatePersonById (int64) error
	// RemovePersonById (int64) error
}
type StudentUseCaseInterface interface {
	//AddStudent
	//DeleteStudent
	//DeleteStudentByStudentId
}
type MentorUseCaseInterface interface {
}

type MentorNoteUseCaseInterface interface{}
type StudentNoteUseCaseInterface interface{}
type CourseStatusUseCaseInterface interface{}
type StudentGroupUseCaseInterface interface{}
type CourseStatusUseVaseInterface interface{}
