use std::fs::File;
use std::io::{self, BufRead, BufReader};

enum State {
    Start,
    Digit,
    Char,
}

fn is_valid(pwd: &str) -> bool {
    use State::*;

    let mut state = Start;
    let mut last_digit = '0';
    let mut last_char = 'a';

    for c in pwd.chars() {
        match c {
            '0'..='9' => {
                if let Char = state {
                    return false;
                }
                state = Digit;
                if c < last_digit {
                    return false;
            }
                last_digit = c;
            }
            'a'..='z' => {
                state = Char;
                if c < last_char {
                    return false;
                }
                last_char = c;
            }
            _ => {
                return false;
            }
        }
    }
    true
}

fn main() -> io::Result<()> {
    let file = File::open("log.txt")?;
    let reader = BufReader::new(file);

    let mut valid = 0;
    let mut invalid = 0;

    for line in reader.lines() {
        let pwd = line?;
        let pwd = pwd.trim();

        if pwd.is_empty() {
            continue;
        }
        if is_valid(pwd) {
            valid += 1;
        } else {
            invalid += 1;
        }
    }
    println!("Submit {}true{}false", valid, invalid);
    Ok(())
}