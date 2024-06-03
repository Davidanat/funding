package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	SaveCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindbyUser(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, err
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) SaveCampaign(input CreateCampaignInput) (Campaign, error) {
	slugName := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign := Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		UserID:           input.User.ID,
		Slug:             slug.Make(slugName),
	}

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != input.User.ID {
		return campaign, errors.New("user not allowed")
	}

	slugName := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.Slug = slug.Make(slugName)

	newCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return CampaignImage{}, errors.New("user not allowed")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkImageNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{
		CampaignID: input.CampaignID,
		IsPrimary:  isPrimary,
		FileName:   fileLocation,
	}

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}
