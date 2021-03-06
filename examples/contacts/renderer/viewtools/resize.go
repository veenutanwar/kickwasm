package viewtools

import (
	"fmt"
	"syscall/js"
)

func (tools *Tools) initializeResize() {
	cb := tools.notjs.RegisterCallBack(func([]js.Value) { tools.SizeApp() })
	tools.Global.Set("onresize", cb)
}

// SizeApp resizes the app
func (tools *Tools) SizeApp() {
	// begin with the height of the inside of the browser where the window is.
	notjs := tools.notjs
	windowWidth := notjs.WindowInnerWidth()
	windowHeight := notjs.WindowInnerHeight()
	// and subtract body measurments
	bodies := notjs.GetElementsByTagName("body")
	body := bodies[0]
	xh := notjs.HeightExtras(body)
	xw := notjs.WidthExtras(body)
	windowHeight -= xh
	windowWidth -= xw
	// size each master view
	tools.sizeTabsMasterView(windowWidth, windowHeight)
	tools.sizeModalMasterView(windowWidth, windowHeight)
	tools.sizeCloserMasterView(windowWidth, windowHeight)
}

func (tools *Tools) sizeTabsMasterView(w, h float64) {
	// now set the masterview height
	notjs := tools.notjs
	if tools.ElementIsShown(tools.tabsMasterview) {
		// tabs masterview is visible
		// subtract extras before setting
		h -= notjs.HeightExtras(tools.tabsMasterview)
		w -= notjs.WidthExtras(tools.tabsMasterview)
		// set master view height, width
		notjs.SetStyleHeight(tools.tabsMasterview, h)
		notjs.SetStyleWidth(tools.tabsMasterview, w)
		// div #tabsMasterview children
		// * H1
		// * div #tabsMasterview-home
		// * div #tabsMasterview-home-slider
		//
		// Process h1
		h1Ht := float64(0)
		children := notjs.ChildrenSlice(tools.tabsMasterview)
		for _, ch := range children {
			if notjs.TagName(ch) == "H1" {
				chwx := notjs.WidthExtras(ch)
				notjs.SetStyleWidth(ch, w-chwx)
				h1Ht = notjs.OuterHeight(ch)
				break
			}
		}
		// h - h1Ht is now the ht for home or slider.
		// home panel or slider is under the h1.
		h -= h1Ht
		// Process Home
		if tools.ElementIsShown(tools.tabsMasterviewHome) {
			// remove extra measurements
			h -= notjs.HeightExtras(tools.tabsMasterviewHome)
			w -= notjs.WidthExtras(tools.tabsMasterviewHome)
			// set the inside height and width
			notjs.SetStyleHeight(tools.tabsMasterviewHome, h)
			notjs.SetStyleWidth(tools.tabsMasterviewHome, w)
			// homepad is the button pad in home.
			h -= notjs.HeightExtras(tools.tabsMasterviewHomeButtonPad)
			w -= notjs.WidthExtras(tools.tabsMasterviewHomeButtonPad)
			//h -= 100
			notjs.SetStyleHeight(tools.tabsMasterviewHomeButtonPad, h)
			notjs.SetStyleWidth(tools.tabsMasterviewHomeButtonPad, w)
			return
		}
		// home is not visible check the slider
		// Process Slider
		if tools.ElementIsShown(tools.tabsMasterviewHomeSlider) {
			// remove extra measurements
			h -= notjs.HeightExtras(tools.tabsMasterviewHomeSlider)
			w -= notjs.WidthExtras(tools.tabsMasterviewHomeSlider)
			// set the inside height and width
			notjs.SetStyleHeight(tools.tabsMasterviewHomeSlider, h)
			notjs.SetStyleWidth(tools.tabsMasterviewHomeSlider, w)
			// slider has a back button
			backOuterWidth := notjs.OuterWidth(tools.tabsMasterviewHomeSliderBack)
			w -= backOuterWidth
			// size slider collection
			h -= notjs.HeightExtras(tools.tabsMasterviewHomeSliderCollection)
			w -= notjs.WidthExtras(tools.tabsMasterviewHomeSliderCollection)
			notjs.SetStyleHeight(tools.tabsMasterviewHomeSliderCollection, h)
			notjs.SetStyleWidth(tools.tabsMasterviewHomeSliderCollection, w)
			// slider collection children
			children := notjs.ChildrenSlice(tools.tabsMasterviewHomeSliderCollection)
			for _, ch := range children {
				if notjs.TagName(ch) == "DIV" && notjs.ClassListContainsAnd(ch, SliderPanelClassName, SeenClassName) {
					tools.sizeSliderPanel(ch, w, h)
					break
				}
			}
		}
	}
}

func (tools *Tools) reSizeSliderBack(height, margintop float64) {
	style := tools.tabsMasterviewHomeSliderBack.Get("style")
	style.Set("height", fmt.Sprintf("%fpx", height))
	style.Set("margin-top", fmt.Sprintf("%fpx", margintop))
}

func (tools *Tools) sizeModalMasterView(w, h float64) {
	// modal master view
	if tools.ElementIsShown(tools.modalMasterView) {
		notjs := tools.notjs
		// modal view is visible
		w -= notjs.WidthExtras(tools.modalMasterView)
		h -= notjs.HeightExtras(tools.modalMasterView)
		notjs.SetStyleWidth(tools.modalMasterView, w)
		notjs.SetStyleHeight(tools.modalMasterView, h)
		// the center div
		w -= notjs.WidthExtras(tools.modalMasterViewCenter)
		h -= notjs.HeightExtras(tools.modalMasterViewCenter)
		notjs.SetStyleWidth(tools.modalMasterViewCenter, w)
		notjs.SetStyleHeight(tools.modalMasterViewCenter, h)
		// subtract ht of h1 and p > button
		children := notjs.ChildrenSlice(tools.modalMasterViewCenter)
		for _, ch := range children {
			tagName := notjs.TagName(ch)
			if tagName == "H1" || tagName == "P" {
				chwx := notjs.WidthExtras(ch)
				notjs.SetStyleWidth(ch, w-chwx)
				oh := notjs.OuterHeight(ch)
				h -= oh
			}
		}
		// message
		w -= notjs.WidthExtras(tools.modalMasterViewMessage)
		h -= notjs.HeightExtras(tools.modalMasterViewMessage)
		notjs.SetStyleWidth(tools.modalMasterViewMessage, w)
		notjs.SetStyleHeight(tools.modalMasterViewMessage, h)
	}
}

func (tools *Tools) sizeCloserMasterView(w, h float64) {
	if tools.ElementIsShown(tools.closerMasterView) {
		tools.notjs.SetStyleHeight(tools.closerMasterView, h)
	}
}

func (tools *Tools) sizeSliderPanel(sliderPanel js.Value, w, h float64) {
	// #tabsMasterView-home-slider-collection is the parant of a slider panel
	// finds and sets the ht of div.slider-panel-inner
	// remove extras.
	notjs := tools.notjs
	w -= notjs.WidthExtras(sliderPanel)
	h -= notjs.HeightExtras(sliderPanel)
	inner := js.Undefined()
	marginHt := float64(0)
	headingHt := float64(0)
	// get height of headings. the ccs are also headings
	// first the optional ccs
	// then the actual heading
	children := notjs.ChildrenSlice(sliderPanel)
	for _, ch := range children {
		if notjs.ClassListContains(ch, PanelHeadingClassName) {
			chwx := notjs.WidthExtras(ch)
			notjs.SetStyleWidth(ch, w-chwx)
			marginHt = headingHt
			headingHt += notjs.OuterHeight(ch)
		}
		if notjs.ClassListContains(ch, SliderPanelInnerClassName) {
			inner = ch
		}
	}
	// size the back button
	tools.reSizeSliderBack(h-marginHt, marginHt)
	// size this slider panel
	h -= headingHt
	notjs.SetStyleWidth(sliderPanel, w)
	notjs.SetStyleHeight(sliderPanel, h)
	// size slider panel's inner panel
	// inside inner panel
	w -= notjs.WidthExtras(inner)
	h -= notjs.HeightExtras(inner)
	notjs.SetStyleWidth(inner, w)
	notjs.SetStyleHeight(inner, h)
	// inside the inner panel will be:
	// * button pad
	// * or user content
	// * or tab bar.
	buttonPad := js.Undefined()
	userContent := js.Undefined()
	tabbar := js.Undefined()
	underTabbar := js.Undefined()
	children = notjs.ChildrenSlice(inner)
	for _, ch := range children {
		if notjs.ClassListContains(ch, SliderButtonPadClassName) {
			buttonPad = ch
			break
		}
		if notjs.ClassListContains(ch, UserContentClassName) {
			userContent = ch
			break
		}
		if notjs.ClassListContains(ch, TabBarClassName) {
			tabbar = ch
			// continue to get the UnderTabBarClassName
		}
		if notjs.ClassListContains(ch, UnderTabBarClassName) {
			underTabbar = ch
			break
		}
	}
	if buttonPad != js.Undefined() {
		// a button pad is inside the inner panel
		w -= notjs.WidthExtras(buttonPad)
		h -= notjs.HeightExtras(buttonPad)
		notjs.SetStyleHeight(buttonPad, h)
		notjs.SetStyleWidth(buttonPad, w)
		return
	}
	if userContent != js.Undefined() {
		// a user content is inside the inner panel
		w -= notjs.WidthExtras(userContent)
		h -= notjs.HeightExtras(userContent)
		notjs.SetStyleHeight(userContent, h)
		notjs.SetStyleWidth(userContent, w)
		return
	}
	if tabbar != js.Undefined() && underTabbar != js.Undefined() {
		// a tab bar is inside the inner panel
		seen := js.Undefined()
		h -= notjs.OuterHeight(tabbar)
		h -= notjs.HeightExtras(underTabbar)
		notjs.SetStyleHeight(underTabbar, h)
		// find the visible panel under the tab bar
		children := notjs.ChildrenSlice(underTabbar)
		for _, ch := range children {
			if notjs.ClassListContains(ch, SeenClassName) {
				seen = ch
				break
			}
		}
		if seen == js.Undefined() {
			// this will only happend in development and testing of kickwasm.
			message := fmt.Sprintf("missing seen div under %s", underTabbar.Get("id"))
			notjs.Alert(message)
			return
		}
		// size the visible panel under the tab bar
		seenwx := notjs.WidthExtras(seen)
		w -= seenwx
		notjs.SetStyleWidth(seen, w)
		notjs.SetStyleHeight(seen, h)

		// the visible panel under the tab bar has a heading over its inner panel
		// the inner panel's height is height of the under tab bar - the heading height.
		children = notjs.ChildrenSlice(seen)
		for _, ch := range children {
			if notjs.ClassListContains(ch, PanelHeadingClassName) {
				chwx := notjs.WidthExtras(ch)
				notjs.SetStyleWidth(ch, w-chwx)
				h -= notjs.OuterHeight(ch)
			} else if notjs.ClassListContains(ch, InnerPanelClassName) {
				notjs.SetStyleHeight(ch, h)
				break
			}
		}
		return
	}
}

func (tools *Tools) sizeTabBarPanel(tabBarPanel js.Value, h float64) {
	// finds and sets the ht of div.under-tab-bar
	// find and subtract the div.tab-bar height
	notjs := tools.notjs
	children := notjs.ChildrenSlice(tabBarPanel)
	for _, ch := range children {
		if notjs.ClassListContains(ch, TabBarClassName) {
			h -= notjs.OuterHeight(ch)
			break
		}
	}
	// subtract top and bottom borders of the div.under-tab-bar
	// find and size the under tab bar panel
	for _, ch := range children {
		if notjs.ClassListContains(ch, UnderTabBarClassName) {
			// style.height is the inner height so remove the extra measurments.
			h -= notjs.HeightExtras(ch)
			notjs.SetStyleHeight(ch, h)
			// size the childs sub panels
			tools.sizeSubPanels(ch, h)
			break
		}
	}
}

func (tools *Tools) sizePanelWithHeading(panelWithHeading js.Value, h float64) {
	// finds and sets the ht of div.inner-panel
	// ht = parent ht - ht of h?.heading-of-panel - size of scroll arrow.
	// find and subtract the size of the h?.heading-of-panel
	notjs := tools.notjs
	children := notjs.ChildrenSlice(panelWithHeading)
	for _, ch := range children {
		if notjs.ClassListContains(ch, PanelHeadingClassName) {
			h -= notjs.OuterHeight(ch)
			break
		}
	}
	// find and size the inner
	for _, ch := range children {
		if notjs.ClassListContains(ch, InnerPanelClassName) {
			if tools.hasSubPanels(ch) {
				h -= notjs.HeightExtras(ch)
				notjs.SetStyleHeight(ch, h)
				tools.sizeSubPanels(ch, h)
			} else {
				// no sub panels so possible scroll arrow at lower right.
				style := ch.Get("style")
				style.Set("padding-bottom", "10px")
				style.Set("height", fmt.Sprintf("%fpx", (h-float64(30))))
			}
			break
		}
	}
}

func (tools *Tools) sizeSubPanels(panel js.Value, h float64) {
	notjs := tools.notjs
	children := notjs.ChildrenSlice(panel)
	for _, ch := range children {
		tagName := notjs.TagName(ch)
		if tagName == "DIV" && !notjs.ClassListContains(ch, UnSeenClassName) {
			if notjs.ClassListContains(ch, PanelWithTabBarClassName) {
				h -= notjs.HeightExtras(ch)
				tools.sizeTabBarPanel(ch, h)
				return
			}
			if notjs.ClassListContains(ch, PanelWithHeadingClassName) {
				h -= notjs.HeightExtras(ch)
				tools.sizePanelWithHeading(ch, h)
				return
			}
		}
	}
}

func (tools *Tools) hasSubPanels(panel js.Value) bool {
	notjs := tools.notjs
	children := notjs.ChildrenSlice(panel)
	for _, ch := range children {
		if notjs.TagName(ch) == "DIV" && notjs.ClassListContainsOr(ch, PanelWithTabBarClassName, PanelWithHeadingClassName) {
			return true
		}
	}
	return false
}
