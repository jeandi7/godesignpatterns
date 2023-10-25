for the fun 

# Contexte

Jeux de test Go

# Options envisagées 💡

# Conseils 
La lecture (même rapide) du livre Design Patterns de Erich Gamma, Richard Helm, Ralph Johnson et John Vlissides


# Décision 🏆
jeux de tests DPO 

# Conséquences 
RAS


Les Design Patterns ci-dessous -->

```mermaid
classDiagram

   DP --> creational : creational 
   DP --> structural : structural
   DP --> behavior : behavior
   structural --> decorator : decorator
   creational --> singleton : singleton

```


Design Pattern Object Decorator ci dessous -->

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

Design Pattern Object Composite ci dessous -->

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
