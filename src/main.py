import pandas as pd
import numpy as np
from tabulate import tabulate
from sympy import sympify

def read_csv(filename):
    try:
        return pd.read_csv(filename)
    except FileNotFoundError:
        print(f"Error: File '{filename}' not found.")
        return None

def get_filled_tanks(tanks_data):
    return tanks_data[tanks_data['Wine'] != '/']

def blend_wines(tanks_data, formula_data):
    # Process Tank Lineup
    filled_tanks = get_filled_tanks(tanks_data)
    wine_totals = filled_tanks.groupby('Wine')['Volume'].sum().reset_index()
    sorted_wines = wine_totals.sort_values('Volume', ascending=False)
    total_volume = sorted_wines['Volume'].sum()
    sorted_wines['Percentage'] = sorted_wines['Volume'] / total_volume * 100

    # Display the sorted lineup
    print(tabulate(sorted_wines, headers='keys', tablefmt='psql'))

    # Process Formula and Generate Output
    formula = formula_data.set_index('Wine')['Percentage'].to_dict()

    blend_instructions = []
    for wine, percentage in formula.items():
        if wine in sorted_wines['Wine'].values:
            wine_percentage = sorted_wines.loc[sorted_wines['Wine'] == wine, 'Percentage'].iloc[0]
            tank_volume = percentage / 100 * total_volume
            tanks_needed = int(round(tank_volume / wine_percentage))

            # Find the tanks to transfer from
            source_tanks = filled_tanks.loc[filled_tanks['Wine'] == wine, 'TankID']
            source_tanks = source_tanks.sample(n=tanks_needed, replace=True)  # Randomly select the tanks if more needed than available

            # Find the tanks to fill
            empty_tanks = tanks_data.loc[tanks_data['Wine'] == '/', 'TankID']
            empty_tanks = empty_tanks.sample(n=tanks_needed, replace=False)  # Randomly select empty tanks if more needed than available

            # Generate blend instructions
            for i in range(tanks_needed):
                blend_instructions.append(f"{wine}: Transfer from Tank {source_tanks.iloc[i]} to Tank {empty_tanks.iloc[i]}")

    return blend_instructions

def generate_output_file(instructions):
    output_file = 'blend_instructions.txt'
    with open(output_file, 'w') as file:
        file.write('\n'.join(instructions))
    print(f"\nBlend instructions have been written to '{output_file}' successfully.")

if __name__ == "__main__":
    tanks_file = ("./random_data.csv")
    formula_file = ("./formula.csv")

    tanks_data = read_csv(tanks_file)
    formula_data = read_csv(formula_file)

    if tanks_data is None or formula_data is None:
        # Error occurred while reading files, exit the program
        exit()

    blend_instructions = blend_wines(tanks_data, formula_data)
    generate_output_file(blend_instructions)
