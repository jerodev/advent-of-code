use std::fs;

fn main() {
    let contents = fs::read_to_string("../input.txt").unwrap();

    let mut calories = contents.as_str()
        .split("\n\n")
        .map(|x| x.split("\n").map(|y| y.parse::<u32>().unwrap_or(0)).sum())
        .collect::<Vec<u32>>();

    calories.sort();
    calories.reverse();

    println!("{:?}", calories[0]);
    println!("{:?}", calories.iter().take(3).sum::<u32>());
}
