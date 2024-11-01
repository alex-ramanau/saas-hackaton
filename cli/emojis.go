package main

// list of allowed emojis using :emoji: syntax
var emojiMap = map[string]string{
	":smile:":        "😊",
	":happy:":        "😄",
	":laugh:":        "😂",
	":love:":         "😍",
	":annoyed:":      "😒",
	":scared:":       "😱",
	":sad:":          "😢",
	":angry:":        "😠",
	":wink:":         "😉",
	":heart:":        "❤️",
	":thumbsup:":     "👍",
	":thumbsdown:":   "👎",
	":clap:":         "👏",
	":yes:":          "👍",
	":no:":           "👎",
	":fire:":         "🔥",
	":star:":         "⭐",
	":sparkles:":     "✨",
	":party:":        "🎉",
	":cry:":          "😭",
	":blush:":        "😊",
	":sunglasses:":   "😎",
	":thinking:":     "🤔",
	":confused:":     "😕",
	":sleepy:":       "😴",
	":tired:":        "😩",
	":cool:":         "😎",
	":nerd:":         "🤓",
	":shy:":          "😳",
	":disappointed:": "😞",
	":pensive:":      "😔",
	":scream:":       "😱",
	":rage:":         "😡",
	":zipper_mouth:": "🤐",
	":alien:":        "👽",
	":robot:":        "🤖",
	":broken_heart:": "💔",
	":unicorn:":      "🦄",
	":swan:":         "🦢",
	":dog:":          "🐶",
	":cat:":          "🐱",
	":tiger:":        "🐯",
	":bear:":         "🐻",
	":panda:":        "🐼",
	":frog:":         "🐸",
	":chicken:":      "🐔",
	":penguin:":      "🐧",
	":turtle:":       "🐢",
	":whale:":        "🐳",
	":fish:":         "🐟",
	":shark:":        "🦈",
	":octopus:":      "🐙",
	":apple:":        "🍎",
	":banana:":       "🍌",
	":grapes:":       "🍇",
	":lemon:":        "🍋",
	":watermelon:":   "🍉",
	":strawberry:":   "🍓",
	":peach:":        "🍑",
	":pineapple:":    "🍍",
	":cake:":         "🎂",
	":cookie:":       "🍪",
	":pizza:":        "🍕",
	":hamburger:":    "🍔",
	":hotdog:":       "🌭",
	":fries:":        "🍟",
	":sushi:":        "🍣",
	":bento:":        "🍱",
	":coffee:":       "☕",
	":tea:":          "🍵",
	":beer:":         "🍺",
	":wine:":         "🍷",
	":cocktail:":     "🍸",
	":tumbler:":      "🥤",
	":popcorn:":      "🍿",
	":chocolate:":    "🍫",
	":gift:":         "🎁",
	":balloon:":      "🎈",
	":ribbon:":       "🎀",
	":candle:":       "🕯️",
	":sparkler:":     "🎆",
	":party_hat:":    "🎉",
	":hourglass:":    "⏳",
	":calendar:":     "📅",
	":clock:":        "🕒",
	":sun:":          "☀️",
	":cloud:":        "☁️",
	":rain:":         "🌧️",
	":snow:":         "❄️",
	":thunder:":      "⛈️",
	":rainbow:":      "🌈",
	":earth:":        "🌍",
	":mountain:":     "🏔️",
	":tree:":         "🌳",
	":flower:":       "🌼",
	":cactus:":       "🌵",
	":ocean:":        "🌊",
	":car:":          "🚗",
	":bike:":         "🚴",
	":plane:":        "✈️",
	":train:":        "🚆",
	":rocket:":       "🚀",
	":boat:":         "⛵",
	":house:":        "🏠",
	":hotel:":        "🏨",
	":school:":       "🏫",
	":church:":       "⛪",
	":hospital:":     "🏥",
	":building:":     "🏢",
	":factory:":      "🏭",
	":office:":       "🏢",
	":love_letter:":  "💌",
	":envelope:":     "✉️",
	":paperclip:":    "📎",
	":book:":         "📚",
	":pen:":          "🖊️",
	":pencil:":       "✏️",
	":scissors:":     "✂️",
	":notebook:":     "📓",
	":chart:":        "📊",
	":mag:":          "🔍",
	":lightbulb:":    "💡",
	":key:":          "🔑",
	":lock:":         "🔒",
	":umbrella:":     "☂️",
	":map:":          "🗺️",
	":compass:":      "🧭",
	":fishing_pole:": "🎣",
	":mushroom:":     "🍄",
	":corn:":         "🌽",
	":eggplant:":     "🍆",
	":potato:":       "🥔",
	":carrot:":       "🥕",
	":onion:":        "🧅",
	":peanuts:":      "🥜",
	":bread:":        "🍞",
	":cucumber:":     "🥒",
	":vampire:":      "🧛",
	":blood:":        "🩸",
	":zombie:":       "🧟",
	":sword:":        "🗡️",
}
