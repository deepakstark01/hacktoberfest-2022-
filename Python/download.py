import urllib.request

def shortcut():
    urllib.request.urlretrieve("http://www.example.com/", "example.html")

def manual():
    with urllib.request.urlopen('http://www.example.com/') as input:
        html = input.read().decode('utf-8')
        with open('example.html', 'w') as output:
            output.write(html)

shortcut()
