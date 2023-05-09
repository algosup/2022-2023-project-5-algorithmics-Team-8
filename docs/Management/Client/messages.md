# Messages.md

This file contains all the messages sent to the client and his answers, for anonymity reasons we will refer to the client as "Client" and to the team as "Group 8".

Group 8
  14 h 15
Hello,
I am reaching out to you regarding the project.
I would greatly appreciate your insights on the following questions:
What are the minimum and maximum sizes of the tanks that we will be working with? Do you have any other information that could be useful to us about the tanks?
Is there a need for an interface between the software and the hardware?
Are there any restrictions or dangers that we need to be aware of in working with the tanks?
What type and format will the input have? Could you provide an example?
What type and format should the output have? Could you provide an example?
Does the input/mixing order matter, and is there a notion of timing that we need to take into account?
Will the software run on specific hardware ?
Does the hardware require a specific software or programming language?
Is there a maximum number of wines that can be blended in a single tank?
If any other specifics come to your mind do not hesitate to tell us about them.
Sincerely,
Team 5-8  


Client
  15 h 57
No minimum or maximum, tanks come in all shapes and sizes, usually expressed in hectoliters.
No, hardware is usually people doing the connections.
No, only danger is for the wine to oxidate if the tank is partially full/empty.
Configuration files with a given syntax, API, etc. there are no globally accepted standards so you should propose a mean of inputting values and getting results as part of the specification.
see 4.
No notion of timing or flow, order doesn't matter.
Portable is better but you can restrict it to one platform if the pros outweigh the cons.
No
No concrete restriction as it is all done with pipes that are connected. In practice, 2-3 makes sense, 5 is probably very high.


Group 8
  16 h 12
Thank you for the answer, we will continue informing ourselves on the process and keep you up to date.


Group 8 
  13 h 43
Hello,
We would like to know how the client currently handles the calculations, do they use specific software ?
After searching and watching videos over the past 2 days we reached the conclusion that they surely used excel.
We would like to base our solution and documentation on the software they  use to ease the user interactions.
Sincerely,
Team 8


Client
  15 h 16
Hello,
They don't use Excel (it would be extremally impractical), they use proprietary software with some of the calculation being done in SQL (again, not a very good choice).
Since the actual transfer from one tank to the next is being done manually, there is very little benefit in mimicking what their current software is doing.
What we expect from you is a proposal on what you think would be the best approach (configuration files/UI/API etc.) regardless on any legacy constraint.


Group 8 
  15 h 38
Ok thank you for the clarification.


Client
  16 h 32
Just to put things in perspective, they change the formula every year and the tank configuration every 10 years maybe. So, if the setup is a bit time consuming and not very user friendly, it's not the end of the world. I guess the people that will get annoyed by it might be you (doing it frequently for testing purpose) a lot more than actual users.


Group 8 
  9 h 12
We will put less focus on the UX then, thank you