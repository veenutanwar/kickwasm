package tap

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestExplore(t *testing.T) {
	builder := NewBuilder()
	builder.Name = "name"
	builder.Title = "title"
	okServices, err := buildShortOkServices()
	if err != nil {
		t.Fatal(err)
	}
	err = builder.BuildFromServices(okServices)
	if err != nil {
		t.Fatal(err)
	}
	// ok service names
	testGenerateServiceNames(t, builder)
	// generate the html
	builder.AddAbout()
	_ = builder.ToHTML("masterid", 2, false)
	//html := builder.ToHTML("masterid", 2, false)
	//t.Error(html)
	testGenerateServiceEmptyPanelIDsMap(t, builder)
	testGenerateServiceEmptyInsidePanelIDsMap(t, builder)
	testGenerateTabBarLevelStartPanelMap(t, builder)
	testGenerateServiceButtonPanelGroups(t, builder)
	testGenerateTabBarIDs(t, builder)
	testgenerateServicePanelNameTemplateMap(t, builder)
	testGenerateServiceTemplatePanelName(t, builder)
	// deep
	deepServices, err := buildDeepServices()
	if err != nil {
		t.Fatal(err)
	}
	builder2 := NewBuilder()
	err = builder2.BuildFromServices(deepServices)
	if err != nil {
		t.Fatal(err)
	}
	_ = builder2.ToHTML("masterid", 2, false)
	testBuilder_GenerateServiceEmptyInsidePanelNamePathMap(t, builder2)
	testBuilder_GenerateServicePanelNamePanelMap(t, builder2)
}

func testGenerateTabBarIDs(t *testing.T, builder *Builder) {
	wants := []string{
		"masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar",
	}
	results := builder.GenerateTabBarIDs()
	if ok := reflect.DeepEqual(results, wants); !ok {
		t.Fatalf(`builder.GenerateTabBarIDs() generated %#v\n\nwant: %#v`, results, wants)
	}
}

func testGenerateServiceButtonPanelGroups(t *testing.T, builder *Builder) {
	wantStr := `
About:
  - buttonname: MasteridHomePadAboutButton
    buttonid: masterid-home-pad-aboutButton
    panelnamesidmap:
      AutoGeneratedAboutPanel:
        id: autoGeneratedAboutPanel
        name: AutoGeneratedAboutPanel
        tabs:
        - id: releasesTab
          label: Releases
          panels:
          - id: releases
            name: AutoGeneratedReleasesPanel
            note: ""
            HTMLID: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-releasesTabPanel-inner-releases
            TabBarHTMLID: ""
            UnderTabBarHTMLID: ""
          note: Auto Generated About | Releases tab.
        - id: contributorsTab
          label: Contributors
          panels:
          - id: contributors
            name: AutoGeneratedContributorsPanel
            note: ""
            HTMLID: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-contributorsTabPanel-inner-contributors
            TabBarHTMLID: ""
            UnderTabBarHTMLID: ""
          note: Auto Generated About | Contributors tab.
        - id: creditsTab
          label: Credits
          panels:
          - id: credits
            name: AutoGeneratedCreditsPanel
            note: ""
            HTMLID: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-creditsTabPanel-inner-credits
            TabBarHTMLID: ""
            UnderTabBarHTMLID: ""
          note: Auto Generated About | Credits tab.
        - id: licensesTab
          label: Licenses
          panels:
          - id: licenses
            name: AutoGeneratedLicensesPanel
            note: ""
            HTMLID: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-licensesTabPanel-inner-licenses
            TabBarHTMLID: ""
            UnderTabBarHTMLID: ""
          note: Auto Generated About | Licenses tab.
        note: This panel is the Auto Generated About tab bar.
        HTMLID: masterid-home-pad-aboutButton-autoGeneratedAboutPanel
        TabBarHTMLID: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar
        UnderTabBarHTMLID: ""
  - buttonname: MasteridHomePadAboutButtonAutoGeneratedAboutPanelTabBarReleasesTab
    buttonid: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-releasesTab
    panelnamesidmap:
      AutoGeneratedReleasesPanel:
        id: releases
        name: AutoGeneratedReleasesPanel
        note: ""
        HTMLID: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-releasesTabPanel-inner-releases
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
  - buttonname: MasteridHomePadAboutButtonAutoGeneratedAboutPanelTabBarContributorsTab
    buttonid: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-contributorsTab
    panelnamesidmap:
      AutoGeneratedContributorsPanel:
        id: contributors
        name: AutoGeneratedContributorsPanel
        note: ""
        HTMLID: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-contributorsTabPanel-inner-contributors
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
  - buttonname: MasteridHomePadAboutButtonAutoGeneratedAboutPanelTabBarCreditsTab
    buttonid: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-creditsTab
    panelnamesidmap:
      AutoGeneratedCreditsPanel:
        id: credits
        name: AutoGeneratedCreditsPanel
        note: ""
        HTMLID: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-creditsTabPanel-inner-credits
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
  - buttonname: MasteridHomePadAboutButtonAutoGeneratedAboutPanelTabBarLicensesTab
    buttonid: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-licensesTab
    panelnamesidmap:
      AutoGeneratedLicensesPanel:
        id: licenses
        name: AutoGeneratedLicensesPanel
        note: ""
        HTMLID: masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-licensesTabPanel-inner-licenses
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
Service1:
  - buttonname: MasteridHomePadOneButton
    buttonid: masterid-home-pad-OneButton
    panelnamesidmap:
      OnePanel:
        id: OnePanel
        name: OnePanel
        note: p1 note
        markup: <p>Panel 1-1</p>
        HTMLID: masterid-home-pad-OneButton-OnePanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
      TwoPanel:
        id: TwoPanel
        name: TwoPanel
        note: p2 note
        markup: <p>Panel 2-1</p>
        HTMLID: masterid-home-pad-OneButton-TwoPanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
Service2:
  - buttonname: MasteridHomePadTwoButton
    buttonid: masterid-home-pad-TwoButton
    panelnamesidmap:
      FourPanel:
        id: fourPanel
        name: FourPanel
        note: p4 note
        markup: <p>Panel 4-1</p>
        HTMLID: masterid-home-pad-TwoButton-fourPanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
      ThreePanel:
        id: threePanel
        name: ThreePanel
        note: p3 note
        markup: <p>Panel 3-1</p>
        HTMLID: masterid-home-pad-TwoButton-threePanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""`
	var wants map[string][]*ButtonPanelGroup
	if err := yaml.Unmarshal([]byte(wantStr), &wants); err != nil {
		t.Fatal(err)
	}
	results := builder.GenerateServiceButtonPanelGroups()
	//bb, _ := yaml.Marshal(results)
	//t.Fatal(string(bb))
	for service, wbpg := range wants {
		rbpg, ok := results[service]
		if !ok {
			t.Errorf(`GenerateServiceButtonPanelGroups service %s is missing`, service)
		} else {
			for i, wbp := range wbpg {
				rbp := rbpg[i]
				if rbp.ButtonName != wbp.ButtonName {
					t.Errorf(`rbpg[%d].ButtonName != wbp.ButtonName: got %s want %s`, i, rbp.ButtonName, wbp.ButtonName)
				}
				if rbpg[i].ButtonID != wbp.ButtonID {
					t.Errorf(`rbpg[%d].ButtonID != wbp.ButtonID: got %s want %s`, i, rbp.ButtonID, wbp.ButtonID)
				}
				for wpname := range wbp.PanelNamesIDMap {
					if _, ok := rbp.PanelNamesIDMap[wpname]; !ok {
						t.Errorf(`%s missing in rbpg[%d].PanelNamesIDMap`, wpname, i)
					}
				}
			}
		}
	}
}

func testGenerateTabBarLevelStartPanelMap(t *testing.T, builder *Builder) {
	wants := map[string]string{
		"masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar": "masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-releasesTabPanel",
	}
	results := builder.GenerateTabBarLevelStartPanelMap()
	if ok := reflect.DeepEqual(results, wants); !ok {
		t.Fatalf(`builder.GenerateTabBarLevelStartPanelMap() generated %#v\n\nwant: %#v`, results, wants)
	}
}

func testGenerateServiceEmptyInsidePanelIDsMap(t *testing.T, builder *Builder) {
	wants := map[string]map[string]string{
		"Service1": map[string]string{
			"OnePanel": "masterid-home-pad-OneButton-OnePanel-inner-user-content",
			"TwoPanel": "masterid-home-pad-OneButton-TwoPanel-inner-user-content",
		},
		"Service2": map[string]string{
			"FourPanel":  "masterid-home-pad-TwoButton-FourPanel-inner-user-content",
			"ThreePanel": "masterid-home-pad-TwoButton-ThreePanel-inner-user-content",
		},
		"About": map[string]string{
			"AutoGeneratedLicensesPanel":     "masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-licensesTabPanel-inner-licenses-inner-user-content",
			"AutoGeneratedReleasesPanel":     "masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-releasesTabPanel-inner-releases-inner-user-content",
			"AutoGeneratedContributorsPanel": "masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-contributorsTabPanel-inner-contributors-inner-user-content",
			"AutoGeneratedCreditsPanel":      "masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-creditsTabPanel-inner-credits-inner-user-content",
		},
	}
	results := builder.GenerateServiceEmptyInsidePanelIDsMap()
	if ok := reflect.DeepEqual(results, wants); !ok {
		t.Fatalf(`builder.GenerateServiceEmptyInsidePanelIDsMap() generated %#v\n\nwant: %#v`, results, wants)
	}
}

func testGenerateServiceEmptyPanelIDsMap(t *testing.T, builder *Builder) {
	wants := map[string][]string{
		"Service1": []string{
			"masterid-home-pad-OneButton-OnePanel",
			"masterid-home-pad-OneButton-TwoPanel",
		},
		"Service2": []string{
			"masterid-home-pad-TwoButton-ThreePanel",
			"masterid-home-pad-TwoButton-FourPanel",
		},
		"About": []string{
			"masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-releasesTabPanel-inner-releases",
			"masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-contributorsTabPanel-inner-contributors",
			"masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-creditsTabPanel-inner-credits",
			"masterid_home_pad_aboutButton_autoGeneratedAboutPanel_tab_bar-licensesTabPanel-inner-licenses",
		},
	}

	results := builder.GenerateServiceEmptyPanelIDsMap()
	if ok := reflect.DeepEqual(results, wants); !ok {
		t.Fatalf("builder.GenerateServiceEmptyPanelIDsMap() generated %#v\n\nwant: %#v", results, wants)
	}
}

func testGenerateServiceNames(t *testing.T, builder *Builder) {
	correctNamesAnswer := []string{"Service1", "Service2"}
	serviceNames := builder.GenerateServiceNames()
	if len(serviceNames) != len(correctNamesAnswer) {
		t.Fatalf(`builder.GenerateServiceNames() len is %d not %d`, len(serviceNames), len(correctNamesAnswer))
	}
	for i, name := range serviceNames {
		if correctNamesAnswer[i] != name {
			t.Errorf(`builder.GenerateServiceNames() [%d] != %q its %q`, i, correctNamesAnswer[i], name)
		}
	}
	serviceNames = builder.GenerateOrganicServiceNames()
	if len(serviceNames) != len(correctNamesAnswer) {
		t.Fatalf(`builder.GenerateOrganicServiceNames() len is %d not %d`, len(serviceNames), len(correctNamesAnswer))
	}
	for i, name := range serviceNames {
		if correctNamesAnswer[i] != name {
			t.Errorf(`builder.GenerateOrganicServiceNames() [%d] != %q its %q`, i, correctNamesAnswer[i], name)
		}
	}
}

func testgenerateServicePanelNameTemplateMap(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string]map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "a",
			want: map[string]map[string]string{
				"Service1": map[string]string{
					"OnePanel": "\n<!--\n\nPanel name: \"OnePanel\"\n\nPanel note: p1 note\n\nThis is one of a group of 2 panels displayed when the \"One\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-OneButton-TwoPanel-inner-user-content\n  * Name: TwoPanel\n  * Note: p2 note\n\n-->\n\n<p>Panel 1-1</p>\n",
					"TwoPanel": "\n<!--\n\nPanel name: \"TwoPanel\"\n\nPanel note: p2 note\n\nThis is one of a group of 2 panels displayed when the \"One\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-OneButton-OnePanel-inner-user-content\n  * Name: OnePanel\n  * Note: p1 note\n\n-->\n\n<p>Panel 2-1</p>\n",
				},
				"Service2": map[string]string{
					"ThreePanel": "\n<!--\n\nPanel name: \"ThreePanel\"\n\nPanel note: p3 note\n\nThis is one of a group of 2 panels displayed when the \"Two\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-TwoButton-FourPanel-inner-user-content\n  * Name: FourPanel\n  * Note: p4 note\n\n-->\n\n<p>Panel 3-1</p>\n",
					"FourPanel":  "\n<!--\n\nPanel name: \"FourPanel\"\n\nPanel note: p4 note\n\nThis is one of a group of 2 panels displayed when the \"Two\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-TwoButton-ThreePanel-inner-user-content\n  * Name: ThreePanel\n  * Note: p3 note\n\n-->\n\n<p>Panel 4-1</p>\n",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.GenerateServicePanelNameTemplateMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.GenerateServicePanelNameTemplateMap() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func testGenerateServiceTemplatePanelName(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string][]string
	}{
		// TODO: Add test cases.
		{
			name: "wtf",
			want: map[string][]string{
				"Service2": []string{
					"ThreePanel",
					"FourPanel",
				},
				"Service1": []string{
					"OnePanel",
					"TwoPanel",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.GenerateServiceTemplatePanelName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.GenerateServiceTemplatePanelName() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func testBuilder_GenerateServiceEmptyInsidePanelNamePathMap(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string]map[string][]string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			want: map[string]map[string][]string{
				"Service1": map[string][]string{
					"OneOnePanel": []string{"OneButton", "OnePanel", "OneOneButton"},
					"TwoOnePanel": []string{"OneButton", "TwoPanel", "TwoOneButton"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.GenerateServiceEmptyInsidePanelNamePathMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.GenerateServiceEmptyInsidePanelNamePathMap() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func testBuilder_GenerateServicePanelNamePanelMap(t *testing.T, builder *Builder) {
	wantStr := `
Service1:
  OneOnePanel:
    id: OneOnePanel
    name: OneOnePanel
    note: ""
    markup: <p>One One Panel</p>
    HTMLID: masterid-home-pad-OneButton-OnePanel-OneOneButton-OneOnePanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
  OnePanel:
    id: OnePanel
    name: OnePanel
    buttons:
    - id: OneOneButton
      label: One One
      heading: One One
      cc: One One
      panels:
      - id: OneOnePanel
        name: OneOnePanel
        note: ""
        markup: <p>One One Panel</p>
        HTMLID: masterid-home-pad-OneButton-OnePanel-OneOneButton-OneOnePanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
    note: p1 note
    HTMLID: masterid-home-pad-OneButton-OnePanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
  TwoOnePanel:
    id: TwoOnePanel
    name: TwoOnePanel
    note: ""
    markup: <p>Two One Panel</p>
    HTMLID: masterid-home-pad-OneButton-TwoPanel-TwoOneButton-TwoOnePanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
  TwoPanel:
    id: TwoPanel
    name: TwoPanel
    buttons:
      - id: TwoOneButton
        label: Two One
        heading: Two One
        cc: Two One
        panels:
        - id: TwoOnePanel
          name: TwoOnePanel
          note: ""
          markup: <p>Two One Panel</p>
          HTMLID: masterid-home-pad-OneButton-TwoPanel-TwoOneButton-TwoOnePanel
          TabBarHTMLID: ""
          UnderTabBarHTMLID: ""
    note: p2 note
    HTMLID: masterid-home-pad-OneButton-TwoPanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""`
	var want map[string]map[string]*Panel
	if err := yaml.Unmarshal([]byte(wantStr), &want); err != nil {
		t.Fatal(err)
	}
	got := builder.GenerateServicePanelNamePanelMap()
	for wService, wPanelNamePanelMap := range want {
		gPanelNamePanelMap, ok := got[wService]
		if !ok {
			t.Errorf("service %q not found in result.", wService)
			return
		}
		for wName, wPanel := range wPanelNamePanelMap {
			gPanel, ok := gPanelNamePanelMap[wName]
			if !ok {
				t.Errorf("panel %q not found in result.", wName)
				return
			}
			if wPanel.ID != gPanel.ID {
				t.Errorf("panel.ID %q not found in result.", wPanel.ID)
				return
			}
		}

	}
}
