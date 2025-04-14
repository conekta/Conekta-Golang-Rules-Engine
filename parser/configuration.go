package parser

type EvaluatorConfigOption func(*EvaluatorConfig)

type EvaluatorConfig struct {
	attrPathsPrefixNilToZeroValue []string
}

func WithAttrPathsPrefixNilToZeroValue(attrPathPrefix ...string) EvaluatorConfigOption {
	return func(e *EvaluatorConfig) {
		e.attrPathsPrefixNilToZeroValue = attrPathPrefix
	}
}
