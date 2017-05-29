package cubecode

// cubecode to unicode conversion table
// cubecode is a small subset of unicode, mainly offering kyrillic and greek letters
// example: server sends a 2, you convert 2 → 193, 193 in unicode is Á
var cubeToUni = [256]rune{
	0, 192, 193, 194, 195, 196, 197, 198, 199, 9, 10, 11, 12, 13, 200, 201,
	202, 203, 204, 205, 206, 207, 209, 210, 211, 212, 213, 214, 216, 217, 218, 219,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63,
	64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 220,
	221, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237,
	238, 239, 241, 242, 243, 244, 245, 246, 248, 249, 250, 251, 252, 253, 255, 0x104,
	0x105, 0x106, 0x107, 0x10C, 0x10D, 0x10E, 0x10F, 0x118, 0x119, 0x11A, 0x11B, 0x11E, 0x11F, 0x130, 0x131, 0x141,
	0x142, 0x143, 0x144, 0x147, 0x148, 0x150, 0x151, 0x152, 0x153, 0x158, 0x159, 0x15A, 0x15B, 0x15E, 0x15F, 0x160,
	0x161, 0x164, 0x165, 0x16E, 0x16F, 0x170, 0x171, 0x178, 0x179, 0x17A, 0x17B, 0x17C, 0x17D, 0x17E, 0x404, 0x411,
	0x413, 0x414, 0x416, 0x417, 0x418, 0x419, 0x41B, 0x41F, 0x423, 0x424, 0x426, 0x427, 0x428, 0x429, 0x42A, 0x42B,
	0x42C, 0x42D, 0x42E, 0x42F, 0x431, 0x432, 0x433, 0x434, 0x436, 0x437, 0x438, 0x439, 0x43A, 0x43B, 0x43C, 0x43D,
	0x43F, 0x442, 0x444, 0x446, 0x447, 0x448, 0x449, 0x44A, 0x44B, 0x44C, 0x44D, 0x44E, 0x44F, 0x454, 0x490, 0x491,
}

var uni2Cube = map[rune]int32{
	'R':    82,
	'т':    241,
	'À':    1,
	'Ë':    17,
	'|':    124,
	'б':    228,
	'â':    132,
	'ê':    140,
	'ą':    160,
	'5':    53,
	'8':    56,
	'K':    75,
	'M':    77,
	'O':    79,
	'Ł':    175,
	'к':    236,
	'Û':    31,
	'/':    47,
	'Ш':    220,
	'я':    252,
	's':    115,
	'æ':    136,
	'ë':    141,
	'ď':    166,
	'2':    50,
	'W':    87,
	'=':    61,
	'ü':    156,
	'Ě':    169,
	'З':    211,
	'и':    234,
	'Ú':    30,
	' ':    32,
	'E':    69,
	'л':    237,
	'x':    120,
	'è':    138,
	'ø':    152,
	'ğ':    172,
	'Š':    191,
	'э':    250,
	'Ô':    25,
	'!':    33,
	'Є':    206,
	'C':    67,
	'V':    86,
	'Ź':    200,
	'3':    51,
	'D':    68,
	'м':    238,
	'Í':    19,
	'Ь':    224,
	'Ń':    177,
	'н':    239,
	'J':    74,
	'ń':    178,
	'v':    118,
	'И':    212,
	'Л':    214,
	'П':    215,
	'Ã':    4,
	'Ê':    16,
	'\n':   10,
	'u':    117,
	'Œ':    183,
	'ž':    205,
	'ù':    153,
	'\x00': 0,
	'?':    63,
	'b':    98,
	'w':    119,
	'f':    102,
	'n':    110,
	'ř':    186,
	'ű':    198,
	'У':    216,
	'д':    231,
	'Å':    6,
	'H':    72,
	'ь':    249,
	'G':    71,
	'c':    99,
	'~':    126,
	'ф':    242,
	'^':    94,
	'o':    111,
	'ý':    157,
	'ň':    180,
	'Ÿ':    199,
	'"':    34,
	'Q':    81,
	'_':    95,
	'ź':    201,
	'á':    131,
	'ę':    168,
	'F':    70,
	'L':    76,
	'}':    125,
	'û':    155,
	'Â':    3,
	',':    44,
	'X':    88,
	'ó':    148,
	'ş':    190,
	'Ó':    24,
	'.':    46,
	'\v':   11,
	'Ñ':    22,
	'Ù':    29,
	'\'':   39,
	'i':    105,
	'Ő':    181,
	'Á':    2,
	'Ç':    8,
	'l':    108,
	'y':    121,
	'{':    123,
	'Ğ':    171,
	'Ż':    202,
	'ґ':    255,
	'6':    54,
	':':    58,
	'Ч':    219,
	'ц':    243,
	'I':    73,
	'e':    101,
	'j':    106,
	'ż':    203,
	'Ю':    226,
	'Я':    227,
	'*':    42,
	'9':    57,
	'(':    40,
	'Ü':    127,
	'Č':    163,
	'Й':    213,
	'Ì':    18,
	'Ö':    27,
	'p':    112,
	'q':    113,
	'ú':    154,
	'ő':    182,
	'0':    48,
	'T':    84,
	'A':    65,
	'Ď':    165,
	'Ů':    195,
	'ě':    170,
	'&':    38,
	'å':    135,
	'ж':    232,
	'ì':    142,
	'Ę':    167,
	'S':    83,
	'ã':    133,
	'Ű':    197,
	'Æ':    7,
	'Î':    20,
	'ï':    145,
	'ö':    151,
	'Ą':    159,
	'ł':    176,
	'œ':    184,
	'ť':    194,
	'#':    35,
	'ç':    137,
	'в':    229,
	'є':    253,
	'Ś':    187,
	'Ъ':    222,
	'ъ':    247,
	'Ï':    21,
	'-':    45,
	'\\':   92,
	'r':    114,
	'í':    143,
	'č':    164,
	'Ř':    185,
	'<':    60,
	'Z':    90,
	'Ť':    193,
	'Ž':    204,
	'É':    15,
	'ò':    147,
	'š':    192,
	'ů':    196,
	'ш':    245,
	'È':    14,
	't':    116,
	'з':    233,
	'ы':    248,
	'Б':    207,
	'Ы':    223,
	'Y':    89,
	'é':    139,
	'ñ':    146,
	'\t':   9,
	'N':    78,
	'î':    144,
	'\f':   12,
	'`':    96,
	'й':    235,
	'щ':    246,
	'Ø':    28,
	'ÿ':    158,
	'à':    130,
	'Г':    208,
	'B':    66,
	'k':    107,
	'z':    122,
	'Ж':    210,
	'ч':    244,
	'Ґ':    254,
	'+':    43,
	']':    93,
	'Э':    225,
	'ä':    134,
	'Ф':    217,
	'@':    64,
	'[':    91,
	';':    59,
	'İ':    173,
	'ś':    188,
	'\r':   13,
	'4':    52,
	'Ć':    161,
	'Д':    209,
	'Ò':    23,
	'm':    109,
	'$':    36,
	'Ц':    218,
	'd':    100,
	'ô':    149,
	'Щ':    221,
	'7':    55,
	'a':    97,
	'ć':    162,
	'Ä':    5,
	'Ý':    128,
	'г':    230,
	'>':    62,
	'g':    103,
	')':    41,
	'1':    49,
	'ı':    174,
	'п':    240,
	'ю':    251,
	'Õ':    26,
	'%':    37,
	'P':    80,
	'õ':    150,
	'U':    85,
	'Ň':    179,
	'Ş':    189,
	'h':    104,
	'ß':    129,
}
