package response

import (
	"encoding/json"
	"io"
)

type SentimentAgreement string

const (
	Agreement    SentimentAgreement = "AGREEMENT"
	Disagreement SentimentAgreement = "DISAGREEMENT"
)

type ScoreTag string

const (
	ScoreTagStrongPositive ScoreTag = "P+"
	ScoreTagPositive       ScoreTag = "P"
	ScoreTagNeutral        ScoreTag = "NEU"
	ScoreTagNegative       ScoreTag = "N"
	ScoreTagStrongNegative ScoreTag = "N+"
	ScoreTagNone           ScoreTag = "NONE"
)

type Subjectivity string

const (
	SubjectivitySubjective Subjectivity = "SUBJECTIVE"
	SubjectivityObjective  Subjectivity = "OBJECTIVE"
)

type Irony string

const (
	IronyIrony    Irony = "IRONIC"
	IronyNonIrony Irony = "NONIRONIC"
)

type Sentiment struct {
	Response
	Model                  string               `json:"model"`
	ScoreTag               ScoreTag             `json:"score_tag"`
	Agreement              SentimentAgreement   `json:"agreement"`
	Subjectivity           Subjectivity         `json:"subjectivity"`
	Confidence             string               `json:"confidence"`
	Irony                  Irony                `json:"irony"`
	SentenceList           []Sentence           `json:"sentence_list"`
	SentimentedEntityList  []SentimentedEntity  `json:"sentimented_entity_list"`
	SentimentedConceptList []SentimentedConcept `json:"sentimented_concept_list"`
}

type Sentence struct {
	Text                   string               `json:"text"`
	Inip                   string               `json:"inip"`
	Endp                   string               `json:"endp"`
	Bop                    string               `json:"bop"`
	Confidence             string               `json:"confidence"`
	ScoreTag               ScoreTag             `json:"score_tag"`
	Agreement              SentimentAgreement   `json:"agreement"`
	SegmentList            []Segment            `json:"segment_list"`
	SentimentedEntityList  []SentimentedEntity  `json:"sentimented_entity_list"`
	SentimentedConceptList []SentimentedConcept `json:"sentimented_concept_list"`
}

type SegmentType string

const (
	SegmentTypeMain      SegmentType = "main"
	SegmentTypeSecondary SegmentType = "secondary"
)

type Segment struct {
	Text                   string               `json:"text"`
	SegmentType            SegmentType          `json:"segment_type"`
	Inip                   string               `json:"inip"`
	Endp                   string               `json:"endp"`
	Confidence             string               `json:"confidence"`
	ScoreTag               ScoreTag             `json:"score_tag"`
	Agreement              SentimentAgreement   `json:"agreement"`
	PolarityTermList       []PolarityTerm       `json:"polarity_term_list,omitempty"`
	SegmentList            []Segment            `json:"segment_list,omitempty"`
	SentimentedConceptList []SentimentedConcept `json:"sentimented_concept_list,omitempty"`
}

type PolarityTerm struct {
	Text                   string               `json:"text"`
	Inip                   string               `json:"inip"`
	Endp                   string               `json:"endp"`
	TagStack               string               `json:"tag_stack"`
	Confidence             string               `json:"confidence"`
	ScoreTag               ScoreTag             `json:"score_tag"`
	SentimentedEntityList  []SentimentedEntity  `json:"sentimented_entity_list,omitempty"`
	SentimentedConceptList []SentimentedConcept `json:"sentimented_concept_list,omitempty"`
}
type SentimentedConcept struct {
	Form     string   `json:"form"`
	ID       string   `json:"id"`
	Variant  string   `json:"variant"`
	Inip     string   `json:"inip"`
	Endp     string   `json:"endp"`
	Type     string   `json:"type"`
	ScoreTag ScoreTag `json:"score_tag"`
}

type SentimentedEntity struct {
	Form     string   `json:"form"`
	ID       string   `json:"id"`
	Variant  string   `json:"variant"`
	Inip     string   `json:"inip"`
	Endp     string   `json:"endp"`
	Type     string   `json:"type"`
	ScoreTag ScoreTag `json:"score_tag"`
}

func (s *Sentiment) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(s)
}
