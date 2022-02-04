package main

import (
	"path"
	"sort"
	"strings"
	"time"
)

type liveRadio struct {
	Slug    string
	Name    string
	Owner   string
	URL     string
	Cards   []string
	Links   []socialLink
	AddedAt time.Time
}

func (r liveRadio) youtubeID() string {
	return path.Base(r.URL)
}

func getLiveRadios() []liveRadio {
	radios := []liveRadio{
		{
			Slug:  "lofi-girl",
			Name:  "Lo-Fi Girl",
			Owner: "LoFi Girl",
			URL:   "https://youtu.be/5qap5aO4i9A",
			Cards: []string{
				"지혜는 듣는 데서 오고 후회는 말하는 데서 온다.",
				"말을 많이 하는 것과 말을 잘하는 것은 다르다.",
				"거짓말을 한 순간부터 뛰어난 기억력이 필요하다. <코르네이유>",
			},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/LofiGirl",
				},
			},
		},
		// {
		// 	Slug:  "everything-fades-to-blue",
		// 	Name:  "Everything Fades To Blue",
		// 	Owner: "Sleepy Fish",
		// 	URL:   "https://youtu.be/PfgS405CdXk",
		// 	Cards: []string{
		// 		"Everything Fades To Blue is a mix of indie/emo and ambient music produced by Sleep Fish, a Pennsylvania-based producer that is also a student in statistics, data science, and machine learning.",
		// 		"Sleepy Fish is one of the few Lo-fi artists who actually sing in its creations.",
		// 		"Everything Fades To Blue is the 3rd episode of a story where a tidal wave destroys an island along with the home where Sleepy Fish used to live.",
		// 		"Toppled into the sea, on its own for the first time, Sleepy Fish uses its glow to search for family, to guide others, and to find its way.",
		// 		`The episode comes after "My Room Becomes the Sea" and "Beneath Your Waves".`,
		// 		"The undersea-themed animation has been made in collaboration with Tristan Gion and Bien à Vous Studio, a French studio based in Nantes.",
		// 	},
		// 	Links: []socialLink{
		// 		{
		// 			Slug: "website",
		// 			URL:  "https://chillhop.com/releases/sleepy-fish-everything-fades-to-blue",
		// 		},
		// 		{
		// 			Slug: "youtube",
		// 			URL:  "https://youtu.be/PfgS405CdXk",
		// 		},
		// 		{
		// 			Slug: "spotify",
		// 			URL:  "https://open.spotify.com/artist/1IJe80moz409PtxW4llPFw",
		// 		},
		// 		{
		// 			Slug: "instagram",
		// 			URL:  "https://www.instagram.com/sleepyfishmusic",
		// 		},
		// 	},
		// },
	}

	sort.Slice(radios, func(a, b int) bool {
		return strings.Compare(radios[a].Name, radios[b].Name) < 0
	})

	for _, r := range radios {
		sort.Slice(r.Links, func(a, b int) bool {
			return strings.Compare(r.Links[a].Slug, r.Links[b].Slug) < 0
		})
	}

	return radios
}

type socialLink struct {
	Slug string
	URL  string
}

func socialIcon(slug string) string {
	switch slug {
	case "youtube":
		return youtubeSVG

	case "reddit":
		return redditSVG

	case "facebook":
		return facebookSVG

	case "instagram":
		return instagramSVG

	case "twitter":
		return twitterSVG

	case "spotify":
		return spotifySVG

	case "discord":
		return discordSVG

	case "website":
		return websiteSVG

	default:
		return linkSVG
	}
}
