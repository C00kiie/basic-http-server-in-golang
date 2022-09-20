#!/usr/bin/env python3

import requests

# configuration

host = '127.0.0.1:8080'

# test 1
# route: /
# method: GET
def test_main_route():
    response = requests.get("http://{}/".format(host))
    if  response.status_code != 200:
        return -1
    print("test_main_route: passed")
    return 0


# test 2
# route : /hello
# method: GET
def test_hello_route():
    response = requests.get("http://{}/Hello".format(host))
    if  response.status_code != 200:
        print("test02: failed")
        return -1
    print("test_hello_route: passed")
    return 0

# test 3
def test_form_route():
    data = {
       "name": "Hassan",
       "age" : 21
    }
    response = requests.get("http://{}/Form".format(host), data=data)
    if response.status_code != 200:
        print("test03: failed")
        return -1
    print("test_form_route: passed")
    return 0

if __name__ == "__main__":

    test_main_route()
    test_form_route()
    test_hello_route()
