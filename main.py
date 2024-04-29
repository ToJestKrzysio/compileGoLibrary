import ctypes

lib = ctypes.CDLL("./hello.so")
hello = lib.Hello()
