main:
    mov a, 4
    mov b, 1
    mov c, string
    la d, length
    syscall

    mov a, 1
    mov b, 0
    syscall

string:
    db "Hello, world!"
    db 10

length:
    db 14
