package ansifonts

import "strings"

// ParseInlineKerning extracts inline kerning markers from the input text and
// returns the clean text (with markers removed) alongside a nested map
// representing per-line, per-character pixel spacing adjustments.
//
// The inline syntax uses \+ to add 1 pixel of space and \- to remove 1 pixel
// of space before the next character. Multiple markers stack, so \+\+ adds
// 2 pixels and \-\- removes 2 pixels.
//
// For example, the input "H\+el\+lo W\-orld" produces the clean text
// "Hello World" and the kerning map {0: {1: 1, 3: 1, 7: -1}}.
//
// The returned kerning map is keyed by line index and then by character index,
// matching the CustomKerning field of RenderOptions. A kerning entry at index N
// adjusts the space before the Nth character (i.e. between characters N-1 and N).
func ParseInlineKerning(text string) (string, map[int]map[int]int) {
	lines := strings.Split(text, "\n")
	cleanLines := make([]string, len(lines))
	kerning := make(map[int]map[int]int)

	for lineIdx, line := range lines {
		runes := []rune(line)
		var cleanRunes []rune
		cleanIdx := 0

		i := 0
		for i < len(runes) {
			// Detect inline kerning markers: \+ (add space) and \- (remove space)
			if runes[i] == '\\' && i+1 < len(runes) {
				if runes[i+1] == '+' {
					if kerning[lineIdx] == nil {
						kerning[lineIdx] = make(map[int]int)
					}
					kerning[lineIdx][cleanIdx]++
					i += 2
					continue
				} else if runes[i+1] == '-' {
					if kerning[lineIdx] == nil {
						kerning[lineIdx] = make(map[int]int)
					}
					kerning[lineIdx][cleanIdx]--
					i += 2
					continue
				}
			}
			cleanRunes = append(cleanRunes, runes[i])
			cleanIdx++
			i++
		}

		cleanLines[lineIdx] = string(cleanRunes)
	}

	return strings.Join(cleanLines, "\n"), kerning
}
