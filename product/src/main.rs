use std::cmp::Ordering;
use std::io;
use rand::Rng;

fn main() {
    guessing_game();
}


use std::collections::HashMap;
pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut mapping: HashMap<i32, i32> = HashMap::new();

    for (i, v) in nums.iter().enumerate() {
        if mapping.contains_key(&(target - v)) {
            return vec![mapping[&(target-v)], i as i32]
        } else {
            mapping.insert(*v, i as i32);
        }
    }
    return Vec::new()
}

fn guessing_game() {
    println!("Guessing the number");
    let mut guess_number: i32 = -1;
    let value = rand::thread_rng().gen_range(1..=100);

    while guess_number != value {
        let mut guess = String::new();
        println!("Please input your guess");

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        guess_number = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => { println!("Invalid number"); continue }
        };

        match guess_number.cmp(&value) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => println!("You win")
        }
    }
}