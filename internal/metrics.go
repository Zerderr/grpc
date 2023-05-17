package internal

import "github.com/prometheus/client_golang/prometheus"

var CreateStudentCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "new_student",
	Help: "New student was created",
})

var GetStudentCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "get_student",
	Help: "Student was received",
})

func init() {
	prometheus.MustRegister(CreateStudentCounter)
	prometheus.MustRegister(GetStudentCounter)
}
