def steps_to_exit(nums: int) -> int:

    pos = 0
    steps = 0
    n = len(nums)

    while (pos >= 0 and pos < n):
        salto = nums[pos]
        nums[pos] += 1
        pos += salto
        steps += 1
    
    return steps

def main():
    total_steps = 0
    steps_last_line = 0

    with open("trace.txt", "r") as file:
        for line in file:
            l = line.strip()
            if not line:
                continue

            nums = [int(x) for x in l.split()]
            steps = steps_to_exit(nums)

            total_steps += steps
            steps_last_line = steps

    print(f"Submit: {total_steps}-{steps_last_line}")

if __name__ == "__main__":
    main()