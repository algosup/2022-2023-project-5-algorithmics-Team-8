# <div align="center">Project 5 - ALGORITHMICS</div>

## <div align="center">Group 5</div>

# <div align="center">Technical Specifications</div>

<div align="right">Created on: 05/05/2023</div>
<div align="right">Last updated on: 23/05/2023</div>

<details>
<summary>Table of Content</summary>

- [Introduction](#introduction)
    - [Overview](#overview)
    - [Goals](#goals)
    - [Technical Requirements](#technical-requirements)
    - [Developpement Environment](#developement-environment)
    - [Out Of Scope](#out-of-scope)
- [Solutions](#solutions)
	- [Bazel](#bazel)
	- [Design](#design-work-in-progress)
	- [Test Plan](#test-plan)
	- [Release and Deployement Plan](#release-and-deployment-plan)
	- [Alternate Design](#alternate-design)
- [Further Considerations](#further-considerations)
	- [Cost Analysis](#cost-analysis)
	- [Security Considerations](#security-considerations)
	- [Privacy Considerations](#privacy-considerations)
	- [Accessibility Considerations](#accessiblity-considerations)
	- [Operational Considerations](#operational-considerations)
	- [Risks](#risks)
- [Success Evaluation](#success-eveluation)
- [Work](#work)
	- [Work Estimate And Timelines](#work-estimates-and-timelines)
	- [Prioritization](#prioritization)
	- [Milestones](#milestones)
- [End Matter](#end-matter)
	- [References](#references)
	- [Acknoledgements](#acknowledgments)
	- [Glossary](#glossary)

</details>

## **Introduction**

### *Overview*

Krug Champagne will open a new winery. They want to hasten and renew the process of blending the wine. They want to create a software that will calculate the blending steps for them. We know that there are 300 tanks available, and we know that the blending can call for more than 400 wines in the process.

### *Goals*

The whole point of this software is to save time during the blending process. In order to do that the software will determine the necessary steps to achieve a specific blend (formula given by the user). 
- We need the final process to be as short as possible (minimum number of steps).
- We need the software to be as fast as possible whatever the number of parameters.
- We need to take into account all requirements related to champagne blending and logistical issues.

### *Technical Requirements*

The software will run on the following minimum configuration:

|           |               |
|-----------|---------------|
|**OS**     | Windows 7     |
|**RAM**    | 2 Go          |
|**MEMORY** | 10 Mo         |

This configuration is subject to change depending on the development process.

### *Developement Environment*

The software will be developped on MAC OS with M1 chips, using the following technologies

| Technology  |  Version   |
|-------------|------------|
| CPP         |  11        |
| Bazel       |  6.2.0	   |
| google test |  1.12.0    |

This configuration is subject to changes if issues were to arise.

### *Out-Of-Scope*

The software will not support OS older than **Windows 7**, nor will it support older versions of CPP.

In the specific case where the input or process does not follow the documentation, troubleshooting won't be done by the software in order to try to format the data and make it work.

The result won't be dynamically updated when the input will change, The user will have to run the software again.

## **Solutions**

### Bazel

For this project we will use Bazel to compile build and test the code (paired with google test).

To setup Bazel on MacOS:
1. Install Homebrew

>```/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"```

2. Install Bazel via Homebrew

>```brew install bazel```

3. You can verify the installation using the following command

>```bazel --version```

4. To update Bazel, use the folling command

>```brew upgrade bazel```

If those steps don't work on your machine you can follow the another method of installation described on [Bazel's official website](https://bazel.build/install/os-x) in the install/os-x section.

To learn how to use Bazel with c++, you can follow this [tutorial](https://bazel.build/start/cpp) on the official website.

### *Design* [```Work in Progress```]

#### **Architecture design**

![architecture_design](/Documents/Image/Architecture_Design.png)

#### **Environment**

Struct Formula containing
- table of maps (intself containing a float and a string representing the percentage and the wine's name)

Struct Tank containing
- string ID/Name
- int Capacity
- bool Empty
- Struct Fomrula

Struct Step containing
- `Work in Progress`

#### **Algorithm**

In order to determine the blending steps the software will follow the algorythm defined below.

Inputs:
- A table containing Tank structs
- Struct Formula

Process:

1. Determine Output Tank(-s)
    1. Calculate total Quantity Of Each Wine
    2. Calculate total possible perfect output using the formula
    3. Determine Output Tank(-s) - closest to the possible maximum output or a division (multiple output tanks)
2. Transfer
    1. For each output Tank
        1. For each wine in the formula
            1. Calculate remaining empty space
            2. Calculate necessary quantity of wine
            3. Search for tank with matching size
				- If there is not tank with matching size, do a division (`work in progress`)
            4. Transfer / Create a step
				- If the remaining empty space can't allow the transfer, transfer the maximum possible quantity to the tank (finishing this one) and transfer the rest to the next output tank
            5. Calculate current formula for this output tank
3. Verify there is no half-full tanks
    1. circle through all tanks
    2. if tanks are half full either
        - Do a final mix with a formula that will be different from the asked one
        - Reunite the same wine in the same tank (if quantities are matching a tank's size)
        - Put those remaining wines in bottles

Output:
- A table containing steps structs

### *Test Plan*

The tests will be executed as described in the [test plan](/Documents/test_plan.md).

### *Release And Deployment Plan*

`To define`

### *Alternate Design*

This is an alternative design as the first one may contain issues or fatal flaws. it differs in environment and algorithm

#### **Environment**

Struct Formula containing:
- ...

Function named *Find_Tank* taking an input of type int which search in the tanks for one with the matching size
```cpp 
Find_Tank(int amount)
```

Function named *Split_Formula* taking an input of type Formula which split the formula (whatever it means)
```cpp
Split_Formula(Formula F)
```

Function named *Solve* taking tow inputs of type Formula and float
```cpp
Solve(Formula F, float a){
	Find_Tank(a);
	formulas = Split_Formula(F);
	foreach(Formula f : formulas)
	{
		Solve(f, ??);
	}
}
```

#### **Algorithm**

Inputs:
- F of type Formula
- a of type (float/double) representing the final quantity of wine

Algorithm:
1. Solve Formula using *Solve*
	2. Search for the output tank using *Find_Tank*
	3. Split the formula Using *Split_Formula*
	4. Foreach new formula
		- Solve the formula using *Solve*

## **Further Considerations**

### *Cost Analysis*

Regarding this software, every library or external work used is free. The only cost will be human and more precisely time.

### *Security Considerations*

No data will be saved direclty by the software, it will simply flow within for the time of the calculation and come out as the asked result. Moreover the data concerning wines quantity and tanks sizes are irrelevent to any people not working in the specific vineyard where the software is being used. The only sensitive data is the formula of the Champagne which won't be saved and will only be in the software for the time of the calculation.

The software will not be connected nor using the internet providing a first defense against leaks threats.

### *Privacy Considerations*

The software won't contain any personal data, as it is just a calculator.

### *Accessiblity Considerations*

The software will come with documents describing the required process to use it. Those documents will be made with accessibilty in mind in order for anyone to be able to use the software.

### *Operational considerations*

`To define`

### *Risks*

The risks evaluated by the team are the following:
- The software won't work because of either:
	- The hardware not being powerful enought
	- A major / critical bug preventing the software from working correctly
- `Work in Progress`

## **Success Eveluation**

Metrics:

- Error margin regarding the output formula(-s)
- Time of execution on minimum required hardware
- `Work in Progress`

## **Work**

### *Work estimates and timelines*

`To define`

### *Prioritization*

The priorities of the team for this project are the following:

1. Algorithm
2. Error Margin
3. Documents (User Manual)
4. User Iterface

### *Milestones*

For this project, the main milestones are the following:
1. Functioning algorithm
2. Algorithm error margin falls under specifications
3. Documents (User Manual)

To go further:

4. User Interface

## **End Matter**

### *References*

- [Functional Specification](/Documents/functional.md)
- [Test plan](/Documents/test_plan.md)
- [Bazel](https://bazel.build/)

`To define`

### *Acknowledgments*

`To define`

### *Glossary*

`To fill`