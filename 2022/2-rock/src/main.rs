use std::fs;

fn main() {
    let content = fs::read_to_string("input.txt").unwrap();

    let score: u32 = content.as_str()
        .split("\n")
        .map(|x| x.split(" ").collect())
        .map(|a: Vec<&str>| {
            let mut score = match a[1] {
                "X" => 1,
                "Y" => 2,
                "Z" => 3,
                &_ => 0,
            };

            score += match a[0] {
                "A" => match a[1] {
                    "X" => 3,
                    "Y" => 6,
                    &_ => 0,
                },
                "B" => match a[1] {
                    "X" => 0,
                    "Y" => 3,
                    &_ => 6,
                },
                "C" => match a[1] {
                    "X" => 6,
                    "Y" => 0,
                    &_ => 3,
                },
                &_ => 0,
            };

            return score;
        })
        .sum();

    println!("{}", score)
}

