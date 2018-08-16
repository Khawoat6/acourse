package app

import (
	"github.com/acoshift/hime"

	"github.com/acoshift/acourse/controller/share"
	"github.com/acoshift/acourse/repository"
	"github.com/acoshift/acourse/view"
)

func index(ctx *hime.Context) error {
	if ctx.Request().URL.Path != "/" {
		return share.NotFound(ctx)
	}

	courses, err := repository.ListPublicCourses(ctx)
	if err != nil {
		return err
	}

	p := view.Page(ctx)
	p["Courses"] = courses
	return ctx.View("app.index", p)
}