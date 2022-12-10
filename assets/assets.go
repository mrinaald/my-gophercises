package assets

import (
	"path/filepath"
	"runtime"
)

// import _ "embed"

// //go:embed quiz-game/problems.csv
// var QuizGameDefaultProblemFile string

func getCurrentFileAbsPath() string {
	_, currentFilePath, _, ok := runtime.Caller(0)

	if !ok {
		return ""
	}

	return filepath.Dir(currentFilePath)
}

var QuizGameDefaultProblemFile, _ = filepath.Abs(filepath.Join(getCurrentFileAbsPath(), "quiz-game/problems.csv"))
