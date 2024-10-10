class Produto:
    def __init__(self, preco):
        self.preco = preco

    def __add__(self, other):
        return Produto(self.preco + other.preco)

produto1 = Produto(50)
produto2 = Produto(30)
produto3 = produto1 + produto2
print(produto3.preco)