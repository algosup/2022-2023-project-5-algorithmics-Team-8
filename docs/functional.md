# Krug Champagne -> Champagne Blending Calculator

## Functional Specifications

<details>
<summary>Table of contents</summary>

- [Krug Champagne -\> Champagne Blending Calculator](#krug-champagne---champagne-blending-calculator)
  - [Functional Specifications](#functional-specifications)
  - [Overview](#overview)
    - [Client](#client)
    - [Evaluation criteria](#evaluation-criteria)
    - [What is Krug Champagne?](#what-is-krug-champagne)
      - [How does making champagne work?](#how-does-making-champagne-work)
    - [Why automate the process?](#why-automate-the-process)
    - [Hardware minimum](#hardware-minimum)
    - [Language selected, Why CPP?](#language-selected-why-cpp)
  - [Audience](#audience)
  - [What is the problem](#what-is-the-problem)
  - [Personas](#personas)
    - [Persona 1](#persona-1)
    - [Persona 2](#persona-2)
    - [Persona 3](#persona-3)
  - [Laws and regulations](#laws-and-regulations)
  - [Ressources](#ressources)
  - [Requirements](#requirements)
  - [Functionalities](#functionalities)
    - [Must have](#must-have)
    - [Compatibility](#compatibility)
    - [Documentation](#documentation)
    - [Testing](#testing)
  - [Cost analysis](#cost-analysis)
    - [Human Resources](#human-resources)
  - [Privacy](#privacy)
  - [Non-functional requirements](#non-functional-requirements)
    - [Security](#security)
    - [Usability](#usability)
    - [Maintainability](#maintainability)
    - [Scalability](#scalability)
  - [Risks and assumptions](#risks-and-assumptions)
    - [Development environment](#development-environment)
      - [What is Bazel?](#what-is-bazel)
    - [What is Google Test?](#what-is-google-test)
    - [CPP Evolution](#cpp-evolution)
  - [Success criteria](#success-criteria)
  - [Out of scope](#out-of-scope)
  - [Glossary](#glossary)

</details>

## Overview

### Client

Krug Champagne (part of LVMH)

### Evaluation criteria

The software will be evaluated on correctness, how close the final product is to the wanted formula, comments and idiomatic style, minimum number of steps to get to the result, speed of the code, the end result does not contain any critical errors.

### What is Krug Champagne?

Krug Champagne is a French champagne house founded by Joseph Krug in 1843. It is based principally in Reims, the main city in France's Champagne region and is one of the famous Champagne houses that formed part of the Grande Marques.

#### How does making champagne work?

- 1) Harvesting: The grapes used to make champagne are typically harvested in September or October. The most commonly used grape varieties are Pinot Noir, Pinot Meunier, and Chardonnay.

- 2) Pressing: The grapes are pressed gently to extract the juice, which is then stored in tanks or barrels for fermentation.
  
- 3) First Fermentation: The first fermentation takes place in stainless steel tanks or oak barrels. Yeast is added to the juice, which converts the sugar into alcohol. This process takes about two weeks, and it produces a still wine with low alcohol content.

`________________________Project Scope_____________________________`

- 4) Blending: The still wines from different vineyards and grape varieties are blended together to create a consistent flavor profile.

`________________________End of  Scope_____________________________`

- 5) Second Fermentation: The blended wine is bottled, and a mixture of sugar and yeast is added to it. The bottles are then sealed with a crown cap and stored horizontally in cool, dark cellars. The second fermentation takes place in the bottle, and the carbon dioxide produced by the yeast is trapped in the wine, creating bubbles.

- 6) Aging: The bottles are aged on their lees (the dead yeast cells) for at least 15 months for non-vintage champagne, and at least three years for vintage champagne. During this time, the wine develops complex flavors and aromas.

- 7) Riddling: After aging, the bottles are gradually rotated and tilted to move the lees towards the neck of the bottle. This process is called riddling, and it is done manually or using a machine.

- 8) Disgorgement: Once the lees have settled in the neck of the bottle, the bottles are opened, and the frozen plug of lees is removed. This process is called disgorgement.

- 9) Dosage: After disgorgement, a mixture of wine and sugar (known as the dosage) is added to the champagne to adjust its sweetness level. The bottles are then corked, wired, and labeled.

### Why automate the process?

The blending process is a very important step in the production of champagne. It is a complex process that requires a lot of time and expertise. The blending process is done by a team of experts who taste the wine and decide which wines to blend together to create the final product.

They then need to reproduce the same blend on a large scale, which is a very difficult task because the process is lossy and the wine is not always the same when done by hand.

### Hardware minimum

Windows

- minimum of 2GB of RAM
- minimum of 10Mo of free disk space

### Language selected, Why CPP?

C++ is a general-purpose programming language created by Bjarne Stroustrup as an extension of the C programming language, or "C with Classes". The language has expanded significantly over time, and modern C++ now has object-oriented, generic, and functional features in addition to facilities for low-level memory manipulation. It is almost always implemented as a compiled language, and many vendors provide C++ compilers, including the Free Software Foundation, LLVM, Microsoft, Intel, Oracle, and IBM, so it is available on many platforms.

C++ is fast and efficient, and it provides a lot of control over the hardware, which makes it a good choice for systems programming and embedded systems. It is also a good choice for applications that require high performance.

C++ is supported by a large number of libraries, which makes it easy to write portable code. It is also supported by a large number of tools, including compilers, debuggers, and profilers.

## Audience

The audience of the solution is the team of experts who are responsible for the blending process. They are the ones who will use the software to create the blending steps.

The audience is any winery that uses excel to calculate the blending steps because this is the most widely used software in the industry.

## What is the problem

The blending process is a very important step in the production of champagne. It is a complex process that requires a lot of time and expertise. The blending process is done by a team of experts who taste the wine and decide which wines to blend together to create the final product. When applying the blending process to a large scale, the team of experts needs to reproduce the same blend on a large scale, which is a very difficult task because the process is lossy and the wine is not always the same when done by hand.

## Personas

### Persona 1

```
Name: Mark Smith
Age: 45
Job: Wine maker
Place: Reims, France

Behaviors: Mark is a wine maker, he creates high quality wine for his clients, to do so he tries multiple combinations of grapes and fermentation processes to create the best wine possible.

Description:
Mark is creating wine made of his 4 different vineyards and he wants to create new wines.

Needs & goals: Mark wants to create a new wine with a new combination of grapes and fermentation processes, he wants to create a new wine that will be the best wine he ever created. He wants to stop losing time and money calculating the blending he found per scale.

Use case: Mark just created his new formula and now wants to mass produce it, he needs to calculate the blending for 10000 bottles of 75cl taking the size of his Tanks into account.
```

### Persona 2

```
Name: Bob Parker
Age: 29
Job: Wine Consultant
Place: Reims, France

Description:
Bob is a wine consultant, he helps wine makers to create new wines and to improve their existing wines.

Needs & goals: Bob wants to help wine makers to create new wines and to improve their existing wines, he wants to help them to create the best wine possible.

Use case: Bob wants to fasten the process of creating new wines in one of his clients winery.
```

### Persona 3

```
Name: Lara Doe
Age: 35
Job: Wine Maker
Place: Reims, France

Description:
Lara has a winery and she wants to create new wines.

Needs & goals: Lara always wants to try new combinations and each year she creates 5 new wines.

Use case: Lara has created 5 new wines and wants to mass produce them.
```

## Laws and regulations

There are no laws or regulations that apply to this project.

## Ressources

- [L'art de l'assemblage entre les cuves](https://youtube.com/watch?v=qyIdO7LEjhc&feature=share)

## Requirements

The Software needs to be focused on :

- never partially filling tanks
- never overfilling tanks
- minimising the number of tanks used
- minimising the number of transfers between tanks
- The closest result to the target blend
- The software needs to be able to intake csv files containing the data of the wine, tank sizes, and the percentage to take from each type of wine.
- The software needs to be able to output a csv file containing the blending steps.

## Functionalities

### Must have

- The software needs to be able to intake csv files containing the data of the wine, tank sizes, and the percentage to take from each type of wine.
- The software needs to be able to calculate the blending steps.
- The software needs to be able to output a csv file containing the blending steps.

### Compatibility

- The software needs to be compatible with Windows 7 minimum and We should extend to as many of the following platforms as possible:
- Windows
- Mac OS
- Linux

### Documentation

The software needs to be heavily documented, we need to make sure that anyone can understand how to use it.
The clients are not supposed to have any real technical knowledge, so we need to make sure that the documentation is easy to understand.

### Testing

The software needs to be tested thoroughtly, we need to make sure that the software is working as intended and that it does not contain any critical errors.

We will use google test to test the software as well as TDD in the development process.

## Cost analysis

The cost of the project is the time spent on it, the project is not supposed to be sold, so there is no cost of production.

### Human Resources

The project is supposed to be developed by 5 people, each person will be spending around 10 hours per week on the project, so the total cost of the project is 50 hours per week.

The project is supposed to be developed during 9 weeks, so the total cost of the project is 450 hours or per person 90 hours.

## Privacy

The software will not be storing any data, so there is no privacy concern.

## Non-functional requirements

### Security

The project is not supposed to be used by anyone outside of the company, so there is no security concern.

### Usability

The software needs to be easy to use, we need to make sure that the user can understand how to use it even if not required.

The software should adapt to the most common use cases, the client uses excel so we need to make sure that the software is easy to use and that it is compatible with excel.

### Maintainability

The software should not need to be adapted since the industry of wine is stable and does not change much.

### Scalability

The software needs to be able to handle large amounts of data, we need to make sure that the software can handle large amounts of data.

## Risks and assumptions

The main risks are:

- The software is not used by the clients
- The software misinterprets the data
- The software is not well documented and is confusing for the user,
- The software creates memory leaks and other memory-related issues during the execution of the program.
- The software uses outdated technologies and encounters critical issues during the execution of the program.

With all of these concerns, the software might miss the deadline or fall out of use.

To prevent this from happening, we need to make sure that the software is well documented and that the software is tested thoroughtly throughout the development process.

### Development environment

The software will be developed only on Mac OS M1, we will use the version 11 of CPP

| Technology | Version |
| --- | --- |
| CPP | 11 |
| QT | 6.5 |
| Bazel | pre-7.0.0 |
| google test | 1.13.0 |

#### What is Bazel?

Bazel is a build system that is used to build C++ projects. It is similar to CMake, but it is more modern and it is easier to use.
It is the compiler used with google test.

### What is Google Test?

Google Test is a unit testing library for the C++ programming language, based on the xUnit architecture. The library is released under the BSD 3-clause license.

### CPP Evolution

The software will be developed using the latest version of CPP, we will use the version 11 of CPP because it is the latest version of CPP and it is the version that supports the most features.
CPP is an old language, so it is not supposed to change much in the future.

## Success criteria

The software will be considered a success if it meets the [requirements](#requirements) and the minimum [functionalities](#functionalities) of the project.

We aim to adapt to the client as much as possible.

## Out of scope

The software will not support hardware older than windows 7, we will not support older versions of CPP, we will not support older versions of QT.

If the process does not follow the documentation, no troobleshooting will be done by the software to try to temper with the data to make it work.

The software will not be dynamically updating the result if the data changes, the user will have to re-run the software to get the new result.

## Glossary

- **CPP**: C++ is a general-purpose programming language created by Bjarne Stroustrup as an extension of the C programming language, or "C with Classes". The language has expanded significantly over time, and modern C++ now has object-oriented, generic, and functional features in addition to facilities for low-level memory manipulation. It is almost always implemented as a compiled language, and many vendors provide C++ compilers, including the Free Software Foundation, LLVM, Microsoft, Intel, Oracle, and IBM, so it is available on many platforms.

- **QT**: Qt is a free and open-source widget toolkit for creating graphical user interfaces as well as cross-platform applications that run on various software and hardware platforms such as Linux, Windows, macOS, Android or embedded systems with little or no change in the underlying codebase while still being a native application with native capabilities and speed.

- **Bazel**: Bazel is a build system that is used to build C++ projects. It is similar to CMake, but it is more modern and it is easier to use.

- **Google Test**: Google Test is a unit testing library for the C++ programming language, based on the xUnit architecture. The library is released under the BSD 3-clause license.

- **LVMH**: LVMH Moët Hennessy Louis Vuitton SE, also known as LVMH, is a French multinational corporation and conglomerate specializing in luxury goods, headquartered in Paris, France. The company was formed in 1987 under the merger of fashion house Louis Vuitton with Moët Hennessy, a company formed after the 1971 merger between the champagne producer Moët & Chandon and Hennessy, the cognac manufacturer.

- **Krug Champagne**: Krug Champagne is a French champagne house founded by Joseph Krug in 1843. It is based principally in Reims, the main city in France's Champagne region and is one of the famous Champagne houses that formed part of the Grande Marques.

- **Champagne**: Champagne is a sparkling wine produced from grapes grown in the Champagne region of France following rules that demand, among other things, secondary fermentation of the wine in the bottle to create carbonation, specific vineyard practices, sourcing of grapes exclusively from specific parcels in the Champagne appellation and specific pressing regimes unique to the region. Some use the term Champagne as a generic term for sparkling wine, but in many countries, it is illegal to label any product Champagne unless it both comes from the Champagne region and is produced under the rules of the appellation.

- **Tank**: A tank is a large container for holding liquids or gases.

- **Excel**: Microsoft Excel is a spreadsheet program included in the Microsoft Office suite of applications. Spreadsheets present tables of values arranged in rows and columns that can be manipulated mathematically using both basic and complex arithmetic operations and functions.