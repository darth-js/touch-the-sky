package service

import (
	"fmt"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/ports"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/validation"
)

var ErrVideoNotFound = fmt.Errorf("video not found")
var ErrAnnotationsNotFound = fmt.Errorf("video not found")

type videoService struct {
	videoRepo       ports.VideoRepository
	annotationsRepo ports.AnnotationRepository
	userRepo        ports.UserRepository
}

func NewVideoService(
	videoRepo ports.VideoRepository,
	annotationsRepo ports.AnnotationRepository,
	userRepo ports.UserRepository) ports.VideoService {
	return &videoService{
		videoRepo:       videoRepo,
		annotationsRepo: annotationsRepo,
		userRepo:        userRepo,
	}
}

func (s *videoService) Create(username string, video *model.Video, annotaions []*model.Annotation) error {
	if err := s.validate(video, annotaions); err != nil {
		return err
	}

	var err error
	var userId int
	if _, err = s.userRepo.FindByUsername(username); err != nil {
		return err
	}

	var videoId int
	if videoId, err = s.videoRepo.Create(video, userId); err != nil {
		return err
	}

	for _, annotation := range annotaions {
		if err := s.annotationsRepo.Create(annotation, videoId, userId); err != nil {
			return err
		}
	}

	return nil
}

func (*videoService) validate(video *model.Video, annotaions []*model.Annotation) error {
	if err := validation.ValidateVideo(video); err != nil {
		return err
	}

	for _, annotation := range annotaions {
		if err := validation.ValidateAnnotation(annotation, video.Duration); err != nil {
			return err
		}
	}
	return nil
}

func (s *videoService) Find(videoId int) (*model.Video, []*model.Annotation, error) {

	video, err := s.videoRepo.FindById(videoId)
	if err != nil {
		return nil, nil, ErrVideoNotFound
	}

	annotations, err := s.annotationsRepo.FindVideoId(videoId)
	if err != nil {
		return nil, nil, ErrAnnotationsNotFound
	}
	return video, annotations, nil
}

func (s *videoService) Update(videoId int, video *model.Video, annotaions []*model.Annotation) error {
	if err := s.validate(video, annotaions); err != nil {
		return err
	}

	if err := s.videoRepo.Update(videoId, video); err != nil {
		return err
	}

	for _, annotation := range annotaions {
		if err := s.annotationsRepo.Update(annotation.ID, annotation); err != nil {
			return err
		}
	}

	return nil
}

func (s *videoService) Remove(id int) error {

	annotations, err := s.annotationsRepo.FindVideoId(id)
	if err != nil {
		return err
	}

	for _, annotation := range annotations {
		if err := s.annotationsRepo.Remove(annotation.ID); err != nil {
			return err
		}
	}

	if err := s.videoRepo.Remove(id); err != nil {
		return err
	}
	return nil
}
