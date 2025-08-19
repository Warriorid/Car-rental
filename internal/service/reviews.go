package service

import (
	"CarRentalService/internal/models"
	"CarRentalService/internal/repository"
	"errors"
)

var (
	ErrReviewPermission = errors.New("You have not permission")
	ErrReviewAlreadySent = errors.New("The review has already been sent")
)

type ReviewServise struct {
	repo repository.Reviews
	rentalRepo repository.RentalCheck 

}

func NewReviewServise(repo repository.Reviews, rentalRepo repository.RentalCheck) *ReviewServise {
	return &ReviewServise{repo: repo, rentalRepo: rentalRepo}
}

func (s *ReviewServise) LeaveReview(userId int, input models.Reviews) (int, error) {
	rental, err := s.rentalRepo.GetRentalInfo(input.RentalId)
	if err != nil {
		return 0, err
	}
	if rental.UserId != userId {
		return 0, ErrReviewPermission
	}
	isReviewExists, err := s.repo.ReviewExists(input.RentalId)
	if err != nil {
		return 0, err
	}
	if isReviewExists {
		return 0, ErrReviewAlreadySent
	}
	return s.repo.LeaveReview(input)
}