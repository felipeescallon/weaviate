//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package vectorizer

import (
	"context"
	"testing"

	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/stretchr/testify/assert"
)

func TestConfigValidator(t *testing.T) {
	t.Run("validate class names", func(t *testing.T) {
		type testCase struct {
			input     string
			valid     bool
			name      string
			vectorize bool
		}

		// for all test cases keep in mind that the word "carrot" is not present in
		// the fake c11y, but every other word is.
		//
		// Additionally, the word "the" is a stopword
		//
		// all inputs represent class names (!)
		tests := []testCase{
			// valid names
			{
				name:      "Single uppercase word present in the c11y",
				input:     "Car",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "Single lowercase word present in the c11y, stored as uppercase",
				input:     "car",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "combination of valid words starting with uppercase letter",
				input:     "CarGarage",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "combination of valid words starting with lowercase letter, stored as uppercase",
				input:     "carGarage",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "combination of valid words and stopwords, starting with uppercase",
				input:     "TheCarGarage",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "combination of valid words and stopwords starting with lowercase letter, stored as uppercase",
				input:     "carTheGarage",
				valid:     true,
				vectorize: true,
			},

			// inavlid names
			{
				name:      "Single uppercase word NOT present in the c11y",
				input:     "Carrot",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "Single lowercase word NOT present in the c11y",
				input:     "carrot",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "Single uppercase stopword",
				input:     "The",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "Single lowercase stopword",
				input:     "the",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of valid and invalid words, valid word first lowercased",
				input:     "potatoCarrot",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of valid and invalid words, valid word first uppercased",
				input:     "PotatoCarrot",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of valid and invalid words, invalid word first lowercased",
				input:     "carrotPotato",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of valid and invalid words, invalid word first uppercased",
				input:     "CarrotPotato",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of only stopwords, starting with lowercase",
				input:     "theThe",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of only stopwords, starting with uppercase",
				input:     "TheThe",
				valid:     false,
				vectorize: true,
			},

			// vectorize turned off
			{
				name:      "non-vectorized: combination of only stopwords, starting with uppercase",
				input:     "TheThe",
				valid:     true,
				vectorize: false,
			},
			{
				name:      "non-vectorized: excluded word",
				input:     "carrot",
				valid:     true,
				vectorize: false,
			},
		}

		for _, test := range tests {
			t.Run(test.name+" object class", func(t *testing.T) {
				class := &models.Class{
					Class: test.input,
					Properties: []*models.Property{{
						Name:     "dummyPropSoWeDontRunIntoAllNoindexedError",
						DataType: []string{"string"},
					}},
				}

				v := NewConfigValidator(&fakeRemote{})
				err := v.Do(context.Background(), class, nil, &fakeIndexChecker{
					vectorizeClassName: test.vectorize,
					propertyIndexed:    true,
				})
				assert.Equal(t, test.valid, err == nil)

				// only proceed if input was supposed to be valid
				if test.valid == false {
					return
				}
			})
		}
	})

	t.Run("validate property names", func(t *testing.T) {
		type testCase struct {
			input     string
			valid     bool
			name      string
			vectorize bool
		}

		// for all test cases keep in mind that the word "carrot" is not present in
		// the fake c11y, but every other word is
		//
		// all inputs represent property names (!)
		tests := []testCase{
			// valid names
			{
				name:      "Single uppercase word present in the c11y, stored as lowercase",
				input:     "Brand",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "Single lowercase word present in the c11y",
				input:     "brand",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "combination of valid words starting with uppercase letter, stored as lowercase",
				input:     "BrandGarage",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "combination of valid words starting with lowercase letter",
				input:     "brandGarage",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "combination of valid words and stop words starting with uppercase letter, stored as lowercase",
				input:     "TheGarage",
				valid:     true,
				vectorize: true,
			},
			{
				name:      "combination of valid words and stop words starting with lowercase letter",
				input:     "theGarage",
				valid:     true,
				vectorize: true,
			},

			// inavlid names
			{
				name:      "Single uppercase word NOT present in the c11y",
				input:     "Carrot",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "Single lowercase word NOT present in the c11y",
				input:     "carrot",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "Single lowercase stop word",
				input:     "the",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of valid and invalid words, valid word first lowercased",
				input:     "potatoCarrot",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of valid and invalid words, valid word first uppercased",
				input:     "PotatoCarrot",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of valid and invalid words, invalid word first lowercased",
				input:     "carrotPotato",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of valid and invalid words, invalid word first uppercased",
				input:     "CarrotPotato",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of only stop words,  first lowercased",
				input:     "theThe",
				valid:     false,
				vectorize: true,
			},
			{
				name:      "combination of only stop words, first uppercased",
				input:     "TheThe",
				valid:     false,
				vectorize: true,
			},

			// without vectorizing
			{
				name:      "non-vectorizing: combination of only stop words, first uppercased",
				input:     "TheThe",
				valid:     true,
				vectorize: false,
			},
			{
				name:      "non-vectorizing: combination of only stop words, first uppercased",
				input:     "carrot",
				valid:     true,
				vectorize: false,
			},
		}

		for _, test := range tests {
			t.Run(test.name+" object class", func(t *testing.T) {
				class := &models.Class{
					Class: "ValidName",
					Properties: []*models.Property{{
						DataType: []string{"string"},
						Name:     test.input,
					}},
				}

				v := NewConfigValidator(&fakeRemote{})
				err := v.Do(context.Background(), class, nil, &fakeIndexChecker{
					vectorizePropertyName: test.vectorize,
					propertyIndexed:       true,
				})
				assert.Equal(t, test.valid, err == nil)
			})
		}
	})

	t.Run("all usable props no-indexed", func(t *testing.T) {
		t.Run("all schema vectorization turned off", func(t *testing.T) {
			class := &models.Class{
				Vectorizer: "text2vec-contextionary",
				Class:      "ValidName",
				Properties: []*models.Property{
					{
						DataType: []string{"text"},
						Name:     "decsription",
					},
					{
						DataType: []string{"string"},
						Name:     "name",
					},
					{
						DataType: []string{"int"},
						Name:     "amount",
					},
				},
			}

			v := NewConfigValidator(&fakeRemote{})
			err := v.Do(context.Background(), class, nil, &fakeIndexChecker{
				vectorizePropertyName: false,
				vectorizeClassName:    false,
				propertyIndexed:       false,
			})
			assert.NotNil(t, err)
		})
	})
}

type fakeIndexChecker struct {
	vectorizeClassName    bool
	vectorizePropertyName bool
	propertyIndexed       bool
}

func (f *fakeIndexChecker) VectorizeClassName() bool {
	return f.vectorizeClassName
}

func (f *fakeIndexChecker) VectorizePropertyName(propName string) bool {
	return f.vectorizePropertyName
}

func (f *fakeIndexChecker) PropertyIndexed(propName string) bool {
	return f.propertyIndexed
}

// Every word in this fake c11y remote client is present except for the word
// Carrot which is not present
type fakeRemote struct{}

func (f *fakeRemote) IsWordPresent(ctx context.Context, word string) (bool, error) {
	if word == "carrot" || word == "the" {
		return false, nil
	}
	return true, nil
}

func (f *fakeRemote) IsStopWord(ctx context.Context, word string) (bool, error) {
	return word == "the", nil
}
