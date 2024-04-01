class Dog:
  def __init__(self, name, age):
    self.name = name
    self.age = 0

  def bark(self):
    print("Woof! Woof!")

ivy = Dog("Ivy", 1.2)
oakley = Dog("Oakley", 3)
juniper = Dog("Juniper", 1.7)

ivy.bark()