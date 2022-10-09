package response

import (
	"io"
)

type TopicsExtraction struct {
	Response
	EntityList          []Entity          `json:"entity_list"`
	ConceptList         []Concept         `json:"concept_list"`
	TimeExpressionList  []TimeExpression  `json:"time_expression_list"`
	MoneyExpressionList []MoneyExpression `json:"money_expression_list"`
	// QuantityExpressionList []QuantityExpression `json:"quantity_expression_list"` beta -> disabled
	OtherExpressionList []OtherExpression `json:"other_expression_list"`
	QuotationList       []Quotation       `json:"quotation_list"`
	RelationList        []Relation        `json:"relation_list"`
}

type Entity struct {
	Form          string     `json:"form"`
	OfficialForm  string     `json:"official_form"`
	Dictionary    string     `json:"dictionary"`
	ID            string     `json:"id"`
	Sementity     Sementity  `json:"sementity"`
	SemgeoList    []Semgeo   `json:"semgeo_list"`
	SemldList     []string   `json:"semld_list"`
	SemreferList  []Semref   `json:"semrefer_list"`
	SemthemeList  []SemTheme `json:"semtheme_list"`
	StandardList  []Standard `json:"standard_list"`
	VariantList   []Variant  `json:"variant_list"`
	Relevance     string     `json:"relevance"`
	SubentityList []Entity   `json:"subentity_list"`
}

type Sementity struct {
	Class      string `json:"class"`
	Fiction    string `json:"fiction"`
	ID         string `json:"id"`
	Type       string `json:"type"`
	Confidence string `json:"confidence"`
}

type Semgeo struct {
	Continent Continent `json:"continent,omitempty"`
	Country   Country   `json:"country,omitempty"`
	Adm2      Adm       `json:"adm2,omitempty"`
	Adm1      Adm       `json:"adm1,omitempty"`
	Adm3      Adm       `json:"adm3,omitempty"`
	// City      *Semgeo   `json:"city,omitempty"` // Sorry, no docs about theses fields
	// District  *Semgeo   `json:"district,omitempty"`
}

type Adm struct {
	ID   string `json:"id"`
	Form string `json:"form"`
}

type Continent struct {
	ID   string `json:"id"`
	Form string `json:"form"`
}

type Country struct {
	Form         string     `json:"form"`
	ID           string     `json:"id"`
	StandardList []Standard `json:"standard_list"`
}

type Concept struct {
	Form          string     `json:"form"`
	OfficialForm  string     `json:"official_form"`
	Dictionary    string     `json:"dictionary"`
	ID            string     `json:"id"`
	Sementity     Sementity  `json:"sementity"`
	SemgeoList    []Semgeo   `json:"semgeo_list"`
	SemldList     []string   `json:"semld_list"`
	SemreferList  []string   `json:"semrefer_list"`
	SemthemeList  []Semtheme `json:"semtheme_list"`
	StandardList  []Standard `json:"standard_list"`
	VariantList   []Variant  `json:"variant_list"`
	Relevance     string     `json:"relevance"`
	SubentityList []string   `json:"subentity_list"`
}

type TimeExpression struct {
	Form           string `json:"form"`
	NormalizedForm string `json:"normalized_form"`
	ActualTime     string `json:"actual_time"`
	Precision      string `json:"precision"`
	Inip           string `json:"inip"`
	Endp           string `json:"endp"`
}

type MoneyExpression struct {
	Form         string `json:"form"`
	AmountForm   string `json:"amount_form"`
	NumericValue string `json:"numeric_value"`
	Currency     string `json:"currency"`
	Inip         string `json:"inip"`
	Endp         string `json:"endp"`
}

type OtherExpression struct {
	Form string `json:"form"`
	Type string `json:"type"`
	Inip string `json:"inip"`
	Endp string `json:"endp"`
}

type Quotation struct {
	Form string `json:"form"`
	Who  string `json:"who"`
	Verb string `json:"verb"`
	Inip string `json:"inip"`
	Endp string `json:"endp"`
}

type Relation struct {
	Form           string       `json:"form"`
	Inip           string       `json:"inip"`
	Endp           string       `json:"endp"`
	Subject        Subject      `json:"subject"`
	Verb           Verb         `json:"verb"`
	ComplementList []Complement `json:"complement_list"`
	Degree         string       `json:"degree"`
}

type Subject struct {
	Form        string   `json:"form"`
	LemmaList   []string `json:"lemma_list"`
	SenseIDList []string `json:"sense_id_list"`
}

type Verb struct {
	Form              string   `json:"form"`
	LemmaList         []string `json:"lemma_list"`
	SenseIDList       []string `json:"sense_id_list"`
	SemanticLemmaList []string `json:"semantic_lemma_list"`
}

type Complement struct {
	Form string `json:"form"`
	Type string `json:"type"`
}

type Variant struct {
	Form string `json:"form"`
	Inip string `json:"inip"`
	Endp string `json:"endp"`
}

type SemTheme struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Semref struct {
	Organization Organization `json:"organization"`
	Affinity     Affinity     `json:"affinity"`
}

type Organization struct {
	Form string `json:"form"`
	ID   string `json:"id"`
}

type Affinity struct {
	Form string `json:"form"`
	ID   string `json:"id"`
}

type Standard struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type Semtheme struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (t *TopicsExtraction) Decode(body io.ReadCloser) error {
	return decode(body, t)
}
