package cmd_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-init/cmd"
	boshdir "github.com/cloudfoundry/bosh-init/director"
	fakedir "github.com/cloudfoundry/bosh-init/director/fakes"
	fakeui "github.com/cloudfoundry/bosh-init/ui/fakes"
	boshtbl "github.com/cloudfoundry/bosh-init/ui/table"
)

var _ = Describe("VMsCmd", func() {
	var (
		ui       *fakeui.FakeUI
		director *fakedir.FakeDirector
		command  VMsCmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		director = &fakedir.FakeDirector{}
		command = NewVMsCmd(ui, director)
	})

	Describe("Run", func() {
		var (
			opts VMsOpts
		)

		BeforeEach(func() {
			opts = VMsOpts{}
		})

		act := func() error { return command.Run(opts) }

		Context("when VMs are successfully retrieved", func() {
			var (
				infos []boshdir.VMInfo
			)

			BeforeEach(func() {
				index1 := 1
				index2 := 2

				infos = []boshdir.VMInfo{
					{
						JobName:      "job-name",
						Index:        &index1,
						State:        "in1-state",
						ResourcePool: "in1-rp",

						IPs: []string{"in1-ip1", "in1-ip2"},
						DNS: []string{"in1-dns1", "in1-dns2"},

						VMID:               "in1-cid",
						AgentID:            "in1-agent-id",
						ResurrectionPaused: false,

						Vitals: boshdir.VMInfoVitals{
							Load: []string{"0.02", "0.06", "0.11"},

							CPU:  boshdir.VMInfoVitalsCPU{Sys: "0.3", User: "1.2", Wait: "2.1"},
							Mem:  boshdir.VMInfoVitalsMemSize{Percent: "20", KB: "2000"},
							Swap: boshdir.VMInfoVitalsMemSize{Percent: "21", KB: "2100"},

							Disk: map[string]boshdir.VMInfoVitalsDiskSize{
								"system":     boshdir.VMInfoVitalsDiskSize{Percent: "35"},
								"ephemeral":  boshdir.VMInfoVitalsDiskSize{Percent: "45"},
								"persistent": boshdir.VMInfoVitalsDiskSize{Percent: "55"},
							},
						},
					},
					{
						JobName:      "job-name",
						Index:        &index2,
						State:        "in2-state",
						AZ:           "in2-az",
						ResourcePool: "in2-rp",

						IPs: []string{"in2-ip1"},
						DNS: []string{"in2-dns1"},

						VMID:               "in2-cid",
						AgentID:            "in2-agent-id",
						ResurrectionPaused: true,

						Vitals: boshdir.VMInfoVitals{
							Load: []string{"0.52", "0.56", "0.51"},

							CPU:  boshdir.VMInfoVitalsCPU{Sys: "50.3", User: "51.2", Wait: "52.1"},
							Mem:  boshdir.VMInfoVitalsMemSize{Percent: "60", KB: "6000"},
							Swap: boshdir.VMInfoVitalsMemSize{Percent: "61", KB: "6100"},

							Disk: map[string]boshdir.VMInfoVitalsDiskSize{
								"system":     boshdir.VMInfoVitalsDiskSize{Percent: "75"},
								"ephemeral":  boshdir.VMInfoVitalsDiskSize{Percent: "85"},
								"persistent": boshdir.VMInfoVitalsDiskSize{Percent: "95"},
							},
						},
					},
					{
						JobName:      "",
						Index:        nil,
						State:        "unresponsive agent",
						ResourcePool: "",
					},
				}

				deployments := []boshdir.Deployment{
					&fakedir.FakeDeployment{
						NameStub:    func() string { return "dep1" },
						VMInfosStub: func() ([]boshdir.VMInfo, error) { return infos, nil },
					},
				}

				director.DeploymentsReturns(deployments, nil)
			})

			It("lists VMs for the deployment", func() {
				Expect(act()).ToNot(HaveOccurred())

				Expect(ui.Table).To(Equal(boshtbl.Table{
					Title: "Deployment 'dep1'",

					Content: "vms",

					HeaderVals: []boshtbl.Value{
						boshtbl.NewValueString("Instance"),
						boshtbl.NewValueString("State"),
						boshtbl.NewValueString("AZ"),
						boshtbl.NewValueString("IPs"),
						boshtbl.NewValueString("VM CID"),
						boshtbl.NewValueString("VM Type"),
					},

					SortBy: []boshtbl.ColumnSort{{Column: 0, Asc: true}},

					Rows: [][]boshtbl.Value{
						{
							boshtbl.NewValueString("job-name/1"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("in1-state"), true),
							boshtbl.ValueString{},
							boshtbl.NewValueStrings([]string{"in1-ip1", "in1-ip2"}),
							boshtbl.NewValueString("in1-cid"),
							boshtbl.NewValueString("in1-rp"),
						},
						{
							boshtbl.NewValueString("job-name/2"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("in2-state"), true),
							boshtbl.NewValueString("in2-az"),
							boshtbl.NewValueStrings([]string{"in2-ip1"}),
							boshtbl.NewValueString("in2-cid"),
							boshtbl.NewValueString("in2-rp"),
						},
						{
							boshtbl.NewValueString("?/?"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("unresponsive agent"), true),
							boshtbl.ValueString{},
							boshtbl.ValueStrings{},
							boshtbl.ValueString{},
							boshtbl.ValueString{},
						},
					},

					Notes: []string{"(*) Bootstrap node"},
				}))
			})

			It("lists VMs for the deployment including details", func() {
				opts.Details = true

				Expect(act()).ToNot(HaveOccurred())

				Expect(ui.Table).To(Equal(boshtbl.Table{
					Title: "Deployment 'dep1'",

					Content: "vms",

					HeaderVals: []boshtbl.Value{
						boshtbl.NewValueString("Instance"),
						boshtbl.NewValueString("State"),
						boshtbl.NewValueString("AZ"),
						boshtbl.NewValueString("IPs"),
						boshtbl.NewValueString("VM CID"),
						boshtbl.NewValueString("VM Type"),
						boshtbl.NewValueString("Disk CID"),
						boshtbl.NewValueString("Agent ID"),
						boshtbl.NewValueString("Resurrection\nPaused"),
					},

					SortBy: []boshtbl.ColumnSort{{Column: 0, Asc: true}},

					Rows: [][]boshtbl.Value{
						{
							boshtbl.NewValueString("job-name/1"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("in1-state"), true),
							boshtbl.ValueString{},
							boshtbl.NewValueStrings([]string{"in1-ip1", "in1-ip2"}),
							boshtbl.NewValueString("in1-cid"),
							boshtbl.NewValueString("in1-rp"),
							boshtbl.ValueString{},
							boshtbl.NewValueString("in1-agent-id"),
							boshtbl.NewValueBool(false),
						},
						{
							boshtbl.NewValueString("job-name/2"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("in2-state"), true),
							boshtbl.NewValueString("in2-az"),
							boshtbl.NewValueStrings([]string{"in2-ip1"}),
							boshtbl.NewValueString("in2-cid"),
							boshtbl.NewValueString("in2-rp"),
							boshtbl.ValueString{},
							boshtbl.NewValueString("in2-agent-id"),
							boshtbl.NewValueBool(true),
						},
						{
							boshtbl.NewValueString("?/?"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("unresponsive agent"), true),
							boshtbl.ValueString{},
							boshtbl.ValueStrings{},
							boshtbl.ValueString{},
							boshtbl.ValueString{},
							boshtbl.ValueString{},
							boshtbl.ValueString{},
							boshtbl.NewValueBool(false),
						},
					},

					Notes: []string{"(*) Bootstrap node"},
				}))
			})

			It("lists VMs for the deployment including dns", func() {
				opts.DNS = true

				Expect(act()).ToNot(HaveOccurred())

				Expect(ui.Table).To(Equal(boshtbl.Table{
					Title: "Deployment 'dep1'",

					Content: "vms",

					HeaderVals: []boshtbl.Value{
						boshtbl.NewValueString("Instance"),
						boshtbl.NewValueString("State"),
						boshtbl.NewValueString("AZ"),
						boshtbl.NewValueString("IPs"),
						boshtbl.NewValueString("VM CID"),
						boshtbl.NewValueString("VM Type"),
						boshtbl.NewValueString("DNS A Records"),
					},

					SortBy: []boshtbl.ColumnSort{{Column: 0, Asc: true}},

					Rows: [][]boshtbl.Value{
						{
							boshtbl.NewValueString("job-name/1"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("in1-state"), true),
							boshtbl.ValueString{},
							boshtbl.NewValueStrings([]string{"in1-ip1", "in1-ip2"}),
							boshtbl.NewValueString("in1-cid"),
							boshtbl.NewValueString("in1-rp"),
							boshtbl.NewValueStrings([]string{"in1-dns1", "in1-dns2"}),
						},
						{
							boshtbl.NewValueString("job-name/2"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("in2-state"), true),
							boshtbl.NewValueString("in2-az"),
							boshtbl.NewValueStrings([]string{"in2-ip1"}),
							boshtbl.NewValueString("in2-cid"),
							boshtbl.NewValueString("in2-rp"),
							boshtbl.NewValueStrings([]string{"in2-dns1"}),
						},
						{
							boshtbl.NewValueString("?/?"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("unresponsive agent"), true),
							boshtbl.ValueString{},
							boshtbl.ValueStrings{},
							boshtbl.ValueString{},
							boshtbl.ValueString{},
							boshtbl.ValueStrings{},
						},
					},

					Notes: []string{"(*) Bootstrap node"},
				}))
			})

			It("lists VMs for the deployment including vitals", func() {
				opts.Vitals = true

				Expect(act()).ToNot(HaveOccurred())

				Expect(ui.Table).To(Equal(boshtbl.Table{
					Title: "Deployment 'dep1'",

					Content: "vms",

					HeaderVals: []boshtbl.Value{
						boshtbl.NewValueString("Instance"),
						boshtbl.NewValueString("State"),
						boshtbl.NewValueString("AZ"),
						boshtbl.NewValueString("IPs"),
						boshtbl.NewValueString("VM CID"),
						boshtbl.NewValueString("VM Type"),
						boshtbl.NewValueString("Uptime"),
						boshtbl.NewValueString("Load\n(1m, 5m, 15m)"),
						boshtbl.NewValueString("CPU\nTotal"),
						boshtbl.NewValueString("CPU\nUser"),
						boshtbl.NewValueString("CPU\nSys"),
						boshtbl.NewValueString("CPU\nWait"),
						boshtbl.NewValueString("Memory\nUsage"),
						boshtbl.NewValueString("Swap\nUsage"),
						boshtbl.NewValueString("System\nDisk Usage"),
						boshtbl.NewValueString("Ephemeral\nDisk Usage"),
						boshtbl.NewValueString("Persistent\nDisk Usage"),
					},

					SortBy: []boshtbl.ColumnSort{{Column: 0, Asc: true}},

					Rows: [][]boshtbl.Value{
						{
							boshtbl.NewValueString("job-name/1"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("in1-state"), true),
							boshtbl.ValueString{},
							boshtbl.NewValueStrings([]string{"in1-ip1", "in1-ip2"}),
							boshtbl.NewValueString("in1-cid"),
							boshtbl.NewValueString("in1-rp"),
							ValueUptime{},
							boshtbl.NewValueString("0.02, 0.06, 0.11"),
							ValueCPUTotal{},
							NewValueStringPercent("1.2"),
							NewValueStringPercent("0.3"),
							NewValueStringPercent("2.1"),
							ValueMemSize{boshdir.VMInfoVitalsMemSize{Percent: "20", KB: "2000"}},
							ValueMemSize{boshdir.VMInfoVitalsMemSize{Percent: "21", KB: "2100"}},
							ValueDiskSize{boshdir.VMInfoVitalsDiskSize{Percent: "35"}},
							ValueDiskSize{boshdir.VMInfoVitalsDiskSize{Percent: "45"}},
							ValueDiskSize{boshdir.VMInfoVitalsDiskSize{Percent: "55"}},
						},
						{
							boshtbl.NewValueString("job-name/2"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("in2-state"), true),
							boshtbl.NewValueString("in2-az"),
							boshtbl.NewValueStrings([]string{"in2-ip1"}),
							boshtbl.NewValueString("in2-cid"),
							boshtbl.NewValueString("in2-rp"),
							ValueUptime{},
							boshtbl.NewValueString("0.52, 0.56, 0.51"),
							ValueCPUTotal{},
							NewValueStringPercent("51.2"),
							NewValueStringPercent("50.3"),
							NewValueStringPercent("52.1"),
							ValueMemSize{boshdir.VMInfoVitalsMemSize{Percent: "60", KB: "6000"}},
							ValueMemSize{boshdir.VMInfoVitalsMemSize{Percent: "61", KB: "6100"}},
							ValueDiskSize{boshdir.VMInfoVitalsDiskSize{Percent: "75"}},
							ValueDiskSize{boshdir.VMInfoVitalsDiskSize{Percent: "85"}},
							ValueDiskSize{boshdir.VMInfoVitalsDiskSize{Percent: "95"}},
						},
						{
							boshtbl.NewValueString("?/?"),
							boshtbl.NewValueFmt(boshtbl.NewValueString("unresponsive agent"), true),
							boshtbl.ValueString{},
							boshtbl.ValueStrings{},
							boshtbl.ValueString{},
							boshtbl.ValueString{},
							ValueUptime{},
							boshtbl.ValueString{},
							ValueCPUTotal{},
							NewValueStringPercent(""),
							NewValueStringPercent(""),
							NewValueStringPercent(""),
							ValueMemSize{},
							ValueMemSize{},
							ValueDiskSize{},
							ValueDiskSize{},
							ValueDiskSize{},
						},
					},

					Notes: []string{"(*) Bootstrap node"},
				}))
			})
		})

		It("returns error if VMs cannot be retrieved", func() {
			deployments := []boshdir.Deployment{
				&fakedir.FakeDeployment{
					NameStub:    func() string { return "dep1" },
					VMInfosStub: func() ([]boshdir.VMInfo, error) { return nil, errors.New("fake-err") },
				},
			}

			director.DeploymentsReturns(deployments, nil)

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("returns error if deployments cannot be retrieved", func() {
			director.DeploymentsReturns(nil, errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
