package parse

import "regexp"

var (
	numberMatcher       = regexp.MustCompile("^[0-9]+(\\.[0-9]+)?$")
	alphanumericMatcher = regexp.MustCompile("[A-Za-z0-9]")
	identifierMatcher   = regexp.MustCompile("^[A-Za-z][A-Za-z0-9]*$")
)

func MatchOperatorToken(input string) bool {
	return input == TOKEN_PLUS ||
		input == TOKEN_MINUS ||
		input == TOKEN_DIVIDE ||
		input == TOKEN_MULTIPLY
}

func MatchSyntaxToken(input string) bool {
	return input == TOKEN_ASSIGN ||
		input == TOKEN_LPAREN ||
		input == TOKEN_RPAREN
}

func MatchWhitespaceTokens(input string) bool {
	return input == TOKEN_SPACE || input == TOKEN_NEWLINE
}

func MatchNotWhitespaceTokens(input string) bool {
	return !MatchWhitespaceTokens(string(input))
}

func MatchNumber(input string) bool {
	return numberMatcher.MatchString(input)
}

func MatchAlphanumeric(input string) bool {
	return alphanumericMatcher.MatchString(input)
}

func MatchIdentifier(input string) bool {
	return identifierMatcher.MatchString(input)
}

func MatchTokenWordDelimiter(input string) bool {
	return !MatchWhitespaceTokens(input) && !MatchSyntaxToken(input)
}
