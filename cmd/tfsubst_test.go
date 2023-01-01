package cmd

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTfSubst(t *testing.T) {
	tests := map[string]struct {
		stateLoc  string
		inputData string
		funcName  string
		expected  string
		error     bool
	}{
		"normal": {
			stateLoc:  "./testdata/terraform.tfstate",
			inputData: `image: {{ tfstate "docker_image.ubuntu.name" }}`,
			funcName:  "tfstate",
			expected:  "image: ubuntu:latest",
		},
		"different func name": {
			stateLoc:  "./testdata/terraform.tfstate",
			inputData: `image: {{ tfstate2 "docker_image.ubuntu.name" }}`,
			funcName:  "tfstate2",
			expected:  "image: ubuntu:latest",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := &tfsubst{}
			in := strings.NewReader(tt.inputData)

			var out bytes.Buffer
			err := c.execute(context.Background(), tt.stateLoc, in, &out, tt.funcName)

			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, out.String())
			}
		})
	}
}
