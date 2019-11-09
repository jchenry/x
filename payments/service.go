package payments

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/jchenry/jchenry/auth"
	_http "github.com/jchenry/jchenry/http"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
	"github.com/stripe/stripe-go/sub"
)

func Service(c Config, auth *auth.ServiceInstance) ServiceInstance {
	stripe.Key = c.StripeKey
	sc := &client.API{}
	sc.Init(c.StripeKey, nil)
	return ServiceInstance{
		c:      c,
		stripe: sc,
		auth:   auth,
	}
}

type ServiceInstance struct {
	c      Config
	stripe *client.API
	auth   *auth.ServiceInstance
}

func (si ServiceInstance) Register(uriBase string, s *_http.Server) {
	s.Get(uriBase+"/subscription", "subscription info endpoint", negroni.New(
		negroni.HandlerFunc(auth.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(si.subscriptionHandler)),
	)).Post(uriBase+"/subscription", "subscription payment endpoint", negroni.New(
		negroni.HandlerFunc(auth.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(si.paymentHandler)),
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
	_http.RenderTemplate(w, "subscription", offering{Product: *prod, Plans: plans})
}

type offering struct {
	Product stripe.Product
	Plans   []stripe.Plan
}

func (si ServiceInstance) paymentHandler(w http.ResponseWriter, r *http.Request) {

	session, err := auth.Store.Get(r, auth.SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if u, ok := session.Values["profile"]; ok {
		user := u.(auth.User)
		r.ParseForm()

		params := &stripe.CustomerParams{
			Email: stripe.String(user.Email),
			Name:  stripe.String(fmt.Sprintf("%s, %s", user.LastName, user.FirstName)),
		}
		params.SetSource(r.PostFormValue("stripeToken"))
		cus, err := customer.New(params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p := &stripe.SubscriptionParams{
			Customer: stripe.String(cus.ID),
			Items: []*stripe.SubscriptionItemsParams{
				{
					Plan: stripe.String(r.PostFormValue("plan")),
				},
			},
		}
		s, err := sub.New(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if si.c.TenantSetup == nil {
			panic("need code to setup the tenant")
		}

		if user.Apps == nil {
			user.Apps = map[string]interface{}{}
		}
		user.Apps[si.c.StripeProductID] = si.c.TenantSetup(s.ID, user.ID)
		err = si.auth.UpdateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, si.c.RedirectURL, http.StatusSeeOther)

	}
}
