package html
// Do Not Edit
// Generated by Blip
// source blip: ./html/index.blip.html

import (
	"context"
	"fmt"
	"github.com/samlotti/blip/blipUtil"
	"io"
	"time"
)



func IndexRender( c context.Context, w io.Writer ) (terror error) {
    start := time.Now()

	var si = blipUtil.Instance()
	var escaper = si.GetEscaperFor( "html") 
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Catch panic %s: %s\n", "IndexRender", err)
			terror = fmt.Errorf("%v", err)
		}
	    si.RenderComplete(escaper, "index", "html", time.Since(start), terror)
	}()
	// Line: 3
	si.Write(w, []byte("\n"))
	// Line: 4
	var ctxL1 = context.WithValue(c, "__Blip__", 1)
		// Line: 5
		var contentF1S1 = func() (terror error) {
			// Line: 11
			si.Write(w, []byte("<style>\n    .spacer {\n        margin: 32px;\n    }\n</style>\n\n    "))
			// End of content block
			return
		}
		ctxL1 = context.WithValue(ctxL1, "styles", contentF1S1)		// Line: 15
		var contentF1S2 = func() (terror error) {
			// Line: 39
			si.Write(w, []byte("\n<p>This is a sample project showing the main features of the Blip Template Engine.\n</p>\n\n    <button class=\"btn btn-lg btn-primary spacer\" type=\"button\" onclick=\"location.href='/users/listAll'\">List Users</button>\n    <button class=\"btn btn-lg btn-success spacer\" type=\"button\" onclick=\"location.href='/users/listActive'\">List Active</button>\n\n<br/>\n\n<div class=\"container spacer\">\n<p>Don't forget that every time the template is changed it will need to be transpiled.</p>\n    <p>One approach to make this easier is to run blip with --watch on the template directory.</p>\n    <p></p>\nThen use <b><a href=\"https://github.com/cosmtrek/air\">air</a></b> to run the application, this will restart the application whenever a template is changed.\n    </p>\n<strong>blip -dir=\"web/template\" --rebuild</strong>\n</div>\n\n<div class=\"container spacer\">\nIn this directory run: <strong>./buildTemplates.sh</strong><br>\nThen in another terminal in this directory run <strong>air</strong>\n</div>\n\n\n        "))
			// Line: 40
			si.Write(w, []byte("    "))
			// End of content block
			return
		}
		ctxL1 = context.WithValue(ctxL1, "mainContent", contentF1S2)
	terror = RootRender("Index Page", ctxL1, w)
	if terror != nil { return }
	return
}