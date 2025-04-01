// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ResourceRuleOrigin struct {
	ResourceMeta *Meta `json:"resourceMeta,omitempty"`
	// index of the to-item in the policy
	RuleIndex *int64 `json:"ruleIndex,omitempty"`
}

func (o *ResourceRuleOrigin) GetResourceMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.ResourceMeta
}

func (o *ResourceRuleOrigin) GetRuleIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.RuleIndex
}
