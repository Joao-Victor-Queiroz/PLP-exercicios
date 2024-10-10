class Animal:
    def som(self):
        pass

class Cachorro(Animal):
    def som(self):
        return "Au Au"

class Gato(Animal):
    def som(self):
        return "Miau"

def fazer_som(animais):
    for animal in animais:
        print(animal.som())

fazer_som([Cachorro(), Gato()])