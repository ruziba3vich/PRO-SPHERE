package handlers

import (
	"github.com/projects/pro-sphere-backend/admin/genproto/genproto/feeds"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
)

// Convert internal FeedCategory to protobuf FeedCategory
func ConvertFeedCategoryToProto(category *models.FeedCategory) *feeds.FeedCategory {
	// Convert translations to protobuf format
	var translations []*feeds.Translation
	for _, translation := range category.Translations {
		translations = append(translations, &feeds.Translation{
			Lang: translation.Lang,
			Name: translation.Name,
		})
	}

	// Return the converted FeedCategory protobuf
	return &feeds.FeedCategory{
		Id:           int32(category.ID),
		IconUrl:      category.IconURL,
		IconId:       category.IconID,
		Translations: translations,
	}
}

// Convert protobuf FeedCategory to internal FeedCategory
func ConvertProtoToFeedCategory(protoCategory *feeds.FeedCategory) *models.FeedCategory {
	// Convert translations to internal format
	translations := ConvertProtoToFeedCategoryTranslation(protoCategory.Translations)
	// Return the converted FeedCategory
	return &models.FeedCategory{
		ID:           int64(protoCategory.Id),
		IconURL:      protoCategory.IconUrl,
		IconID:       protoCategory.IconId,
		Translations: translations,
	}
}

func ConvertProtoToFeedCategoryTranslation(protoTranslations []*feeds.Translation) []models.FeedCategoryTranslation {
	var translations []models.FeedCategoryTranslation
	for _, t := range protoTranslations {
		translations = append(translations, models.FeedCategoryTranslation{
			Lang: t.Lang,
			Name: t.Name,
		})
	}
	return translations
}
