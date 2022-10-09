package response

import (
	"encoding/json"
	"io"
)

type Parser struct {
	Response
	TokenList       []Token         `json:"token_list"`
	GlobalSentiment GlobalSentiment `json:"global_sentiment"`
}

type Token struct {
	Type                      string                  `json:"type"`
	Form                      string                  `json:"form"`
	NormalizedForm            string                  `json:"normalized_form"`
	ID                        string                  `json:"id"`
	Inip                      string                  `json:"inip"`
	Endp                      string                  `json:"endp"`
	Style                     Style                   `json:"style"`
	Separation                string                  `json:"separation"`
	QuoteLevel                string                  `json:"quote_level"`
	AffectedByNegotiation     string                  `json:"affected_by_negotiation"`
	Head                      string                  `json:"head"`
	SyntacticTreeRelationList []SyntacticTreeRelation `json:"syntactic_tree_relation_list"`
	AnalysisList              []Analysis              `json:"analysis_list"`
	SenseList                 []Sense                 `json:"sense_list"`
	Sentiment                 TokenSentiment          `json:"sentiment"`
	TopicList                 TopicsExtraction        `json:"topic_list"` // Uh, this is not a list!
	TokenList                 []Token                 `json:"token_list"`
}

type Style struct {
	IsBold       string `json:"is_bold"`
	IsItalics    string `json:"is_italics"`
	IsUnderlined string `json:"is_underlined"`
	IsTitle      string `json:"is_title"`
}

type SyntacticTreeRelation struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Analysis struct {
	Origin                string    `json:"origin"`
	VarietyDictionary     string    `json:"variety_dictionary"`
	Tag                   string    `json:"tag"`
	Lemma                 string    `json:"lemma"`
	OriginalForm          string    `json:"original_form"`
	TagInfo               string    `json:"tag_info"`
	VarietyDictionaryInfo string    `json:"variety_dictionary_info"`
	CheckInfo             string    `json:"check_info"`
	Remission             string    `json:"remission"`
	SenseIDList           []SenseID `json:"sense_id_list"`
}

type SenseID struct {
	SenseID string `json:"sense_id"`
}

type Sense struct {
	ID           string `json:"id"`
	Info         string `json:"info"`
	Form         string `json:"form"`
	OfficialForm string `json:"official_form"`
}

type TokenSentiment struct {
	SelfSentiment      TokenSelfSentiment      `json:"self_sentiment"`
	InheritedSentiment TokenInheritedSentiment `json:"inherited_sentiment"`
}

type TokenSelfSentiment struct {
	Text       string `json:"text"`
	Inip       string `json:"inip"`
	Endp       string `json:"endp"`
	TagStack   string `json:"tag_stack"`
	Confidence string `json:"confidence"`
	ScoreTag   string `json:"score_tag"`
}

type TokenInheritedSentiment struct {
	RelationList []Relation `json:"relation_list"`
	ScoreTag     string     `json:"score_tag"`
}

type GlobalSentiment struct {
	Model        string `json:"model"`
	ScoreTag     string `json:"score_tag"`
	Agreement    string `json:"agreement"`
	Subjectivity string `json:"subjectivity"`
	Confidence   string `json:"confidence"`
	Irony        string `json:"irony"`
}

func (p *Parser) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(p)
}
