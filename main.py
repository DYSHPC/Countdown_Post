#!/usr/bin/python3
#大家好我們是程式設計社~8/20 大園高中社團博覽會-茫茫人海遇見你 倒數4天！ 一定要來喔！
import time
import random
"""
for i in range(10):
    print(i,end="\r",flush=True)
    time.sleep(0.1)
"""
a="大家好我們是程式設計社~8/20 大園高中社團博覽會-茫茫人海遇見你 倒數4天！ 一定要來喔！"#47
b=["h","e","l","o","w","r","d"]
g=["大","家","好","我","們","是","程","式","設","計","社","~","8","/","2","0"," ","大","園","高","中","社","團","博","覽","會","-","茫","茫","人","海","遇","見","你"," ","倒","數","4","天","！"," ","一","定","要","來","喔","！"]
"""
for i in range(len(a)):
    print(a[i],end="\r",flush=True)
    time.sleep(0.1)
"""
c=""
gc=""
def cool(x,y,f):
    for i in range(20):
        #print(x,end="")
        #print(random.choice(y),end="\r",flush=True)
        k=random.choice(y)
        z=f"{x+k}"
        print(z,end="\r",flush=True)        
        time.sleep(0.01)
    print(x+f,end="\r",flush=True)
d="hello world"
for i in d:
    cool(c,b,i)
    c=c+i
print()
for i in a:
    cool(gc,g,i)
    gc=gc+i
print()

print("╱╱╭╮")
print("╱╱┃┃")
print("╭━╯┣╮╱╭┳━━┳━━╮")
print("┃╭╮┃┃╱┃┃╭╮┃╭━╯")
print("┃╰╯┃╰━╯┃╰╯┃╰━╮")
print("╰━━┻━╮╭┫╭━┻━━╯")
print("╱╱╱╭━╯┃┃┃")
print("╱╱╱╰━━╯╰╯")
print()
print("     ***      ***")
print("    *****    *****")
print("     ***      ***")
print("*                    *")
print(" *                  *")
print("  *                *")
print("   *              *")
print("    **************")
