package day7

var nb = map[string]map[string]int{
	"bright beige": {
		"drab red":      4,
		"dull silver":   4,
		"pale violet":   2,
		"vibrant brown": 5,
	},
	"bright black": {
		"plaid tan": 5,
	},
	"bright blue": {
		"faded violet":  2,
		"light aqua":    3,
		"light lime":    5,
		"vibrant coral": 4,
	},
	"bright bronze": {
		"clear gray":    3,
		"muted silver":  3,
		"pale bronze":   1,
		"striped coral": 2,
	},
	"bright brown": {
		"bright coral": 1,
		"drab tan":     4,
		"posh brown":   3,
	},
	"bright chartreuse": {
		"light violet": 2,
		"vibrant gray": 3,
	},
	"bright coral": {
		"plaid indigo": 2,
		"shiny indigo": 3,
		"vibrant tan":  2,
	},
	"bright crimson": {
		"muted blue":    3,
		"striped olive": 2,
	},
	"bright cyan": {
		"bright aqua":   1,
		"dark white":    3,
		"light crimson": 2,
	},
	"bright fuchsia": {
		"dull orange": 4,
	},
	"bright gold": {
		"wavy coral": 5,
		"wavy red":   1,
	},
	"bright gray": {
		"clear cyan":    4,
		"dim green":     1,
		"mirrored teal": 4,
		"striped brown": 4,
	},
	"bright green": {
		"bright aqua": 3,
		"dark brown":  4,
		"dull maroon": 3,
	},
	"bright indigo": {
		"bright blue": 3,
	},
	"bright lavender": {
		"muted fuchsia": 4,
	},
	"bright magenta": {
		"bright blue":      3,
		"drab gold":        2,
		"mirrored magenta": 1,
		"vibrant blue":     1,
	},
	"bright maroon": {
		"light teal":      4,
		"mirrored yellow": 1,
	},
	"bright olive": {
		"drab red":       5,
		"dull crimson":   5,
		"posh turquoise": 5,
		"shiny fuchsia":  5,
	},
	"bright plum": {
		"dim aqua":     4,
		"dull crimson": 2,
		"wavy maroon":  1,
	},
	"bright purple": {
		"bright plum":       3,
		"dotted chartreuse": 2,
		"posh cyan":         4,
	},
	"bright red": {
		"dull silver": 3,
	},
	"bright salmon": {
		"drab red":       1,
		"muted red":      1,
		"plaid lavender": 4,
		"striped coral":  1,
	},
	"bright silver": {
		"dim teal":      3,
		"light purple":  5,
		"vibrant coral": 1,
	},
	"bright tan": {
		"drab teal":  5,
		"light blue": 2,
		"pale gold":  5,
	},
	"bright teal": {
		"dark silver": 4,
		"shiny brown": 4,
	},
	"bright tomato": {
		"muted blue": 5,
	},
	"bright turquoise": {
		"clear plum":    3,
		"mirrored aqua": 1,
	},
	"bright violet": {
		"drab olive":     3,
		"mirrored olive": 5,
		"wavy bronze":    2,
		"wavy lavender":  2,
	},
	"bright white": {
		"drab green":    3,
		"mirrored cyan": 2,
		"wavy red":      5,
	},
	"bright yellow": {
		"dotted gold":     2,
		"plaid lime":      3,
		"vibrant crimson": 2,
	},
	"clear aqua": {
		"clear gray":      5,
		"drab silver":     4,
		"posh chartreuse": 3,
	},
	"clear beige": {
		"dull coral":  3,
		"posh gray":   2,
		"posh silver": 4,
	},
	"clear black": {
		"light black":   2,
		"muted coral":   1,
		"muted fuchsia": 5,
	},
	"clear blue": {
		"dim fuchsia": 4,
		"dull gold":   4,
		"pale coral":  4,
		"pale tomato": 2,
	},
	"clear bronze": {
		"plaid beige": 2,
	},
	"clear brown": {
		"bright green": 4,
		"drab cyan":    2,
		"light lime":   4,
		"pale yellow":  5,
	},
	"clear chartreuse": {
		"dark orange":  3,
		"plaid purple": 4,
		"wavy bronze":  3,
	},
	"clear coral": {
		"shiny gray": 2,
	},
	"clear crimson": {
		"dim chartreuse": 5,
		"light crimson":  3,
		"shiny aqua":     4,
	},
	"clear cyan": {
		"bright indigo":   1,
		"dark chartreuse": 5,
		"pale yellow":     4,
		"vibrant coral":   2,
	},
	"clear fuchsia": {
		"muted bronze": 1,
	},
	"clear gold": {
		"drab indigo":  3,
		"light brown":  4,
		"shiny salmon": 2,
		"wavy violet":  4,
	},
	"clear gray": {
		"dotted cyan":  4,
		"vibrant plum": 4,
	},
	"clear green": {
		"mirrored olive": 3,
		"shiny tomato":   3,
	},
	"clear indigo": {
		"dull coral": 4,
		"light lime": 3,
	},
	"clear lavender": {
		"bright aqua":   2,
		"bright green":  3,
		"bright orange": 2,
		"dull silver":   3,
	},
	"clear lime": {
		"dim magenta":  3,
		"drab magenta": 1,
		"plaid tomato": 3,
		"shiny purple": 3,
	},
	"clear magenta": {
		"muted chartreuse": 2,
		"shiny magenta":    2,
	},
	"clear maroon": {
		"bright orange": 5,
	},
	"clear olive": {
		"faded turquoise": 5,
	},
	"clear orange": {
		"clear violet":    5,
		"dark chartreuse": 2,
		"dull beige":      2,
	},
	"clear plum": {
		"dim coral":      3,
		"faded black":    1,
		"mirrored brown": 4,
		"vibrant brown":  5,
	},
	"clear purple": {
		"vibrant white": 3,
	},
	"clear red": {
		"dotted salmon": 1,
		"dotted silver": 1,
		"dull teal":     1,
		"faded maroon":  1,
	},
	"clear salmon": {
		"clear plum":  5,
		"dull blue":   5,
		"posh bronze": 4,
	},
	"clear silver": {
		"clear salmon": 2,
		"drab indigo":  4,
	},
	"clear tan": {
		"striped orange": 3,
		"striped silver": 3,
		"wavy lavender":  4,
	},
	"clear teal": {
		"dark white": 3,
	},
	"clear tomato": {
		"mirrored tomato": 4,
		"muted indigo":    3,
		"striped tan":     1,
	},
	"clear turquoise": {
		"drab gold":   4,
		"drab tomato": 4,
		"shiny gold":  3,
	},
	"clear violet": {
		"light crimson": 3,
	},
	"clear white": {
		"drab indigo":   4,
		"dull magenta":  4,
		"muted magenta": 5,
		"pale plum":     3,
	},
	"clear yellow": {
		"vibrant teal": 2,
	},
	"dark aqua": {
		"dull red": 2,
	},
	"dark beige": {
		"posh lime":     2,
		"striped black": 3,
	},
	"dark black": {
		"dim silver":    1,
		"plaid magenta": 3,
	},
	"dark blue": {
		"plaid indigo": 3,
		"posh black":   2,
	},
	"dark bronze": {
		"dim green":       1,
		"posh chartreuse": 5,
	},
	"dark brown": {
		"light purple": 5,
	},
	"dark chartreuse": {
		"bright lime": 5,
	},
	"dark coral": {
		"dull tan": 3,
	},
	"dark crimson": {
		"bright white": 3,
		"faded purple": 3,
		"shiny orange": 4,
	},
	"dark cyan": {
		"dim chartreuse":   1,
		"dotted magenta":   4,
		"light chartreuse": 4,
		"shiny beige":      2,
	},
	"dark fuchsia": {
		"plaid beige":      4,
		"plaid chartreuse": 2,
	},
	"dark gold": {
		"dotted indigo": 1,
		"dotted silver": 4,
		"shiny teal":    4,
	},
	"dark gray": {
		"faded lavender": 4,
		"muted purple":   5,
		"striped blue":   2,
	},
	"dark green": {
		"dim cyan": 2,
	},
	"dark indigo": {
		"drab lavender":  3,
		"dull beige":     5,
		"mirrored black": 1,
		"shiny beige":    4,
	},
	"dark lavender": {
		"drab red":       1,
		"drab turquoise": 1,
		"faded beige":    4,
		"shiny indigo":   3,
	},
	"dark lime": {
		"bright indigo":    5,
		"dotted turquoise": 5,
		"drab maroon":      2,
		"shiny black":      4,
	},
	"dark magenta": {
		"bright teal":  2,
		"dotted brown": 5,
		"drab crimson": 1,
		"faded maroon": 4,
	},
	"dark maroon": {
		"drab turquoise": 3,
	},
	"dark olive": {
		"clear bronze": 5,
		"drab plum":    2,
	},
	"dark orange": {
		"bright beige":  5,
		"dim teal":      1,
		"drab tomato":   1,
		"striped brown": 3,
	},
	"dark plum": {
		"vibrant tan": 2,
	},
	"dark purple": {
		"bright black": 3,
		"dotted teal":  3,
		"wavy indigo":  1,
	},
	"dark red": {
		"dull maroon": 2,
		"dull plum":   4,
		"muted plum":  4,
	},
	"dark salmon": {
		"drab coral":   4,
		"drab tan":     5,
		"dull tan":     4,
		"vibrant plum": 1,
	},
	"dark silver": {
		"dark cyan":    3,
		"plaid salmon": 4,
	},
	"dark tan": {
		"faded gray":   2,
		"shiny salmon": 4,
	},
	"dark teal": {
		"muted fuchsia": 5,
	},
	"dark tomato": {
		"drab cyan":     2,
		"dull coral":    5,
		"posh blue":     1,
		"vibrant coral": 3,
	},
	"dark turquoise": {
		"dull plum": 5,
	},
	"dark violet": {
		"light purple": 2,
	},
	"dark white": {
		"dim teal":      2,
		"light aqua":    5,
		"light purple":  5,
		"muted fuchsia": 2,
	},
	"dark yellow": {
		"bright teal": 5,
		"pale cyan":   4,
		"pale olive":  2,
	},
	"dim aqua": {
		"dull tan":      1,
		"muted fuchsia": 1,
	},
	"dim beige": {
		"wavy red": 5,
	},
	"dim black": {
		"clear gold":   3,
		"muted violet": 4,
	},
	"dim blue": {
		"vibrant orange": 4,
	},
	"dim bronze": {
		"dim indigo":   2,
		"dim tomato":   4,
		"light silver": 1,
		"shiny silver": 2,
	},
	"dim brown": {
		"drab green":     5,
		"mirrored olive": 1,
		"muted olive":    3,
	},
	"dim chartreuse": {
		"bright blue": 5,
		"faded cyan":  5,
		"muted gold":  5,
	},
	"dim coral": {
		"dull coral":    3,
		"striped brown": 3,
	},
	"dim crimson": {
		"faded violet":   4,
		"striped silver": 1,
	},
	"dim cyan": {
		"drab purple":     5,
		"mirrored bronze": 3,
		"wavy olive":      3,
	},
	"dim fuchsia": {
		"mirrored brown": 4,
		"mirrored cyan":  3,
		"vibrant yellow": 3,
	},
	"dim gold": {
		"dim olive":   2,
		"plaid olive": 5,
		"posh orange": 2,
	},
	"dim gray": {
		"faded brown":  1,
		"light purple": 3,
		"posh crimson": 1,
	},
	"dim green": {
		"light lime":  1,
		"shiny beige": 1,
	},
	"dim indigo": {
		"bright aqua":  1,
		"clear cyan":   4,
		"dull tomato":  4,
		"shiny salmon": 2,
	},
	"dim lavender": {
		"muted tan": 4,
	},
	"dim lime": {
		"clear gray": 5,
		"wavy gray":  2,
	},
	"dim magenta": {
		"dark violet": 3,
	},
	"dim maroon": {
		"clear violet": 5,
		"light bronze": 4,
	},
	"dim olive": {
		"clear cyan":     4,
		"light lavender": 5,
		"shiny tomato":   4,
	},
	"dim orange": {
		"dull olive": 2,
	},
	"dim plum": {
		"clear purple": 1,
		"light brown":  4,
	},
	"dim purple": {
		"drab purple": 5,
		"muted white": 4,
	},
	"dim red": {
		"clear bronze": 3,
	},
	"dim salmon": {
		"faded gray":     5,
		"striped salmon": 1,
	},
	"dim silver": {
		"pale coral": 5,
		"wavy coral": 3,
	},
	"dim tan": {
		"shiny tomato":     1,
		"striped lavender": 1,
	},
	"dim teal": {
		"dull magenta":   1,
		"light purple":   4,
		"plaid lavender": 4,
	},
	"dim tomato": {
		"drab gold":  4,
		"shiny aqua": 2,
	},
	"dim turquoise": {
		"pale violet": 1,
	},
	"dim violet": {
		"posh chartreuse": 2,
		"posh lime":       1,
		"shiny orange":    4,
	},
	"dim white": {
		"dull magenta":  2,
		"plaid tan":     5,
		"striped brown": 5,
	},
	"dim yellow": {
		"dark black": 4,
		"pale tan":   5,
	},
	"dotted aqua": {
		"light gray":    4,
		"light purple":  2,
		"mirrored lime": 5,
	},
	"dotted beige": {
		"drab indigo": 3,
	},
	"dotted black": {
		"mirrored aqua": 1,
		"plaid purple":  2,
		"posh beige":    5,
	},
	"dotted blue": {
		"muted purple": 3,
		"wavy beige":   2,
	},
	"dotted bronze": {
		"light orange": 5,
	},
	"dotted brown": {
		"dark purple": 3,
		"drab green":  5,
		"wavy maroon": 4,
	},
	"dotted chartreuse": {
		"bright blue": 3,
	},
	"dotted coral": {
		"bright red":        2,
		"dotted chartreuse": 4,
		"vibrant brown":     1,
		"vibrant white":     1,
	},
	"dotted crimson": {
		"dark chartreuse": 5,
	},
	"dotted cyan": {
		"shiny black": 4,
	},
	"dotted fuchsia": {
		"clear olive":     1,
		"faded turquoise": 1,
		"posh salmon":     4,
		"shiny lavender":  2,
	},
	"dotted gold": {
		"dull fuchsia": 3,
		"shiny red":    3,
	},
	"dotted gray": {
		"posh plum": 3,
	},
	"dotted green": {
		"bright salmon":  4,
		"dim crimson":    4,
		"plaid gray":     1,
		"shiny lavender": 1,
	},
	"dotted indigo": {
		"clear tomato": 2,
	},
	"dotted lavender": {
		"dull cyan": 3,
		"posh aqua": 2,
	},
	"dotted lime": {
		"pale yellow":   2,
		"shiny crimson": 3,
	},
	"dotted magenta": {
		"bright lime":  2,
		"dull maroon":  4,
		"faded cyan":   4,
		"plaid indigo": 4,
	},
	"dotted maroon": {
		"dull silver": 3,
		"posh aqua":   5,
	},
	"dotted olive": {
		"dotted bronze":    5,
		"dotted turquoise": 4,
		"pale red":         4,
		"pale yellow":      3,
	},
	"dotted orange": {
		"drab silver":    5,
		"mirrored olive": 1,
	},
	"dotted plum": {
		"clear brown":  3,
		"muted indigo": 3,
		"shiny bronze": 1,
	},
	"dotted purple": {
		"dim teal": 2,
	},
	"dotted red": {
		"light blue": 1,
	},
	"dotted salmon": {
		"dull beige":        2,
		"light turquoise":   3,
		"vibrant turquoise": 3,
	},
	"dotted silver": {
		"posh coral": 1,
	},
	"dotted tan": {
		"dim aqua":       5,
		"striped violet": 4,
	},
	"dotted teal": {
		"bright red": 5,
		"clear plum": 3,
		"dull coral": 1,
		"wavy red":   3,
	},
	"dotted tomato": {
		"mirrored lime": 1,
		"vibrant white": 2,
	},
	"dotted turquoise": {
		"clear plum":  5,
		"dark violet": 4,
		"muted gold":  3,
	},
	"dotted violet": {
		"bright black": 3,
	},
	"dotted white": {
		"mirrored fuchsia": 1,
		"mirrored white":   5,
		"muted crimson":    3,
	},
	"dotted yellow": {
		"plaid turquoise": 3,
		"posh black":      5,
		"posh yellow":     2,
		"striped blue":    4,
	},
	"drab aqua": {
		"dark silver":  1,
		"plaid purple": 4,
		"posh orange":  5,
		"wavy teal":    2,
	},
	"drab beige": {
		"bright tomato": 3,
		"pale red":      4,
		"pale violet":   4,
		"shiny bronze":  2,
	},
	"drab black": {
		"dotted chartreuse": 3,
		"faded tomato":      5,
	},
	"drab blue": {
		"mirrored blue": 1,
		"mirrored lime": 3,
	},
	"drab bronze": {
		"plaid maroon": 4,
	},
	"drab brown": {
		"dotted tomato":  4,
		"striped orange": 1,
	},
	"drab chartreuse": {
		"clear turquoise": 4,
		"dark brown":      2,
		"drab salmon":     2,
	},
	"drab coral": {
		"posh yellow":      3,
		"shiny chartreuse": 4,
		"wavy indigo":      3,
	},
	"drab crimson": {
		"dark bronze":  2,
		"dull tan":     3,
		"shiny indigo": 4,
	},
	"drab cyan": {
		"bright orange": 3,
		"dark white":    3,
		"drab teal":     1,
		"light aqua":    1,
	},
	"drab fuchsia": {
		"faded crimson":   3,
		"mirrored salmon": 2,
		"pale plum":       3,
		"posh orange":     4,
	},
	"drab gold": {
		"bright green":  2,
		"bright indigo": 5,
	},
	"drab gray": {
		"faded aqua":     5,
		"light green":    1,
		"mirrored white": 5,
		"shiny lavender": 5,
	},
	"drab green": {
		"vibrant yellow": 1,
	},
	"drab indigo": {
		"dim coral":      2,
		"drab green":     1,
		"shiny lavender": 2,
	},
	"drab lavender": {
		"drab crimson": 3,
		"drab cyan":    5,
		"muted purple": 1,
		"wavy red":     1,
	},
	"drab lime": {
		"bright green":  4,
		"drab cyan":     1,
		"drab lavender": 3,
		"pale crimson":  3,
	},
	"drab magenta": {
		"bright indigo": 4,
		"clear beige":   2,
		"drab cyan":     2,
		"wavy lavender": 3,
	},
	"drab maroon": {
		"clear beige":  2,
		"faded purple": 5,
		"wavy beige":   3,
	},
	"drab olive": {
		"dark aqua":     2,
		"striped brown": 4,
	},
	"drab orange": {
		"plaid salmon":  5,
		"vibrant beige": 2,
	},
	"drab plum": {
		"muted brown": 4,
		"muted plum":  5,
	},
	"drab purple": {
		"shiny indigo":   1,
		"striped yellow": 4,
	},
	"drab red": {
		"bright red":    5,
		"vibrant brown": 2,
	},
	"drab salmon": {
		"dark olive":       5,
		"dotted indigo":    1,
		"muted chartreuse": 3,
	},
	"drab silver": {
		"dim teal":     1,
		"faded cyan":   3,
		"shiny indigo": 1,
	},
	"drab tan": {
		"bright red":  5,
		"light green": 2,
		"shiny beige": 4,
	},
	"drab teal": {
		"light crimson": 2,
		"vibrant brown": 3,
		"vibrant tan":   3,
	},
	"drab tomato": {
		"bright red":     3,
		"dim teal":       3,
		"striped yellow": 4,
	},
	"drab turquoise": {
		"drab red":  5,
		"posh gray": 3,
	},
	"drab violet": {
		"faded red": 4,
	},
	"drab white": {
		"dull fuchsia":   2,
		"vibrant bronze": 1,
	},
	"drab yellow": {
		"dotted green":     2,
		"dotted turquoise": 3,
		"shiny teal":       4,
		"vibrant silver":   5,
	},
	"dull aqua": {
		"dark salmon":     1,
		"dotted coral":    5,
		"posh indigo":     3,
		"striped crimson": 3,
	},
	"dull beige": {
		"dark bronze":   4,
		"dull tan":      5,
		"wavy lavender": 1,
	},
	"dull black": {
		"light lavender": 3,
		"striped indigo": 2,
		"wavy teal":      1,
	},
	"dull blue": {
		"dark orange":   1,
		"light crimson": 4,
		"light indigo":  5,
		"light lime":    3,
	},
	"dull bronze": {
		"mirrored beige": 1,
		"muted violet":   5,
		"wavy lime":      1,
	},
	"dull brown": {
		"light lime": 4,
	},
	"dull chartreuse": {
		"bright silver": 2,
		"dotted lime":   4,
		"dull red":      3,
		"wavy maroon":   4,
	},
	"dull coral": {
		"dim teal":     2,
		"faded bronze": 3,
	},
	"dull crimson": {
		"dull cyan": 3,
	},
	"dull cyan": {
		"bright coral":  4,
		"light black":   3,
		"mirrored aqua": 5,
		"shiny aqua":    2,
	},
	"dull fuchsia": {
		"dotted olive": 5,
		"muted purple": 2,
	},
	"dull gold": {
		"dark indigo": 3,
	},
	"dull gray": {
		"dark gray":     1,
		"light indigo":  2,
		"striped green": 3,
		"wavy coral":    3,
	},
	"dull green": {
		"bright blue": 2,
		"dull red":    5,
	},
	"dull indigo": {
		"dotted tan": 3,
		"dull coral": 4,
		"light gray": 4,
	},
	"dull lavender": {
		"dotted maroon": 4,
		"drab black":    4,
		"dull cyan":     4,
		"muted brown":   3,
	},
	"dull lime": {
		"muted chartreuse": 3,
	},
	"dull maroon": {
		"dull magenta": 3,
		"dull silver":  5,
		"dull tan":     1,
		"faded cyan":   1,
	},
	"dull olive": {
		"striped black":   1,
		"vibrant magenta": 1,
	},
	"dull orange": {
		"dark bronze":    3,
		"mirrored brown": 2,
		"pale gold":      4,
		"posh brown":     5,
	},
	"dull plum": {
		"bright red":   5,
		"dull magenta": 4,
	},
	"dull purple": {
		"light gray":  1,
		"wavy salmon": 1,
		"wavy yellow": 3,
	},
	"dull red": {
		"clear brown": 5,
		"dark orange": 2,
		"pale salmon": 3,
		"plaid blue":  5,
	},
	"dull salmon": {
		"light indigo": 2,
		"plaid gold":   2,
	},
	"dull tan": {
		"faded cyan":   4,
		"light purple": 2,
		"shiny indigo": 4,
	},
	"dull teal": {
		"bright tomato": 2,
	},
	"dull tomato": {
		"dim aqua":      2,
		"faded red":     1,
		"posh lavender": 4,
	},
	"dull turquoise": {
		"light lavender": 1,
		"mirrored brown": 1,
		"plaid bronze":   5,
		"striped red":    4,
	},
	"dull violet": {
		"dim black":     4,
		"light blue":    4,
		"striped white": 5,
	},
	"dull white": {
		"clear tan":    3,
		"drab crimson": 2,
		"shiny gold":   5,
	},
	"dull yellow": {
		"dark tomato":    1,
		"pale tan":       5,
		"vibrant orange": 4,
	},
	"faded aqua": {
		"muted brown": 5,
	},
	"faded beige": {
		"drab lavender": 1,
		"drab red":      4,
		"vibrant white": 1,
	},
	"faded black": {
		"drab tomato":   1,
		"muted fuchsia": 4,
		"wavy coral":    4,
	},
	"faded blue": {
		"striped black":  5,
		"vibrant yellow": 1,
	},
	"faded bronze": {
		"faded cyan":  4,
		"muted red":   5,
		"shiny beige": 5,
	},
	"faded brown": {
		"faded turquoise": 5,
	},
	"faded chartreuse": {
		"dim coral":       3,
		"mirrored bronze": 1,
		"posh bronze":     3,
	},
	"faded coral": {
		"clear white":    5,
		"drab turquoise": 3,
		"mirrored green": 5,
		"shiny plum":     1,
	},
	"faded crimson": {
		"dotted olive": 4,
		"drab silver":  5,
		"muted purple": 2,
	},
	"faded fuchsia": {
		"bright lime": 4,
		"light lime":  2,
		"pale violet": 3,
	},
	"faded gold": {
		"dotted purple":  4,
		"vibrant maroon": 4,
	},
	"faded green": {
		"dull turquoise":   5,
		"mirrored crimson": 1,
		"pale orange":      3,
	},
	"faded indigo": {
		"dark chartreuse":  1,
		"dull green":       4,
		"mirrored magenta": 1,
	},
	"faded lavender": {
		"posh lavender":  1,
		"striped yellow": 2,
	},
	"faded lime": {
		"dark indigo":   3,
		"dim turquoise": 5,
		"muted silver":  1,
		"vibrant gray":  2,
	},
	"faded magenta": {
		"dim crimson": 4,
		"dull indigo": 1,
		"plaid aqua":  4,
	},
	"faded maroon": {
		"dark indigo":  5,
		"posh fuchsia": 1,
	},
	"faded olive": {
		"drab plum":  1,
		"light blue": 3,
		"pale aqua":  1,
	},
	"faded orange": {
		"dark bronze": 2,
	},
	"faded plum": {
		"bright aqua": 3,
		"clear teal":  1,
	},
	"faded purple": {
		"bright orange": 4,
		"faded violet":  2,
	},
	"faded red": {
		"clear cyan":       5,
		"dull coral":       1,
		"plaid chartreuse": 4,
		"plaid orange":     3,
	},
	"faded salmon": {
		"wavy gray": 5,
	},
	"faded silver": {
		"mirrored cyan": 1,
	},
	"faded tan": {
		"bright fuchsia": 2,
		"mirrored brown": 1,
		"muted fuchsia":  1,
		"shiny fuchsia":  5,
	},
	"faded teal": {
		"mirrored white": 4,
	},
	"faded tomato": {
		"muted red":      4,
		"striped yellow": 2,
	},
	"faded turquoise": {
		"dark white": 4,
		"light lime": 5,
	},
	"faded violet": {
		"bright orange": 1,
		"dull silver":   1,
		"dull tan":      2,
		"wavy coral":    2,
	},
	"faded white": {
		"dull cyan":   3,
		"wavy maroon": 2,
	},
	"faded yellow": {
		"bright turquoise": 2,
		"dark silver":      5,
		"faded gold":       4,
		"wavy coral":       3,
	},
	"light aqua": {
		"plaid lavender": 1,
		"shiny indigo":   5,
		"wavy coral":     3,
	},
	"light beige": {
		"dotted olive": 4,
		"dull olive":   5,
		"faded orange": 1,
	},
	"light black": {
		"dark chartreuse": 1,
		"faded tomato":    3,
	},
	"light blue": {
		"dull white":     4,
		"posh bronze":    1,
		"striped bronze": 2,
	},
	"light bronze": {
		"clear plum": 4,
	},
	"light brown": {
		"mirrored brown": 3,
	},
	"light chartreuse": {
		"drab cyan":     4,
		"dull plum":     3,
		"plaid purple":  5,
		"striped brown": 4,
	},
	"light coral": {
		"bright cyan":   2,
		"muted silver":  1,
		"pale red":      2,
		"shiny fuchsia": 4,
	},
	"light crimson": {
		"dull maroon": 5,
		"muted gold":  4,
	},
	"light cyan": {
		"mirrored silver": 4,
		"mirrored tomato": 2,
		"vibrant coral":   2,
		"vibrant violet":  3,
	},
	"light fuchsia": {
		"bright lavender": 3,
		"clear brown":     3,
	},
	"light gold": {
		"pale plum": 2,
	},
	"light gray": {
		"pale red":       2,
		"vibrant yellow": 4,
	},
	"light green": {
		"bright red":    4,
		"vibrant coral": 1,
	},
	"light indigo": {
		"bright aqua":  5,
		"faded cyan":   1,
		"shiny indigo": 1,
	},
	"light lavender": {
		"muted violet": 1,
	},
	"light lime": {
		"plaid indigo": 4,
	},
	"light magenta": {
		"pale salmon":  5,
		"plaid maroon": 2,
		"posh beige":   1,
	},
	"light maroon": {
		"dull lavender":      4,
		"mirrored turquoise": 2,
		"muted yellow":       4,
		"pale cyan":          1,
	},
	"light olive": {
		"dim turquoise": 3,
		"drab white":    4,
		"dull gold":     4,
		"shiny coral":   1,
	},
	"light orange": {
		"bright red":     2,
		"dark white":     4,
		"dim chartreuse": 1,
		"faded fuchsia":  5,
	},
	"light plum": {
		"bright bronze": 1,
		"drab bronze":   3,
		"posh yellow":   1,
	},
	"light red": {
		"shiny chartreuse": 4,
	},
	"light salmon": {
		"dark silver": 5,
	},
	"light silver": {
		"dotted lavender": 3,
		"wavy lavender":   2,
	},
	"light tan": {
		"light black":    1,
		"posh turquoise": 2,
		"vibrant white":  3,
	},
	"light teal": {
		"dark olive":         5,
		"light yellow":       5,
		"plaid salmon":       3,
		"vibrant chartreuse": 5,
	},
	"light tomato": {
		"dotted bronze": 1,
		"dotted gold":   2,
	},
	"light turquoise": {
		"dotted tan":  5,
		"dull salmon": 2,
	},
	"light violet": {
		"dim silver":      4,
		"mirrored bronze": 2,
		"pale red":        1,
	},
	"light white": {
		"bright green": 5,
	},
	"light yellow": {
		"shiny crimson": 4,
		"striped red":   2,
	},
	"mirrored aqua": {
		"dotted purple": 5,
		"faded violet":  1,
	},
	"mirrored beige": {
		"dark cyan":  4,
		"dull green": 5,
	},
	"mirrored black": {
		"plaid chartreuse": 3,
		"shiny beige":      2,
		"shiny indigo":     2,
	},
	"mirrored blue": {
		"vibrant white": 2,
	},
	"mirrored bronze": {
		"faded fuchsia": 2,
	},
	"mirrored brown": {
		"dotted magenta": 5,
		"dull maroon":    4,
		"faded violet":   1,
	},
	"mirrored chartreuse": {
		"bright silver":   4,
		"mirrored tomato": 3,
		"wavy orange":     4,
	},
	"mirrored coral": {
		"clear indigo": 2,
		"drab crimson": 4,
		"light salmon": 2,
	},
	"mirrored crimson": {
		"mirrored magenta": 3,
		"posh maroon":      3,
		"striped olive":    5,
	},
	"mirrored cyan": {
		"clear cyan": 2,
		"drab blue":  5,
		"drab gold":  1,
		"light aqua": 4,
	},
	"mirrored fuchsia": {
		"faded beige": 1,
		"muted gold":  3,
	},
	"mirrored gold": {
		"mirrored white": 4,
	},
	"mirrored gray": {
		"dim indigo": 5,
	},
	"mirrored green": {
		"dotted lime": 5,
		"pale cyan":   4,
	},
	"mirrored indigo": {
		"light tomato":  3,
		"shiny crimson": 2,
	},
	"mirrored lavender": {
		"light purple":  4,
		"light white":   1,
		"shiny lime":    2,
		"wavy lavender": 3,
	},
	"mirrored lime": {
		"dull tan":    2,
		"shiny beige": 3,
	},
	"mirrored magenta": {
		"bright violet": 1,
		"drab lavender": 2,
		"pale violet":   3,
		"plaid salmon":  2,
	},
	"mirrored maroon": {
		"dull blue": 3,
	},
	"mirrored olive": {
		"dull plum": 1,
	},
	"mirrored orange": {
		"vibrant gray": 4,
	},
	"mirrored plum": {
		"clear gold": 1,
		"pale green": 5,
	},
	"mirrored purple": {
		"dim magenta": 3,
	},
	"mirrored red": {
		"dark silver":    1,
		"dull red":       1,
		"mirrored olive": 4,
	},
	"mirrored salmon": {
		"bright red":  5,
		"clear black": 3,
		"light green": 2,
		"posh brown":  5,
	},
	"mirrored silver": {
		"drab silver": 1,
		"shiny beige": 3,
	},
	"mirrored tan": {
		"clear teal":      4,
		"dim tomato":      2,
		"dotted olive":    5,
		"mirrored tomato": 2,
	},
	"mirrored teal": {
		"mirrored aqua": 4,
	},
	"mirrored tomato": {
		"dotted chartreuse": 2,
		"light aqua":        3,
		"posh beige":        3,
	},
	"mirrored turquoise": {
		"dark fuchsia":  3,
		"dim turquoise": 2,
		"plaid white":   5,
		"wavy indigo":   5,
	},
	"mirrored violet": {
		"dim silver":   3,
		"light salmon": 1,
		"posh tomato":  1,
	},
	"mirrored white": {
		"clear gold":     5,
		"dotted bronze":  4,
		"drab tomato":    3,
		"striped orange": 3,
	},
	"mirrored yellow": {
		"clear tomato":   5,
		"dotted gray":    5,
		"dull maroon":    4,
		"striped violet": 2,
	},
	"muted aqua": {
		"dull beige": 4,
		"dull red":   4,
	},
	"muted beige": {
		"dull blue":  2,
		"posh lime":  2,
		"shiny gray": 5,
	},
	"muted black": {
		"dotted chartreuse": 2,
		"faded cyan":        5,
		"pale tan":          5,
	},
	"muted blue": {
		"dim indigo":       2,
		"dotted silver":    4,
		"vibrant lavender": 1,
	},
	"muted bronze": {
		"dim green":     4,
		"mirrored aqua": 5,
	},
	"muted brown": {
		"dotted magenta": 4,
		"dull coral":     3,
		"posh blue":      1,
	},
	"muted chartreuse": {
		"clear tan": 2,
	},
	"muted coral": {
		"dotted chartreuse": 1,
		"shiny beige":       1,
		"wavy violet":       3,
	},
	"muted crimson": {
		"muted red":      1,
		"pale yellow":    5,
		"posh blue":      3,
		"striped orange": 4,
	},
	"muted cyan": {
		"shiny tomato": 3,
	},
	"muted fuchsia": {
		"dull tan": 1,
	},
	"muted gold": {
		"dull maroon":  4,
		"dull tan":     2,
		"faded bronze": 1,
	},
	"muted gray": {
		"clear salmon":   5,
		"faded aqua":     3,
		"shiny olive":    1,
		"vibrant violet": 1,
	},
	"muted green": {
		"clear cyan":   5,
		"dull magenta": 2,
	},
	"muted indigo": {
		"dotted teal":      2,
		"muted chartreuse": 4,
	},
	"muted lavender": {
		"light green":    1,
		"posh crimson":   3,
		"striped silver": 4,
	},
	"muted lime": {
		"muted tan": 1,
	},
	"muted magenta": {
		"pale violet":     4,
		"posh chartreuse": 2,
	},
	"muted maroon": {
		"bright aqua":    3,
		"dim chartreuse": 4,
		"dull salmon":    2,
		"faded indigo":   2,
	},
	"muted olive": {
		"muted magenta": 1,
		"plaid orange":  5,
	},
	"muted orange": {
		"dull coral":      2,
		"dull magenta":    4,
		"faded lime":      4,
		"vibrant magenta": 5,
	},
	"muted plum": {
		"bright black":  2,
		"dotted tomato": 3,
		"vibrant brown": 2,
	},
	"muted purple": {
		"clear cyan": 4,
	},
	"muted red": {
		"bright aqua":  1,
		"dull magenta": 4,
		"dull tan":     3,
	},
	"muted salmon": {
		"drab tan":    2,
		"wavy tomato": 5,
	},
	"muted silver": {
		"dim teal":     2,
		"muted bronze": 5,
		"wavy blue":    3,
	},
	"muted tan": {
		"drab gold":    1,
		"plaid purple": 1,
		"shiny beige":  3,
	},
	"muted teal": {
		"dull silver": 4,
		"pale violet": 1,
	},
	"muted tomato": {
		"clear fuchsia":  4,
		"drab red":       4,
		"vibrant silver": 3,
		"wavy white":     3,
	},
	"muted turquoise": {
		"bright teal":   2,
		"clear crimson": 4,
		"faded beige":   5,
	},
	"muted violet": {
		"light chartreuse": 2,
	},
	"muted white": {
		"dull olive": 5,
	},
	"muted yellow": {
		"dotted silver":   3,
		"faded turquoise": 3,
		"light purple":    5,
	},
	"no other bags": {},
	"pale aqua": {
		"dark lavender": 3,
		"posh white":    3,
		"shiny yellow":  5,
	},
	"pale beige": {
		"bright silver": 5,
	},
	"pale black": {
		"mirrored gold": 1,
	},
	"pale blue": {
		"light lavender": 1,
		"muted gold":     2,
		"vibrant brown":  4,
	},
	"pale bronze": {
		"clear crimson": 1,
	},
	"pale brown": {
		"light aqua": 5,
	},
	"pale chartreuse": {
		"clear bronze":  2,
		"shiny crimson": 4,
	},
	"pale coral": {
		"posh gray": 5,
	},
	"pale crimson": {
		"dark violet":    1,
		"plaid lavender": 3,
		"plaid orange":   1,
	},
	"pale cyan": {
		"wavy salmon":    5,
		"wavy turquoise": 2,
	},
	"pale fuchsia": {
		"dark brown": 3,
		"posh gold":  5,
	},
	"pale gold": {
		"striped brown": 4,
	},
	"pale gray": {
		"dim green":     3,
		"faded black":   1,
		"posh brown":    2,
		"wavy lavender": 4,
	},
	"pale green": {
		"light chartreuse": 4,
	},
	"pale indigo": {
		"dotted lavender": 2,
		"mirrored beige":  2,
		"striped green":   4,
		"wavy turquoise":  5,
	},
	"pale lavender": {
		"plaid magenta": 4,
		"striped blue":  4,
	},
	"pale lime": {
		"bright lavender": 4,
	},
	"pale magenta": {
		"posh silver": 4,
	},
	"pale maroon": {
		"pale bronze": 2,
	},
	"pale olive": {
		"plaid lime":  4,
		"posh salmon": 1,
	},
	"pale orange": {
		"bright green": 2,
		"clear salmon": 2,
		"posh fuchsia": 2,
	},
	"pale plum": {
		"bright coral":   1,
		"dotted fuchsia": 1,
		"drab tan":       1,
		"plaid fuchsia":  1,
	},
	"pale purple": {
		"posh silver": 4,
		"wavy yellow": 4,
	},
	"pale red": {
		"faded gray": 2,
	},
	"pale salmon": {
		"dull magenta":  3,
		"light crimson": 1,
	},
	"pale silver": {
		"clear beige":      4,
		"striped gold":     4,
		"vibrant lavender": 4,
	},
	"pale tan": {
		"dotted chartreuse": 2,
		"striped green":     4,
	},
	"pale teal": {
		"dotted olive":   2,
		"drab turquoise": 2,
		"shiny white":    4,
		"wavy maroon":    2,
	},
	"pale tomato": {
		"dim chartreuse": 2,
		"dim teal":       1,
		"drab cyan":      5,
		"muted coral":    3,
	},
	"pale turquoise": {
		"dark tan": 4,
	},
	"pale violet": {
		"dull tan": 2,
	},
	"pale white": {
		"bright brown":   3,
		"mirrored beige": 4,
	},
	"pale yellow": {
		"dim aqua":      5,
		"vibrant brown": 1,
		"vibrant tan":   2,
		"wavy lavender": 3,
	},
	"plaid aqua": {
		"drab green": 4,
		"muted cyan": 3,
		"pale beige": 4,
		"wavy gray":  2,
	},
	"plaid beige": {
		"dull green": 3,
	},
	"plaid black": {
		"posh gray": 2,
	},
	"plaid blue": {
		"muted gold":    4,
		"plaid purple":  1,
		"striped green": 5,
	},
	"plaid bronze": {
		"drab indigo":  2,
		"plaid orange": 5,
	},
	"plaid brown": {
		"vibrant turquoise": 5,
	},
	"plaid chartreuse": {
		"dim chartreuse": 5,
		"dim coral":      2,
		"drab silver":    4,
		"light purple":   3,
	},
	"plaid coral": {
		"dotted green": 2,
		"light beige":  3,
		"striped tan":  1,
		"vibrant cyan": 3,
	},
	"plaid crimson": {
		"dull fuchsia":  2,
		"mirrored teal": 2,
	},
	"plaid cyan": {
		"dim red":        3,
		"dotted fuchsia": 3,
		"plaid fuchsia":  4,
		"shiny teal":     2,
	},
	"plaid fuchsia": {
		"dark coral": 3,
	},
	"plaid gold": {
		"bright beige":   1,
		"clear lavender": 1,
	},
	"plaid gray": {
		"dull coral": 2,
		"shiny gold": 3,
	},
	"plaid green": {
		"mirrored plum":   1,
		"mirrored tan":    4,
		"mirrored tomato": 1,
		"pale indigo":     5,
	},
	"plaid indigo": {
		"dark brown":   3,
		"light aqua":   4,
		"light green":  3,
		"light purple": 5,
	},
	"plaid lavender": {
		"dull magenta": 1,
		"dull silver":  2,
		"shiny indigo": 1,
	},
	"plaid lime": {
		"dotted turquoise": 1,
		"posh beige":       5,
	},
	"plaid magenta": {
		"clear lavender": 3,
		"mirrored brown": 5,
		"shiny bronze":   5,
	},
	"plaid maroon": {
		"bright beige":   2,
		"muted teal":     2,
		"pale yellow":    2,
		"striped yellow": 5,
	},
	"plaid olive": {
		"dim green":       4,
		"faded turquoise": 5,
		"striped beige":   2,
	},
	"plaid orange": {
		"dull green":       1,
		"light crimson":    4,
		"plaid chartreuse": 4,
		"vibrant coral":    5,
	},
	"plaid plum": {
		"clear brown": 4,
		"dotted lime": 1,
		"dull red":    4,
	},
	"plaid purple": {
		"light lime": 1,
	},
	"plaid red": {
		"faded fuchsia": 3,
		"light crimson": 4,
		"pale blue":     5,
	},
	"plaid salmon": {
		"dim coral":   5,
		"wavy bronze": 2,
	},
	"plaid silver": {
		"clear yellow": 4,
		"dotted brown": 1,
		"drab gold":    3,
	},
	"plaid tan": {
		"dim crimson": 5,
	},
	"plaid teal": {
		"dark tan":  2,
		"pale lime": 3,
	},
	"plaid tomato": {
		"faded chartreuse": 3,
		"faded white":      1,
		"mirrored maroon":  3,
		"wavy salmon":      2,
	},
	"plaid turquoise": {
		"dotted teal": 5,
	},
	"plaid violet": {
		"drab violet":     2,
		"muted indigo":    5,
		"muted turquoise": 4,
	},
	"plaid white": {
		"drab turquoise": 4,
	},
	"plaid yellow": {
		"bright red": 2,
		"clear teal": 4,
	},
	"posh aqua": {
		"drab lime":   4,
		"dull tan":    4,
		"vibrant tan": 5,
	},
	"posh beige": {
		"dotted magenta": 3,
		"drab blue":      3,
		"muted coral":    2,
		"shiny black":    5,
	},
	"posh black": {
		"bright silver": 3,
		"dim teal":      5,
		"posh beige":    5,
		"striped black": 2,
	},
	"posh blue": {
		"dull maroon":   4,
		"vibrant coral": 3,
	},
	"posh bronze": {
		"bright indigo": 5,
		"dark orange":   2,
		"dark violet":   1,
		"dotted purple": 5,
	},
	"posh brown": {
		"bright green":  3,
		"posh lavender": 5,
	},
	"posh chartreuse": {
		"dim teal":   3,
		"light aqua": 5,
	},
	"posh coral": {
		"dark aqua":   3,
		"dim aqua":    4,
		"pale yellow": 2,
		"plaid blue":  5,
	},
	"posh crimson": {
		"bright coral": 2,
		"dull magenta": 5,
		"striped blue": 4,
	},
	"posh cyan": {
		"bright violet": 4,
	},
	"posh fuchsia": {
		"dotted olive":   2,
		"mirrored beige": 2,
		"shiny bronze":   3,
	},
	"posh gold": {
		"dark aqua":    5,
		"dark tan":     5,
		"dotted olive": 2,
	},
	"posh gray": {
		"bright blue":   5,
		"bright orange": 5,
	},
	"posh green": {
		"clear plum":    4,
		"dark gray":     2,
		"dull lavender": 2,
	},
	"posh indigo": {
		"dotted coral": 2,
		"faded violet": 5,
		"shiny orange": 2,
		"wavy green":   3,
	},
	"posh lavender": {
		"dark brown": 4,
	},
	"posh lime": {
		"posh chartreuse": 3,
	},
	"posh magenta": {
		"mirrored black": 5,
	},
	"posh maroon": {
		"dull red": 2,
	},
	"posh olive": {
		"bright red": 3,
	},
	"posh orange": {
		"pale plum": 4,
		"posh aqua": 5,
	},
	"posh plum": {
		"dark chartreuse":  3,
		"dotted turquoise": 3,
		"dull maroon":      4,
		"posh bronze":      1,
	},
	"posh purple": {
		"dim gray":    5,
		"faded brown": 3,
	},
	"posh red": {
		"bright black": 3,
	},
	"posh salmon": {
		"bright red":    5,
		"dark brown":    3,
		"striped green": 3,
	},
	"posh silver": {
		"pale violet":  3,
		"plaid salmon": 1,
		"posh coral":   1,
	},
	"posh tan": {
		"drab magenta": 3,
		"shiny yellow": 1,
	},
	"posh teal": {
		"clear plum": 4,
	},
	"posh tomato": {
		"mirrored teal":      4,
		"vibrant chartreuse": 4,
	},
	"posh turquoise": {
		"dull salmon":        3,
		"muted gold":         1,
		"striped lavender":   5,
		"vibrant chartreuse": 5,
	},
	"posh violet": {
		"dim aqua": 2,
	},
	"posh white": {
		"dotted silver": 2,
		"drab violet":   4,
	},
	"posh yellow": {
		"dark bronze":    1,
		"light crimson":  2,
		"mirrored white": 2,
	},
	"shiny aqua": {
		"dark white": 1,
	},
	"shiny beige": {
		"bright lime":  3,
		"dull magenta": 5,
		"dull silver":  2,
	},
	"shiny black": {
		"plaid chartreuse": 3,
	},
	"shiny blue": {
		"dull tan":      3,
		"muted magenta": 5,
	},
	"shiny bronze": {
		"light green":     1,
		"vibrant magenta": 4,
	},
	"shiny brown": {
		"bright brown": 5,
		"clear tan":    5,
		"faded cyan":   3,
		"plaid maroon": 2,
	},
	"shiny chartreuse": {
		"dotted purple": 2,
		"plaid beige":   3,
		"wavy red":      5,
	},
	"shiny coral": {
		"clear plum":       1,
		"muted chartreuse": 5,
		"muted violet":     1,
		"striped yellow":   5,
	},
	"shiny crimson": {
		"faded fuchsia": 3,
		"light green":   3,
		"striped brown": 5,
	},
	"shiny cyan": {
		"clear gray":   2,
		"pale crimson": 5,
	},
	"shiny fuchsia": {
		"drab indigo": 3,
	},
	"shiny gold": {
		"dark white":      5,
		"dull magenta":    1,
		"faded turquoise": 4,
	},
	"shiny gray": {
		"bright gold":   1,
		"muted bronze":  1,
		"shiny teal":    5,
		"striped green": 5,
	},
	"shiny green": {
		"dull maroon": 5,
	},
	"shiny lavender": {
		"bright aqua": 5,
	},
	"shiny lime": {
		"dark aqua": 4,
	},
	"shiny magenta": {
		"drab silver": 5,
		"muted gold":  2,
	},
	"shiny maroon": {
		"posh gold": 1,
	},
	"shiny olive": {
		"dark violet":   3,
		"mirrored lime": 2,
		"striped gold":  1,
	},
	"shiny orange": {
		"muted purple": 2,
	},
	"shiny plum": {
		"bright blue":  3,
		"dark tomato":  5,
		"vibrant gray": 5,
	},
	"shiny purple": {
		"clear tan": 2,
	},
	"shiny red": {
		"dull chartreuse": 4,
		"pale tomato":     5,
	},
	"shiny salmon": {
		"drab turquoise": 5,
		"shiny gold":     1,
	},
	"shiny silver": {
		"dim green":  2,
		"dull red":   3,
		"plaid gold": 2,
		"plaid lime": 5,
	},
	"shiny tan": {
		"dull brown":    1,
		"striped coral": 5,
	},
	"shiny teal": {
		"faded tomato": 1,
		"muted violet": 4,
	},
	"shiny tomato": {
		"dull beige": 5,
		"pale coral": 5,
	},
	"shiny turquoise": {
		"dull red":        2,
		"faded gray":      5,
		"muted turquoise": 1,
	},
	"shiny violet": {
		"dotted tomato": 3,
		"drab coral":    1,
	},
	"shiny white": {
		"bright tan": 4,
	},
	"shiny yellow": {
		"dim teal":     3,
		"plaid tan":    1,
		"posh plum":    5,
		"vibrant plum": 2,
	},
	"striped aqua": {
		"dull red":    3,
		"posh salmon": 3,
	},
	"striped beige": {
		"shiny magenta": 2,
	},
	"striped black": {
		"dotted teal":     4,
		"plaid beige":     2,
		"plaid salmon":    1,
		"posh chartreuse": 2,
	},
	"striped blue": {
		"clear turquoise": 1,
		"dotted fuchsia":  1,
		"dull bronze":     1,
		"light lavender":  1,
	},
	"striped bronze": {
		"clear tan":     5,
		"drab lavender": 1,
		"pale crimson":  2,
	},
	"striped brown": {
		"bright orange": 3,
	},
	"striped chartreuse": {
		"dull magenta": 2,
		"posh tomato":  1,
	},
	"striped coral": {
		"dim aqua": 2,
	},
	"striped crimson": {
		"faded gray": 2,
	},
	"striped cyan": {
		"dark white":   2,
		"drab red":     4,
		"plaid salmon": 2,
	},
	"striped fuchsia": {
		"posh gold":      4,
		"wavy green":     5,
		"wavy turquoise": 1,
	},
	"striped gold": {
		"bright lavender": 5,
		"dark coral":      5,
		"dim teal":        5,
		"dull silver":     2,
	},
	"striped gray": {
		"plaid orange": 5,
		"wavy coral":   1,
	},
	"striped green": {
		"dark brown":    3,
		"shiny gold":    4,
		"vibrant brown": 5,
		"wavy coral":    4,
	},
	"striped indigo": {
		"shiny green":  1,
		"vibrant gold": 1,
	},
	"striped lavender": {
		"striped brown": 3,
	},
	"striped lime": {
		"light tan": 4,
	},
	"striped magenta": {
		"dark brown":  5,
		"posh silver": 1,
	},
	"striped maroon": {
		"shiny yellow": 2,
	},
	"striped olive": {
		"clear lavender": 4,
		"dim chartreuse": 5,
		"plaid indigo":   4,
	},
	"striped orange": {
		"dotted coral":  5,
		"light crimson": 5,
		"muted tan":     3,
		"plaid blue":    3,
	},
	"striped plum": {
		"clear tan":     3,
		"drab beige":    4,
		"light aqua":    5,
		"shiny fuchsia": 1,
	},
	"striped purple": {
		"dim indigo": 1,
	},
	"striped red": {
		"clear cyan": 1,
		"drab green": 4,
	},
	"striped salmon": {
		"bright red": 3,
		"light aqua": 1,
		"wavy gold":  4,
	},
	"striped silver": {
		"drab tomato": 3,
	},
	"striped tan": {
		"dim teal":         5,
		"dotted turquoise": 4,
		"drab lime":        4,
	},
	"striped teal": {
		"clear black": 2,
		"pale coral":  3,
	},
	"striped tomato": {
		"drab tan":    4,
		"light blue":  4,
		"striped tan": 2,
	},
	"striped turquoise": {
		"muted indigo": 3,
	},
	"striped violet": {
		"dim teal":         3,
		"mirrored fuchsia": 4,
		"vibrant white":    4,
	},
	"striped white": {
		"pale coral":  2,
		"shiny green": 2,
	},
	"striped yellow": {
		"bright red":      5,
		"dull maroon":     2,
		"posh chartreuse": 2,
	},
	"vibrant aqua": {
		"faded lavender": 2,
	},
	"vibrant beige": {
		"pale teal":       3,
		"posh chartreuse": 4,
		"wavy lavender":   2,
		"wavy lime":       5,
	},
	"vibrant black": {
		"faded cyan":      5,
		"vibrant magenta": 2,
	},
	"vibrant blue": {
		"dark tomato": 3,
		"faded cyan":  3,
		"shiny lime":  1,
	},
	"vibrant bronze": {
		"bright tan":  5,
		"drab tomato": 3,
	},
	"vibrant brown": {
		"muted red":     4,
		"shiny indigo":  2,
		"striped brown": 3,
	},
	"vibrant chartreuse": {
		"striped orange": 4,
	},
	"vibrant coral": {
		"bright orange": 1,
	},
	"vibrant crimson": {
		"dotted fuchsia": 4,
	},
	"vibrant cyan": {
		"wavy coral": 5,
	},
	"vibrant fuchsia": {
		"shiny magenta": 5,
		"vibrant beige": 1,
	},
	"vibrant gold": {
		"clear chartreuse": 1,
		"dull orange":      2,
	},
	"vibrant gray": {
		"posh brown": 3,
	},
	"vibrant green": {
		"shiny fuchsia": 3,
	},
	"vibrant indigo": {
		"dim black": 2,
		"dim blue":  4,
		"dim white": 3,
	},
	"vibrant lavender": {
		"drab gold":     3,
		"faded red":     4,
		"plaid tan":     1,
		"vibrant brown": 3,
	},
	"vibrant lime": {
		"posh coral": 5,
	},
	"vibrant magenta": {
		"dim green": 1,
	},
	"vibrant maroon": {
		"bright indigo":   4,
		"dark lavender":   1,
		"mirrored maroon": 5,
		"pale crimson":    1,
	},
	"vibrant olive": {
		"vibrant coral": 2,
	},
	"vibrant orange": {
		"light chartreuse": 4,
		"posh crimson":     5,
	},
	"vibrant plum": {
		"muted chartreuse": 4,
		"posh silver":      4,
	},
	"vibrant purple": {
		"faded orange": 5,
	},
	"vibrant red": {
		"clear brown":     1,
		"drab orange":     3,
		"plaid turquoise": 1,
		"striped fuchsia": 4,
	},
	"vibrant salmon": {
		"muted olive": 2,
	},
	"vibrant silver": {
		"bright fuchsia": 1,
		"dotted teal":    3,
		"drab lavender":  3,
		"drab olive":     2,
	},
	"vibrant tan": {
		"dull plum":     2,
		"striped brown": 1,
		"vibrant coral": 4,
	},
	"vibrant teal": {
		"dotted turquoise": 5,
		"drab tan":         1,
		"striped tan":      1,
	},
	"vibrant tomato": {
		"dark silver":  2,
		"plaid purple": 3,
	},
	"vibrant turquoise": {
		"dim teal":      5,
		"shiny beige":   3,
		"striped brown": 3,
	},
	"vibrant violet": {
		"bright red": 5,
	},
	"vibrant white": {
		"drab red":       1,
		"plaid lavender": 5,
		"vibrant brown":  4,
	},
	"vibrant yellow": {
		"light chartreuse": 4,
		"shiny aqua":       2,
	},
	"wavy aqua": {
		"dark yellow": 1,
		"drab indigo": 3,
	},
	"wavy beige": {
		"dark bronze":     1,
		"dull plum":       4,
		"mirrored silver": 4,
	},
	"wavy black": {
		"clear orange": 2,
		"muted purple": 4,
	},
	"wavy blue": {
		"dull blue": 1,
		"wavy teal": 4,
	},
	"wavy bronze": {
		"bright silver": 5,
		"dark brown":    2,
		"dark orange":   3,
	},
	"wavy brown": {
		"drab crimson": 1,
		"wavy indigo":  2,
	},
	"wavy chartreuse": {
		"shiny yellow": 4,
		"vibrant blue": 4,
	},
	"wavy crimson": {
		"bright blue":  4,
		"dull coral":   2,
		"light purple": 1,
		"wavy beige":   1,
	},
	"wavy cyan": {
		"bright fuchsia":  3,
		"clear gold":      3,
		"light bronze":    2,
		"mirrored salmon": 1,
	},
	"wavy fuchsia": {
		"bright aqua":  4,
		"bright brown": 4,
		"light orange": 5,
	},
	"wavy gold": {
		"bright tan":  3,
		"shiny olive": 4,
	},
	"wavy gray": {
		"muted gray":    4,
		"posh teal":     1,
		"shiny magenta": 3,
	},
	"wavy green": {
		"drab blue":   4,
		"dull blue":   1,
		"pale coral":  5,
		"striped tan": 1,
	},
	"wavy indigo": {
		"clear gold":      3,
		"mirrored bronze": 5,
	},
	"wavy lavender": {
		"faded bronze": 4,
		"light lime":   3,
		"muted gold":   4,
		"muted red":    4,
	},
	"wavy lime": {
		"dark coral": 2,
		"shiny gold": 1,
	},
	"wavy magenta": {
		"muted silver": 3,
		"pale bronze":  5,
		"pale plum":    5,
	},
	"wavy maroon": {
		"dull salmon":     2,
		"mirrored bronze": 4,
		"shiny crimson":   3,
		"striped olive":   1,
	},
	"wavy olive": {
		"bright gold":      1,
		"dotted turquoise": 4,
		"dull silver":      4,
	},
	"wavy orange": {
		"dark bronze": 5,
	},
	"wavy plum": {
		"clear lavender": 1,
		"faded brown":    2,
	},
	"wavy purple": {
		"clear maroon":    3,
		"mirrored green":  1,
		"pale black":      1,
		"striped magenta": 5,
	},
	"wavy red": {
		"bright indigo": 4,
		"vibrant brown": 3,
	},
	"wavy salmon": {
		"bright aqua": 4,
	},
	"wavy silver": {
		"dotted white": 4,
	},
	"wavy tan": {
		"faded gray": 3,
	},
	"wavy teal": {
		"drab cyan":      2,
		"striped orange": 4,
	},
	"wavy tomato": {
		"bright brown":   1,
		"striped red":    3,
		"vibrant maroon": 2,
	},
	"wavy turquoise": {
		"dim white":   4,
		"dotted gray": 1,
		"drab olive":  1,
		"shiny lime":  5,
	},
	"wavy violet": {
		"bright aqua":    1,
		"dull coral":     3,
		"shiny lavender": 3,
	},
	"wavy white": {
		"drab lime":       5,
		"plaid lime":      4,
		"posh maroon":     4,
		"vibrant crimson": 1,
	},
	"wavy yellow": {
		"dark gray":    5,
		"faded red":    4,
		"light indigo": 5,
		"plaid indigo": 5,
	},
}
