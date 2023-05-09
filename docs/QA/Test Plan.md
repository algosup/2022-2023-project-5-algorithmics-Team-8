# Test Plan --- 

This document describes the tests that were, are and will be performed for an application that produces an list of instructions used for the blending or "assemblage" of champagne for Krug Champagne.

It provides instructions on how to reproduce the bugs and the environment in which they were found, as well as the inputs given, the expected and gotten result.

Since we are planing to create a portable application, we will test it on Windows, GNU/Linux and MacOS.

There also is the Test Report that contains a list of all the bugs found during the development process.

## Test report template

The test report is a csv file that contains the following information:
| Bug ID | Bug Description | Severity | Platform | Steps to reproduce | Expected result | Gotten result | Status |

Bug ID: The ID of the bug found
Bug Description: A short description of the bug
Severity: The severity of the bug is rated as Low, Severe or Fatal
Platform: The platform on which the bug can be reproduced
Steps to reproduce: The steps to reproduce the bug
Expected result: The actual result of the bug
Gotten result: The expected result of the action
Status: The status of the bug (Fixed/Not Fixed)

### Severity level

| Severity | Description |
| --- | --- |
| Low | The bug is a minor bug that does only affects minor parts of the Software  |
| Severe | The bug is a major bug that affects the final output and results in erronous results being given |
| Critical | The bug is a critical bug that cause a fatal exit of the Software |

It will be updated as we find new bugs and fix them during the development process.

On top of that, we will be using the [Bug Report Template](Bug%20Report%20Template.md) to report the bugs discovered directly on GitHub.


## Table of Contents

<details>
<summary>
Click to expand
</summary>

- [Test Plan ---](#test-plan----)
  - [Test report template](#test-report-template)
    - [Severity level](#severity-level)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
    - [Portability](#portability)
    - [Software Health](#software-health)
    - [Accuracy](#accuracy)
    - [Execution time](#execution-time)
    - [Input Table](#input-table)
    - [Data Processing](#data-processing)
    - [Main Algorithmic Loop](#main-algorithmic-loop)
    - [Instructions Generation](#instructions-generation)
    - [Output Table](#output-table)
</details>

## Requirements

### Portability

The following requirements are to be tested:

- [ ] The game must be able to run on different platforms
  - [ ] Windows
  - [ ] MacOS
  - [ ] GNU/Linux

The game fully runs on Windows and Mac os machines; however, we only tried to run it on a Nintendo Switch emulator and not on a real one.

### Software Health

- [ ] The Program should never crash, even with an erronous input
 - [ ] Correct Input given
 - [ ] Incorrect Input given
- [ ] There should not be any memory leaks

### Accuracy

- [ ] The Software should output a valid set of instructions that respect the following conditions
  - [ ] A tank should never be left partially empty
  - [ ] A tank should never be overfilled
  - [ ] The number of steps should be minimized
  - [ ] The amount of different inputs mixed in a tank at once should be minimized
  - [ ] The final mixed proportions should correspond to the given ones
  - [ ] The production of waste should be avoided


### Execution time 

- [ ] The Software should execute as fast as possible

Further information on execution time depending on platform and given inputs will be provided in the Test

### Input Table 

- [ ] The Software should be able to read a csv file given as input

### Data Processing

- [ ] The Software should be able to handle the resulting data read on the csv file
    - [ ] The Barrels should be listed with their ID, capacity, and the presence of wine if there is inside
    - [ ] The wines should be listed with their ID, and the tank they are in
    - [ ] The formula should be available with each composing wine and it's part of the final volume, either in quantity or percentage
    - [ ] If given, the final tanks or result quantity should be available

### Main Algorithmic Loop

- [ ] TODO

### Instructions Generation

- [ ] TODO

### Output Table

- [ ] TODO