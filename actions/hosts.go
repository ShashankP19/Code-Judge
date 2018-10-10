package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/shashankp/cpjudge/models"
)

// HostRegisterGet displays a register form
func HostsRegisterGet(c buffalo.Context) error {
	// Make host available inside the html template
	c.Set("host", &models.Host{})
	return c.Render(200, r.HTML("hosts/register.html"))
}

func HostHomePage(c buffalo.Context) error {
	c.Set("host", &models.Host{})
	return c.Render(200, r.HTML("hosts/home.html"))
}

// HostsRegisterPost adds a host to the DB. This function is mapped to the
// path POST /accounts/register
func HostsRegisterPost(c buffalo.Context) error {
	// Allocate an empty Host
	host := &models.Host{}
	// Bind host to the html form elements
	if err := c.Bind(host); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	verrs, err := host.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make host available inside the html template
		c.Set("host", host)
		// Make the errors available inside the html template
		c.Set("errors", verrs.Errors)
		// Render again the register.html template that the host can
		// correct the input.
		return c.Render(422, r.HTML("hosts/register.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Account created successfully.")
	// and redirect to the home page
	return c.Redirect(302, "/")
}
