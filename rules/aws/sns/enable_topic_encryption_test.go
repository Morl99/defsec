package sns

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/aws/sns"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnableTopicEncryption(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    sns.SNS
		expected bool
	}{
		{
			name:     "positive result",
			input:    sns.SNS{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    sns.SNS{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.SNS = test.input
			results := CheckEnableTopicEncryption.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Rule().LongID() == CheckEnableTopicEncryption.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}