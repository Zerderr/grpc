package server

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"homework-5/internal"
	"homework-5/internal/pkg/pb"
	"homework-5/internal/pkg/repository"
)

type Implementation struct {
	pb.UnimplementedServiceServer
	studentRepo    repository.StudentRepo
	universityRepo repository.UniversityRepo
}

func NewImplementation(sr repository.StudentRepo, ur repository.UniversityRepo) *Implementation {
	return &Implementation{
		studentRepo:    sr,
		universityRepo: ur,
	}
}
func (i *Implementation) ListStudent(ctx context.Context, studentRequest *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {
	tr := otel.Tracer("DeleteTodo")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String("studentRequest.String())"))
	defer span.End()
	post, err := i.studentRepo.GetById(ctx, studentRequest.Id)
	if err != nil {
		return nil, err
	}
	internal.GetStudentCounter.Add(1)
	internal.Tracing(ctx, &internal.TracingOpts{
		Name:       "Student Get",
		SpanName:   "Student Get",
		Attributes: studentRequest.String(),
	})
	return &pb.GetStudentResponse{
		Id:     post.ID,
		Name:   post.Name,
		Grades: int32(post.Grades),
		UnivId: post.UnivID,
	}, nil
}
func (i *Implementation) CreateUniversity(ctx context.Context, university *pb.CreateUniversityRequest) (*pb.CreateUniversityResponse, error) {
	id, err := i.universityRepo.Add(ctx, &repository.University{
		Name:     university.Name,
		Facility: university.Facility,
	})
	tr := otel.Tracer("DeleteStudent")
	ctx, span := tr.Start(ctx, "JFDKLFJDLKFJLDKFJLDKF request")
	span.SetAttributes(attribute.Key("params").String(university.String()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateUniversityResponse{Id: uint64(id)}, nil
}

func (i *Implementation) ListUniversity(ctx context.Context, universityRequest *pb.GetUniversityRequest) (*pb.GetUniversityResponse, error) {
	tr := otel.Tracer("DeleteTodo")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String("BLYAD"))
	defer span.End()

	university, err := i.universityRepo.GetById(ctx, universityRequest.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUniversityResponse{
		Id:       university.ID,
		Name:     university.Name,
		Facility: university.Facility,
	}, nil
}

func (i *Implementation) UpdateUniversity(ctx context.Context, universityRequest *pb.UpdateUniversityRequest) (*pb.UpdateUniversityResponse, error) {
	updated, err := i.universityRepo.Update(ctx, &repository.University{
		ID:       universityRequest.Id,
		Facility: universityRequest.Facility,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUniversityResponse{Ok: updated}, nil
}

func (i *Implementation) DeleteUniversity(ctx context.Context, universityRequest *pb.DeleteUniversityRequest) (*pb.DeleteUniversityResponse, error) {
	deleted, err := i.universityRepo.Delete(ctx, universityRequest.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUniversityResponse{Ok: deleted}, nil
}

func (i *Implementation) CreateStudent(ctx context.Context, studentRequest *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {
	id, err := i.studentRepo.Add(ctx, &repository.Student{
		Name:   studentRequest.Name,
		Grades: int16(studentRequest.Grades),
		UnivID: studentRequest.UnivId,
	})
	internal.Tracing(ctx, &internal.TracingOpts{
		Name:       "Student Create",
		SpanName:   "Student Create",
		Attributes: studentRequest.String(),
	})
	internal.CreateStudentCounter.Add(1)
	if err != nil {
		return nil, err
	}
	internal.CreateStudentCounter.Add(1)
	return &pb.CreateStudentResponse{Id: id}, nil
}

func (i *Implementation) UpdateStudent(ctx context.Context, studentRequest *pb.UpdateStudentRequest) (*pb.UpdateStudentResponse, error) {
	updated, err := i.studentRepo.Update(ctx, &repository.Student{
		ID:   studentRequest.Id,
		Name: studentRequest.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateStudentResponse{Ok: updated}, nil
}

func (i *Implementation) DeleteStudent(ctx context.Context, studentRequest *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	tr := otel.Tracer("DeleteStudent")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(studentRequest.String()))
	defer span.End()
	deleted, err := i.studentRepo.Delete(ctx, studentRequest.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteStudentResponse{Ok: deleted}, nil
}
