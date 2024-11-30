<h1 align="center">Hi 👋, I'm João</h1>
<h3 align="center">A passionate software developer from Brazil</h3>


- 🌱 I’m currently learning **Go**

- 💬 Ask me about **Flutter, Dart**

- 📫 How to reach me **joaopedrorafael19@gmail.com**

- 📄 Know about my experiences [http://www.linkedin.com/in/joaopedrorafael](http://www.linkedin.com/in/joaopedrorafael)

```dart
class Joao extends Human implements SoftwareDeveloper {
  late final String? nationality;
  late final String? localization;

  @override
  Joao(String name, {String? nationality, String? localization}) : super(name) {
    nationality ??= "Brazilian";
    localization ??= "Belo Horizonte, MG";
  }

  final List<String> languages = [
    "Portuguese",
    "English",
  ];

  final Map<String, dynamic> technologies = {
    "Flutter": ["GetX", "Provider", "MobX", "Bloc"],
    "GoLang": ["GinGonic", "gorm"],
    "Dart": ["Shelf", "Dio"],
    "JavaScript": ["React.js", "Express"],
    "Python": ["Django", "Flask", "Qt", "Pygame"],
  };

  final Map<String, dynamic> databases = {
    "SQL": ["PostgreSQL", "MySQL", "SQLite"],
    "NoSQL": ["MongoDB", "Redis"],
  };

  String getTechologies(String stack) {
    switch (stack) {
      case "Flutter":
        return "GetX, Provider, MobX, Bloc, Firebase, Supabase, Dio";
      case "Frontend":
        return "Flutter, React.js, Tailwind";
      case "Backend":
        return "GoLang, Django, Flask, Node.js, Shelf";
      case "Mobile":
        return "Flutter";
    }
    return technologies[Random().nextInt(10)];
  }
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
