def append_to_file(data: list, filename:str):
    with open(filename, 'w+') as file:
        count = 0
        file.read()
        for soz in data:
            count += file.write(soz) 

    return count

print(append_to_file(['Qanday', 'siz', "Salom", "yaxsh", "rahmat"], "test.txt"))

def hello() ->str:
    return "fewfkjnewj"

hello()
