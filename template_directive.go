package alexa

const DisplayTemplateDirectiveType = "Display.RenderTemplate"

type displayTemplateDirective struct {
	Type            string          `json:"type"`
	DisplayTemplate displayTemplate `json:"template"`
}

// Make this a directive
func (d displayTemplateDirective) DirectiveType() string {
	return d.Type
}

// DisplayTemplate
// https://developer.amazon.com/docs/custom-skills/display-interface-reference.html#display-template-elements
type displayTemplate struct {
	Type            string                      `json:"type"`
	Token           string                      `json:"token,omitempty"`
	BackButton      string                      `json:"backButton"` // HIDDEN | VISIBLE
	BackgroundImage *displayTemplateImage       `json:"backgroundImage,omitempty"`
	Title           *string                     `json:"title,omitempty"`
	TextContent     *displayTemplateTextContent `json:"textContent,omitempty"`
	Image           *displayTemplateImage       `json:"image,omitempty"`
}

type displayTemplateImage struct {
	ContentDescription string                       `json:"contentDescription"`
	Sources            []displayTemplateImageSource `json:"sources"`
}

type displayTemplateImageSource struct {
	URL          string  `json:"url"`
	Size         *string `json:"size,omitempty"`
	WidthPixels  int     `json:"widthPixels"`
	HeightPixels int     `json:"heightPixels"`
}

type displayTemplateTextContent struct {
	PrimaryText   displayTemplateText `json:"primaryText"`
	SecondaryText displayTemplateText `json:"secondaryText"`
	TertiaryText  displayTemplateText `json:"tertiaryText"`
}

type displayTemplateText struct {
	Text string `json:"text"`
	Type string `json:"type"` // PlainText | RichText
}

// ===== Body Template 7 =====

// https://developer.amazon.com/docs/custom-skills/display-interface-reference.html#bodytemplate7
type bodyTemplate7 struct {
	Token              string
	BackButtonVisible  bool
	Title              string
	BackgroundImageURL string
	BackgroundImageAlt string
	ForegroundImageURL string
	ForegroundImageAlt string
}

func (b bodyTemplate7) asDirective() directive {
	visibleString := "HIDDEN"
	if b.BackButtonVisible {
		visibleString = "VISIBLE"
	}

	directive := displayTemplateDirective{
		Type: DisplayTemplateDirectiveType,
		DisplayTemplate: displayTemplate{
			Type:       "BodyTemplate7",
			Token:      b.Token,
			BackButton: visibleString,
			Title:      &b.Title,
		},
	}

	// Background Image
	if b.BackgroundImageURL != "" {
		directive.DisplayTemplate.BackgroundImage = &displayTemplateImage{
			ContentDescription: b.BackgroundImageAlt,
			Sources: []displayTemplateImageSource{
				{
					URL: b.BackgroundImageURL,
				},
			},
		}
	}

	// Foreground Image
	if b.ForegroundImageURL != "" {
		directive.DisplayTemplate.Image = &displayTemplateImage{
			ContentDescription: b.ForegroundImageAlt,
			Sources: []displayTemplateImageSource{
				{
					URL: b.ForegroundImageURL,
				},
			},
		}
	}

	return directive
}
