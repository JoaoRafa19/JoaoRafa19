<h1 align="center">Hi 👋, I'm João</h1>
<h3 align="center">A passionate software developer from Brazil</h3>


- 🌱 I’m currently learning **Go**

- 💬 Ask me about **Flutter, Dart**

- 📫 How to reach me **joaopedrorafael19@gmail.com**

- 📄 Know about my experiences [http://www.linkedin.com/in/joaopedrorafael](http://www.linkedin.com/in/joaopedrorafael)


## Meet João, the Software Developer 👨‍💻

```go

package main

import "fmt"

// Joao represents a skilled software developer.
type Joao struct {
	Name         string
	Nationality  string
	Localization string
	Languages    []string
	Technologies map[string][]string
}

// NewJoao initializes João with default values if none are provided.
func NewJoao(name, nationality, localization string) Joao {
	if nationality == "" {
		nationality = "Brazilian"
	}
	if localization == "" {
		localization = "Belo Horizonte, MG"
	}

	return Joao{
		Name:         name,
		Nationality:  nationality,
		Localization: localization,
		Languages:    []string{"Portuguese", "English"},
		Technologies: map[string][]string{
			"Frontend": {"React.js", "Flutter", "Tailwind"},
			"Backend":  {"GoLang", "Django", "Node.js"},
			"Mobile":   {"Flutter"},
		},
	}
}

// Introduce prints a brief introduction about João.
func (j Joao) Introduce() {
	fmt.Printf("Hi, I'm %s! A %s developer from %s.\n", j.Name, j.Nationality, j.Localization)
	fmt.Printf("I speak %v and specialize in:\n", j.Languages)
	for stack, tools := range j.Technologies {
		fmt.Printf("- %s: %v\n", stack, tools)
	}
}

func main() {
	joao := NewJoao("João", "", "")
	joao.Introduce()
}


```


<h3 align="left">Connect with me:</h3>
<p align="left">
<a href="https://dev.to/joaorafa19" target="blank"><img align="center" src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/devto.svg" alt="joaorafa19" height="30" width="40" /></a>
<a href="https://linkedin.com/in/joaopedrorafael" target="blank"><img align="center" src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/linked-in-alt.svg" alt="joaopedrorafael" height="30" width="40" /></a>
<a href="https://instagram.com/joaorafa19" target="blank"><img align="center" src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/instagram.svg" alt="joaorafa19" height="30" width="40" /></a>
</p>

<h3 align="left">Languages and Tools:</h3>
<p align="left">
<a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a>
<a href="https://flutter.dev" target="_blank" rel="noreferrer">
  <img src="https://www.vectorlogo.zone/logos/flutterio/flutterio-icon.svg" alt="flutter" width="40" height="40"/> </a>


  <a href="https://git-scm.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg" alt="git" width="40" height="40"/> </a>

[](https://github-readme-stats.vercel.app/api?username=joaorafa19&show=reviews,discussions_started,discussions_answered,prs_merged,prs_merged_percentage&show_icons=true)

<p><img align="center" src="https://github-readme-streak-stats.herokuapp.com/?user=joaorafa19&" alt="joaorafa19" /></p>
