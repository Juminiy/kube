package docker_api

import (
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

var (
	allowURL = []string{
		"jammy-env",              // only repo
		"jammy-env:v1.6",         // only refStr-artifact
		"library/jammy-env:v1.6", // only refStr
		"10.112.121.243:8111/library/jammy-env:v1.6",                                                                    // absRefStr
		"sha256:305243c734571da2d100c8c8b3c3167a098cab6049c9a5b066b6021a60fcb966",                                       // only digest
		"jammy-env@sha256:305243c734571da2d100c8c8b3c3167a098cab6049c9a5b066b6021a60fcb966",                             // only digest-artifact
		"library/jammy-env@sha256:305243c734571da2d100c8c8b3c3167a098cab6049c9a5b066b6021a60fcb966",                     // only digestStr
		"10.112.121.243:8111/library/jammy-env@sha256:305243c734571da2d100c8c8b3c3167a098cab6049c9a5b066b6021a60fcb966", // absDigestStr
	}
)

func TestParseToArtifact(t *testing.T) {
	var prettyEnc = safe_json.Pretty
	for _, s := range allowURL {
		t.Log(prettyEnc(ParseToArtifact(s)))
	}
}

func TestArtifact(t *testing.T) {
	for i, s := range allowURL {
		arti := ParseToArtifact(s)
		switch i {
		case 0, 1, 2:
			t.Log(arti.RefStr())
		case 3:
			t.Log(arti.AbsRefStr())
		case 4, 5, 6:
			t.Log(arti.DigestStr())
		case 7:
			t.Log(arti.AbsRefStr())
		}
	}
}
