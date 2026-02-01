package course

import (
	"reflect"
	"testing"

	"github.com/goccy/go-yaml"
)

func TestCourseParsedFromYaml(t *testing.T) {

	tests := []struct {
		yamlSource string
		want       ActivitySpec
	}{
		{
			yamlSource: `
title: HW1
max-score: 20
uploads:
  files:
  - name: "*.c"
    required: true
    max-matches: 10
    max-size: 10MB
grader:
  rubric:
    items:
    - name: Some Name
      description: Some description
      max-score: 10
    - name: Some Name 2
      description: Some description 2
      max-score: 10`,
			want: ActivitySpec{
				Title:     "HW1",
				ScoreSpec: &ScoreSpec{20},
				UploadSpec: &UploadSpec{
					Files: []FileSpec{
						{
							Name:       "*.c",
							Required:   true,
							MaxSize:    FileSize("10MB"),
							MaxMatches: 10,
						},
					},
				},
				GraderSpec: &GraderSpec{
					Rubric: &Rubric{
						Items: []RubricItem{
							{
								Name:        "Some Name",
								Description: "Some description",
								ScoreSpec:   &ScoreSpec{10},
							},
							{
								Name:        "Some Name 2",
								Description: "Some description 2",
								ScoreSpec:   &ScoreSpec{10},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		var got ActivitySpec
		if err := yaml.UnmarshalWithOptions([]byte(test.yamlSource), &got, yaml.DisallowUnknownField()); err != nil {
			t.Fatalf("%s", err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %#v, want %#v", got, test.want)
		}
	}
}
