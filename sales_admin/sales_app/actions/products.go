package actions

import (

  "fmt"
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop/v5"
  "github.com/gobuffalo/x/responder"
  "github.com/henry-jackson/challenges/sales_admin/sales_app/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Product)
// DB Table: Plural (products)
// Resource: Plural (Products)
// Path: Plural (/products)
// View Template Folder: Plural (/templates/products/)

// ProductsResource is the resource for the Product model
type ProductsResource struct{
  buffalo.Resource
}

// List gets all Products. This function is mapped to the path
// GET /products
func (v ProductsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  products := &models.Products{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Products from the DB
  if err := q.All(products); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)

    c.Set("products", products)
    return c.Render(http.StatusOK, r.HTML("/products/index.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(products))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(products))
  }).Respond(c)
}

// Show gets the data for one Product. This function is mapped to
// the path GET /products/{product_id}
func (v ProductsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Product
  product := &models.Product{}

  // To find the Product the parameter product_id is used.
  if err := tx.Find(product, c.Param("product_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    c.Set("product", product)

    return c.Render(http.StatusOK, r.HTML("/products/show.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(product))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(product))
  }).Respond(c)
}

// New renders the form for creating a new Product.
// This function is mapped to the path GET /products/new
func (v ProductsResource) New(c buffalo.Context) error {
  c.Set("product", &models.Product{})

  return c.Render(http.StatusOK, r.HTML("/products/new.plush.html"))
}
// Create adds a Product to the DB. This function is mapped to the
// path POST /products
func (v ProductsResource) Create(c buffalo.Context) error {
  // Allocate an empty Product
  product := &models.Product{}

  // Bind product to the html form elements
  if err := c.Bind(product); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(product)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the new.html template that the user can
      // correct the input.
      c.Set("product", product)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/products/new.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "product.created.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/products/%v", product.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.JSON(product))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.XML(product))
  }).Respond(c)
}

// Edit renders a edit form for a Product. This function is
// mapped to the path GET /products/{product_id}/edit
func (v ProductsResource) Edit(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Product
  product := &models.Product{}

  if err := tx.Find(product, c.Param("product_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  c.Set("product", product)
  return c.Render(http.StatusOK, r.HTML("/products/edit.plush.html"))
}
// Update changes a Product in the DB. This function is mapped to
// the path PUT /products/{product_id}
func (v ProductsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Product
  product := &models.Product{}

  if err := tx.Find(product, c.Param("product_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Product to the html form elements
  if err := c.Bind(product); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(product)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the edit.html template that the user can
      // correct the input.
      c.Set("product", product)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/products/edit.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "product.updated.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/products/%v", product.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(product))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(product))
  }).Respond(c)
}

// Destroy deletes a Product from the DB. This function is mapped
// to the path DELETE /products/{product_id}
func (v ProductsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Product
  product := &models.Product{}

  // To find the Product the parameter product_id is used.
  if err := tx.Find(product, c.Param("product_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(product); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a flash message
    c.Flash().Add("success", T.Translate(c, "product.destroyed.success"))

    // Redirect to the index page
    return c.Redirect(http.StatusSeeOther, "/products")
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(product))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(product))
  }).Respond(c)
}
