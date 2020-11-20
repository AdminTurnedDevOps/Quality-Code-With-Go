class Cars:
    def __init__(self, make, model, year):
        self.make = make
        self.model = model
        self.year = year

myCar = Cars("Ford", "F150", "2020")

print(myCar.make)
