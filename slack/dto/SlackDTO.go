package dto

//https://api.slack.com/docs/message-attachments
type SlackDTO struct {
	Attachments []AttachmentDTO `json:"attachments,omitempty"`
}

type AttachmentDTO struct {
	Pretext        string                 `json:"pretext,omitempty"`
	Fallback       string                 `json:"fallback,omitempty"`
	Color          string                 `json:"color,omitempty"`
	AuthorName     string                 `json:"author_name,omitempty"`
	AuthorLink     string                 `json:"author_link,omitempty"`
	AuthorIcon     string                 `json:"author_icon,omitempty"`
	Title          string                 `json:"title,omitempty"`
	TitleLink      string                 `json:"title_link,omitempty"`
	Text           string                 `json:"text,omitempty"`
	AttachmentType string                 `json:"attachment_type,omitempty"`
	CallbackId     string                 `json:"callback_id,omitempty"`
	Fields         []AttachmentFieldsDTO  `json:"fields,omitempty"`
	Actions        []AttachmentActionsDTO `json:"actions,omitempty"`
	ImageURL       string                 `json:"image_url,omitempty"`
	ThumbURL       string                 `json:"thumb_url,omitempty"`
	Footer         string                 `json:"footer,omitempty"`
	FooterIcon     string                 `json:"footer_icon,omitempty"`
	Timestamp      int64                  `json:"ts,omitempty"`
}

type AttachmentFieldsDTO struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

type AttachmentActionsDTO struct {
	Name  string `json:"name,omitempty"`
	Text  string `json:"text,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	URL   string `json:"url,omitempty"`
}

/**

{
    "attachments": [
        {
        	"pretext": "Optional text that appears above the attachment block",
            "fallback": "Required plain-text summary of the attachment.",
            "color": "#2eb886",
            "author_name": "SillyHat Xu",
            "author_link": "http://sillyhatxu.com",
            "author_icon": "https://www.staffsprep.com/software/flat_faces_icons/png/flat_faces_icons_circle/flat-faces-icons-circle-4.png",
            "title": "Slack API Documentation",
            "title_link": "https://api.slack.com/",
            "text": "Optional text that appears within the attachment",
            "fields": [
                {
                    "title": "Priority",
                    "value": "High",
                    "short": true
                },
                {
                    "title": "test2",
                    "value": "value2",
                    "short": true
                },
                {
                    "title": "test3",
                    "value": "haha",
                    "short": true
                }
            ],
            "image_url": "http://my-website.com/path/to/image.jpg",
            "thumb_url": "http://example.com/path/to/thumb.png",
            "footer": "Slack API",
            "footer_icon": "https://platform.slack-edge.com/img/default_application_icon.png",
            "ts": 123456789
        }
    ]
}

*/
