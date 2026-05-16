package schema

// TagCondition represents a condition with an operation type, type, key, and values.
type TagCondition struct {
	OpType OpType   `json:"opType"`
	Type   string   `json:"type"`
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

const (
	TagConditionTypeHeader = "HEADER"
	TagConditionTypeQuery  = "QUERY"
	TagConditionTypeCookie = "COOKIE"
)

// OpType defines the operation types for a TagCondition.
type OpType string

const (
	OpType_EQUAL     OpType = "EQUAL"
	OpType_NOT_EQUAL OpType = "NOT_EQUAL"
	OpType_IN        OpType = "IN"
	OpType_NOT_IN    OpType = "NOT_IN"
	OpType_REGULAR   OpType = "REGULAR"
	OpType_PREFIX    OpType = "PREFIX"
)

// TagGroup represents a group of conditions with a relation type and order.
type TagGroup struct {
	RelationType RelationType   `json:"relationType"`
	Conditions   []TagCondition `json:"conditions"`
	Order        int32          `json:"order"`
}

// RelationType defines the relation types for a TagGroup.
type RelationType string

const (
	RelationType_AND RelationType = "AND"
	RelationType_OR  RelationType = "OR"
)

// TagDestination represents a destination with weight, relation type, conditions, and order.
type TagDestination struct {
	Weight       int32          `json:"weight"`
	RelationType RelationType   `json:"relationType"`
	Conditions   []TagCondition `json:"conditions"`
}

// TagRule represents a rule with destinations, relation type, conditions, and order.
type TagRule struct {
	TagGroup
	Destinations []TagDestination `json:"destinations"`
}

// Defining the SlidingWindow of `LimitForm` struct.
type SlidingWindow struct {
	Threshold      int   `json:"threshold"`
	TimeWindowInMs int64 `json:"timeWindowInMs"`
}

// ErrorParserPolicy represents an error parser policy with a parser and expression.
type ErrorParserPolicy struct {
	Parser     string `json:"parser"`
	Expression string `json:"expression"`
}

// DegradeConfig represents a degrade configuration.
type DegradeConfig struct {
	ResponseCode int               `json:"responseCode"`
	Attributes   map[string]string `json:"attributes"`
	ResponseBody string            `json:"responseBody"`
}

// RetryPolicy represents a retry policy configuration.
type RetryPolicy struct {
	Retry          int                `json:"retry"`
	Interval       int                `json:"interval"`
	Timeout        int                `json:"timeout"`
	ErrorCodes     []string           `json:"errorCodes"`
	CodePolicy     *ErrorParserPolicy `json:"codePolicy"`
	ErrorMessages  []string           `json:"errorMessages"`
	MessagePolicy  *ErrorParserPolicy `json:"messagePolicy"`
	Exceptions     []string           `json:"exceptions"`
	Methods        []string           `json:"methods"`
	MethodPrefixes []string           `json:"methodPrefixes"`
	Version        int64              `json:"version"`
}

// Define default values for policy configuration.
const (
	DefaultGroup                 = "default"
	DefaultLimitType             = "Rate"
	DefaultAgentLimitRealizeType = "Resilience4j"
)
