package campaign

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
