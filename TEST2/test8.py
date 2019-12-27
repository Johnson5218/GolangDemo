class ClassA(object):

    """Docstring for ClassA. """

    def __init__(self):
        print('this is __init__')

    def add(self, x, y):
        return x + y
    def __del__(self):
        print('this is __del__')


a = ClassA()

print(a.add(1,2))
