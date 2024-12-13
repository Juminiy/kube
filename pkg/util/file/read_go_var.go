package file

import (
	"github.com/Juminiy/kube/pkg/util"
	"strings"
)

var Xxx int
var Xxx2 = "xxx2"
var ()
var (
	Xxx3 int64
	Xxx4 = "xxx4"
)

// ReadGoVar
/* read check exported global var
var Xxx type
var Xxx2 = value
var ()
var (
\tXxx3 type
\tXxx4 = value
)
*/
func ReadGoVar(filePath string) *ExportedVar {
	if strings.HasSuffix(filePath, "_test.go") {
		return nil
	}
	rd := NewReader(filePath)
	defer rd.Done()

	lineCounts := make([]int, 0, 8)
	missCounts := make([]int, 0, 2)
	var line string
	var lineCount int
	for rd.NextLine(&line) {
		lineCount++
		if !strings.HasPrefix(line, "var ") {
			continue
		}
		trimVar := strings.TrimPrefix(line, "var ")
		if strings.Contains(trimVar, "()") || isNotExported(trimVar) {
			continue
		} else if isExported(trimVar) {
			lineCounts = append(lineCounts, lineCount)
		} else if strings.Contains(trimVar, "(") {
			var newLine string
			for rd.NextLine(&newLine) {
				lineCount++
				newLine = strings.TrimLeft(newLine, "\t")
				// a complex problem not to write myself, try to find tool, no tool
				if len(newLine) > 0 && newLine[0] == ')' {
					break
				} else if isExported(newLine) {
					lineCounts = append(lineCounts, lineCount)
				}
			}
		} else {
			missCounts = append(missCounts, lineCount)
		}
	}

	if len(lineCounts) > 0 ||
		len(missCounts) > 0 {
		return (&ExportedVar{
			LineCount:     lineCounts,
			MissLineCount: missCounts,
		}).parseFilePath(filePath)
	}
	return nil
}

func isExported(s string) bool {
	return len(s) > 0 && util.InRange(s[0], 'A', 'Z')
}

func isNotExported(s string) bool { return len(s) > 0 && (util.InRange(s[0], 'a', 'z') || s[0] == '_') }

type ExportedVar struct {
	FilePath      string
	FileName      string
	LineCount     []int
	MissLineCount []int
}

func (v *ExportedVar) parseFilePath(s string) *ExportedVar {
	v.FilePath = s
	seps := strings.Split(s, "/")
	if len(seps) > 0 {
		v.FileName = seps[len(seps)-1]
	}
	return v
}
