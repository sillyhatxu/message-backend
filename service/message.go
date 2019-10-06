package service

import (
	"github.com/sillyhatxu/message-backend/grpc/message"
	"github.com/sillyhatxu/message-backend/slack/dto"
)

func SendMessage(request *message.AttachmentRequest) error {
	//gotItAction := dto.AttachmentActionsDTO{
	//	Name:  "ACTION_001",
	//	Text:  "I got it",
	//	Type:  constants.AttachmentActionButton,
	//	Value: "0",
	//}
	//oneHourAction := dto.AttachmentActionsDTO{
	//	Name:  "ACTION_002",
	//	Text:  "An hour later",
	//	Type:  constants.AttachmentActionButton,
	//	Value: "3600000",
	//}
	//tomorrowAction := dto.AttachmentActionsDTO{
	//	Name:  "ACTION_003",
	//	Text:  "Remind me tomorrow",
	//	Type:  constants.AttachmentActionButton,
	//	Value: "86400000",
	//}
	//attachment := dto.AttachmentDTO{
	//	Fallback:       request.Fallback,
	//	Text:           request.Text,
	//	Color:          constants.Pink,
	//	AuthorName:     request.AuthorName,
	//	AuthorIcon:     constants.AuthorIconBirthday,
	//	AttachmentType: constants.AttachmentType,
	//	CallbackId:     constants.CallbackIdBirthday,
	//	Actions:        []dto.AttachmentActionsDTO{gotItAction, oneHourAction, tomorrowAction},
	//	Timestamp:      time.Now().Unix(),
	//}
	attachment := dto.AttachmentDTO{
		Pretext:        request.Pretext,
		Fallback:       request.Fallback,
		Color:          request.Color,
		AuthorName:     request.AuthorName,
		AuthorLink:     request.AuthorLink,
		AuthorIcon:     request.AuthorIcon,
		Title:          request.Title,
		TitleLink:      request.TitleLink,
		Text:           request.Text,
		AttachmentType: request.AttachmentType,
		CallbackId:     request.CallbackId,
		Fields:         getAttachmentFields(request.Fields),
		Actions:        getAttachmentActions(request.Actions),
		ImageURL:       request.ImageURL,
		ThumbURL:       request.ThumbURL,
		Footer:         request.Footer,
		FooterIcon:     request.FooterIcon,
		Timestamp:      request.Timestamp.Seconds,
	}
	return birthdaySlackClient.Send(dto.SlackDTO{Attachments: []dto.AttachmentDTO{attachment}})
}

func getAttachmentFields(attachmentFields []*message.AttachmentField) []dto.AttachmentFieldsDTO {
	if attachmentFields == nil {
		return nil
	}
	var array []dto.AttachmentFieldsDTO
	for _, attachmentField := range attachmentFields {
		array = append(array, dto.AttachmentFieldsDTO{
			Title: attachmentField.Title,
			Value: attachmentField.Value,
			Short: attachmentField.Short,
		})
	}
	return array
}

func getAttachmentActions(actions []*message.AttachmentAction) []dto.AttachmentActionsDTO {
	if actions == nil {
		return nil
	}
	var array []dto.AttachmentActionsDTO
	for _, action := range actions {
		array = append(array, dto.AttachmentActionsDTO{
			Name:  action.Name,
			Text:  action.Text,
			Type:  action.Type,
			Value: action.Value,
			//URL:   action.URL,
		})
	}
	return array
}
