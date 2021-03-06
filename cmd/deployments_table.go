package cmd

import (
	boshdir "github.com/cloudfoundry/bosh-init/director"
	boshui "github.com/cloudfoundry/bosh-init/ui"
	boshtbl "github.com/cloudfoundry/bosh-init/ui/table"
)

type DeploymentsTable struct {
	Deployments []boshdir.Deployment
	UI          boshui.UI
}

func (t DeploymentsTable) Print() error {
	table := boshtbl.Table{
		Content: "deployments",

		Header: []string{"Name", "Release(s)", "Stemcell(s)", "Cloud Config"},

		SortBy: []boshtbl.ColumnSort{
			{Column: 0, Asc: true},
		},
	}

	for _, d := range t.Deployments {
		releases, err := d.Releases()
		if err != nil {
			return err
		}

		stemcells, err := d.Stemcells()
		if err != nil {
			return err
		}

		config, err := d.CloudConfig()
		if err != nil {
			return err
		}

		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueString(d.Name()),
			boshtbl.NewValueStrings(t.takeReleases(releases)),
			boshtbl.NewValueStrings(t.takeStemcells(stemcells)),
			boshtbl.NewValueString(config),
		})
	}

	t.UI.PrintTable(table)

	return nil
}

func (t DeploymentsTable) takeReleases(rels []boshdir.Release) []string {
	var names []string
	for _, r := range rels {
		names = append(names, r.Name()+"/"+r.Version().String())
	}
	return names
}

func (t DeploymentsTable) takeStemcells(stemcells []boshdir.Stemcell) []string {
	var names []string
	for _, s := range stemcells {
		names = append(names, s.Name()+"/"+s.Version().String())
	}
	return names
}
