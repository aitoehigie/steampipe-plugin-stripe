package stripe

import (
	"context"

	"github.com/stripe/stripe-go/v76"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableStripePrices(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "stripe_prices",
		Description: "Prices in Stripe represent the unit cost for a product, specifying the amount, currency, and billing frequency.",
		List: &plugin.ListConfig{
			Hydrate: listPrices,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the price."},
			{Name: "active", Type: proto.ColumnType_BOOL, Description: "Whether the price is currently active or not."},
			{Name: "currency", Type: proto.ColumnType_STRING, Description: "The currency the price is in."},
			{Name: "product", Type: proto.ColumnType_STRING, Transform: transform.FromField("Product.ID"), Description: "The ID of the product this price is associated with."},
			{Name: "unit_amount", Type: proto.ColumnType_INT, Description: "The unit amount in cents (or other currency equivalent) for the price."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the price, one of `recurring`, `one_time`, or `usage_based`."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "The time the price was created."},
			{Name: "attrs", Type: proto.ColumnType_JSON, Description: "Additional attributes of the price."},
		},
	}
}

// listPrices lists all prices
func listPrices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("stripe_prices.listPrices", "connection_error", err)
		return nil, err
	}

	params := &stripe.PriceListParams{}
	i := conn.Prices.List(params)
	for i.Next() {
		d.StreamListItem(ctx, i.Price())
	}

	return nil, nil
}