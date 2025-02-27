//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package websearch

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/websearch"

const (
	DefaultEndpoint = original.DefaultEndpoint
)

type AnswerType = original.AnswerType

const (
	AnswerTypeComputation      AnswerType = original.AnswerTypeComputation
	AnswerTypeImages           AnswerType = original.AnswerTypeImages
	AnswerTypeNews             AnswerType = original.AnswerTypeNews
	AnswerTypeRelatedSearches  AnswerType = original.AnswerTypeRelatedSearches
	AnswerTypeSpellSuggestions AnswerType = original.AnswerTypeSpellSuggestions
	AnswerTypeTimeZone         AnswerType = original.AnswerTypeTimeZone
	AnswerTypeVideos           AnswerType = original.AnswerTypeVideos
	AnswerTypeWebPages         AnswerType = original.AnswerTypeWebPages
)

type ErrorCode = original.ErrorCode

const (
	InsufficientAuthorization ErrorCode = original.InsufficientAuthorization
	InvalidAuthorization      ErrorCode = original.InvalidAuthorization
	InvalidRequest            ErrorCode = original.InvalidRequest
	None                      ErrorCode = original.None
	RateLimitExceeded         ErrorCode = original.RateLimitExceeded
	ServerError               ErrorCode = original.ServerError
)

type ErrorSubCode = original.ErrorSubCode

const (
	AuthorizationDisabled   ErrorSubCode = original.AuthorizationDisabled
	AuthorizationExpired    ErrorSubCode = original.AuthorizationExpired
	AuthorizationMissing    ErrorSubCode = original.AuthorizationMissing
	AuthorizationRedundancy ErrorSubCode = original.AuthorizationRedundancy
	Blocked                 ErrorSubCode = original.Blocked
	HTTPNotAllowed          ErrorSubCode = original.HTTPNotAllowed
	NotImplemented          ErrorSubCode = original.NotImplemented
	ParameterInvalidValue   ErrorSubCode = original.ParameterInvalidValue
	ParameterMissing        ErrorSubCode = original.ParameterMissing
	ResourceError           ErrorSubCode = original.ResourceError
	UnexpectedError         ErrorSubCode = original.UnexpectedError
)

type Freshness = original.Freshness

const (
	Day   Freshness = original.Day
	Month Freshness = original.Month
	Week  Freshness = original.Week
)

type SafeSearch = original.SafeSearch

const (
	Moderate SafeSearch = original.Moderate
	Off      SafeSearch = original.Off
	Strict   SafeSearch = original.Strict
)

type TextFormat = original.TextFormat

const (
	HTML TextFormat = original.HTML
	Raw  TextFormat = original.Raw
)

type Type = original.Type

const (
	TypeWebWebGrouping Type = original.TypeWebWebGrouping
)

type TypeBasicResponseBase = original.TypeBasicResponseBase

const (
	TypeAnswer                             TypeBasicResponseBase = original.TypeAnswer
	TypeArticle                            TypeBasicResponseBase = original.TypeArticle
	TypeComputation                        TypeBasicResponseBase = original.TypeComputation
	TypeCreativeWork                       TypeBasicResponseBase = original.TypeCreativeWork
	TypeErrorResponse                      TypeBasicResponseBase = original.TypeErrorResponse
	TypeIdentifiable                       TypeBasicResponseBase = original.TypeIdentifiable
	TypeImageObject                        TypeBasicResponseBase = original.TypeImageObject
	TypeImages                             TypeBasicResponseBase = original.TypeImages
	TypeIntangible                         TypeBasicResponseBase = original.TypeIntangible
	TypeMediaObject                        TypeBasicResponseBase = original.TypeMediaObject
	TypeNews                               TypeBasicResponseBase = original.TypeNews
	TypeNewsArticle                        TypeBasicResponseBase = original.TypeNewsArticle
	TypePlaces                             TypeBasicResponseBase = original.TypePlaces
	TypeRelatedSearchesRelatedSearchAnswer TypeBasicResponseBase = original.TypeRelatedSearchesRelatedSearchAnswer
	TypeResponse                           TypeBasicResponseBase = original.TypeResponse
	TypeResponseBase                       TypeBasicResponseBase = original.TypeResponseBase
	TypeSearchResponse                     TypeBasicResponseBase = original.TypeSearchResponse
	TypeSearchResultsAnswer                TypeBasicResponseBase = original.TypeSearchResultsAnswer
	TypeSpellSuggestions                   TypeBasicResponseBase = original.TypeSpellSuggestions
	TypeStructuredValue                    TypeBasicResponseBase = original.TypeStructuredValue
	TypeThing                              TypeBasicResponseBase = original.TypeThing
	TypeTimeZone                           TypeBasicResponseBase = original.TypeTimeZone
	TypeVideoObject                        TypeBasicResponseBase = original.TypeVideoObject
	TypeVideos                             TypeBasicResponseBase = original.TypeVideos
	TypeWebPage                            TypeBasicResponseBase = original.TypeWebPage
	TypeWebWebAnswer                       TypeBasicResponseBase = original.TypeWebWebAnswer
)

type Answer = original.Answer
type Article = original.Article
type BaseClient = original.BaseClient
type BasicAnswer = original.BasicAnswer
type BasicArticle = original.BasicArticle
type BasicCreativeWork = original.BasicCreativeWork
type BasicIdentifiable = original.BasicIdentifiable
type BasicIntangible = original.BasicIntangible
type BasicMediaObject = original.BasicMediaObject
type BasicResponse = original.BasicResponse
type BasicResponseBase = original.BasicResponseBase
type BasicSearchResultsAnswer = original.BasicSearchResultsAnswer
type BasicThing = original.BasicThing
type BasicWebWebGrouping = original.BasicWebWebGrouping
type Computation = original.Computation
type CreativeWork = original.CreativeWork
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type Identifiable = original.Identifiable
type ImageObject = original.ImageObject
type Images = original.Images
type Intangible = original.Intangible
type MediaObject = original.MediaObject
type News = original.News
type NewsArticle = original.NewsArticle
type Places = original.Places
type Query = original.Query
type QueryContext = original.QueryContext
type RankingRankingGroup = original.RankingRankingGroup
type RankingRankingItem = original.RankingRankingItem
type RankingRankingResponse = original.RankingRankingResponse
type RelatedSearchesRelatedSearchAnswer = original.RelatedSearchesRelatedSearchAnswer
type Response = original.Response
type ResponseBase = original.ResponseBase
type SearchResponse = original.SearchResponse
type SearchResultsAnswer = original.SearchResultsAnswer
type SpellSuggestions = original.SpellSuggestions
type StructuredValue = original.StructuredValue
type Thing = original.Thing
type TimeZone = original.TimeZone
type TimeZoneTimeZoneInformation = original.TimeZoneTimeZoneInformation
type VideoObject = original.VideoObject
type Videos = original.Videos
type WebClient = original.WebClient
type WebMetaTag = original.WebMetaTag
type WebPage = original.WebPage
type WebWebAnswer = original.WebWebAnswer
type WebWebGrouping = original.WebWebGrouping

func New() BaseClient {
	return original.New()
}
func NewWebClient() WebClient {
	return original.NewWebClient()
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleAnswerTypeValues() []AnswerType {
	return original.PossibleAnswerTypeValues()
}
func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}
func PossibleErrorSubCodeValues() []ErrorSubCode {
	return original.PossibleErrorSubCodeValues()
}
func PossibleFreshnessValues() []Freshness {
	return original.PossibleFreshnessValues()
}
func PossibleSafeSearchValues() []SafeSearch {
	return original.PossibleSafeSearchValues()
}
func PossibleTextFormatValues() []TextFormat {
	return original.PossibleTextFormatValues()
}
func PossibleTypeBasicResponseBaseValues() []TypeBasicResponseBase {
	return original.PossibleTypeBasicResponseBaseValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
