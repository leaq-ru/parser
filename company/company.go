package company

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Direct link .jpg
type link = string

type Company struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	URL         string             `bson:"u,omitempty"`
	Slug        string             `bson:"s,omitempty"`
	Title       string             `bson:"t,omitempty"`
	Type        string             `bson:"ty,omitempty"`
	Email       string             `bson:"e,omitempty"`
	Description string             `bson:"d,omitempty"`
	Online      bool               `bson:"o,omitempty"`
	Phone       int                `bson:"p,omitempty"`
	INN         int                `bson:"i,omitempty"`
	KPP         int                `bson:"k,omitempty"`
	OGRN        int                `bson:"og,omitempty"`
	Domain      *domain            `bson:"do,omitempty"`
	Avatar      link               `bson:"a,omitempty"` // TODO download first image from site
	Location    *location          `bson:"l,omitempty"`
	App         *app               `bson:"ap,omitempty"`
	Social      *social            `bson:"so,omitempty"`
	People      []*peopleItem      `bson:"pe,omitempty"`
}

type peopleItem struct {
	VkID        int    `bson:"v,omitempty"`
	FirstName   string `bson:"f,omitempty"`
	LastName    string `bson:"l,omitempty"`
	VkIsClosed  bool   `bson:"vc,omitempty"`
	Sex         int8   `bson:"s,omitempty"`
	Photo200    link   `bson:"ph,omitempty"`
	Phone       int    `bson:"p,omitempty"`
	Email       string `bson:"e,omitempty"`
	Description string `bson:"d,omitempty"`
}

type location struct {
	CityID       primitive.ObjectID `bson:"c,omitempty"` // $lookup
	Address      string             `bson:"a,omitempty"`
	AddressTitle string             `bson:"at,omitempty"`
}

type domain struct {
	Address          string    `bson:"a,omitempty"`
	Registrar        string    `bson:"r,omitempty"`
	RegistrationDate time.Time `bson:"rd,omitempty"`
}

type social struct {
	Vk        *vkItem `bson:"v,omitempty"`
	Instagram *item   `bson:"i,omitempty"`
	Twitter   *item   `bson:"t,omitempty"`
	Youtube   *item   `bson:"y,omitempty"`
	Facebook  *item   `bson:"f,omitempty"`
}

type vkItem struct {
	URL          string `bson:"u"`
	GroupID      int    `bson:"g,omitempty"`
	Name         string `bson:"n,omitempty"`
	ScreenName   string `bson:"s,omitempty"`
	IsClosed     int8   `bson:"i,omitempty"`
	Description  string `bson:"d,omitempty"`
	MembersCount int    `bson:"m,omitempty"`
	Photo200     link   `bson:"p,omitempty"`
}

type app struct {
	AppStore   *item `bson:"a,omitempty"`
	GooglePlay *item `bson:"g,omitempty"`
}

type item struct {
	URL string `bson:"u,omitempty"`
}

func (c Company) validate() error {
	err := validation.ValidateStruct(
		&c,
		validation.Field(&c.ID, validation.Required),
		validation.Field(&c.URL, validation.Required),
		validation.Field(&c.Slug, validation.Required),
		validation.Field(&c.Online, validation.Required),
	)
	if err != nil {
		return err
	}

	return validation.ValidateStruct(
		&c.Domain,
		validation.Field(&c.Domain.Address, validation.Required),
		validation.Field(&c.Domain.Registrar, validation.Required),
		validation.Field(&c.Domain.RegistrationDate, validation.Required),
	)
}
