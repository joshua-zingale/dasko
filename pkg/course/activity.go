package course

type ActivitySpec struct {
	Title string `json:"title"`

	*ScoreSpec `json:",inline,omitempty"`
	// *ScorePolicy `json:",inline,omitempty"`

	// *Visibility `json:",inline,omitempty"`

	// *Initializaton `json:"initialization,omitempty"`

	UploadSpec *UploadSpec `json:"uploads,omitempty"`

	// *AutoGrader `json:"autograder,omitempty"`

	// Readings []ReadingSpec `json:"readings,omitempty"`

	GraderSpec *GraderSpec `json:"grader,omitempty"`
}

type ScoreSpec struct {
	MaxScore float64 `json:"max-score"`
}

type FileSize string

type UploadSpec struct {
	Files []FileSpec `json:"files,omitempty"`
}

type FileSpec struct {
	Name       string   `json:"name,omitempty"`
	Required   bool     `json:"required,omitempty"`
	MaxSize    FileSize `json:"max-size,omitempty"`
	MaxMatches int      `json:"max-matches,omitempty"`
}

// type Visibility struct {
// 	Groups []string `json:"visibility,omitempty"`
// }

// type ScorePolicy struct {
// 	ExactScoring    bool `json:"exact-scoring,omitempty"`
// 	OverflowScoring bool `json:"overflow-scoring,omitempty"`
// }

// type ReadingSpec struct {
// 	Source string `json:"source"`
// }

// type Initializaton struct {
// 	Copy           []CopyAction `json:"copy,omitempty"`
// 	*ExecutionSpec `json:"exec,omitempty"`
// }

// type CopyAction struct {
// 	Src  string `json:"src"`
// 	Dest string `json:"dest"`
// }

// type ExecutionSpec struct {
// 	Exec string `json:"exec,omitempty"`
// }

// type AutoGrader struct {
// 	*ExecutionSpec `json:",inline"`
// 	*PointValue    `json:",inline"`
// 	*ScorePolicy   `json:",inline"`
// }

// type ReadingActivity struct {
// 	*ReadingSpec `json:",inline"`
// 	*ScorePolicy `json:",inline"`
// 	*PointValue  `json:",inline"`
// }

type RubricItem struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	*ScoreSpec  `json:",inline"`
}

type Rubric struct {
	Items []RubricItem `json:"items"`
}

type GraderSpec struct {
	Rubric *Rubric `json:"rubric"`
	// Setup        ExecutionSpec `json:"setup"`
	// *ScorePolicy `json:",inline"`
}
