// Package mockv2 was generated
package mockv2

import (
	"github.com/Juminiy/kube/pkg/util/zerobuf"
)

const lowerStr = "abcdefghijklmnopqrstuvwxyz"
const upperStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphaStr = lowerStr + upperStr
const numericStr = "0123456789"
const alphaNumericStr = alphaStr + numericStr
const binStr = "01"
const octStr = "01234567"
const hexStr = "01234567890abcdefABCDEF"
const specialStr = "@#$%&?|!(){}<>=*+-_:;,."
const specialSafeStr = "!@.-_*" // https://github.com/1Password/spg/pull/22
const spaceStr = " "
const allStr = lowerStr + upperStr + numericStr + specialStr + spaceStr
const vowels = "aeiou"
const hashtag = '#'
const questionmark = '?'
const dash = '-'
const base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const minUint = 0
const maxUint = ^uint(0)
const minInt = -maxInt - 1
const maxInt = int(^uint(0) >> 1)
const is32bit = ^uint(0)>>32 == 0

func randStringByRunes(r []rune, size int) string {
	buf := zerobuf.Get()
	defer buf.Free()

	for range size {
		buf.WriteString(string(randT(r...)))
	}

	return buf.String()
}
