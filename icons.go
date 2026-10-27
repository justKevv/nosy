package main

import "github.com/egoist/mygo/ui"

// lineIcon wraps stroke-drawn icon paths in the stroke style the icons share.
func lineIcon(body string) *ui.SVG {
	return ui.MustParseSVG([]byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">` + body + `</g></svg>`))
}

var (
	databaseIcon = lineIcon(`<ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5v14a9 3 0 0 0 18 0V5"/><path d="M3 12a9 3 0 0 0 18 0"/>`)
	chevronIcon  = lineIcon(`<path d="m6 9 6 6 6-6"/>`)
	activityIcon = lineIcon(`<path d="M22 12h-4l-3 9L9 3l-3 9H2"/>`)
	folderIcon   = lineIcon(`<path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/>`)
	chevronsIcon = lineIcon(`<path d="m7 15 5 5 5-5"/><path d="m7 9 5-5 5 5"/>`)
	filterIcon   = lineIcon(`<path d="M3 6h18M7 12h10M10 18h4"/>`)
	plusIcon     = lineIcon(`<path d="M12 5v14M5 12h14"/>`)
	tableIcon    = lineIcon(`<rect x="3" y="3" width="18" height="18" rx="2"/><path d="M3 9h18M3 15h18M9 3v18"/>`)
	keyIcon      = lineIcon(`<circle cx="7.5" cy="15.5" r="5.5"/><path d="m21 2-9.6 9.6M15.5 7.5l3 3L22 7l-3-3"/>`)
	linkIcon     = lineIcon(`<path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>`)
	hashIcon     = lineIcon(`<path d="M4 9h16M4 15h16M10 3 8 21M16 3l-2 18"/>`)
	eyeIcon      = lineIcon(`<path d="M2 12s3.5-7 10-7 10 7 10 7-3.5 7-10 7S2 12 2 12Z"/><circle cx="12" cy="12" r="3"/>`)
	closeIcon    = lineIcon(`<path d="M18 6 6 18M6 6l12 12"/>`)
	terminalIcon = lineIcon(`<path d="m4 17 6-6-6-6"/><path d="M12 19h8"/>`)
	diagramIcon  = lineIcon(`<circle cx="18" cy="5" r="3"/><circle cx="6" cy="12" r="3"/><circle cx="18" cy="19" r="3"/><path d="m8.59 13.51 6.83 3.98M15.41 6.51l-6.82 3.98"/>`)
	boltIcon     = lineIcon(`<path d="M4 14a1 1 0 0 1-.78-1.63l9.9-10.2a.5.5 0 0 1 .86.46l-1.92 6.02A1 1 0 0 0 13 10h7a1 1 0 0 1 .78 1.63l-9.9 10.2a.5.5 0 0 1-.86-.46l1.92-6.02A1 1 0 0 0 11 14z"/>`)
	downloadIcon = lineIcon(`<path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><path d="m7 10 5 5 5-5"/><path d="M12 15V3"/>`)
	functionIcon = lineIcon(`<path d="M9 17c2 0 2.8-1 2.8-2.8V10.8C11.8 9 12.6 8 14.6 8"/><path d="M7 11h8"/>`)
)
