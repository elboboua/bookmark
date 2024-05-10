package bookservice

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)
const GOOGLE_BOOK_API_URL = "https://www.googleapis.com/books/v1"

type GoogleBookService struct {
	apiKey string
	client http.Client
}

func (gbs *GoogleBookService) SearchBooks(ctx context.Context, query string) ([]Book, error) {
	urlEscapedQuery := url.QueryEscape(query)
	status, response, err := gbs.request(ctx, "GET", "/volumes?q="+urlEscapedQuery)
	if err != nil {
		return nil, errors.New("failed searching for books "+err.Error())
	}
	if status != 200 {
		return nil, errors.New("status code searching for books not 200")
	}

	var gbRes GoogleBookQueryResponse
	err = json.Unmarshal(response, &gbRes)
	if err != nil {
		return nil, err
	}

	books := []Book{}
	for _, gbook := range gbRes.Items {
		books = append(books, Book{
			ID: gbook.ID,
			Title: gbook.VolumeInfo.Title,
			Subtitle: gbook.VolumeInfo.Subtitle,
			Description: gbook.VolumeInfo.Description,
			Authors: gbook.VolumeInfo.Authors,
			PurchaseLink: gbook.SaleInfo.BuyLink,
			PublishedDate: gbook.VolumeInfo.PublishedDate,
			Language: gbook.VolumeInfo.Language,
			PageCount: gbook.VolumeInfo.PageCount,
			Categories: gbook.VolumeInfo.Categories,
		})
	}

	return books, nil
}

func (gbs *GoogleBookService) request(_ context.Context, method string, path string) (status int, res []byte, err error) {
	url := GOOGLE_BOOK_API_URL + path
	if strings.Contains(url, "?") {
		url += "&" + gbs.apiKey
	} else {
		url += "?" + gbs.apiKey
	}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, nil, errors.New("unable to send request to google book api: "+err.Error())
	}

	response, err := gbs.client.Do(request)
	if err != nil {
		return 0, nil, errors.New("unable to perform request to google book api: "+err.Error())
	}
	defer response.Body.Close()

	res, err = io.ReadAll(response.Body)	

	return response.StatusCode, res, nil
}

type GoogleBookQueryResponse struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []struct {
		Kind       string `json:"kind"`
		ID         string `json:"id"`
		Etag       string `json:"etag"`
		SelfLink   string `json:"selfLink"`
		VolumeInfo struct {
			Title               string   `json:"title"`
			Subtitle            string   `json:"subtitle"`
			Authors             []string `json:"authors"`
			Publisher           string   `json:"publisher"`
			PublishedDate       string   `json:"publishedDate"`
			Description         string   `json:"description"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
			ReadingModes struct {
				Text  bool `json:"text"`
				Image bool `json:"image"`
			} `json:"readingModes"`
			PageCount           int      `json:"pageCount"`
			PrintType           string   `json:"printType"`
			Categories          []string `json:"categories"`
			MaturityRating      string   `json:"maturityRating"`
			AllowAnonLogging    bool     `json:"allowAnonLogging"`
			ContentVersion      string   `json:"contentVersion"`
			PanelizationSummary struct {
				ContainsEpubBubbles  bool `json:"containsEpubBubbles"`
				ContainsImageBubbles bool `json:"containsImageBubbles"`
			} `json:"panelizationSummary"`
			ImageLinks struct {
				SmallThumbnail string `json:"smallThumbnail"`
				Thumbnail      string `json:"thumbnail"`
			} `json:"imageLinks"`
			Language            string `json:"language"`
			PreviewLink         string `json:"previewLink"`
			InfoLink            string `json:"infoLink"`
			CanonicalVolumeLink string `json:"canonicalVolumeLink"`
		} `json:"volumeInfo"`
		SaleInfo struct {
			Country     string `json:"country"`
			Saleability string `json:"saleability"`
			IsEbook     bool   `json:"isEbook"`
			ListPrice   struct {
				Amount       float64 `json:"amount"`
				CurrencyCode string  `json:"currencyCode"`
			} `json:"listPrice"`
			RetailPrice struct {
				Amount       float64 `json:"amount"`
				CurrencyCode string  `json:"currencyCode"`
			} `json:"retailPrice"`
			BuyLink string `json:"buyLink"`
			Offers  []struct {
				FinskyOfferType int `json:"finskyOfferType"`
				ListPrice       struct {
					AmountInMicros int    `json:"amountInMicros"`
					CurrencyCode   string `json:"currencyCode"`
				} `json:"listPrice"`
				RetailPrice struct {
					AmountInMicros int    `json:"amountInMicros"`
					CurrencyCode   string `json:"currencyCode"`
				} `json:"retailPrice"`
				Giftable bool `json:"giftable"`
			} `json:"offers"`
		} `json:"saleInfo"`
		AccessInfo struct {
			Country                string `json:"country"`
			Viewability            string `json:"viewability"`
			Embeddable             bool   `json:"embeddable"`
			PublicDomain           bool   `json:"publicDomain"`
			TextToSpeechPermission string `json:"textToSpeechPermission"`
			Epub                   struct {
				IsAvailable  bool   `json:"isAvailable"`
				AcsTokenLink string `json:"acsTokenLink"`
			} `json:"epub"`
			Pdf struct {
				IsAvailable bool `json:"isAvailable"`
			} `json:"pdf"`
			WebReaderLink       string `json:"webReaderLink"`
			AccessViewStatus    string `json:"accessViewStatus"`
			QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
		} `json:"accessInfo"`
		SearchInfo struct {
			TextSnippet string `json:"textSnippet"`
		} `json:"searchInfo"`
	} `json:"items"`
}