# goDesignPatterns

jeux de tests en go sur les designs patterns.
for the fun car golang n'est pas vraiment oriente objet 


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

