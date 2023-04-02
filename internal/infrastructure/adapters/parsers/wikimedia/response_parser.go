package wikimedia

import (
	"regexp"
	"strings"

	"github.com/wikimedia/pkg/config"
)

type ShortDescriptionParser struct {
	Cfg *config.Config
}

func (p *ShortDescriptionParser) Parse(content string) (string, error) {
	var re = regexp.MustCompile(`{{[Ss]hort [Dd]escription\|(?P<description>[^|]+)}}`)
	match := re.FindStringSubmatch(content)

	shortDescription := ""
	// Extract the value of the Short description field
	if len(match) > 1 {
		index := re.SubexpIndex("description")
		stringMatch := match[index]
		shortDescription = strings.TrimSpace(stringMatch)
	}

	return shortDescription, nil
}
