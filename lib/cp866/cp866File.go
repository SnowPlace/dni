package cp866

import (
	"strings"
)

const version = "1.0.11.05"

func Decode(text []byte) (message string) {
	table := map[byte]string{
		byte(10): "\n", byte(13): "\r",
		byte(32): " ", byte(33): "!", byte(34): string('"'), byte(35): "#",
		byte(36): "$", byte(37): "%", byte(38): "&", byte(39): "'", byte(40): "(",
		byte(41): ")", byte(42): "*", byte(43): "+", byte(44): ",", byte(45): "-",
		byte(46): ".", byte(47): "/", byte(48): "0", byte(49): "1", byte(50): "2",
		byte(51): "3", byte(52): "4", byte(53): "5", byte(54): "6", byte(55): "7",
		byte(56): "8", byte(57): "9", byte(58): ":", byte(59): ";", byte(60): "<",
		byte(61): "=", byte(62): ">", byte(63): "?", byte(64): "@", byte(65): "A",
		byte(66): "B", byte(67): "C", byte(68): "D", byte(69): "E", byte(70): "F",
		byte(71): "G", byte(72): "H", byte(73): "I", byte(74): "J", byte(75): "K",
		byte(76): "L", byte(77): "M", byte(78): "N", byte(79): "O", byte(80): "P",
		byte(81): "Q", byte(82): "R", byte(83): "S", byte(84): "T", byte(85): "U",
		byte(86): "V", byte(87): "W", byte(88): "X", byte(89): "Y", byte(90): "Z",
		byte(91): "[", byte(92): "\\", byte(93): "]", byte(94): "^", byte(95): "_",
		byte(96): "`", byte(97): "a", byte(98): "b", byte(99): "c", byte(100): "d",
		byte(101): "e", byte(102): "f", byte(103): "g", byte(104): "h", byte(105): "i",
		byte(106): "j", byte(107): "k", byte(108): "l", byte(109): "m", byte(110): "n",
		byte(111): "o", byte(112): "p", byte(113): "q", byte(114): "r", byte(115): "s",
		byte(116): "t", byte(117): "u", byte(118): "v", byte(119): "w", byte(120): "x",
		byte(121): "y", byte(122): "z", byte(123): "{", byte(124): "|", byte(125): "}",
		byte(126): "~",
		byte(128): "A", byte(129): "Б", byte(130): "В", byte(131): "Г", byte(132): "Д",
		byte(133): "Е", byte(134): "Ж", byte(135): "З", byte(136): "И", byte(137): "Й",
		byte(138): "К", byte(139): "Л", byte(140): "М", byte(141): "Н", byte(142): "О",
		byte(143): "П", byte(144): "Р", byte(145): "С", byte(146): "Т", byte(147): "У",
		byte(148): "Ф", byte(149): "Х", byte(150): "Ц", byte(151): "Ч", byte(152): "Ш",
		byte(153): "Щ", byte(154): "Ъ", byte(155): "Ы", byte(156): "Ь", byte(157): "Э",
		byte(158): "Ю", byte(159): "Я", byte(160): "а", byte(161): "б", byte(162): "в",
		byte(163): "г", byte(164): "д", byte(165): "е", byte(166): "ж", byte(167): "з",
		byte(168): "и", byte(169): "й", byte(170): "к", byte(171): "л", byte(172): "м",
		byte(173): "н", byte(174): "о", byte(175): "п",
		byte(224): "р", byte(225): "с", byte(226): "т", byte(227): "у", byte(228): "ф",
		byte(229): "х", byte(230): "ц", byte(231): "ч", byte(232): "ш", byte(233): "щ",
		byte(234): "ъ", byte(235): "ы", byte(236): "ь", byte(237): "э", byte(238): "ю",
		byte(239): "я",
	}
	var array []string
	for b := range text {
		array = append(array, table[text[b]])
	}
	return strings.Join(array, "")
}
