package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		if actualUrl, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, actualUrl, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

type pathAndUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func parseYAML(yml []byte) ([]pathAndUrl, error) {
	var pathAndUrls []pathAndUrl
	err := yaml.Unmarshal(yml, &pathAndUrls)
	if err != nil {
		return nil, err
	}
	return pathAndUrls, nil
}

func convertToMap(parsedYaml []pathAndUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range parsedYaml {
		pathsToUrls[pu.Path] = pu.Url
	}
	return pathsToUrls
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathsToUrls := convertToMap(parsedYaml)
	return MapHandler(pathsToUrls, fallback), nil
}
