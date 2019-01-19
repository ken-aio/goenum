// Copyright © 2019 @ken-aio <suguru.akiho@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

func defaultTemplateFile() string {
	return `# This file is goenum template file auto generated. if you see more detail, https://github.com/ken-aio/goenume
package {{.PackageName}}

// {{.Name}} {{.Description}}
type {{.EnumName}} int

// {{.EnumName}} enum
const (
{{range $i, $v := .Enums}}
	{{if eq $i 0}}
	{{$v}} {{.Name}} = iota
	{{else}}
	{{$v}}
	{{end}}
{{end}}
)

func (e {{Name}}) String() string {
	names := [...]string{
	{{range .EnumValues}}
		"{{.}}",
	{{end}}
	}
	return names[e]
}
`
}
