// Code generated by re2dfa (https://github.com/opennota/re2dfa). DO NOT EDIT

package benchmarks

import "unicode/utf8"

func match1(s string) (end int) {
	end = -1
	var r rune
	var rlen int
	i := 0
	lazy := false
	type jmp struct{ s, i int }
	var lazyArr [4]jmp
	lazyStack := lazyArr[:0]
	_, _, _ = r, rlen, i
	switch {
	case i == 0:
		goto s2
	}
	goto bt
s2:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 60:
		goto s3
	}
	goto bt
s3:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 33:
		goto s4
	case r == 47:
		goto s32
	case r == 63:
		goto s37
	case r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s42
	}
	goto bt
s4:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 45:
		goto s5
	case r >= 65 && r <= 90:
		goto s15
	case r == 91:
		goto s20
	}
	goto bt
s5:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 45:
		goto s6
	}
	goto bt
s6:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 44 || r >= 46 && r <= 61 || r >= 63:
		goto s7
	case r == 45:
		goto s12
	}
	goto bt
s7:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 44 || r >= 46:
		goto s8
	case r == 45:
		goto s9
	}
	goto bt
s8:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 44 || r >= 46:
		goto s8
	case r == 45:
		goto s9
	}
	goto bt
s9:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 44 || r >= 46:
		goto s8
	case r == 45:
		goto s10
	}
	goto bt
s10:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 62:
		end = i
	}
	goto bt
s12:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 44 || r >= 46 && r <= 61 || r >= 63:
		goto s7
	case r == 45:
		goto s13
	}
	goto bt
s13:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 62:
		end = i
	}
	goto bt
s15:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s16
	case r >= 65 && r <= 90:
		goto s15
	}
	goto bt
s16:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s17
	case r >= 33 && r <= 61 || r >= 63:
		goto s18
	case r == 62:
		end = i
	}
	goto bt
s17:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 8 || r == 11 || r >= 14 && r <= 31 || r >= 33 && r <= 61 || r >= 63:
		goto s18
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s17
	case r == 62:
		end = i
	}
	goto bt
s18:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 61 || r >= 63:
		goto s18
	case r == 62:
		end = i
	}
	goto bt
s20:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 67:
		goto s21
	}
	goto bt
s21:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 68:
		goto s22
	}
	goto bt
s22:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 65:
		goto s23
	}
	goto bt
s23:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 84:
		goto s24
	}
	goto bt
s24:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 65:
		goto s25
	}
	goto bt
s25:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 91:
		goto s26
	}
	goto bt
s26:
	if lazy {
		lazy = false
		goto s27
	}
	lazyStack = append(lazyStack, jmp{s: 26, i: i})
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 93:
		goto s29
	}
	goto bt
s27:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 1114111:
		goto s28
	}
	goto bt
s28:
	if lazy {
		lazy = false
		goto s27
	}
	lazyStack = append(lazyStack, jmp{s: 28, i: i})
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 93:
		goto s29
	}
	goto bt
s29:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 93:
		goto s30
	}
	goto bt
s30:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 62:
		end = i
	}
	goto bt
s32:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s33
	}
	goto bt
s33:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 62:
		end = i
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s34
	case r == 45 || r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s36
	}
	goto bt
s34:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s34
	case r == 62:
		end = i
	}
	goto bt
s36:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s34
	case r == 45 || r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s36
	case r == 62:
		end = i
	}
	goto bt
s37:
	if lazy {
		lazy = false
		goto s38
	}
	lazyStack = append(lazyStack, jmp{s: 37, i: i})
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 63:
		goto s40
	}
	goto bt
s38:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 9 || r >= 11:
		goto s39
	}
	goto bt
s39:
	if lazy {
		lazy = false
		goto s38
	}
	lazyStack = append(lazyStack, jmp{s: 39, i: i})
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 63:
		goto s40
	}
	goto bt
s40:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 62:
		end = i
	}
	goto bt
s42:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s43
	case r == 45 || r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s59
	case r == 47:
		goto s44
	case r == 62:
		end = i
	}
	goto bt
s43:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s43
	case r == 47:
		goto s44
	case r == 58 || r >= 65 && r <= 90 || r == 95 || r >= 97 && r <= 122:
		goto s46
	case r == 62:
		end = i
	}
	goto bt
s44:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 62:
		end = i
	}
	goto bt
s46:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s47
	case r >= 45 && r <= 46 || r >= 48 && r <= 58 || r >= 65 && r <= 90 || r == 95 || r >= 97 && r <= 122:
		goto s58
	case r == 47:
		goto s44
	case r == 61:
		goto s48
	case r == 62:
		end = i
	}
	goto bt
s47:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 61:
		goto s48
	case r == 62:
		end = i
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s47
	case r == 47:
		goto s44
	case r == 58 || r >= 65 && r <= 90 || r == 95 || r >= 97 && r <= 122:
		goto s46
	}
	goto bt
s48:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s49
	case r == 33 || r >= 35 && r <= 38 || r >= 40 && r <= 59 || r >= 63 && r <= 95 || r >= 97:
		goto s50
	case r == 34:
		goto s52
	case r == 39:
		goto s55
	}
	goto bt
s49:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s49
	case r == 33 || r >= 35 && r <= 38 || r >= 40 && r <= 59 || r >= 63 && r <= 95 || r >= 97:
		goto s50
	case r == 34:
		goto s52
	case r == 39:
		goto s55
	}
	goto bt
s50:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s43
	case r == 33 || r >= 35 && r <= 38 || r >= 40 && r <= 46 || r >= 48 && r <= 59 || r >= 63 && r <= 95 || r >= 97:
		goto s50
	case r == 47:
		goto s51
	case r == 62:
		end = i
	}
	goto bt
s51:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s43
	case r == 33 || r >= 35 && r <= 38 || r >= 48 && r <= 59 || r >= 63 && r <= 95 || r >= 97:
		goto s50
	case r == 47:
		goto s51
	case r == 62:
		end = i
	}
	goto bt
s52:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 33 || r >= 35:
		goto s53
	case r == 34:
		goto s54
	}
	goto bt
s53:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 33 || r >= 35:
		goto s53
	case r == 34:
		goto s54
	}
	goto bt
s54:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s43
	case r == 47:
		goto s44
	case r == 62:
		end = i
	}
	goto bt
s55:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 38 || r >= 40:
		goto s56
	case r == 39:
		goto s57
	}
	goto bt
s56:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r <= 38 || r >= 40:
		goto s56
	case r == 39:
		goto s57
	}
	goto bt
s57:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s43
	case r == 47:
		goto s44
	case r == 62:
		end = i
	}
	goto bt
s58:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 62:
		end = i
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s47
	case r >= 45 && r <= 46 || r >= 48 && r <= 58 || r >= 65 && r <= 90 || r == 95 || r >= 97 && r <= 122:
		goto s58
	case r == 47:
		goto s44
	case r == 61:
		goto s48
	}
	goto bt
s59:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r >= 9 && r <= 10 || r >= 12 && r <= 13 || r == 32:
		goto s43
	case r == 45 || r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s59
	case r == 47:
		goto s44
	case r == 62:
		end = i
	}
bt:
	if end >= 0 || len(lazyStack) == 0 {
		return
	}
	var to jmp
	to, lazyStack = lazyStack[len(lazyStack)-1], lazyStack[:len(lazyStack)-1]
	lazy = true
	i = to.i
	switch to.s {
	case 26:
		goto s26
	case 28:
		goto s28
	case 37:
		goto s37
	case 39:
		goto s39
	}
	return
}
