package miro

import (
	"context"

	"github.com/Miro-Ecosystem/go-miro/miro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBoard() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the Board",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the Board",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
		CreateContext: resourceBoardCreate,
		ReadContext:   resourceBoardRead,
		UpdateContext: resourceBoardUpdate,
		DeleteContext: resourceBoardDelete,
	}
}

func resourceBoardCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*miro.Client)
	name := data.Get("name").(string)
	desc := data.Get("description").(string)

	req := &miro.CreateBoardRequest{
		Name:        name,
		Description: desc,
	}

	board, err := c.Boards.Create(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(board.ID)

	return resourceBoardRead(ctx, data, meta)
}

func resourceBoardRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*miro.Client)

	var diags diag.Diagnostics
	board, err := c.Boards.Get(ctx, data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if board == nil {
		data.SetId("")
		return diags
	}

	if err := data.Set("boards", board); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(board.ID)
	return diags
}

func resourceBoardUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO
	return nil
}

func resourceBoardDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO
	return nil
}
