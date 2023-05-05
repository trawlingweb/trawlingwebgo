package models

import "time"

type Tweet struct {
	ID                      string    `json:"id"`
	Crawled                 int64     `json:"crawled"`
	Published               time.Time `json:"published"`
	PostID                  int64     `json:"post_id"`
	URL                     string    `json:"url"`
	Text                    string    `json:"text"`
	Lang                    string    `json:"lang"`
	IsQuoteStatus           bool      `json:"is_quote_status"`
	RetweetCount            int       `json:"retweet_count"`
	ReplyCount              int       `json:"reply_count"`
	QuoteCount              int       `json:"quote_count"`
	FavoriteCount           int       `json:"favorite_count"`
	Retweet                 bool      `json:"retweet"`
	UserID                  int64     `json:"user_id"`
	UserName                string    `json:"user_name"`
	UserScreenName          string    `json:"user_screen_name"`
	UserLocation            string    `json:"user_location"`
	UserURL                 string    `json:"user_url"`
	UserSite                string    `json:"user_site"`
	UserDescription         string    `json:"user_description"`
	UserProtected           bool      `json:"user_protected"`
	UserVerified            bool      `json:"user_verified"`
	UserFollowersCount      int       `json:"user_followers_count"`
	UserFriendsCount        int       `json:"user_friends_count"`
	UserListedCount         int       `json:"user_listed_count"`
	UserFavouritesCount     int       `json:"user_favourites_count"`
	UserStatusesCount       int       `json:"user_statuses_count"`
	UserCreatedAt           time.Time `json:"user_created_at"`
	UserGeoEnabled          bool      `json:"user_geo_enabled"`
	UserContributorsEnabled bool      `json:"user_contributors_enabled"`
	UserProfileImageURL     string    `json:"user_profile_image_url"`
	UserProfileBannerURL    string    `json:"user_profile_banner_url"`
	Hashtags                []string  `json:"hashtags"`
	UserMentions            []string  `json:"user_mentions"`
}

type Article struct {
	Crawled      int64     `json:"crawled"`
	Author       string    `json:"author"`
	Language     string    `json:"language"`
	Published    time.Time `json:"published"`
	Title        string    `json:"title"`
	URL          string    `json:"url"`
	Section      string    `json:"section"`
	Value        float64   `json:"value"`
	Rank         int64     `json:"rank"`
	Visitors     int64     `json:"unique_visitors"`
	Site         string    `json:"site"`
	SiteType     string    `json:"site_type"`
	SiteSection  string    `json:"site_section"`
	SiteLanguage string    `json:"site_language"`
	SiteRegion   string    `json:"site_region"`
	SiteCountry  string    `json:"site_country"`
	Domain       string    `json:"domain"`
	Text         string    `json:"text"`
	ID           string    `json:"id"`
}
