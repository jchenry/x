package payments

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/jchenry/jchenry/auth"
	jch_http "github.com/jchenry/jchenry/http"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
)

func Service(c Config) ServiceInstance {
	stripe.Key = c.StripeKey
	sc := &client.API{}
	sc.Init(c.StripeKey, nil)
	return ServiceInstance{
		c:      c,
		stripe: sc,
	}
}

type ServiceInstance struct {
	c      Config
	stripe *client.API
}

func (si ServiceInstance) Register(uriBase string, s *jch_http.Server) {
	s.GET(uriBase+"/subscription", "subscription info endpoint", negroni.New(
		negroni.HandlerFunc(auth.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(si.subscriptionHandler)),
	))
}

func (si ServiceInstance) subscriptionHandler(w http.ResponseWriter, r *http.Request) {

	prod, _ := product.Get(si.c.StripeProductID, nil)

	params := &stripe.PlanListParams{
		Product: &si.c.StripeProductID,
	}

	it := plan.List(params)
	var plans []stripe.Plan
	for it.Next() {
		plans = append(plans, *it.Plan())
	}
	jch_http.RenderTemplate(w, "subscription", offering{Product: *prod, Plans: plans})
}

type offering struct {
	Product stripe.Product
	Plans   []stripe.Plan
}
