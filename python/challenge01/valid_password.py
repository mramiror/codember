def is_valid(pwd: str) -> bool:

    state = "start"
    last_digit = "0"
    last_char = "a"

    for c in pwd:
        if "0" <= c <= "9":
            if state == "char":
                return False
            if c < last_digit:
                return False
            last_digit = c
        elif "a" <= c <= "z":
            state = "char"
            if c < last_char:
                return False
            last_char = c
        else:
            return False
    return True


def main():
    valid = 0
    invalid = 0

    with open("log.txt", "r") as file:
        for line in file:
            pwd = line.strip()
            if not pwd:
                continue
            if is_valid(pwd):
                valid += 1
            else:
                invalid += 1
    
    print(f"Submit: {valid}true,{invalid}false")

if __name__ == "__main__":
    main()