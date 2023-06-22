import os 
import random 



def generate_random_volume():
    volume = random.randint(1, 140) * 5
    return volume

def fill_file(filename, content):
    with open(filename, 'w') as file:
        file.write(content)

random_hectoliters = random.randint(1, 800) * 5

wine_list = ['Chardonnay', 'Pinot noir', 'Pinot meunier', 'Champagne', 'Merlot', 'Cabernet', 'Sauvignon blanc', 'Prosecco']

def random_formula(wine_list):
    number_of_wines = random.randint(2, 7)
    wine_list_copy = wine_list.copy()
    random.shuffle(wine_list_copy)
    wine_list_copy = wine_list_copy[:number_of_wines]
    percentages = []
    remaining_percentage = 100
    for i in range(number_of_wines - 1):
        if remaining_percentage == 0:
            break
        max_percentage = min(remaining_percentage, 100 - (number_of_wines - i - 1))
        percentage = random.randint(1, max_percentage - 1)  # Subtract 1 to exclude 0%
        percentages.append(percentage)
        remaining_percentage -= percentage
    percentages.append(remaining_percentage)  # Assign remaining percentage to the last wine within the loop
    formula = ''
    for i in range(number_of_wines):
        if i < len(percentages):  # Check the index against the length of percentages
            formula += wine_list_copy[i] + ';' + str(percentages[i]) + '\n'
    return formula

def random_tank():
    for i in range(1, 331):
        tank_id = '#' + str(i)
        volume = generate_random_volume()
        wine = random.choice(wine_list + ['/'] * 15 )
        yield tank_id + ';' + str(volume) + ';' + wine + '\n'



def content():
    instructions = [f"""! <- this is a comment, line starting with ! are ignored
! <- ceci est un commentaire, les lignes commençant par ! sont ignorées


! EN
! This is the config file for Sebastien, make a copy of it and follow the instructions to fill it out and
! prepare the execution.
! To run the program, drag your config file on the executable and wait for the program to finish, it will create
! a file called "result.txt" in the same folder as the executable, this file contains the result of the program.
! The inputs you enter should not have any unit or symbol, just the number
! follow the examples below but don't add "ex:" or "!" in your inputs

! FR
! Ceci est le fichier de configuration pour Sebastien, faites une copie et suivez les instructions pour le remplir et
! préparer l'exécution.
! Pour lancer le programme, glissez votre fichier de configuration sur l'exécutable.
! Attendez que le programme se termine, il créera un fichier appelé "result.txt" 
! dans le même dossier que l'exécutable, ce fichier contient le résultat du programme.
! Les entrées que vous entrez ne doivent pas avoir d'unité ou de symbole, uniquement le nombre
! suivez les exemples ci-dessous mais n'ajoutez pas "ex:" ou "!" dans vos entrées




! Total (hL), the total amount of wine you want to make, this is an upperbound
! Total (hL), la quantité totale de vin que vous voulez faire, c'est une borne supérieure

! ex: 200
                    
{random_hectoliters}

! Wines and their proportions need to be positive(%)  in mix 
! Name; percentage
! Vins et leur proportions qui ont besoin d'être positive(%) dans le mélange final
! Nom; pourcentage

! ex: Chardonnay; 80

{random_formula(wine_list)}

! Tanks; capacities and wines contained
! Tank ID (#number); capacity (hL); wine contained ('/' = empty)
! Cuves; capacités et vins contenus
! ID de la cuve (nombre); capacité (hL); vin contenu ('/' = vide)


! ex: #12; 300; Chardonnay
! ex2: #A3ER; 34.5; /

{"".join(random_tank())}"""]
    return ''.join(instructions)
        

if __name__ == '__main__':
    fill_file('./test.config', content())
    
