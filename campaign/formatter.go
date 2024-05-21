package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"title"`
	ShortDescription string `json:"short_description"`
	CampaignImage    string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserID           int    `json:"user_id"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	images := ""
	if len(campaign.CampaignImages) > 0 {
		images = campaign.CampaignImages[0].FileName
	}

	formatter := CampaignFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		CampaignImage:    images,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormater := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormater := FormatCampaign(campaign)
		campaignsFormater = append(campaignsFormater, campaignFormater)
	}

	return campaignsFormater
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"title"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	CampaignImage    string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	image := ""
	if len(campaign.CampaignImages) > 0 {
		image = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	formatUser := CampaignUserFormatter{
		Name:     campaign.User.Name,
		ImageURL: campaign.User.Avatar,
	}

	var images []CampaignImageFormatter

	for _, i := range campaign.CampaignImages {
		image := CampaignImageFormatter{
			ImageURL:  i.FileName,
			IsPrimary: i.IsPrimary,
		}
		images = append(images, image)
	}

	formatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		CampaignImage:    image,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
		Perks:            perks,
		User:             formatUser,
		Images:           images,
	}

	return formatter
}
