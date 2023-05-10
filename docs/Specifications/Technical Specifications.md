# <div align="center">Project 5 - ALGORITHMICS</div>

## <div align="center">Group 5</div>

- Project Manager: [Lucas AUBARD]()
- Program Manager: [Louis DE LAVENNE](https://github.com/Louis-de-Lavenne-de-Choulot)
- Technical Leader: [Arthur LEMOINE](https://github.com/arthur-lemo1ne)
- Quality Ensurance: [Laurent BOUQUIN](https://github.com/laurentbouquin)
- Software Engineer: [Robin DEBRY](https://github.com/robin-debry)

<div align="right">Created on: 05/05/2023</div>
<div align="right">Last updated on: 10/05/2023</div>

# <div align="center">Technical Specifications</div>

<details>
<summary>Table of Content</summary>

- [Introduction](#introduction)
    - [Overview](#overview)
    - [Goals](#goals)
    - [Technical Requirements](#technical-requirements)
    - [Developpement Environment](#developement-environment)
    - [Out Of Scope](#out-of-scope)
- [Solutions](#solutions)
- [Further Considerations](#further-considerations)
- [Success Evaluation](#success-eveluation)
- [Work](#work)
- [End Matter](#end-matter)

</details>

## **Introduction**

### *Overview*

Krug Champagne will open a new winery. They want to hasten and renew the process of blending the wine. They want to create a software that will calculate the blending steps for them. We know that there are 300 tanks available, and we know that the blending can call for more than 400 wines in the process.

### *Goals*

The whole point of this software is to save time during the blending process. In order to do that the software will determine the necessary steps to acheive a specific blend (formula given by the user). 
- We need the final process to be as short as possible (minimum numer of steps).
- We need the software to be as fast as possible whatever the number of parameters.
- We need to take into account all requirements related to champagne blending and logistical issues.

### *Technical Requirements*

The software will run on the following configuration:

|           |               |
|-----------|---------------|
|**OS**     | Windows 7     |
|**RAM**    | 2 Go          |
|**MEMORY** | 10 Mo         |

This configuration is subject to change in the future depending on the developement process.

### *Developement Environment*

The software will be developped on MAC OS with M1 chips, using the following technologies

| Technology  |  Version   |
|-------------|------------|
| CPP         |  11        |
| Bazel       |  pre-7.0.0 |
| google test |  1.13.0    |

This configuration is subject to changes if issues were to arise.

### *Out-Of-Scope*

The software will not support OS older than **Windows 7**, nor will it support older version of CPP.

In the case that the input or process does not follow documentation, troubleshooting won't be done by the software in order to try to format the data and make it work.

Result won't be dynamically updated when the input will change, user will have to run the software again.

## **Solutions**

### *Design* [```Work in Progress```]

#### <u>Environment</u>

Struct Tank containing
- string ID/Name
- int Capacity
- bool Empty
- string Wine_contained

Struct Formula containing
- double Ch (Chardonai)
- double M (pinot Meniers)
- double N (Pinot Noir)

Struct Step containing
- Nothing for the moment

#### <u>Algorythm</u>

In order to determine the blending steps the software will follow the algorythm defined bellow.

Inputs:
- A table containing Tank structs
- Struct Formula

Process:

1. Determine Output Tank(-s)
    1. Total Quantity Of Each Wine
    2. Total Possible Output
    3. Determine Output Tank(-s)
2. Transer
    1. Select Tank Output1
    2. Determine Tank Origin1
    3. Determine Tank Origin2
    4. Determine Tank Origin3
    5. Repeat for each Output Tank
3. Verify there is no half full tanks
    1. circle through all tanks
    2. if tanks are half full either
        - Do a final mix with a formula that will be wrong
        - Reunite the same wine in the same tank (if quantities are good)
        - Put those wines in bottle

Output:
- A table containing steps structs

### *Test Plan*

The tests will be executed as descirbed in the [test plan](/Documents/test_plan.md).

### *Release And Deployement Plan*

`To define`

### *Alternate Design*

Second algorithm differing in parameters:

`To define`

## **Further Considerations**

### *Cost Analysis*

`To define`

### *Security Considerations*

`To define`

### *Privacy Considerations*

`To define`

### *Accessiblity Considerations*

`To define`

### *Operational considerations*

`To define`

### *Risks*

`To define`

## **Success Eveluation**

Metrics:

`To define`

## **Work**

### *Work estimates and timelines*

`To define`

### *Prioritization*

`To define`

### *Milestones*

`To define`

## **End Matter**

### *References*

`To define`

### *Acknowledgments*

`To define`