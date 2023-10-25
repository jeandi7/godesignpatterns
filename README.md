# goDesignPatterns

jeux de tests en go sur les designs patterns.

voir le cours Oodrive : Concevoir avec les Designs Patterns.

![image info](./ADR/DPO.png)

Le jeu de test porte sur les design patterns suivants :

## DPO : Singleton 

```mermaid

classDiagram 
note "On utilise le Singleton lorsque :
• il n'y a qu'une unique instance d'une classe et qu'elle
doit être accessible de manière connue
• une instance unique peut être sous-classée et que
les clients peuvent référencer cette extension sans
avoir à modifier leur code"

```

```mermaid
classDiagram 

class Singleton {
   +int static Instance
}

```
Deux implémentations sont proposées : 
- une implémetation naive
- une implémentation safe 

## DPO : Composite

```mermaid
classDiagram
note "On utilise Composite lorsque on veut :
• représenter une hiérarchie d'objets
• ignorer la différence entre un composant simple et un
composant en contenant d'autres. (interface
uniforme)"

```


```mermaid
classDiagram
   Component <|--Feuille
   Component <|--Composite
   Composite o-- Component


class Component {
    +Operation()*
    +add()*
}


class Composite {
    +Operation()*
    +add()*
}

```

## DPO : Decorator

Design Pattern Object Decorator ci dessous -->

```mermaid
classDiagram
note "On utilise Decorator lorsque :
• il faut ajouter des responsabilités dynamiquement et
de manière transparente
• il existe des responsabilités dont on peut se passer
• des extensions sont indépendantes et qu'il serait
impraticable d'implanter par sous-classage"

```


```mermaid
classDiagram
   Component <|--ConcreteComponent
   Decorator --|>  Component
   Decorator <| -- DecoratorB
   Decorator <| -- DecoratorA
   Component --o Decorator : comp

class Component {
    +Operation()*
}

%% note for DecoratorA "dans Operation : {\n ...\n super.Operation()\n addedBehavior()\n ...\n}"
class DecoratorA {
    +Operation()*
    +addedBehavior()*
}

```

## DPO : Interpreter

Design Pattern Object Interpreter ci dessous -->

```mermaid
classDiagram
note "On utilise Interpreter lorsqu'il faut interpréter un langage
et que :
• la grammaire est simple
• l'efficacité n'est pas un paramètre critique "
```

```mermaid
classDiagram

class Client  { }
class Context { }

class AbstractExpression {
    +Eval(Context)
 }

class TerminalExpression {
    +Eval(Context)
 }

class NonTerminalExpression { 
    +Eval(Context)
}

Client --> Context
Client --> AbstractExpression

TerminalExpression --|> AbstractExpression
NonTerminalExpression --|> AbstractExpression

NonTerminalExpression o-- AbstractExpression

```
![image info](./ADR/HP25.png)

- L'exemple fourni en golang calcule l'expression "3 4 + 2 -" d'où la photo
Vous comprenez que je ne pouvais pas mettre ma Ti58

## DPO : Adapter

```mermaid
classDiagram
note "On utilise l'Adapter lorsque on veut utiliser :
• une classe existante dont l'interface ne convient pas
• plusieurs sous-classes mais il est est coûteux de
redéfinir l'interface de chaque sous-classe en les
sous-classant. Un adapter peut adapter l'interface au
niveau du parent."

```

```mermaid
classDiagram

class Client  { }
class Cible { }


class Cible {
    +Request()
 }

class Adapter {
    +Request()
 }

note "Adapter.Request() = Adapté.methodSpec()"

class Adapté {
    +methodeSpec()
 }


Client --> Cible
Adapter --|> Cible
Adapter --> Adapté : adapté

```

## DPO : Visitor

Design Pattern Object Visitor ci dessous -->
```mermaid
classDiagram

note "On utilise Visitor lorsque :
• une structure d'objets contient de nombreuses
classes avec des interfaces différentes et on veut
appliquer des operations diverses sur ces objets.
• les structures sont assez stables, et les opération sur
leurs objets évolutives."
```

```mermaid
classDiagram

class Visitor {
    +visitConcretElementA()
 }

class ConcreteVisitorA {
    +visitConcretElementA()
 }

class ConcreteVisitorB {
    +visitConcretElementA()
 }

ConcreteVisitorA --|>  Visitor
ConcreteVisitorB --|>  Visitor

note for ConcreteElementA  "accept(Visitor v) { v.visitConcretElementA(this) }"

class Element {
    +accept(Visitor v) 
}


class ConcreteElementA {
    +accept(Visitor v)
 }

class ConcreteElementB {
    +accept(Visitor v)
}

ConcreteElementA --|>  Element
ConcreteElementB --|>  Element


```

## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin https://oogit.oodrive.net/architecture/godesignpatterns.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

- [ ] [Set up project integrations](https://oogit.oodrive.net/architecture/godesignpatterns/-/settings/integrations)

## Collaborate with your team

- [ ] [Invite team members and collaborators](https://docs.gitlab.com/ee/user/project/members/)
- [ ] [Create a new merge request](https://docs.gitlab.com/ee/user/project/merge_requests/creating_merge_requests.html)
- [ ] [Automatically close issues from merge requests](https://docs.gitlab.com/ee/user/project/issues/managing_issues.html#closing-issues-automatically)
- [ ] [Enable merge request approvals](https://docs.gitlab.com/ee/user/project/merge_requests/approvals/)
- [ ] [Set auto-merge](https://docs.gitlab.com/ee/user/project/merge_requests/merge_when_pipeline_succeeds.html)

## Test and Deploy

Use the built-in continuous integration in GitLab.

- [ ] [Get started with GitLab CI/CD](https://docs.gitlab.com/ee/ci/quick_start/index.html)
- [ ] [Analyze your code for known vulnerabilities with Static Application Security Testing(SAST)](https://docs.gitlab.com/ee/user/application_security/sast/)
- [ ] [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/ee/topics/autodevops/requirements.html)
- [ ] [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/ee/user/clusters/agent/)
- [ ] [Set up protected environments](https://docs.gitlab.com/ee/ci/environments/protected_environments.html)

***

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thank you to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README
Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a self-explaining name for your project.

## Description
Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.
