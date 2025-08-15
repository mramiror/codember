use std::fs::File;
use std::io::{self, BufRead, BufReader};

fn steps_to_exit(mut nums: Vec<i32>) -> i32 {
    let mut pos: i32 = 0;
    let mut steps: i32 = 0;
    let n: i32 = nums.len() as i32;

    while pos >= 0 && pos < n {
        let salto = nums[pos as usize];
        nums[pos as usize] += 1;
        pos += salto;
        steps += 1;
    }

    steps
}

fn main() -> io::Result<()> {
    let file = File::open("trace.txt")?;
    let reader = BufReader::new(file);

    let mut total_steps: i32 = 0;
    let mut steps_last_line: i32 = 0;

    for line in reader.lines() {
        let l = line?;
        let line = l.trim();

        if line.is_empty() {
            continue;
        }

        let nums: Vec<i32> = line
            .split_whitespace()
            .filter_map(|s| s.parse::<i32>().ok())
            .collect();

        let steps = steps_to_exit(nums);

        total_steps += steps;
        steps_last_line = steps;
    }

    println!("submit {}-{}", total_steps, steps_last_line);

    Ok(())
}
