package company

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Direct link .jpg
type link string

type Company struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	CategoryID      primitive.ObjectID   `bson:"c,omitempty"`
	TechnologyIDs   []primitive.ObjectID `bson:"ti,omitempty"`
	DNSIDs          []primitive.ObjectID `bson:"dn,omitempty"`
	URL             string               `bson:"u,omitempty"`
	Slug            string               `bson:"s,omitempty"`
	Title           string               `bson:"t,omitempty"`
	Email           string               `bson:"e,omitempty"`
	Description     string               `bson:"d,omitempty"`
	Phone           int                  `bson:"p,omitempty"`
	INN             int                  `bson:"i,omitempty"`
	KPP             int                  `bson:"k,omitempty"`
	OGRN            int                  `bson:"og,omitempty"`
	Domain          *domain              `bson:"do,omitempty"`
	Avatar          link                 `bson:"a,omitempty"`
	Location        *location            `bson:"l,omitempty"`
	App             *app                 `bson:"ap,omitempty"`
	Social          *social              `bson:"so,omitempty"`
	People          []*peopleItem        `bson:"pe,omitempty"`
	UpdatedAt       time.Time            `bson:"ua,omitempty"`
	PageSpeed       uint32               `bson:"ps,omitempty"` // milliseconds
	Verified        bool                 `bson:"v,omitempty"`
	Premium         bool                 `bson:"pr,omitempty"`
	PremiumDeadline time.Time            `bson:"pd,omitempty"`
	Hash            string               `bson:"has,omitempty"`
	Hidden          *bool                `bson:"h,omitempty"`
	Online          *bool                `bson:"o,omitempty"`
	HasEmail        *bool                `bson:"he,omitempty"`
	HasPhone        *bool                `bson:"hp,omitempty"`
	HasVk           *bool                `bson:"hv,omitempty"`
	HasInstagram    *bool                `bson:"hi,omitempty"`
	HasTwitter      *bool                `bson:"ht,omitempty"`
	HasYoutube      *bool                `bson:"hy,omitempty"`
	HasFacebook     *bool                `bson:"hf,omitempty"`
	HasAppStore     *bool                `bson:"ha,omitempty"`
	HasGooglePlay   *bool                `bson:"hg,omitempty"`
	HasINN          *bool                `bson:"hin,omitempty"`
	HasKPP          *bool                `bson:"hk,omitempty"`
	HasOGRN         *bool                `bson:"ho,omitempty"`
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
	CityID       primitive.ObjectID `bson:"c,omitempty"`
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

func (c *Company) GetSocial() *social {
	if c != nil {
		return c.Social
	}
	return nil
}

func (s *social) GetVk() *vkItem {
	if s != nil {
		return s.Vk
	}
	return nil
}

func (v *vkItem) GetGroupId() int {
	if v != nil {
		return v.GroupID
	}
	return 0
}

func (c *Company) validate() error {
	err := validation.ValidateStruct(
		c,
		validation.Field(&c.URL, validation.Required),
		validation.Field(&c.Slug, validation.Required),
	)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}
	return err
}

func getIDByURL(ctx context.Context, url string) (primitive.ObjectID, error) {
	var comp Company
	err := mongo.Companies.FindOne(ctx, Company{
		URL: url,
	}, options.FindOne().SetProjection(bson.M{
		"_id": 1,
	})).Decode(&comp)
	return comp.ID, err
}

func SetTechIDs(ctx context.Context, companyID primitive.ObjectID, techIDs []primitive.ObjectID) error {
	_, err := mongo.Companies.UpdateOne(ctx, Company{
		ID: companyID,
	}, bson.M{
		"$set": bson.M{
			"ti": techIDs,
		},
	})
	return err
}
