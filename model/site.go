package model

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"strings"
	"time"
)

type Site struct {
	ID      primitive.ObjectID `bson:"_id"`
	URL     string             `bson:"u"`
	Address string             `bson:"a"`
	HTML    string             `bson:"h"`
	Online  bool               `bson:"o"`
	Domain  domain             `bson:"d"`
}

type domain struct {
	Registrar        string    `bson:"r"`
	RegistrationDate time.Time `bson:"rd"`
}

func (s Site) validate() error {
	err := validation.ValidateStruct(
		&s,
		validation.Field(&s.ID, validation.Required),
		validation.Field(&s.URL, validation.Required),
		validation.Field(&s.Address, validation.Required),
		validation.Field(&s.HTML, validation.Required),
		validation.Field(&s.Online, validation.Required),
	)
	if err != nil {
		return err
	}

	return validation.ValidateStruct(
		&s.Domain,
		validation.Field(&s.Domain.Registrar, validation.Required),
		validation.Field(&s.Domain.RegistrationDate, validation.Required),
	)
}

func (s Site) Create(url, registrar string, registrationDate time.Time) {
	uri := strings.Join([]string{
		"http://",
		url,
	}, "")

	doc := Site{
		ID:  primitive.NewObjectID(),
		URL: uri,
		Domain: domain{
			Registrar:        registrar,
			RegistrationDate: registrationDate,
		},
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)
	res := fasthttp.AcquireResponse()

	client := fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		ReadTimeout:              3 * time.Second,
		WriteTimeout:             3 * time.Second,
		MaxConnWaitTimeout:       3 * time.Second,
		MaxResponseBodySize:      math.MaxInt64,
	}
	err := client.DoRedirects(req, res, 25)
	if err != nil {
		_, err := mongo.Sites.InsertOne(context.Background(), doc)
		logger.Must(err)

		logger.Log.Info().Err(err).
			Bool("online", doc.Online).
			Str("url", doc.URL).
			Msg("website saved")
		return
	}

	location := uri
	if l := string(res.Header.Peek("location")); l != "" {
		location = l
	}
	if l := string(res.Header.Peek("Location")); l != "" {
		location = l
	}

	doc.Online = true
	doc.URL = location
	doc.HTML = string(res.Body())
	doc.Address = res.RemoteAddr().String()
	err = doc.validate()
	logger.Must(err)

	_, err = mongo.Sites.InsertOne(context.Background(), doc)
	logger.Must(err)
	logger.Log.Info().
		Bool("online", doc.Online).
		Str("url", doc.URL).
		Msg("website saved")
	return
}
