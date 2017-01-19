package mcsm

// Run ...Run this program.
// Execute mackerel-agent-plugin command and formatting to be able POST api.
// And HttpPost these formatting data as ServiceMetric in Mackerel.
func Run(conf string) error {
	var config Config

	if err := LoadConfig(conf, &config); err != nil {
		return err
	}

	for _, v := range config.Rule {
		url := postURLString(v.ServiceName)
		out, err := execCommandOutput(v.Cmd, v.MetricName)
		if err != nil {
			return err
		}
		for _, d := range out {
			if err := config.HTTPPost(d, url); err != nil {
				return err
			}
		}
	}

	return nil
}
