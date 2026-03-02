package main

import (
	"time"
)

type Launch struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		ID               string `json:"id"`
		URL              string `json:"url"`
		Name             string `json:"name"`
		ResponseMode     string `json:"response_mode"`
		Slug             string `json:"slug"`
		LaunchDesignator any    `json:"launch_designator"`
		Status           struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Abbrev      string `json:"abbrev"`
			Description string `json:"description"`
		} `json:"status"`
		LastUpdated  time.Time `json:"last_updated"`
		Net          time.Time `json:"net"`
		NetPrecision struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Abbrev      string `json:"abbrev"`
			Description string `json:"description"`
		} `json:"net_precision"`
		WindowEnd   time.Time `json:"window_end"`
		WindowStart time.Time `json:"window_start"`
		Image       struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			ImageURL     string `json:"image_url"`
			ThumbnailURL string `json:"thumbnail_url"`
			Credit       string `json:"credit"`
			License      struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Priority int    `json:"priority"`
				Link     string `json:"link"`
			} `json:"license"`
			SingleUse bool  `json:"single_use"`
			Variants  []any `json:"variants"`
		} `json:"image"`
		Infographic           any    `json:"infographic"`
		Probability           any    `json:"probability"`
		WeatherConcerns       any    `json:"weather_concerns"`
		Failreason            string `json:"failreason"`
		Hashtag               any    `json:"hashtag"`
		LaunchServiceProvider struct {
			ResponseMode string `json:"response_mode"`
			ID           int    `json:"id"`
			URL          string `json:"url"`
			Name         string `json:"name"`
			Abbrev       string `json:"abbrev"`
			Type         struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"type"`
		} `json:"launch_service_provider"`
		Rocket struct {
			ID            int `json:"id"`
			Configuration struct {
				ResponseMode string `json:"response_mode"`
				ID           int    `json:"id"`
				URL          string `json:"url"`
				Name         string `json:"name"`
				Families     []struct {
					ResponseMode string `json:"response_mode"`
					ID           int    `json:"id"`
					Name         string `json:"name"`
				} `json:"families"`
				FullName string `json:"full_name"`
				Variant  string `json:"variant"`
			} `json:"configuration"`
		} `json:"rocket"`
		Mission struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
			Image       any    `json:"image"`
			Orbit       struct {
				ID            int    `json:"id"`
				Name          string `json:"name"`
				Abbrev        string `json:"abbrev"`
				CelestialBody struct {
					ResponseMode string `json:"response_mode"`
					ID           int    `json:"id"`
					Name         string `json:"name"`
				} `json:"celestial_body"`
			} `json:"orbit"`
			Agencies []struct {
				ResponseMode string `json:"response_mode"`
				ID           int    `json:"id"`
				URL          string `json:"url"`
				Name         string `json:"name"`
				Abbrev       string `json:"abbrev"`
				Type         struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"type"`
				Featured bool `json:"featured"`
				Country  []struct {
					ID                      int    `json:"id"`
					Name                    string `json:"name"`
					Alpha2Code              string `json:"alpha_2_code"`
					Alpha3Code              string `json:"alpha_3_code"`
					NationalityName         string `json:"nationality_name"`
					NationalityNameComposed string `json:"nationality_name_composed"`
				} `json:"country"`
				Description   string `json:"description"`
				Administrator string `json:"administrator"`
				FoundingYear  int    `json:"founding_year"`
				Launchers     string `json:"launchers"`
				Spacecraft    string `json:"spacecraft"`
				Parent        any    `json:"parent"`
				Image         struct {
					ID           int    `json:"id"`
					Name         string `json:"name"`
					ImageURL     string `json:"image_url"`
					ThumbnailURL string `json:"thumbnail_url"`
					Credit       any    `json:"credit"`
					License      struct {
						ID       int    `json:"id"`
						Name     string `json:"name"`
						Priority int    `json:"priority"`
						Link     any    `json:"link"`
					} `json:"license"`
					SingleUse bool  `json:"single_use"`
					Variants  []any `json:"variants"`
				} `json:"image"`
				Logo struct {
					ID           int    `json:"id"`
					Name         string `json:"name"`
					ImageURL     string `json:"image_url"`
					ThumbnailURL string `json:"thumbnail_url"`
					Credit       any    `json:"credit"`
					License      struct {
						ID       int    `json:"id"`
						Name     string `json:"name"`
						Priority int    `json:"priority"`
						Link     any    `json:"link"`
					} `json:"license"`
					SingleUse bool  `json:"single_use"`
					Variants  []any `json:"variants"`
				} `json:"logo"`
				SocialLogo struct {
					ID           int    `json:"id"`
					Name         string `json:"name"`
					ImageURL     string `json:"image_url"`
					ThumbnailURL string `json:"thumbnail_url"`
					Credit       any    `json:"credit"`
					License      struct {
						ID       int    `json:"id"`
						Name     string `json:"name"`
						Priority int    `json:"priority"`
						Link     any    `json:"link"`
					} `json:"license"`
					SingleUse bool  `json:"single_use"`
					Variants  []any `json:"variants"`
				} `json:"social_logo"`
				TotalLaunchCount              int    `json:"total_launch_count"`
				ConsecutiveSuccessfulLaunches int    `json:"consecutive_successful_launches"`
				SuccessfulLaunches            int    `json:"successful_launches"`
				FailedLaunches                int    `json:"failed_launches"`
				PendingLaunches               int    `json:"pending_launches"`
				ConsecutiveSuccessfulLandings int    `json:"consecutive_successful_landings"`
				SuccessfulLandings            int    `json:"successful_landings"`
				FailedLandings                int    `json:"failed_landings"`
				AttemptedLandings             int    `json:"attempted_landings"`
				SuccessfulLandingsSpacecraft  int    `json:"successful_landings_spacecraft"`
				FailedLandingsSpacecraft      int    `json:"failed_landings_spacecraft"`
				AttemptedLandingsSpacecraft   int    `json:"attempted_landings_spacecraft"`
				SuccessfulLandingsPayload     int    `json:"successful_landings_payload"`
				FailedLandingsPayload         int    `json:"failed_landings_payload"`
				AttemptedLandingsPayload      int    `json:"attempted_landings_payload"`
				InfoURL                       string `json:"info_url"`
				WikiURL                       string `json:"wiki_url"`
				SocialMediaLinks              []any  `json:"social_media_links"`
			} `json:"agencies"`
			InfoUrls []any `json:"info_urls"`
			VidUrls  []any `json:"vid_urls"`
		} `json:"mission"`
		Pad struct {
			ID       int    `json:"id"`
			URL      string `json:"url"`
			Active   bool   `json:"active"`
			Agencies []any  `json:"agencies"`
			Name     string `json:"name"`
			Image    struct {
				ID           int    `json:"id"`
				Name         string `json:"name"`
				ImageURL     string `json:"image_url"`
				ThumbnailURL string `json:"thumbnail_url"`
				Credit       string `json:"credit"`
				License      struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Priority int    `json:"priority"`
					Link     string `json:"link"`
				} `json:"license"`
				SingleUse bool  `json:"single_use"`
				Variants  []any `json:"variants"`
			} `json:"image"`
			Description string  `json:"description"`
			InfoURL     any     `json:"info_url"`
			WikiURL     string  `json:"wiki_url"`
			MapURL      string  `json:"map_url"`
			Latitude    float64 `json:"latitude"`
			Longitude   float64 `json:"longitude"`
			Country     struct {
				ID                      int    `json:"id"`
				Name                    string `json:"name"`
				Alpha2Code              string `json:"alpha_2_code"`
				Alpha3Code              string `json:"alpha_3_code"`
				NationalityName         string `json:"nationality_name"`
				NationalityNameComposed string `json:"nationality_name_composed"`
			} `json:"country"`
			MapImage                  string `json:"map_image"`
			TotalLaunchCount          int    `json:"total_launch_count"`
			OrbitalLaunchAttemptCount int    `json:"orbital_launch_attempt_count"`
			FastestTurnaround         string `json:"fastest_turnaround"`
			Location                  struct {
				ResponseMode  string `json:"response_mode"`
				ID            int    `json:"id"`
				URL           string `json:"url"`
				Name          string `json:"name"`
				CelestialBody struct {
					ResponseMode string `json:"response_mode"`
					ID           int    `json:"id"`
					Name         string `json:"name"`
					Type         struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"type"`
					Diameter    float64 `json:"diameter"`
					Mass        float64 `json:"mass"`
					Gravity     float64 `json:"gravity"`
					LengthOfDay string  `json:"length_of_day"`
					Atmosphere  bool    `json:"atmosphere"`
					Image       struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						ImageURL     string `json:"image_url"`
						ThumbnailURL string `json:"thumbnail_url"`
						Credit       string `json:"credit"`
						License      struct {
							ID       int    `json:"id"`
							Name     string `json:"name"`
							Priority int    `json:"priority"`
							Link     string `json:"link"`
						} `json:"license"`
						SingleUse bool  `json:"single_use"`
						Variants  []any `json:"variants"`
					} `json:"image"`
					Description            string `json:"description"`
					WikiURL                string `json:"wiki_url"`
					TotalAttemptedLaunches int    `json:"total_attempted_launches"`
					SuccessfulLaunches     int    `json:"successful_launches"`
					FailedLaunches         int    `json:"failed_launches"`
					TotalAttemptedLandings int    `json:"total_attempted_landings"`
					SuccessfulLandings     int    `json:"successful_landings"`
					FailedLandings         int    `json:"failed_landings"`
				} `json:"celestial_body"`
				Active  bool `json:"active"`
				Country struct {
					ID                      int    `json:"id"`
					Name                    string `json:"name"`
					Alpha2Code              string `json:"alpha_2_code"`
					Alpha3Code              string `json:"alpha_3_code"`
					NationalityName         string `json:"nationality_name"`
					NationalityNameComposed string `json:"nationality_name_composed"`
				} `json:"country"`
				Description string `json:"description"`
				Image       struct {
					ID           int    `json:"id"`
					Name         string `json:"name"`
					ImageURL     string `json:"image_url"`
					ThumbnailURL string `json:"thumbnail_url"`
					Credit       string `json:"credit"`
					License      struct {
						ID       int    `json:"id"`
						Name     string `json:"name"`
						Priority int    `json:"priority"`
						Link     string `json:"link"`
					} `json:"license"`
					SingleUse bool  `json:"single_use"`
					Variants  []any `json:"variants"`
				} `json:"image"`
				MapImage          string  `json:"map_image"`
				Longitude         float64 `json:"longitude"`
				Latitude          float64 `json:"latitude"`
				TimezoneName      string  `json:"timezone_name"`
				TotalLaunchCount  int     `json:"total_launch_count"`
				TotalLandingCount int     `json:"total_landing_count"`
			} `json:"location"`
		} `json:"pad"`
		WebcastLive bool `json:"webcast_live"`
		Program     []struct {
			ResponseMode string `json:"response_mode"`
			ID           int    `json:"id"`
			URL          string `json:"url"`
			Name         string `json:"name"`
			Image        struct {
				ID           int    `json:"id"`
				Name         string `json:"name"`
				ImageURL     string `json:"image_url"`
				ThumbnailURL string `json:"thumbnail_url"`
				Credit       string `json:"credit"`
				License      struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Priority int    `json:"priority"`
					Link     string `json:"link"`
				} `json:"license"`
				SingleUse bool  `json:"single_use"`
				Variants  []any `json:"variants"`
			} `json:"image"`
			InfoURL     string `json:"info_url"`
			WikiURL     string `json:"wiki_url"`
			Description string `json:"description"`
			Agencies    []struct {
				ResponseMode string `json:"response_mode"`
				ID           int    `json:"id"`
				URL          string `json:"url"`
				Name         string `json:"name"`
				Abbrev       string `json:"abbrev"`
				Type         struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"type"`
			} `json:"agencies"`
			StartDate      time.Time `json:"start_date"`
			EndDate        any       `json:"end_date"`
			MissionPatches []struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Priority int    `json:"priority"`
				ImageURL string `json:"image_url"`
				Agency   struct {
					ResponseMode string `json:"response_mode"`
					ID           int    `json:"id"`
					URL          string `json:"url"`
					Name         string `json:"name"`
					Abbrev       string `json:"abbrev"`
					Type         struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"type"`
				} `json:"agency"`
				ResponseMode string `json:"response_mode"`
			} `json:"mission_patches"`
			Type struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"type"`
		} `json:"program"`
		OrbitalLaunchAttemptCount      int `json:"orbital_launch_attempt_count"`
		LocationLaunchAttemptCount     int `json:"location_launch_attempt_count"`
		PadLaunchAttemptCount          int `json:"pad_launch_attempt_count"`
		AgencyLaunchAttemptCount       int `json:"agency_launch_attempt_count"`
		OrbitalLaunchAttemptCountYear  int `json:"orbital_launch_attempt_count_year"`
		LocationLaunchAttemptCountYear int `json:"location_launch_attempt_count_year"`
		PadLaunchAttemptCountYear      int `json:"pad_launch_attempt_count_year"`
		AgencyLaunchAttemptCountYear   int `json:"agency_launch_attempt_count_year"`
	} `json:"results"`
}
