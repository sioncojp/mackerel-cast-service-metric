package mcsm

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mattn/go-shellwords"
	"github.com/pkg/errors"
)

// execCommandOutput ...Execute mackerel command.
func execCommandOutput(cmd, metricname string) ([]Data, error) {
	var out []byte
	c, err := shellwords.Parse(cmd)
	if err != nil {
		return nil, errors.Wrap(err, "parse command")
	}

	switch len(c) {
	case 0:
		return nil, nil
	case 1:
		out, err = exec.Command(c[0]).Output()
	default:
		out, err = exec.Command(c[0], c[1:]...).Output()
	}
	if err != nil {
		return nil, errors.Wrap(err, "exec command")
	}

	return formatString(string(out), metricname), nil
}

// formatString ...Format string from execCommandOutput.
// These string gets into Data type struct array.
func formatString(s, metricname string) []Data {
	var data []Data

	lines := strings.Split(s, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		word := strings.Fields(line)

		name := nameString(word[0], metricname)
		value, _ := strconv.ParseFloat(word[1], 64)
		time, _ := strconv.ParseInt(word[2], 10, 64)

		item := Data{
			Name:  name,
			Value: value,
			Time:  time,
		}

		data = append(data, item)
	}

	return data
}

// nameString ...Format string from name in Data type.
func nameString(s, metricname string) string {
	sub := strings.Split(s, ".")

	name := strings.Join(sub[:len(sub)-1], "-")
	label := sub[len(sub)-1]

	if len(metricname) == 0 {
		return fmt.Sprintf("%s.%s", name, label)
	}

	return fmt.Sprintf("%s-%s.%s", metricname, name, label)
}
