
def pow(x,n):
    ret = 1
    cur = x
    while n:
        if n&1:
            ret *= cur
        cur *= cur
        n=n>>1
    return ret




# if __name__ == "__main__":
#     print(pow(3,15))