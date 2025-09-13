
class Vendendor ():
    def __init__ (self, nome):
        self.nome = nome
        self.vendas = 0

    def vendeu (self, vendas):
        self.vendas = vendas
    
    def bat_meta(self, meta ):
        if self.vendas >= meta:
            print (self.nome, 'a meta foi batida!')
        else:
            print (self.nome, 'A meta nao foi batida')

vendedor1 = Vendendor('Ikaro')
vendedor1.vendeu(1000)
vendedor1.bat_meta(500)