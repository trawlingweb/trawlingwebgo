package models

import "time"

type Tweet struct {
	ID                   string  `json:"id"`
	Hash                 string  `json:"hash"`
	Published            string  `json:"published"`
	Crawled              int64   `json:"crawled"`
	Updated              int64   `json:"updated"`
	PostID               string  `json:"post_id"`
	URL                  string  `json:"url"`
	Text                 string  `json:"text"`
	Lang                 string  `json:"lang"`
	RetweetCount         int64   `json:"retweet_count"`
	ReplyCount           int64   `json:"reply_count"`
	FavoriteCount        int64   `json:"favorite_count"`
	ReproductionsCount   int64   `json:"reproductions_count"`
	EntitiesURL          string  `json:"entities_url"`
	URLImage             string  `json:"url_image"`
	Hashtags             string  `json:"hashtags"`
	UserMentions         string  `json:"user_mentions"`
	TimeDistance         float64 `json:"time_distance"`
	Reply                bool    `json:"reply"`
	UserID                  int64     `json:"user_id"`
	UserName             string  `json:"user_name"`
	UserScreenName       string  `json:"user_screen_name"`
	UserCreationDate     string  `json:"user_creation_date"`
	UserURL              string  `json:"user_url"`
	UserProfileImageURL  string  `json:"user_profile_image_url"`
	UserProfileBannerURL string  `json:"user_profile_banner_url"`
	UserDescription      string  `json:"user_description"`
	UserExternalURL      string  `json:"user_external_url"`
	UserLocation         string  `json:"user_location"`
	UserFollowerCount    int64   `json:"user_follower_count"`
	UserFollowingCount   int64   `json:"user_following_count"`
	UserFavouritesCount  int64   `json:"user_favourites_count"`
	UserIsPrivate        bool    `json:"user_is_private"`
	UserIsVerified       bool    `json:"user_is_verified"`
	UserIsBlueVerified   bool    `json:"user_is_blue_verified"`
	UserNumberOfTweets   int64   `json:"user_number_of_tweets"`
}

type SCRTweet struct {
   
	ID                   string  `json:"id"`
	Hash                 string  `json:"hash"`
	Published            string  `json:"published"`
	Crawled              int64   `json:"crawled"`
	Updated              int64   `json:"updated"`
	PostID               string  `json:"post_id"`
	URL                  string  `json:"url"`
	Text                 string  `json:"text"`
	Lang                 string  `json:"lang"`
	RetweetCount         int64   `json:"retweet_count"`
	ReplyCount           int64   `json:"reply_count"`
	FavoriteCount        int64   `json:"favorite_count"`
	ReproductionsCount   int64   `json:"reproductions_count"`
	EntitiesURL          string  `json:"entities_url"`
	URLImage             string  `json:"url_image"`
	Hashtags             string  `json:"hashtags"`
	UserMentions         string  `json:"user_mentions"`
	TimeDistance         float64 `json:"time_distance"`
	Reply                bool    `json:"reply"`
	UserID                  int64     `json:"user_id"`
	UserName             string  `json:"user_name"`
	UserScreenName       string  `json:"user_screen_name"`
	UserCreationDate     string  `json:"user_creation_date"`
	UserURL              string  `json:"user_url"`
	UserProfileImageURL  string  `json:"user_profile_image_url"`
	UserProfileBannerURL string  `json:"user_profile_banner_url"`
	UserDescription      string  `json:"user_description"`
	UserExternalURL      string  `json:"user_external_url"`
	UserLocation         string  `json:"user_location"`
	UserFollowerCount    int64   `json:"user_follower_count"`
	UserFollowingCount   int64   `json:"user_following_count"`
	UserFavouritesCount  int64   `json:"user_favourites_count"`
	UserIsPrivate        bool    `json:"user_is_private"`
	UserIsVerified       bool    `json:"user_is_verified"`
	UserIsBlueVerified   bool    `json:"user_is_blue_verified"`
	UserNumberOfTweets   int64   `json:"user_number_of_tweets"`
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
