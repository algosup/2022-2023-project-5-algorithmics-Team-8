import csv
import random

def generate_random_volume():
    volume = random.randint(1, 100) * 5
    return volume

def fill_csv(filename, rows):
    with open(filename, 'w', newline='') as csvfile:
        fieldnames = ['TankID', 'Volume', 'Wine']
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
        writer.writeheader()

        for row in rows:
            writer.writerow(row)

if __name__ == '__main__':
    data = [
        {'TankID': 1, 'Volume': 100, 'Wine': 'Cabernet'},
        {'TankID': 2, 'Volume': 120, 'Wine': 'Merlot'},
        {'TankID': 3, 'Volume': 80, 'Wine': 'Chardonnay'},
        {'TankID': 4, 'Volume': 90, 'Wine': 'Pinot Noir'},
        {'TankID': 5, 'Volume': 110, 'Wine': 'Sauvignon Blanc'},
        {'TankID': 6, 'Volume': 100, 'Wine': 'Pinot Grigio'},
        {'TankID': 7, 'Volume': 100, 'Wine': 'Champagne'},
        {'TankID': 8, 'Volume': 100, 'Wine': 'Prosecco'}
    ]

    for i in range(9, 250):
        data.append({'TankID': i, 'Volume': generate_random_volume(), 'Wine': '/'})

    fill_csv('random_data.csv', data)

    print("CSV file created successfully.")
