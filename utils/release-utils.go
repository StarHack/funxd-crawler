package utils

import (
    "regexp"
)

var pattern1 = regexp.MustCompile("Release:.*?([A-Za-z0-9\\.-]+-FuN).*?(https://www\\.filecrypt\\.cc/[a-zA-Z0-9/\\.]+)")
var pattern2 = regexp.MustCompile("Release:.*?(https://www\\.filecrypt\\.cc/[a-zA-Z0-9/\\.]+).*?([A-Za-z0-9\\.-]+-FuN)")

type ReleaseUtils struct {

}

func NewReleaseUtils() *ReleaseUtils {
    var n ReleaseUtils
    return &n
}

func (instance *ReleaseUtils) ExtractAllReleases(html string) map[string]string {
    releases := make(map[string]string)
    
    matchedPattern1 := pattern1.FindAllStringSubmatch(html, -1)
    for _, match := range(matchedPattern1) {
        releaseName := match[1]
        releaseLink := match[2]
        releases[releaseName] = releaseLink
    }

    matchedPattern2 := pattern2.FindAllStringSubmatch(html, -1)
    for _, match := range(matchedPattern2) {
        releaseLink := match[1]
        releaseName := match[2]
        releases[releaseName] = releaseLink
    }

    return releases
}
