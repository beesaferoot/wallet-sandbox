package tests

import (
	"reflect"
	"testing"
	"wallet-sandbox/utils"
)



func TestTokenize(t *testing.T){
	for _, tc := range []struct {
		Name        string 
		srcString 	string
		expectedTokenMap  map[string]string
	}{
		{
			Name: "Test one variable",
			srcString: `
			VAR=VAR_2
			`,
			expectedTokenMap: map[string]string{"VAR": "VAR_2"},
		},
		{
			Name: "Test empty variable",
			srcString: ``,
			expectedTokenMap: map[string]string{},
		},{
			Name: "Test comment",
			srcString: `
			# this is a comment
			`,
			expectedTokenMap: map[string]string{},
		},{
			Name: "Test two variables",
			srcString: `
			DATABASE_HOST=value
			DATABASE_PORT=value
			`,
			expectedTokenMap: map[string]string{"DATABASE_HOST": "value", "DATABASE_PORT": "value"},
		},{
			Name: "Test empty values",
			srcString: `
			DATABASE_HOST=
			DATABASE_PORT=
			`,
			expectedTokenMap: map[string]string{"DATABASE_HOST": "", "DATABASE_PORT": ""},
		},
		}{
			t.Run(tc.Name, func(t *testing.T) {
				kvPair, err := utils.Tokenize([]byte(tc.srcString))
				if err != nil {
					t.Errorf("%v", err.Error())
					return 
				}
				if !reflect.DeepEqual(tc.expectedTokenMap, kvPair) {
					t.Errorf("expected %v as value, got %v instead", tc.expectedTokenMap, kvPair)
				}
			})
		}

}