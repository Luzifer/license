package main

func init() {
	addLicense("apache", license{
		Name:        "Apache v2.0",
		Description: "Apache License, Version 2.0",
		URL:         "http://www.apache.org/licenses/LICENSE-2.0",
		Template: `Copyright {{ .Year }} {{ .Author }}

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.`,
	})
}
