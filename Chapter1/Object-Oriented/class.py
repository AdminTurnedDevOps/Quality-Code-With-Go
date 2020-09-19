class Cars:
    def __init__(self, Make, Model, Year):
        self.Make = Make
        self.Model = Model
        self.Year = Year

myCar = Cars("Ford", "F150", "2020")

print(myCar.Make)