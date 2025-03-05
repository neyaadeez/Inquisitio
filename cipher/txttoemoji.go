package cipher

import (
	"strings"
	"unicode"
)

func TextToEmoji(txt string) string {
	flags := []string{
		"\U0001F1E6\U0001F1E8", // 🇦🇨 (A)
		"\U0001F1E6\U0001F1E9", // 🇦🇩 (B)
		"\U0001F1E6\U0001F1EA", // 🇦🇪 (C)
		"\U0001F1E6\U0001F1EB", // 🇦🇫 (D)
		"\U0001F1E6\U0001F1EC", // 🇦🇬 (E)
		"\U0001F1E6\U0001F1ED", // 🇦🇭 (F)
		"\U0001F1E6\U0001F1EE", // 🇦🇮 (G)
		"\U0001F1E6\U0001F1EF", // 🇦🇯 (H)
		"\U0001F1E6\U0001F1F0", // 🇦🇰 (I)
		"\U0001F1E6\U0001F1F1", // 🇦🇱 (J)
		"\U0001F1E6\U0001F1F2", // 🇦🇲 (K)
		"\U0001F1E6\U0001F1F3", // 🇦🇳 (L)
		"\U0001F1E6\U0001F1F4", // 🇦🇴 (M)
		"\U0001F1E6\U0001F1F5", // 🇦🇵 (N)
		"\U0001F1E6\U0001F1F6", // 🇦🇶 (O)
		"\U0001F1E6\U0001F1F7", // 🇦🇷 (P)
		"\U0001F1E6\U0001F1F8", // 🇦🇸 (Q)
		"\U0001F1E6\U0001F1F9", // 🇦🇹 (R)
		"\U0001F1E6\U0001F1FA", // 🇦🇺 (S)
		"\U0001F1E6\U0001F1FB", // 🇦🇼 (T)
		"\U0001F1E6\U0001F1FC", // 🇦🇽 (U)
		"\U0001F1E6\U0001F1FD", // 🇦🇾 (V)
		"\U0001F1E6\U0001F1FE", // 🇦🇿 (W)
		"\U0001F1E6\U0001F1FF", // 🇧🇦 (X)
		"\U0001F1E7\U0001F1F7", // 🇧🇧 (Y)
		"\U0001F1E7\U0001F1F9", // 🇧🇸 (Z)
	}

	txt = strings.ToLower(txt)
	var cipher []string
	for _, ch := range txt {
		if unicode.IsLetter(ch) {
			ind := ch - 'a'
			if ind >= 0 && ind < 26 {
				cipher = append(cipher, flags[ind])
			}
		}
	}

	return strings.Join(cipher, " ")
}

func EmojiToText(emoji string) string {
	parts := strings.Split(emoji, " ")
	flags := map[string]string{
		"\U0001F1E6\U0001F1E8": "A",
		"\U0001F1E6\U0001F1E9": "B",
		"\U0001F1E6\U0001F1EA": "C",
		"\U0001F1E6\U0001F1EB": "D",
		"\U0001F1E6\U0001F1EC": "E",
		"\U0001F1E6\U0001F1ED": "F",
		"\U0001F1E6\U0001F1EE": "G",
		"\U0001F1E6\U0001F1EF": "H",
		"\U0001F1E6\U0001F1F0": "I",
		"\U0001F1E6\U0001F1F1": "J",
		"\U0001F1E6\U0001F1F2": "K",
		"\U0001F1E6\U0001F1F3": "L",
		"\U0001F1E6\U0001F1F4": "M",
		"\U0001F1E6\U0001F1F5": "N",
		"\U0001F1E6\U0001F1F6": "O",
		"\U0001F1E6\U0001F1F7": "P",
		"\U0001F1E6\U0001F1F8": "Q",
		"\U0001F1E6\U0001F1F9": "R",
		"\U0001F1E6\U0001F1FA": "S",
		"\U0001F1E6\U0001F1FB": "T",
		"\U0001F1E6\U0001F1FC": "U",
		"\U0001F1E6\U0001F1FD": "V",
		"\U0001F1E6\U0001F1FE": "W",
		"\U0001F1E6\U0001F1FF": "X",
		"\U0001F1E7\U0001F1F7": "Y",
		"\U0001F1E7\U0001F1F9": "Z",
	}

	var res []string
	for _, val := range parts {
		if key, ok := flags[val]; ok {
			res = append(res, key)
		}
	}

	return strings.Join(res, "")
}
