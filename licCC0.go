package main

func init() {
	addLicense("cc0", license{
		Name:        "CC0 1.0",
		Description: "CC0 1.0",
		URL:         "http://creativecommons.org/publicdomain/zero/1.0/",
		Template: `<WORK'S NAME> by {{ .Author }}

To the extent possible under law, the person who associated CC0 with
<WORK'S NAME> has waived all copyright and related or neighboring rights
to <WORK'S NAME>.

You should have received a copy of the CC0 legalcode along with this
work.  If not, see <http://creativecommons.org/publicdomain/zero/1.0/>.`,
	})
}
